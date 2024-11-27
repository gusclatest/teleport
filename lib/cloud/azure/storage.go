package azure

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	"github.com/gravitational/trace"
)

type StorageAccountsClient struct {
	cli *armstorage.AccountsClient
}

func NewStorageAccountsClient(subscriptionID string, cred azcore.TokenCredential, options *arm.ClientOptions) (*StorageAccountsClient, error) {
	cli, err := armstorage.NewAccountsClient(subscriptionID, cred, options)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	return &StorageAccountsClient{cli: cli}, nil
}

func (c *StorageAccountsClient) ListStorageAccounts(
	ctx context.Context,
	resourceGroup string,
	opts *armstorage.AccountsClientListByResourceGroupOptions,
) ([]*armstorage.Account, error) {
	pager := c.cli.NewListByResourceGroupPager(resourceGroup, opts)
	accounts := make([]*armstorage.Account, 0, 10)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, trace.Wrap(err)
		}
		for _, account := range page.Value {
			accounts = append(accounts, account)
		}
	}
	return accounts, nil
}
