package azure

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/gravitational/trace"
)

type ResourceGroupsClient struct {
	cli *armresources.ResourceGroupsClient
}

func NewResourceGroupsClient(
	subscriptionID string,
	cred azcore.TokenCredential,
	opts *arm.ClientOptions,
) (*ResourceGroupsClient, error) {
	cli, err := armresources.NewResourceGroupsClient(subscriptionID, cred, opts)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	return &ResourceGroupsClient{cli: cli}, nil
}

func (c *ResourceGroupsClient) ListResourceGroups(
	ctx context.Context,
	opts *armresources.ResourceGroupsClientListOptions,
) ([]*armresources.ResourceGroup, error) {
	pager := c.cli.NewListPager(opts)
	rgs := make([]*armresources.ResourceGroup, 0)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, trace.Wrap(err)
		}
		for _, rg := range page.Value {
			rgs = append(rgs, rg)
		}
	}
	return rgs, nil
}
