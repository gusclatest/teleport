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
	"fmt"
	"github.com/gravitational/teleport/lib/defaults"
	"strings"

	"github.com/alecthomas/kingpin/v2"
	"github.com/gravitational/teleport"
	"github.com/gravitational/teleport/lib/asciitable"
	"github.com/gravitational/trace"

	apiclient "github.com/gravitational/teleport/api/client"
	"github.com/gravitational/teleport/api/client/proto"
	"github.com/gravitational/teleport/api/types"
	"github.com/gravitational/teleport/lib/client"
)

type gitListCommand struct {
	*kingpin.CmdClause

	format string
}

func newGitListCommand(parent *kingpin.CmdClause) *gitListCommand {
	cmd := &gitListCommand{
		CmdClause: parent.Command("ls", "List Git servers"),
	}
	cmd.Flag("format", defaults.FormatFlagDescription(defaults.DefaultFormats...)).
		Short('f').
		Default(teleport.Text).
		EnumVar(&cmd.format, defaults.DefaultFormats...)
	return cmd
}

func (c *gitListCommand) run(cf *CLIConf) error {
	tc, err := makeClient(cf)
	if err != nil {
		return trace.Wrap(err)
	}

	var resources []*types.EnrichedResource
	err = client.RetryWithRelogin(cf.Context, tc, func() error {
		client, err := tc.ConnectToCluster(cf.Context)
		if err != nil {
			return trace.Wrap(err)
		}
		defer client.Close()

		resources, err = apiclient.GetAllUnifiedResources(cf.Context, client.AuthClient, &proto.ListUnifiedResourcesRequest{
			Kinds:               []string{types.KindGitServer},
			SortBy:              types.SortBy{Field: types.ResourceMetadataName},
			SearchKeywords:      tc.SearchKeywords,
			PredicateExpression: tc.PredicateExpression,
			IncludeLogins:       true,
		})
		return trace.Wrap(err)
	})
	if err != nil {
		return trace.Wrap(err)
	}

	servers := make(types.Servers, 0, len(resources))
	for _, resource := range resources {
		server, ok := resource.ResourceWithLabels.(types.Server)
		if !ok {
			return trace.BadParameter("expecting types.Server but got %T", server)
		}
		servers = append(servers, server)
	}

	return c.printGitServers(cf, servers)
}

func (c *gitListCommand) printGitServers(cf *CLIConf, servers []types.Server) error {
	format := strings.ToLower(c.format)
	switch format {
	case teleport.Text, "":
		return printGitServersAsText(cf, servers)
	case teleport.JSON, teleport.YAML:
		out, err := serializeNodes(servers, format)
		if err != nil {
			return trace.Wrap(err)
		}
		if _, err := fmt.Fprintln(cf.Stdout(), out); err != nil {
			return trace.Wrap(err)
		}
		return nil
	default:
		return trace.BadParameter("unsupported format %q", format)
	}
}

func printGitServersAsText(cf *CLIConf, servers []types.Server) error {
	var rows [][]string
	var showLoginNote bool
	for _, server := range servers {
		// TODO(greedy52) fill in GitHub login when available from Profile.
		login := "(n/a)*"
		showLoginNote = true

		if github := server.GetGitHub(); github != nil {
			rows = append(rows, []string{
				"GitHub",
				github.Organization,
				login,
				fmt.Sprintf("https://github.com/%s", github.Organization),
			})
		} else {
			return trace.BadParameter("expecting Git server but got %v", server.GetKind())
		}
	}

	t := asciitable.MakeTable([]string{"Type", "Organization", "Username", "URL"}, rows...)
	if _, err := fmt.Fprintln(cf.Stdout(), t.AsBuffer().String()); err != nil {
		return trace.Wrap(err)
	}

	if showLoginNote {
		fmt.Fprintln(cf.Stdout(), gitLoginNote)
	}

	fmt.Fprintln(cf.Stdout(), gitCommandsGeneralHint)
	return nil
}

const gitLoginNote = "" +
	"(n/a)*: Usernames will be retrieved automatically upon git commands.\n" +
	"        Alternatively, run `tsh git login --github-org <org>`.\n"

const gitCommandsGeneralHint = "" +
	"hint: use 'tsh git clone <git-clone-ssh-url>' to clone a new repository\n" +
	"      use 'tsh git config update' to configure an existing repository to use Teleport\n" +
	"      once the repository is cloned or configured, use 'git' as normal\n"
