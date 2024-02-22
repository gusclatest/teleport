/**
 * Teleport
 * Copyright (C) 2024 Gravitational, Inc.
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

import { useCallback, useState } from 'react';
import { useTheme } from 'styled-components';
import { useAsync } from 'shared/hooks/useAsync';
import { Flex, Box, Alert, Text, ButtonPrimary, ButtonSecondary } from 'design';
import * as icons from 'design/SVGIcon';

import { StyledTable, StyledTableWrapper } from 'design/DataTable/StyledTable';

import Document from 'teleterm/ui/Document';

import { useAppContext } from 'teleterm/ui/appContextProvider';

import type * as docTypes from 'teleterm/ui/services/workspacesService';

export function DocumentVnet(props: {
  visible: boolean;
  doc: docTypes.DocumentVnet;
}) {
  const { doc } = props;
  const { vnet } = useAppContext();
  const theme = useTheme();
  const [status, setStatus] = useState<'running' | 'stopped'>('stopped');

  const [startAttempt, startVnet] = useAsync(
    useCallback(async () => {
      await vnet.start(doc.rootClusterUri);
      setStatus('running');
    }, [vnet, doc.rootClusterUri])
  );

  const [stopAttempt, stopVnet] = useAsync(
    useCallback(async () => {
      await vnet.stop(doc.rootClusterUri);
      setStatus('stopped');
    }, [vnet, doc.rootClusterUri])
  );

  return (
    <Document visible={props.visible}>
      <Flex
        flexDirection="column"
        alignItems="flex-start"
        gap={2}
        p={4}
        maxWidth="680px"
        mx="auto"
        width="100%"
      >
        <Flex width="100%" justifyContent="space-between" alignItems="baseline">
          <Text typography="h3">VNet</Text>
          {status === 'running' && (
            <ButtonSecondary
              onClick={stopVnet}
              title="Stop VNet for teleport-local.dev"
              disabled={stopAttempt.status === 'processing'}
            >
              Stop VNet
            </ButtonSecondary>
          )}
          {status === 'stopped' && (
            <ButtonPrimary
              onClick={startVnet}
              disabled={startAttempt.status === 'processing'}
            >
              Start VNet
            </ButtonPrimary>
          )}
        </Flex>

        {startAttempt.status === 'error' && (
          <Alert>{startAttempt.statusText}</Alert>
        )}
        {stopAttempt.status === 'error' && (
          <Alert>{stopAttempt.statusText}</Alert>
        )}

        <Text>
          Proxying connections made to .teleport-local.dev.internal,
          .company.private
        </Text>

        <Flex width="100%" flexDirection="column" gap={1}>
          <Text typography="h4">Recent connections</Text>
          <StyledTableWrapper borderRadius={1}>
            <StyledTable>
              <tbody>
                <tr>
                  <td>
                    <Flex gap={2} alignItems="center">
                      <Flex
                        width="12px"
                        height="12px"
                        bg="success.main"
                        borderRadius="50%"
                        justifyContent="center"
                        alignItems="center"
                        css={`
                          flex-shrink: 0;
                        `}
                      ></Flex>{' '}
                      httpbin.company.private
                    </Flex>
                  </td>
                  <td></td>
                </tr>

                <tr>
                  <td>
                    <Flex gap={2} alignItems="center">
                      <Flex
                        width="12px"
                        height="12px"
                        bg="success.main"
                        borderRadius="50%"
                        justifyContent="center"
                        alignItems="center"
                        css={`
                          flex-shrink: 0;
                        `}
                      ></Flex>{' '}
                      tcp-postgres.teleport-local.dev.internal
                    </Flex>
                  </td>
                  <td></td>
                </tr>
                <tr>
                  <td>
                    <Flex gap={2} alignItems="center">
                      <Flex
                        width="12px"
                        height="12px"
                        bg="error.main"
                        borderRadius="50%"
                        justifyContent="center"
                        alignItems="center"
                      >
                        {/* TODO(ravicious): Make SVGIcon support passing color strings as fill. */}
                        <icons.ErrorIcon
                          size={12}
                          fill={theme.colors.levels.surface}
                        />
                      </Flex>{' '}
                      grafana.teleport-local.dev.internal
                    </Flex>
                  </td>

                  {/*
                      TODO(ravicious): Solve this without using an arbitrary max-width if possible,
                      perhaps switch to a flexbox instead of using a table?
                    */}
                  <td
                    css={`
                      max-width: 320px;
                      overflow: hidden;
                      text-overflow: ellipsis;
                      white-space: nowrap;
                    `}
                  >
                    DNS query for "grafana.teleport-local.dev.internal" in
                    custom DNS zone failed: no matching Teleport app and
                    upstream nameserver did not respond
                  </td>
                </tr>
                <tr>
                  <td>
                    <Flex gap={2} alignItems="center">
                      <Flex
                        width="12px"
                        height="12px"
                        bg="unset"
                        borderRadius="50%"
                        justifyContent="center"
                        alignItems="center"
                        css={`
                          flex-shrink: 0;
                        `}
                      ></Flex>{' '}
                      dumper.teleport-local.dev.internal
                    </Flex>
                  </td>
                  <td></td>
                </tr>
              </tbody>
            </StyledTable>
          </StyledTableWrapper>
        </Flex>
      </Flex>
    </Document>
  );
}
