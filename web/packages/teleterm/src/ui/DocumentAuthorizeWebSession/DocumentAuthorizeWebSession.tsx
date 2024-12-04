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

import { Text, Alert, ButtonPrimary, H1, ButtonText } from 'design';
import Flex from 'design/Flex';
import { useAsync, Attempt } from 'shared/hooks/useAsync';
import { processRedirectUri } from 'shared/redirects';
import { Cluster } from 'gen-proto-ts/teleport/lib/teleterm/v1/cluster_pb';
import { DeviceConfirmationToken } from 'gen-proto-ts/teleport/devicetrust/v1/device_confirmation_token_pb';

import Document from 'teleterm/ui/Document';
import { routing } from 'teleterm/ui/uri';
import { retryWithRelogin } from 'teleterm/ui/utils';
import { useAppContext } from 'teleterm/ui/appContextProvider';
import * as types from 'teleterm/ui/services/workspacesService';
import { WebSessionRequest } from 'teleterm/ui/services/workspacesService';
import { useWorkspaceContext } from 'teleterm/ui/Documents';

export function DocumentAuthorizeWebSession(props: {
  doc: types.DocumentAuthorizeWebSession;
  visible: boolean;
}) {
  const ctx = useAppContext();
  const { documentsService } = useWorkspaceContext();
  const rootCluster = ctx.clustersService.findCluster(props.doc.rootClusterUri);
  const [attempt, authorize] = useAsync(async () => {
    const {
      response: { confirmationToken },
    } = await retryWithRelogin(ctx, props.doc.rootClusterUri, () =>
      ctx.clustersService.authenticateWebDevice(
        props.doc.rootClusterUri,
        props.doc.webSessionRequest
      )
    );
    const url = buildAuthorizedSessionUrl(
      rootCluster,
      props.doc.webSessionRequest,
      confirmationToken
    );
    // This endpoint verifies the token and "upgrades" the web session and redirects to "/web".
    window.open(url);
  });
  const clusterName = routing.parseClusterName(props.doc.rootClusterUri);
  const canAuthorize = rootCluster.loggedInUser?.isDeviceTrusted;

  async function authorizeAndCloseDocument() {
    const [, error] = await authorize();
    if (!error) {
      closeAndNotify();
    }
  }

  function openUnauthorizedAndCloseDocument() {
    const url = buildUnauthorizedSessionUrl(
      rootCluster,
      props.doc.webSessionRequest
    );
    window.open(url);
    closeAndNotify();
  }

  function closeAndNotify() {
    documentsService.close(props.doc.uri);
    ctx.notificationsService.notifyInfo(
      'Web session has been opened in the browser'
    );
  }

  return (
    <Document visible={props.visible}>
      <Flex
        flexDirection="column"
        maxWidth="680px"
        width="100%"
        mx="auto"
        mt="4"
        px="5"
      >
        <H1 mb="4">Authorize Web Session</H1>
        <Flex flexDirection="column" gap={3}>
          {!canAuthorize && (
            <Alert
              mb={0}
              details={
                <>
                  To authorize a web session, you must first{' '}
                  <a
                    href="https://goteleport.com/docs/admin-guides/access-controls/device-trust/guide/#step-22-enroll-device"
                    target="_blank"
                  >
                    enroll your device.
                  </a>{' '}
                  Then log out of the app, log back in, and try again.
                </>
              }
            >
              This device is not trusted
            </Alert>
          )}
          {attempt.status === 'error' && (
            <Alert mb={0} details={attempt.statusText}>
              Could not authorize the session
            </Alert>
          )}
          <Text>
            Would you like to authorize a device trust web session for{' '}
            <b>{clusterName}</b>?
            <br />
            The session will automatically open in a new browser tab.
          </Text>
          <Flex flexDirection="column" gap={2}>
            <ButtonPrimary
              disabled={
                !canAuthorize ||
                attempt.status === 'processing' ||
                attempt.status === 'success'
              }
              size="large"
              onClick={authorizeAndCloseDocument}
            >
              {getButtonText(attempt)}
            </ButtonPrimary>
            <ButtonText
              disabled={
                attempt.status === 'processing' || attempt.status === 'success'
              }
              onClick={openUnauthorizedAndCloseDocument}
            >
              Open Session Without Device Trust
            </ButtonText>
          </Flex>
        </Flex>
      </Flex>
    </Document>
  );
}

const confirmPath = 'webapi/devices/webconfirm';

function buildAuthorizedSessionUrl(
  rootCluster: Cluster,
  webSessionRequest: WebSessionRequest,
  confirmationToken: DeviceConfirmationToken
): string {
  const { redirectUri } = webSessionRequest;
  let url = `https://${rootCluster.proxyHost}/${confirmPath}?id=${confirmationToken.id}&token=${confirmationToken.token}`;
  if (redirectUri) {
    url = `${url}&redirect_uri=${redirectUri}`;
  }
  return url;
}

function buildUnauthorizedSessionUrl(
  rootCluster: Cluster,
  webSessionRequest: WebSessionRequest
): string {
  const processedRedirectUri = processRedirectUri(
    webSessionRequest.redirectUri
  );
  return `https://${rootCluster.proxyHost}${processedRedirectUri}`;
}

function getButtonText(attempt: Attempt<void>): string {
  switch (attempt.status) {
    case '':
    case 'error':
      return 'Authorize Session';
    case 'processing':
      return 'Authorizing Session…';
    case 'success':
      return 'Session Authorized';
  }
}
