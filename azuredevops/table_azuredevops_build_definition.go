package azuredevops

import (
	"context"
	"github.com/google/uuid"
	"net/http"
	"net/url"
	"strconv"

	"github.com/microsoft/azure-devops-go-api/azuredevops/v6"
	builds "github.com/microsoft/azure-devops-go-api/azuredevops/v6/build"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableAzureDevOpsBuildDefinition(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azuredevops_build_definition",
		Description: "Represents an Azure DevOps build definition.",

		List: &plugin.ListConfig{
			Hydrate: listBuildDefinitions,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "project_id", Require: plugin.Required},
			},
		},

		Columns: []*plugin.Column{
			{
				Name:        "created_date",
				Description: "The date this version of the definition was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("CreatedDate.Time"),
			},
			{
				Name:        "id",
				Description: "The ID of the referenced definition.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Id"),
			},
			{
				Name:        "name",
				Description: "The name of the referenced definition.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
			{
				Name:        "path",
				Description: "The folder path of the definition.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Path"),
			},
			{
				Name:        "project_id",
				Description: "ID of the project.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Project.Id"),
			},
			{
				Name:        "queue_status",
				Description: "A value that indicates whether builds can be queued against this definition.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("QueueStatus"),
			},
			{
				Name:        "revision",
				Description: "The definition revision number.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Revision"),
			},
			{
				Name:        "type",
				Description: "The type of the definition.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Type"),
			},
			{
				Name:        "uri",
				Description: "The definition's URI.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Uri"),
			},
			{
				Name:        "url",
				Description: "The REST URL of the definition.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Url"),
			},
			{
				Name:      "links",
				Type:      proto.ColumnType_JSON,
				Transform: transform.FromField("Links"),
			},
			{
				Name:        "authored_by",
				Description: "The author of the definition.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("AuthoredBy"),
			},
			{
				Name:        "draft_of",
				Description: "A reference to the definition that this definition is a draft of, if this is a draft definition.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("DraftOf"),
			},
			{
				Name:        "drafts",
				Description: "The list of drafts associated with this definition, if this is not a draft definition.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Drafts"),
			},
			{
				Name:        "latest_build",
				Description: "Data representation of a build.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("LatestBuild"),
			},
			{
				Name:        "latest_completed_build",
				Description: "Data representation of a build.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("LatestCompletedBuild"),
			},
			{
				Name:        "metrics",
				Description: "Data representation of a build.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Metrics"),
			},
			{
				Name:        "quality",
				Description: "The quality of the definition document (draft, etc.)",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Quality"),
			},
			{
				Name:        "queue",
				Description: "The default queue for builds run against this definition.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Queue"),
			},
			{
				Name:        "badge_enabled",
				Description: "Indicates whether badges are enabled for this definition.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("BadgeEnabled"),
			},
			{
				Name:        "build_number_format",
				Description: "The build number format.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("BuildNumberFormat"),
			},
			{
				Name:        "comment",
				Description: "A save-time comment for the definition.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Comment"),
			},
			{
				Name:        "demands",
				Description: "A list of demands that represents the agent capabilities required by this build.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Demands"),
			},
			{
				Name:        "description",
				Description: "The description.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description"),
			},
			{
				Name:        "drop_location",
				Description: "The drop location for the definition.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DropLocation"),
			},
			{
				Name:        "job_authorization_scope",
				Description: "The job authorization scope for builds queued against this definition.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("JobAuthorizationScope"),
			},
			{
				Name:        "job_cancel_timeout_in_minutes",
				Description: "The job cancel timeout (in minutes) for builds cancelled by user for this definition.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("JobCancelTimeoutInMinutes"),
			},
			{
				Name:        "job_timeout_in_minutes",
				Description: "The job execution timeout (in minutes) for builds queued against this definition.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("JobTimeoutInMinutes"),
			},
			{
				Name:        "options",
				Description: "Represents the application of an optional behavior to a build definition.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Options"),
			},
			{
				Name:        "process",
				Description: "The build process.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Process"),
			},
			{
				Name:        "process_parameters",
				Description: "The process parameters for this definition.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ProcessParameters"),
			},
			{
				Name:        "properties",
				Description: "The class represents a property bag as a collection of key-value pairs. Values of all primitive types (any type with a TypeCode != TypeCode.Object) except for DBNull are accepted. Values of type Byte[], Int32, Double, DateType and String preserve their type, other primitives are retuned as a String. Byte[] expected as base64 encoded string.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Properties"),
			},
			{
				Name:        "repository",
				Description: "The repository.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Repository"),
			},
			{
				Name:        "retention_rules",
				Description: "Represents a retention policy for a build definition.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("RetentionRules"),
			},
			{
				Name:      "tags",
				Type:      proto.ColumnType_JSON,
				Transform: transform.FromField("Tags"),
			},
			{
				Name:        "triggers",
				Description: "Represents a trigger for a buld definition.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Triggers"),
			},
			{
				Name:        "variable_groups",
				Description: "Represents a variable group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("VariableGroups"),
			},
			{
				Name:      "variables",
				Type:      proto.ColumnType_JSON,
				Transform: transform.FromField("Variables"),
			},
		},
	}
}

