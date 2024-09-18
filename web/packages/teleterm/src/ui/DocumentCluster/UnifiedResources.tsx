/**
 * Teleport
 * Copyright (C) 2023  Gravitational, Inc.
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

import { useCallback, useEffect, useMemo, memo } from 'react';

import {
  UnifiedResources as SharedUnifiedResources,
  useUnifiedResourcesFetch,
  UnifiedResourcesQueryParams,
  SharedUnifiedResource,
  UnifiedResourcesPinning,
  getResourceAvailabilityFilter,
  ResourceAvailabilityFilter,
} from 'shared/components/UnifiedResources';
import {
  DbProtocol,
  formatDatabaseInfo,
  DbType,
} from 'shared/services/databases';

import { Flex, ButtonPrimary, Text, Link, H1 } from 'design';

import * as icons from 'design/Icon';
import Image from 'design/Image';
import stack from 'design/assets/resources/stack.png';

import { Attempt } from 'shared/hooks/useAsync';

import { ShowResources } from 'gen-proto-ts/teleport/lib/teleterm/v1/cluster_pb';
import { DefaultTab } from 'gen-proto-ts/teleport/userpreferences/v1/unified_resource_preferences_pb';

import { NodeSubKind } from 'shared/services';
import { waitForever } from 'shared/utils/wait';
import {
  RequestableResourceKind,
  Option,
} from 'shared/components/AccessRequests/NewRequest/resource';

import {
  UserPreferences,
  ListUnifiedResourcesRequest,
} from 'gen-proto-ts/teleport/lib/teleterm/v1/service_pb';

import { UnifiedResourceResponse } from 'teleterm/ui/services/resources';
import { useAppContext } from 'teleterm/ui/appContextProvider';
import * as uri from 'teleterm/ui/uri';
import { useWorkspaceContext } from 'teleterm/ui/Documents';
import { useWorkspaceLoggedInUser } from 'teleterm/ui/hooks/useLoggedInUser';
import { useConnectMyComputerContext } from 'teleterm/ui/ConnectMyComputer';

import { retryWithRelogin } from 'teleterm/ui/utils';
import {
  DocumentClusterQueryParams,
  DocumentCluster,
  DocumentClusterResourceKind,
} from 'teleterm/ui/services/workspacesService';
import { getAppAddrWithProtocol } from 'teleterm/services/tshd/app';

import { useStoreSelector } from 'teleterm/ui/hooks/useStoreSelector';

import {
  ConnectServerActionButton,
  ConnectKubeActionButton,
  ConnectDatabaseActionButton,
  ConnectAppActionButton,
  AccessRequestButton,
} from './ActionButtons';
import { useResourcesContext, ResourcesContext } from './resourcesContext';
import { useUserPreferences } from './useUserPreferences';
import { KubeRequestButton } from 'shared/components/AccessRequests/NewRequest/KubeRequestButton';
import {
  extractResourceRequestProperties,
  toResourceRequest,
} from '../services/workspacesService/accessRequestsService';
import {
  getKubeNamespaceId,
  KubeNamespaceRequest,
} from 'shared/components/AccessRequests/NewRequest/kube';

export function UnifiedResources(props: {
  clusterUri: uri.ClusterUri;
  docUri: uri.DocumentUri;
  queryParams: DocumentClusterQueryParams;
}) {
  const { clustersService, tshd } = useAppContext();
  const { userPreferencesAttempt, updateUserPreferences, userPreferences } =
    useUserPreferences(props.clusterUri);
  const { documentsService, rootClusterUri, accessRequestsService } =
    useWorkspaceContext();
  const rootCluster = clustersService.findCluster(rootClusterUri);

  const addedResources = useStoreSelector(
    'workspacesService',
    useCallback(
      state => {
        const pending =
          state.workspaces[rootClusterUri]?.accessRequests.pending;
        if (pending?.kind === 'resource') {
          return pending.resources;
        }
      },
      [rootClusterUri]
    )
  );
  const { onResourcesRefreshRequest } = useResourcesContext();
  const loggedInUser = useWorkspaceLoggedInUser();

  const { unifiedResourcePreferences } = userPreferences;

  const mergedParams = useMemo<UnifiedResourcesQueryParams>(
    () => ({
      kinds: props.queryParams.resourceKinds,
      sort: props.queryParams.sort,
      pinnedOnly: unifiedResourcePreferences.defaultTab === DefaultTab.PINNED,
      search: props.queryParams.advancedSearchEnabled
        ? ''
        : props.queryParams.search,
      query: props.queryParams.advancedSearchEnabled
        ? props.queryParams.search
        : '',
    }),
    [
      props.queryParams.advancedSearchEnabled,
      props.queryParams.resourceKinds,
      props.queryParams.search,
      props.queryParams.sort,
      unifiedResourcePreferences.defaultTab,
    ]
  );

  const integratedAccessRequests = useMemo<IntegratedAccessRequests>(() => {
    // Ideally, we would have a cluster loading status that would tell us,
    // whether the cluster data from the auth server has been loaded.
    // However, since we don't have that,
    // we use the `showResources` status as an indicator.
    if (rootCluster.showResources === ShowResources.UNSPECIFIED) {
      return { supported: 'unknown' };
    }
    if (!rootCluster.features?.advancedAccessWorkflows) {
      return { supported: 'no' };
    }
    return {
      supported: 'yes',
      availabilityFilter: getResourceAvailabilityFilter(
        userPreferences.unifiedResourcePreferences.availableResourceMode,
        rootCluster.showResources === ShowResources.REQUESTABLE
      ),
    };
  }, [
    rootCluster.features?.advancedAccessWorkflows,
    rootCluster.showResources,
    userPreferences.unifiedResourcePreferences.availableResourceMode,
  ]);

  const { canUse: hasPermissionsForConnectMyComputer, agentCompatibility } =
    useConnectMyComputerContext();

  const isRootCluster = props.clusterUri === rootClusterUri;
  const canAddResources = isRootCluster && loggedInUser?.acl?.tokens.create;
  let discoverUrl: string;
  if (isRootCluster) {
    discoverUrl = `https://${rootCluster.proxyHost}/web/discover`;
  }

  const canUseConnectMyComputer =
    isRootCluster &&
    hasPermissionsForConnectMyComputer &&
    agentCompatibility === 'compatible';

  const openConnectMyComputerDocument = useCallback(() => {
    documentsService.openConnectMyComputerDocument({ rootClusterUri });
  }, [documentsService, rootClusterUri]);

  const onParamsChange = useCallback(
    (newParams: UnifiedResourcesQueryParams): void => {
      documentsService.update(props.docUri, (draft: DocumentCluster) => {
        const { queryParams } = draft;
        queryParams.sort = newParams.sort;
        queryParams.resourceKinds =
          newParams.kinds as DocumentClusterResourceKind[];
        queryParams.search = newParams.search || newParams.query;
        queryParams.advancedSearchEnabled = !!newParams.query;
      });
    },
    [documentsService, props.docUri]
  );

  const requestStarted = accessRequestsService.getAddedItemsCount() > 0;

  const getAddedItemsCount = useCallback(() => {
    return accessRequestsService.getAddedItemsCount();
  }, [accessRequestsService]);

  async function fetchKubeNamespaces({
    kubeCluster,
    search,
  }: KubeNamespaceRequest): Promise<Option[]> {
    console.log('--- cluster uri : ', props.clusterUri);
    const { response } = await tshd.listKubernetesResources({
      kubeCluster,
      search,
      limit: 5,
      searchAsRoles: true,
      startKey: '',
      kubeResourceType: 'namespace',
      clusterUri: props.clusterUri,
    });
    return response.items.map(i => {
      const resourceRequest = toResourceRequest({
        kind: 'namespace',
        clusterUri: props.clusterUri,
        resourceId: i.name,
        resourceName: kubeCluster,
      });
      return {
        kind: 'namespace',
        value: resourceRequest.resource.uri,
        label: i.name,
      };
    });
  }

  const getAccessRequestButton = useCallback(
    (resource: UnifiedResourceResponse) => {
      const isResourceAdded = addedResources?.has(resource.resource.uri);

      const showRequestButton =
        integratedAccessRequests.supported === 'yes' &&
        (integratedAccessRequests.availabilityFilter.mode === 'requestable' ||
          resource.requiresRequest ||
          // If we are currently making an access request, all buttons change to
          // add to request.
          requestStarted);

      if (showRequestButton) {
        if (resource.kind === 'kube') {
          return (
            <KubeRequestButton
              resource={resource}
              addOrRemoveResource={resource =>
                accessRequestsService.addOrRemoveResource(resource)
              }
              fetchNamespaces={fetchKubeNamespaces}
              RequestButton={({ onClick }) => {
                return (
                  <AccessRequestButton
                    isResourceAdded={isResourceAdded}
                    requestStarted={requestStarted}
                    onClick={onClick}
                  />
                );
              }}
              addedResources={addedResources}
              clusterUri={props.clusterUri}
            />
          );
        }
        return (
          <AccessRequestButton
            isResourceAdded={isResourceAdded}
            requestStarted={requestStarted}
            onClick={() => accessRequestsService.addOrRemoveResource(resource)}
          />
        );
      }
    },
    [
      accessRequestsService,
      addedResources,
      requestStarted,
      integratedAccessRequests,
    ]
  );

  const bulkAddResources = useCallback(
    (resources: UnifiedResourceResponse[]) => {
      accessRequestsService.addOrRemoveResources(resources);
    },
    [accessRequestsService]
  );

  return (
    <Resources
      getAccessRequestButton={getAccessRequestButton}
      queryParams={mergedParams}
      onParamsChange={onParamsChange}
      clusterUri={props.clusterUri}
      userPreferencesAttempt={userPreferencesAttempt}
      updateUserPreferences={updateUserPreferences}
      userPreferences={userPreferences}
      canAddResources={canAddResources}
      canUseConnectMyComputer={canUseConnectMyComputer}
      openConnectMyComputerDocument={openConnectMyComputerDocument}
      onResourcesRefreshRequest={onResourcesRefreshRequest}
      bulkAddResources={bulkAddResources}
      getAddedItemsCount={getAddedItemsCount}
      discoverUrl={discoverUrl}
      integratedAccessRequests={integratedAccessRequests}
      // Reset the component state when query params object change.
      // JSON.stringify on the same object will always produce the same string.
      key={`${JSON.stringify(mergedParams)}-${JSON.stringify(integratedAccessRequests)}`}
    />
  );
}

const Resources = memo(
  (props: {
    clusterUri: uri.ClusterUri;
    queryParams: UnifiedResourcesQueryParams;
    onParamsChange(params: UnifiedResourcesQueryParams): void;
    userPreferencesAttempt?: Attempt<void>;
    userPreferences: UserPreferences;
    updateUserPreferences(u: UserPreferences): Promise<void>;
    canAddResources: boolean;
    canUseConnectMyComputer: boolean;
    openConnectMyComputerDocument(): void;
    onResourcesRefreshRequest: ResourcesContext['onResourcesRefreshRequest'];
    discoverUrl: string;
    getAccessRequestButton: (resource: UnifiedResourceResponse) => JSX.Element;
    getAddedItemsCount: () => number;
    bulkAddResources: (resources: UnifiedResourceResponse[]) => void;
    integratedAccessRequests: IntegratedAccessRequests;
  }) => {
    const appContext = useAppContext();

    const { fetch, resources, attempt } = useUnifiedResourcesFetch({
      fetchFunc: useCallback(
        async (paginationParams, signal) => {
          // Block the call if we don't know yet what resources to show.
          // We will remount the component and do the call when integratedAccessRequests changes.
          if (props.integratedAccessRequests.supported === 'unknown') {
            await waitForever(signal);
          }

          const { searchAsRoles, includeRequestable } =
            getRequestableResourcesParams(props.integratedAccessRequests);
          const response = await retryWithRelogin(
            appContext,
            props.clusterUri,
            () =>
              appContext.resourcesService.listUnifiedResources(
                {
                  clusterUri: props.clusterUri,
                  sortBy: {
                    isDesc: props.queryParams.sort.dir === 'DESC',
                    field: props.queryParams.sort.fieldName,
                  },
                  search: props.queryParams.search,
                  kinds: props.queryParams.kinds,
                  query: props.queryParams.query,
                  pinnedOnly: props.queryParams.pinnedOnly,
                  startKey: paginationParams.startKey,
                  limit: paginationParams.limit,
                  searchAsRoles,
                  includeRequestable,
                },
                signal
              )
          );

          return {
            startKey: response.nextKey,
            agents: response.resources,
          };
        },
        [
          appContext,
          props.queryParams.kinds,
          props.queryParams.pinnedOnly,
          props.queryParams.query,
          props.queryParams.search,
          props.queryParams.sort.dir,
          props.queryParams.sort.fieldName,
          props.clusterUri,
          props.integratedAccessRequests,
        ]
      ),
    });

    const { onResourcesRefreshRequest } = props;
    useEffect(() => {
      const { cleanup } = onResourcesRefreshRequest(() => {
        fetch({ clear: true });
      });
      return cleanup;
    }, [onResourcesRefreshRequest, fetch]);

    const { getAccessRequestButton } = props;
    // The action callback in the requestAccess action has access to
    // `SharedUnifiedResource['resource']`, but `props.bulkAddResources` accepts
    // `UnifiedResourceResponse`. Because of that, we need to to have the
    // getUnifiedResourceFromSharedResource function.
    const { sharedResources, getUnifiedResourceFromSharedResource } =
      useMemo(() => {
        const sharedResources: SharedUnifiedResource[] = [];
        const sharedResourceToUnifiedResource = new Map<
          SharedUnifiedResource['resource'],
          UnifiedResourceResponse
        >();

        resources.forEach(resource => {
          let sharedResource = mapToSharedResource(resource);
          const accessRequestButton = getAccessRequestButton(resource);
          if (accessRequestButton) {
            sharedResource.ui.ActionButton = accessRequestButton;
          }

          sharedResources.push(sharedResource);
          sharedResourceToUnifiedResource.set(
            sharedResource.resource,
            resource
          );
        });

        const getUnifiedResourceFromSharedResource =
          sharedResourceToUnifiedResource.get.bind(
            sharedResourceToUnifiedResource
          );

        return {
          sharedResources,
          getUnifiedResourceFromSharedResource,
        };
      }, [resources, getAccessRequestButton]);

    const resourceIds =
      props.userPreferences.clusterPreferences?.pinnedResources?.resourceIds;
    const { updateUserPreferences } = props;
    const pinning: UnifiedResourcesPinning = {
      kind: 'supported',
      getClusterPinnedResources: async () => resourceIds,
      updateClusterPinnedResources: pinnedIds =>
        updateUserPreferences({
          clusterPreferences: {
            pinnedResources: { resourceIds: pinnedIds },
          },
        }),
    };

    return (
      <SharedUnifiedResources
        params={props.queryParams}
        setParams={props.onParamsChange}
        unifiedResourcePreferencesAttempt={props.userPreferencesAttempt}
        bulkActions={
          props.integratedAccessRequests.supported === 'yes'
            ? [
                {
                  key: 'requestAccess',
                  Icon: icons.AddCircle,
                  text:
                    props.getAddedItemsCount() > 0
                      ? 'Add/Remove to Request'
                      : 'Request Access', // access request
                  disabled: false,
                  action: selectedResources =>
                    props.bulkAddResources(
                      selectedResources.map(sharedResource =>
                        getUnifiedResourceFromSharedResource(
                          sharedResource.resource
                        )
                      )
                    ),
                },
              ]
            : []
        }
        unifiedResourcePreferences={
          props.userPreferences.unifiedResourcePreferences
        }
        updateUnifiedResourcesPreferences={unifiedResourcePreferences =>
          props.updateUserPreferences({ unifiedResourcePreferences })
        }
        pinning={pinning}
        availabilityFilter={
          props.integratedAccessRequests.supported === 'yes'
            ? props.integratedAccessRequests.availabilityFilter
            : undefined
        }
        resources={sharedResources}
        resourcesFetchAttempt={attempt}
        fetchResources={fetch}
        availableKinds={[
          {
            kind: 'node',
            disabled: false,
          },
          {
            kind: 'app',
            disabled: false,
          },
          {
            kind: 'db',
            disabled: false,
          },
          {
            kind: 'kube_cluster',
            disabled: false,
          },
        ]}
        NoResources={
          <NoResources
            canCreate={props.canAddResources}
            discoverUrl={props.discoverUrl}
            canUseConnectMyComputer={props.canUseConnectMyComputer}
            onConnectMyComputerCtaClick={props.openConnectMyComputerDocument}
          />
        }
      />
    );
  }
);

const mapToSharedResource = (
  resource: UnifiedResourceResponse
): SharedUnifiedResource => {
  switch (resource.kind) {
    case 'server': {
      const { resource: server } = resource;
      return {
        resource: {
          kind: 'node' as const,
          labels: server.labels,
          id: server.name,
          hostname: server.hostname,
          addr: server.addr,
          tunnel: server.tunnel,
          subKind: server.subKind as NodeSubKind,
          requiresRequest: resource.requiresRequest,
        },
        ui: {
          ActionButton: <ConnectServerActionButton server={server} />,
        },
      };
    }
    case 'database': {
      const { resource: database } = resource;
      return {
        resource: {
          kind: 'db' as const,
          labels: database.labels,
          description: database.desc,
          name: database.name,
          type: formatDatabaseInfo(
            database.type as DbType,
            database.protocol as DbProtocol
          ).title,
          protocol: database.protocol as DbProtocol,
          requiresRequest: resource.requiresRequest,
        },
        ui: {
          ActionButton: <ConnectDatabaseActionButton database={database} />,
        },
      };
    }
    case 'kube': {
      const { resource: kube } = resource;

      return {
        resource: {
          kind: 'kube_cluster' as const,
          labels: kube.labels,
          name: kube.name,
          requiresRequest: resource.requiresRequest,
        },
        ui: {
          ActionButton: <ConnectKubeActionButton kube={kube} />,
        },
      };
    }
    case 'app': {
      const { resource: app } = resource;

      return {
        resource: {
          kind: 'app' as const,
          labels: app.labels,
          name: app.name,
          id: app.name,
          addrWithProtocol: getAppAddrWithProtocol(app),
          awsConsole: app.awsConsole,
          description: app.desc,
          friendlyName: app.friendlyName,
          samlApp: app.samlApp,
          requiresRequest: resource.requiresRequest,
        },
        ui: {
          ActionButton: <ConnectAppActionButton app={app} />,
        },
      };
    }
  }
};

function NoResources(props: {
  canCreate: boolean;
  discoverUrl: string | undefined;
  canUseConnectMyComputer: boolean;
  onConnectMyComputerCtaClick(): void;
}) {
  let $content: React.ReactElement;
  if (!props.canCreate) {
    $content = (
      <>
        <H1 mb="2">No Resources Found</H1>
        <Text>
          Either there are no resources in the cluster, or your roles don't
          grant you access.
        </Text>
      </>
    );
  } else {
    const $discoverLink = (
      <Link href={props.discoverUrl} target="_blank">
        the&nbsp;Teleport Web UI
      </Link>
    );
    $content = (
      <>
        <Image src={stack} ml="auto" mr="auto" mb={4} height="100px" />
        <H1 mb={2}>Add your first resource to Teleport</H1>
        <Text color="text.slightlyMuted">
          {props.canUseConnectMyComputer ? (
            <>
              You can add it in {$discoverLink} or by connecting your computer
              to the cluster.
            </>
          ) : (
            <>
              Connect SSH servers, Kubernetes clusters, Databases and more from{' '}
              {$discoverLink}.
            </>
          )}
        </Text>
        {props.canUseConnectMyComputer && (
          <ButtonPrimary
            type="button"
            mt={3}
            gap={2}
            onClick={props.onConnectMyComputerCtaClick}
          >
            <icons.Laptop size={'medium'} />
            Connect My Computer
          </ButtonPrimary>
        )}
      </>
    );
  }

  return (
    <Flex
      maxWidth={600}
      p={8}
      pt={5}
      width="100%"
      mx="auto"
      flexDirection="column"
      alignItems="center"
      justifyContent="center"
    >
      {$content}
    </Flex>
  );
}

/**
 * Describes availability of integrated access requests
 * (requesting resources from the unified resources view).
 *
 * If `supported` is `'no'` it basically means that the cluster doesn't support
 * access requests at all.
 */
type IntegratedAccessRequests =
  | {
      supported: 'unknown';
    }
  | {
      supported: 'no';
    }
  | {
      supported: 'yes';
      availabilityFilter: ResourceAvailabilityFilter;
    };

/**
 * When `includeRequestable` is true,
 * all resources (accessible and requestable) are returned.
 * When only `searchAsRoles` is true, only requestable resources are returned.
 * When both are false, only accessible resources are returned.
 */
function getRequestableResourcesParams(
  integratedAccessRequests: IntegratedAccessRequests
): Pick<ListUnifiedResourcesRequest, 'searchAsRoles' | 'includeRequestable'> {
  if (integratedAccessRequests.supported === 'yes') {
    switch (integratedAccessRequests.availabilityFilter.mode) {
      case 'all':
      case 'none':
        return {
          searchAsRoles: false,
          includeRequestable: true,
        };
      case 'requestable':
        return {
          searchAsRoles: true,
          includeRequestable: false,
        };
    }
  }

  return {
    searchAsRoles: false,
    includeRequestable: false,
  };
}
