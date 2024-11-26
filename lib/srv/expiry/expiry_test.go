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
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/gravitational/teleport/api/types"
	"github.com/gravitational/teleport/lib/auth"
	"github.com/gravitational/teleport/lib/events/eventstest"
	"github.com/gravitational/teleport/lib/utils"
	"github.com/jonboulle/clockwork"
	"github.com/stretchr/testify/require"
)

func TestExpiry(t *testing.T) {
	logger := utils.NewSlogLoggerForTests()
	mockEmitter := &eventstest.MockRecorderEmitter{}
	clock := clockwork.NewFakeClock()

	dataDir := t.TempDir()
	// hostUUID := uuid.New().String()
	// login := uuid.New().String()

	authServer, err := auth.NewTestAuthServer(auth.TestAuthServerConfig{
		ClusterName: "root.example.com",
		Dir:         dataDir,
		Clock:       clock,
	})

	require.NoError(t, err)
	t.Cleanup(func() { authServer.Close() })

	cfg := &Config{
		Log:         logger,
		Emitter:     mockEmitter,
		AccessPoint: authServer.AuthServer,
		Clock:       clock,
	}

	ctx := context.Background()

	expiry, err := New(ctx, cfg)
	require.NoError(t, err)

	go func() {
		err := expiry.Start()
		require.NoError(t, err)
	}()

	req1Name := uuid.New().String()
	req1, err := types.NewAccessRequest(req1Name, "someUser", "someRole")
	require.NoError(t, err)
	req1.SetExpiry(clock.Now().Add(scanInterval))
	err = authServer.AuthServer.CreateAccessRequest(ctx, req1)
	require.NoError(t, err)

	req2Name := uuid.New().String()
	req2, err := types.NewAccessRequest(req2Name, "someUser", "someRole")
	require.NoError(t, err)
	req2.SetExpiry(clock.Now().Add(scanInterval * 10))
	err = authServer.AuthServer.CreateAccessRequest(ctx, req2)
	require.NoError(t, err)

	// Wait until expiry scanner picks up first request and expires it
	clock.BlockUntil(2)
	clock.Advance(scanInterval*5 + expiryInterval*2) // Scanner will have seen request by here

	require.Eventually(t, func() bool {
		return len(mockEmitter.Events()) == 1
	}, time.Second*5, time.Second/2)

	// Wait until expiry scanner picks up second request and expires it
	clock.Advance(scanInterval*20 + expiryInterval*2) // Scanner will have seen request by here

	// Ensure second request has expired
	require.Eventually(t, func() bool {
		clock.BlockUntil(2)
		return len(mockEmitter.Events()) == 2
	}, time.Second*5, time.Second/2)
}