func listBuildDefinitions(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	projectID := d.KeyColumnQuals["project_id"].GetStringValue()

	connection, err := GetAzureDevOpsConnection(ctx, d)
	if err != nil {
		logger.Error("azuredevops_project.listBuildDefinitions", "connection_error", err)
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
	queryParams.Add("includeAllProperties", "true")

	locationId, _ := uuid.Parse("dbeaf647-6167-421a-bda9-c9327b25e2e6")

	for {
		response, err := client.Send(ctx, http.MethodGet, locationId, "6.0", routeValues, queryParams, nil, "", "application/json", nil)
		if err != nil {
			logger.Error("listBuildDefinitions", "list_build_definitions_error", err)
			return nil, err
		}

		var definitions []builds.BuildDefinition
		err = client.UnmarshalCollectionBody(response, &definitions)
		if err != nil {
			logger.Error("listBuildDefinitions", "unmarshal_error", err)
			return nil, err
		}

		for _, definition := range definitions {
			d.StreamListItem(ctx, definition)

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

/*
func listBuildDefinitions(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	projectID := d.KeyColumnQuals["project_id"].GetStringValue()

	connection, err := GetAzureDevOpsConnection(ctx, d)
	if err != nil {
		logger.Error("azuredevops_project.listBuildDefinitions", "connection_error", err)
		return nil, err
	}

	client, err := builds.NewClient(ctx, connection)
	if err != nil {
		logger.Error("azuredevops_project.listBuildDefinitions", "client_error", err)
		return nil, err
	}

	top := 999
	input := builds.GetDefinitionsArgs{
		Project: &projectID,
		Top:     &top,
	}

	limit := d.QueryContext.Limit
	if limit != nil {
		if *limit > 0 && *limit < 999 {
			top = int(*limit)
		}
	}

	for {
		response, err := client.GetDefinitions(ctx, input)
		if err != nil {
			logger.Error("listBuildDefinitions", "list_build_definitions_error", err)
			return nil, err
		}

		for _, ref := range (*response).Value {
			definition, err := getAzureDevOpsBuildDefinition(ctx, client, ref)
			if err != nil {
				logger.Error("listBuildDefinitions", "get_build_definition_error", err)
				return nil, err
			}

			d.StreamListItem(ctx, definition)

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

func getAzureDevOpsBuildDefinition(ctx context.Context, client builds.Client, ref builds.BuildDefinitionReference) (*builds.BuildDefinition, error) {
	projectID := ref.Project.Id.String()

	input := builds.GetDefinitionArgs{
		Project:      &projectID,
		DefinitionId: ref.Id,
	}

	return client.GetDefinition(ctx, input)
}
*/
