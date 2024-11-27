package azure_sync

import (
	"context"
	accessgraphv1alpha "github.com/gravitational/teleport/gen/proto/go/accessgraph/v1alpha"
	"github.com/gravitational/trace"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type storageLocation struct {
	resourceGroupName string
	accountName       string
}

func (a *Fetcher) fetchBlobContainers(ctx context.Context) ([]*accessgraphv1alpha.AzureBlobContainer, error) {
	// Fetch the resource groups and associated storage accounts
	resGrps, err := a.resourceGroupsClient.ListResourceGroups(ctx, nil)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	storageLocations := make([]storageLocation, 0, len(resGrps))
	for _, resGrp := range resGrps {
		accts, err := a.storageAccountsClient.ListStorageAccounts(ctx, *resGrp.Name, nil)
		if err != nil {
			return nil, trace.Wrap(err)
		}
		for _, acct := range accts {
			storageLocations = append(storageLocations, storageLocation{*resGrp.Name, *acct.Name})
		}
	}

	// List all the blob containers with the resource group and storage account pairs
	pbContainers := make([]*accessgraphv1alpha.AzureBlobContainer, 0, len(storageLocations))
	for _, location := range storageLocations {
		containers, err := a.blobContainersClient.ListBlobContainers(ctx, location.resourceGroupName, location.accountName, nil)
		if err != nil {
			return nil, trace.Wrap(err)
		}
		for _, container := range containers {
			pbContainer := &accessgraphv1alpha.AzureBlobContainer{
				Id:                *container.ID,
				SubscriptionId:    a.SubscriptionID,
				LastSyncTime:      timestamppb.Now(),
				Name:              *container.Name,
				AccountName:       location.accountName,
				ResourceGroupName: location.resourceGroupName,
			}
			pbContainers = append(pbContainers, pbContainer)
		}
	}
	return pbContainers, nil
}
