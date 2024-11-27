package azure

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/gravitational/trace"
)

type BlobContainersClient struct {
	cli     *armstorage.BlobContainersClient
	blobCli *azblob.Client
}

const defaultServiceURL = "https://%s.blob.core.windows.net/"

func NewBlobContainersClient(subscriptionID string, cred azcore.TokenCredential, opts *arm.ClientOptions) (*BlobContainersClient, error) {
	cli, err := armstorage.NewBlobContainersClient(subscriptionID, cred, opts)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	return &BlobContainersClient{cli: cli}, nil
}

func (c *BlobContainersClient) ListBlobContainers(
	ctx context.Context,
	resourceGroupName string,
	accountName string,
	opts *armstorage.BlobContainersClientListOptions,
) ([]*armstorage.ListContainerItem, error) {
	pager := c.cli.NewListPager(resourceGroupName, accountName, opts)
	containers := make([]*armstorage.ListContainerItem, 0, 10)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, trace.Wrap(err)
		}
		for _, container := range page.Value {
			containers = append(containers, container)
		}
	}
	return containers, nil
}
