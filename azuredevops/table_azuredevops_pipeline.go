package azuredevops

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6/pipelines"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableAzureDevOpsPipeline(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azuredevops_pipeline",
		Description: "Represents an Azure DevOps pipeline.",

		List: &plugin.ListConfig{
			Hydrate: listPipelines,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "project_id", Require: plugin.Required},
			},
		},

		Columns: []*plugin.Column{
			{
				Name:        "project_id",
				Description: "ID of the project this pipeline belongs to.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getProjectId,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "folder",
				Description: "Pipeline folder",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Folder"),
			},
			{
				Name:        "id",
				Description: "Pipeline ID",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Id"),
			},
			{
				Name:        "name",
				Description: "Pipeline name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
			{
				Name:        "revision",
				Description: "Revision number",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Revision"),
			},
			{
				Name:        "links",
				Description: "The class to represent a collection of REST reference links.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Links"),
			},
			{
				Name:        "configuration_type",
				Description: "Type of configuration>",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Configuration.Type"),
			},
			{
				Name:        "url",
				Description: "URL of the pipeline",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Url"),
			},
		},
	}
}

func listPipelines(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	projectID := d.KeyColumnQuals["project_id"].GetStringValue()

	connection, err := GetAzureDevOpsConnection(ctx, d)
	if err != nil {
		logger.Error("azuredevops_project.listPipelines", "connection_error", err)
		return nil, err
	}

	client := connection.GetClientByUrl(connection.BaseUrl)

	top := 999
	limit := d.QueryContext.Limit
	if limit != nil {
		if *limit > 0 && *limit < 999 {
			top = int(*limit)
		}
	}

	routeValues := make(map[string]string)
	routeValues["project"] = projectID

	queryParams := url.Values{}
	queryParams.Set("$top", strconv.Itoa(top))

	locationId, _ := uuid.Parse("28e1305e-2afe-47bf-abaf-cbb0e6a91988")

	for {
		response, err := client.Send(ctx, http.MethodGet, locationId, "6.0-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
		if err != nil {
			logger.Error("listPipelines", "list_pipelines_error", err)
			return nil, err
		}

		var pipes []pipelines.Pipeline
		err = client.UnmarshalCollectionBody(response, &pipes)
		if err != nil {
			logger.Error("listPipelines", "unmarshal_error", err)
			return nil, err
		}

		for _, pipeline := range pipes {
			d.StreamListItem(ctx, pipeline)

			if d.QueryStatus.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		continuationToken := response.Header.Get(azuredevops.HeaderKeyContinuationToken)
		if continuationToken == "" {
			break
		}

		queryParams.Set("continuationToken", continuationToken)
	}

	return nil, nil
}

func getProjectId(_ context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	return d.KeyColumnQuals["project_id"].GetStringValue(), nil
}
