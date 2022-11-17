package azuredevops

import (
	"context"

	"github.com/microsoft/azure-devops-go-api/azuredevops/v6/core"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableAzureDevOpsProject(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azuredevops_project",
		Description: "Represents an Azure DevOps project.",

		List: &plugin.ListConfig{
			Hydrate: listProjects,
		},

		Columns: []*plugin.Column{
			{
				Name:        "abbreviation",
				Description: "Project abbreviation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Abbreviation"),
			},
			{
				Name:        "default_team_image_url",
				Description: "Url to default team identity image.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DefaultTeamImageUrl"),
			},
			{
				Name:        "description",
				Description: "The project's description (if any).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description"),
			},
			{
				Name:        "id",
				Description: "Project identifier.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Id"),
			},
			{
				Name:        "last_update_time",
				Description: "Project last update time.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("LastUpdateTime.Time"),
			},
			{
				Name:        "name",
				Description: "Project name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
			{
				Name:        "revision",
				Description: "Project revision.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Revision"),
			},
			{
				Name:        "state",
				Description: "Project state.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("State"),
			},
			{
				Name:        "url",
				Description: "Url to the full version of the object.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Url"),
			},
			{
				Name:        "visibility",
				Description: "Project visibility.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Visibility"),
			},
		},
	}
}

func listProjects(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	connection, err := GetAzureDevOpsConnection(ctx, d)
	if err != nil {
		logger.Error("azuredevops_project.listProjects", "connection_error", err)
		return nil, err
	}

	client, err := core.NewClient(ctx, connection)
	if err != nil {
		logger.Error("azuredevops_project.listProjects", "client_error", err)
		return nil, err
	}

	top := 999
	input := core.GetProjectsArgs{
		Top: &top,
	}

	limit := d.QueryContext.Limit
	if limit != nil {
		if *limit > 0 && *limit < 999 {
			top = int(*limit)
		}
	}

	for {
		response, err := client.GetProjects(ctx, input)
		if err != nil {
			logger.Error("listProjects", "list_projects_error", err)
			return nil, err
		}

		for _, project := range (*response).Value {
			d.StreamListItem(ctx, project)

			if d.QueryStatus.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		input.ContinuationToken = &response.ContinuationToken
		if *input.ContinuationToken == "" {
			break
		}
	}

	return nil, nil
}
