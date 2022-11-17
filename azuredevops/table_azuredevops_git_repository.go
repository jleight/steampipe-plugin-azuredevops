package azuredevops

import (
	"context"

	"github.com/microsoft/azure-devops-go-api/azuredevops/v6/git"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableAzureDevOpsGetRepository(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azuredevops_git_repository",
		Description: "Represents an Azure DevOps git repository.",

		List: &plugin.ListConfig{
			Hydrate: listGitRepositories,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "project_id", Require: plugin.Required},
			},
		},

		Columns: []*plugin.Column{
			{
				Name:        "links",
				Description: "The class to represent a collection of REST reference links.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Links"),
			},
			{
				Name:      "default_branch",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("DefaultBranch"),
			},
			{
				Name:      "id",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Id"),
			},
			{
				Name:        "is_fork",
				Description: "True if the repository was created as a fork.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("IsFork"),
			},
			{
				Name:      "name",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Name"),
			},
			{
				Name:      "parent_git_repository_id",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("ParentRepository.Id"),
			},
			{
				Name:      "project_id",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Project.Id"),
			},
			{
				Name:      "remote_url",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("RemoteUrl"),
			},
			{
				Name:        "size",
				Description: "Compressed size (bytes) of the repository.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Size"),
			},
			{
				Name:      "ssh_url",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("SshUrl"),
			},
			{
				Name:      "url",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Url"),
			},
			{
				Name:      "valid_remote_urls",
				Type:      proto.ColumnType_JSON,
				Transform: transform.FromField("ValidRemoteUrls"),
			},
			{
				Name:      "web_url",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("WebUrl"),
			},
		},
	}
}

func listGitRepositories(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	projectID := d.KeyColumnQuals["project_id"].GetStringValue()

	connection, err := GetAzureDevOpsConnection(ctx, d)
	if err != nil {
		logger.Error("azuredevops_project.listGitRepositories", "connection_error", err)
		return nil, err
	}

	client, err := git.NewClient(ctx, connection)
	if err != nil {
		logger.Error("azuredevops_project.listGitRepositories", "client_error", err)
		return nil, err
	}

	input := git.GetRepositoriesArgs{
		Project: &projectID,
	}

	response, err := client.GetRepositories(ctx, input)
	if err != nil {
		logger.Error("listGitRepositories", "list_git_repositories_error", err)
		return nil, err
	}

	for _, repository := range *response {
		d.StreamListItem(ctx, repository)

		if d.QueryStatus.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
