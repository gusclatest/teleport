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
	"errors"
	"log/slog"
	"sync"
	"time"

	"github.com/gravitational/teleport"
	"github.com/gravitational/teleport/api/types"
	apievents "github.com/gravitational/teleport/api/types/events"
	"github.com/gravitational/teleport/lib/auth/authclient"
	"github.com/gravitational/teleport/lib/services"
	"github.com/gravitational/trace"
)

const (
	// handlerTimeout is used to bound the execution time of watcher event handler.
	handlerTimeout = time.Second * 5
)

// Config provides configuration for the expiry server.
type Config struct {
	// Log is the logger.
	Log *slog.Logger
	// Emitter is events emitter, used to submit discrete events
	Emitter apievents.Emitter
	// AccessPoint is a expiry access point
	AccessPoint authclient.ExpiryAccessPoint
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

	resources map[string]types.Resource

	// accessRequestWatcher is a accessRequest watcher.
	accessRequestWatcher *services.AccessRequestWatcher

	accessRequests *accessRequestSyncMap
}

// New initializes a expiry service
func New(ctx context.Context, cfg *Config) (*Service, error) {
	if err := cfg.CheckAndSetDefaults(); err != nil {
		return nil, trace.Wrap(err)
	}

	localCtx, cancelfn := context.WithCancel(ctx)
	s := &Service{
		Config:         cfg,
		ctx:            localCtx,
		cancelfn:       cancelfn,
		accessRequests: newAccessRequestSyncMap(),
	}
	return s, nil
}

// Start starts the expiry service.
func (s *Service) Start() error {
	s.Log.InfoContext(s.ctx, "Starting expiry service")
	if err := s.initWatcher(s.ctx); err != nil {
		return trace.Wrap(err)
	}
	go func() {
		for {
			select {
			case <-s.ctx.Done():
				return
			}
		}
	}()
	return nil
}

func (s *Service) initWatcher(ctx context.Context) error {
	var err error
	s.accessRequestWatcher, err = services.NewAccessRequestWatcher(s.ctx, services.AccessRequestWatcherConfig{
		ResourceWatcherConfig: services.ResourceWatcherConfig{
			Component:    teleport.ComponentExpiry,
			Logger:       s.Log,
			Client:       s.AccessPoint,
			MaxStaleness: time.Minute,
		},
		AccessRequestGetter: s.AccessPoint,
	})
	if err != nil {
		return trace.Wrap(err)
	}

	go func() {
		defer func() {
			s.Log.DebugContext(ctx, "Expiry service access request resource watcher finished")
		}()

		for {
			select {
			case accessRequestChanges := <-s.accessRequestWatcher.AccessRequestsC:
				if err := s.handleAccessRequestChanges(ctx, accessRequestChanges); err != nil {
					s.Log.ErrorContext(ctx, "new ar changes", "error", err.Error())
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return nil
}

func (s *Service) handleAccessRequestChanges(ctx context.Context, requests types.AccessRequests) error {
	for _, ar := range requests {
		s.accessRequests.Insert(ar)
		go func() {
			select {
			case <-ctx.Done():
				return
			case <-time.After(time.Until(ar.Expiry())):
				s.tryExpireRequest(ctx, ar)
			}
		}()
	}
	return nil
}

func (s *Service) tryExpireRequest(ctx context.Context, req types.AccessRequest) error {
	if !time.Now().Before(s.accessRequests.Get(req.GetName()).Expiry()) {
		// Request expiry has been extended since it was first seen.
		return nil
	}
	s.accessRequests.Delete(req)
	if err := s.AccessPoint.DeleteAccessRequest(ctx, req.GetName()); err != nil {
		return trace.Wrap(err)
	}
	return nil
}

// Stop stops the expiry service.
func (s *Service) Stop() {
	s.accessRequestWatcher.Close()
	s.cancelfn()
}

// Wait will block while the service is running.
func (s *Service) Wait() error {
	<-s.ctx.Done()
	if err := s.ctx.Err(); err != nil && !errors.Is(err, context.Canceled) {
		return trace.Wrap(err)
	}
	return nil
}

type accessRequestSyncMap struct {
	mu             sync.RWMutex
	accessRequests map[string]types.AccessRequest
}

func newAccessRequestSyncMap() *accessRequestSyncMap {
	return &accessRequestSyncMap{
		accessRequests: make(map[string]types.AccessRequest),
	}
}

func (m *accessRequestSyncMap) Insert(v types.AccessRequest) {
	m.mu.Lock()
	m.accessRequests[v.GetName()] = v
	m.mu.Unlock()
}

func (m *accessRequestSyncMap) Delete(v types.AccessRequest) {
	m.mu.Lock()
	delete(m.accessRequests, v.GetName())
	m.mu.Unlock()
}

func (m *accessRequestSyncMap) Get(accessRequestName string) types.AccessRequest {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.accessRequests[accessRequestName]
}
