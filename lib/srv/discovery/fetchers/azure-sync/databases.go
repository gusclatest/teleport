package azure_sync

import (
	"context"
	"github.com/gravitational/trace"
)

func (a *Fetcher) fetchDatabases(ctx context.Context) error {

}

func (a *Fetcher) fetchSinglePostgresServer(ctx context.Context) error {
	cli, err := a.CloudClients.GetAzurePostgresClient(a.SubscriptionID)
	if err != nil {
		return trace.Wrap(err)
	}
}

func (a *Fetcher) fetchFlexPostgresServer(ctx context.Context) error {

}

func (a *Fetcher) fetchMySQLDatabases(ctx context.Context) error {

}

func (a *Fetcher) fetchManagedSQLDatabases(ctx context.Context) error {

}
