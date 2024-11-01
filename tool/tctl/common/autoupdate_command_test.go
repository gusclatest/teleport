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

package common

import (
	"bytes"
	"context"
	"github.com/gravitational/trace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/gravitational/teleport/lib/auth/authclient"
	"github.com/gravitational/teleport/lib/utils"
	"github.com/gravitational/teleport/tool/teleport/testenv"
)

// TestClientToolsAutoUpdateCommands verifies all commands related to client auto updates, by
// enabling/disabling auto update, setting the target version and retrieve it.
func TestClientToolsAutoUpdateCommands(t *testing.T) {
	ctx := context.Background()
	log := utils.NewSlogLoggerForTests()
	process := testenv.MakeTestServer(t, testenv.WithLogger(log))
	authClient := testenv.MakeDefaultAuthClient(t, process)

	// Enable mode to check that resources were modified.
	_, err := runAutoUpdateCommand(t, ctx, authClient, []string{"client-tools", "configure", "--set-mode=enabled"})
	require.NoError(t, err)

	config, err := authClient.GetAutoUpdateConfig(ctx)
	require.NoError(t, err)
	assert.Equal(t, "enabled", config.Spec.Tools.Mode)

	// Disable mode to check that resources were modified.
	_, err = runAutoUpdateCommand(t, ctx, authClient, []string{"client-tools", "configure", "--set-mode=disabled"})
	require.NoError(t, err)

	config, err = authClient.GetAutoUpdateConfig(ctx)
	require.NoError(t, err)
	assert.Equal(t, "disabled", config.Spec.Tools.Mode)

	// Set target version for auto update.
	_, err = runAutoUpdateCommand(t, ctx, authClient, []string{"client-tools", "get"})
	require.Error(t, err)
	assert.True(t, trace.IsNotFound(err))

	_, err = runAutoUpdateCommand(t, ctx, authClient, []string{"client-tools", "configure", "--set-target-version=1.2.3"})
	require.NoError(t, err)

	version, err := authClient.GetAutoUpdateVersion(ctx)
	require.NoError(t, err)
	assert.Equal(t, "1.2.3", version.Spec.Tools.TargetVersion)

	getBuf, err := runAutoUpdateCommand(t, ctx, authClient, []string{"client-tools", "get"})
	require.NoError(t, err)
	response := mustDecodeJSON[versionResponse](t, getBuf)
	assert.Equal(t, "1.2.3", response.TargetVersion)
}

func runAutoUpdateCommand(t *testing.T, ctx context.Context, client *authclient.Client, args []string) (*bytes.Buffer, error) {
	var stdoutBuff bytes.Buffer
	command := &AutoUpdateCommand{
		stdout: &stdoutBuff,
	}
	args = append([]string{"autoupdate"}, args...)
	return &stdoutBuff, runCommandWithContext(t, ctx, client, command, args)
}
