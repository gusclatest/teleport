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

package service

import (
	"github.com/gravitational/teleport"
	"github.com/gravitational/teleport/api/types"
	"github.com/gravitational/teleport/lib/srv/expiry"
	"github.com/gravitational/trace"
)

// Expiry is a service that manages explicit experation of resources for cases
// where actions e.g. audit events are expected on expiry.
type Expiry interface{}

func (process *TeleportProcess) initExpiry() {
	process.RegisterWithAuthServer(types.RoleExpiry, ExpiryIdentityEvent)
	process.RegisterCriticalFunc("expiry.init", process.initExpiryService)
}

func (process *TeleportProcess) initExpiryService() error {
	logger := process.logger.With(teleport.ComponentKey,
		teleport.Component(teleport.ComponentExpiry, process.id))

	conn, err := process.WaitForConnector(ExpiryIdentityEvent, logger)
	if conn == nil {
		return trace.Wrap(err)
	}

	accessPoint, err := process.newLocalCacheForExpiry(conn.Client,
		[]string{teleport.ComponentExpiry})
	if err != nil {
		return trace.Wrap(err)
	}

	// asyncEmitter makes sure that sessions do not block
	// in case if connections are slow
	asyncEmitter, err := process.NewAsyncEmitter(conn.Client)
	if err != nil {
		return trace.Wrap(err)
	}

	expiryService, err := expiry.New(process.ExitContext(), &expiry.Config{
		Log:         logger,
		Emitter:     asyncEmitter,
		AccessPoint: accessPoint,
	})
	if err != nil {
		return trace.Wrap(err)
	}

	process.OnExit("expiry.stop", func(payload interface{}) {
		logger.InfoContext(process.ExitContext(), "Shutting down.")
		if expiryService != nil {
			expiryService.Stop()
		}
		if asyncEmitter != nil {
			warnOnErr(process.ExitContext(), asyncEmitter.Close(), logger)
		}
		warnOnErr(process.ExitContext(), conn.Close(), logger)
		logger.InfoContext(process.ExitContext(), "Exited.")
	})

	process.BroadcastEvent(Event{Name: ExpiryReady, Payload: nil})

	if err := expiryService.Start(); err != nil {
		return trace.Wrap(err)
	}
	logger.InfoContext(process.ExitContext(), "Expiry service has successfully started")

	// The Expiry service doesn't have heartbeats so we cannot use them to check health.
	// For now, we just mark ourselves ready all the time on startup.
	// If we don't, a process only running the Expiry service will never report ready.
	process.OnHeartbeat(teleport.ComponentExpiry)(nil)

	if err := expiryService.Wait(); err != nil {
		return trace.Wrap(err)
	}
	return nil
}
