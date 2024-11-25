/*
 * Teleport
 * Copyright (C) 2024  Gravitational, Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package expiry

import (
	"context"
	"log/slog"
	"sync"
	"time"

	"github.com/gravitational/teleport/api/types"
	apievents "github.com/gravitational/teleport/api/types/events"
	"github.com/gravitational/teleport/api/utils/retryutils"
	"github.com/gravitational/teleport/lib/auth/authclient"
	"github.com/gravitational/teleport/lib/events"
	"github.com/gravitational/teleport/lib/services"
	"github.com/gravitational/trace"
	"github.com/jonboulle/clockwork"
)

const (
	semaphoreName       = "expiry"
	semaphoreExpiration = time.Minute
	semaphoreJitter     = time.Minute

	// scanInterval is the interval at which the expiry checker scans for access requests
	scanInterval = time.Minute * 5
	// expiryInterval is the interval at which access request deletions are rate limited to
	expiryInterval = time.Second * 10
)

// Config provides configuration for the expiry server.
type Config struct {
	// Log is the logger.
	Log *slog.Logger
	// Emitter is an events emitter, used to submit discrete events.
	Emitter apievents.Emitter
	// AccessPoint is a expiry access point.
	AccessPoint authclient.ExpiryAccessPoint
	// Clock is the clock used to calculate expiry.
	Clock clockwork.Clock
	// HostID is a unique ID of this host.
	HostID string
}

// CheckAndSetDefaults checks required fields and sets default values.
func (c *Config) CheckAndSetDefaults() error {
	if c.Emitter == nil {
		return trace.BadParameter("no Emitter configured for expiry")
	}
	if c.AccessPoint == nil {
		return trace.BadParameter("no AccessPoint configured for expiry")
	}
	return nil
}

// Service is a expiry service.
type Service struct {
	*Config

	ctx      context.Context
	cancelfn context.CancelFunc

	toExpireMut         sync.RWMutex
	requestsToExpire    []types.AccessRequest
	requestsToExpireSet map[string]bool
}

// New initializes a expiry service
func New(ctx context.Context, cfg *Config) (*Service, error) {
	if err := cfg.CheckAndSetDefaults(); err != nil {
		return nil, trace.Wrap(err)
	}

	s := &Service{
		Config:              cfg,
		ctx:                 ctx,
		requestsToExpireSet: make(map[string]bool),
	}
	return s, nil
}

// Start starts the expiry service.
func (s *Service) Start() error {

	semCfg := services.SemaphoreLockConfigWithRetry{
		SemaphoreLockConfig: services.SemaphoreLockConfig{
			Service: s.AccessPoint,
			Params: types.AcquireSemaphoreRequest{
				SemaphoreKind: types.KindAccessRequest,
				SemaphoreName: semaphoreName,
				MaxLeases:     1,
				Holder:        s.HostID,
			},
			Expiry: semaphoreExpiration,
			Clock:  s.Clock,
		},
		Retry: retryutils.LinearConfig{
			Clock:  s.Clock,
			First:  time.Second,
			Step:   semaphoreExpiration / 2,
			Max:    semaphoreExpiration,
			Jitter: retryutils.DefaultJitter,
		},
	}

	// Work through the expired requests at a rate of 1 every expiryInterval
	go func() {
		for {
			select {
			case <-s.ctx.Done():
				return
			case <-time.After(retryutils.HalfJitter(expiryInterval)):
				lease, err := services.AcquireSemaphoreLockWithRetry(
					s.ctx,
					semCfg,
				)
				if err != nil {
					s.Log.WarnContext(s.ctx, "aquiring semaphore", "error", err)
					return
				}

				s.toExpireMut.Lock()
				if len(s.requestsToExpire) > 0 {
					req := s.requestsToExpire[0]
					s.requestsToExpire = s.requestsToExpire[1:]
					delete(s.requestsToExpireSet, req.GetName())
					if err := s.expireRequest(s.ctx, req); err != nil {
						s.Log.WarnContext(s.ctx, "error expiring access request", "error", err)
						s.requestsToExpire = append(s.requestsToExpire, req)
						s.requestsToExpireSet[req.GetName()] = true
					}
				}
				s.toExpireMut.Unlock()

				ctx, cancel := context.WithCancel(lease)
				defer cancel()
				lease.Stop()
				if err := lease.Wait(); err != nil {
					s.Log.WarnContext(ctx, "error cleaning up semaphore", "error", err)
				}
			}
		}
	}()

	for {
		select {
		case <-s.ctx.Done():
			return nil
		case <-time.After(retryutils.HalfJitter(scanInterval)):
			lease, err := services.AcquireSemaphoreLockWithRetry(
				s.ctx,
				semCfg,
			)
			if err != nil {
				return trace.Wrap(err)
			}

			requests, err := s.AccessPoint.GetAccessRequests(s.ctx, types.AccessRequestFilter{})
			if err != nil {
				return trace.Wrap(err)
			}

			if err := s.processAccessRequests(s.ctx, requests); err != nil {
				s.Log.WarnContext(s.ctx, "error expiring access requests", "error", err)
			}

			ctx, cancel := context.WithCancel(lease)
			defer cancel()
			lease.Stop()
			if err := lease.Wait(); err != nil {
				s.Log.WarnContext(ctx, "error cleaning up semaphore", "error", err)
			}
		}
	}
}

func (s *Service) processAccessRequests(ctx context.Context, requests []types.AccessRequest) error {
	for _, req := range requests {
		if s.Clock.Now().Before(req.Expiry()) {
			continue
		}
		s.addToExpiryQueue(ctx, req)
	}
	return nil
}

func (s *Service) addToExpiryQueue(ctx context.Context, req types.AccessRequest) {
	s.toExpireMut.Lock()
	defer s.toExpireMut.Unlock()
	if _, ok := s.requestsToExpireSet[req.GetName()]; !ok {
		s.requestsToExpireSet[req.GetName()] = true
		s.requestsToExpire = append(s.requestsToExpire, req)
	}
}

func (s *Service) expireRequest(ctx context.Context, req types.AccessRequest) error {
	var annotations *apievents.Struct
	if sa := req.GetSystemAnnotations(); len(sa) > 0 {
		var err error
		annotations, err = apievents.EncodeMapStrings(sa)
		if err != nil {
			return trace.Wrap(err)
		}
	}
	expiry := req.Expiry()
	event := &apievents.AccessRequestExpire{
		Metadata: apievents.Metadata{
			Type: events.AccessRequestExpireEvent,
			Code: events.AccessRequestExpireCode,
		},
		ResourceMetadata: apievents.ResourceMetadata{
			Expires: req.GetAccessExpiry(),
		},
		Roles:                req.GetRoles(),
		RequestedResourceIDs: apievents.ResourceIDs(req.GetRequestedResourceIDs()),
		RequestID:            req.GetName(),
		RequestState:         req.GetState().String(),
		Reason:               req.GetRequestReason(),
		MaxDuration:          req.GetMaxDuration(),
		Annotations:          annotations,
		ResourceExpiry:       &expiry,
	}
	if err := s.Emitter.EmitAuditEvent(ctx, event); err != nil {
		return trace.Wrap(err)
	}

	if err := s.AccessPoint.DeleteAccessRequest(ctx, req.GetName()); err != nil {
		return trace.Wrap(err)
	}
	return nil
}
