package azuredevops

import (
	"context"

	builds "github.com/microsoft/azure-devops-go-api/azuredevops/v6/build"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableAzureDevOpsBuild(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azuredevops_build",
		Description: "Represents an Azure DevOps build.",

		List: &plugin.ListConfig{
			Hydrate: listBuilds,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "project_id", Require: plugin.Required},
			},
		},

		Columns: []*plugin.Column{
			{
				Name:      "links",
				Type:      proto.ColumnType_JSON,
				Transform: transform.FromField("Links"),
			},
			{
				Name:        "agent_specification_id",
				Description: "Agent specification unique identifier.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("AgentSpecification.Identifier"),
			},
			{
				Name:        "build_number",
				Description: "The build number/name of the build.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("BuildNumber"),
			},
			{
				Name:        "build_number_revision",
				Description: "The build number revision.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("BuildNumberRevision"),
			},
			{
				Name:        "controller",
				Description: "The build controller. This is only set if the definition type is Xaml.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Controller"),
			},
			{
				Name:        "definition",
				Description: "The definition associated with the build.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Definition"),
			},
			{
				Name:        "deleted",
				Description: "Indicates whether the build has been deleted.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Deleted"),
			},
			{
				Name:        "deleted_by",
				Description: "The identity of the process or person that deleted the build.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("DeletedBy"),
			},
			{
				Name:        "deleted_date",
				Description: "The date the build was deleted.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("DeletedDate.Time"),
			},
			{
				Name:        "deleted_reason",
				Description: "The description of how the build was deleted.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DeletedReason"),
			},
			{
				Name:        "demands",
				Description: "A list of demands that represents the agent capabilities required by this build.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Demands"),
			},
			{
				Name:        "finish_time",
				Description: "The time that the build was completed.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("FinishTime.Time"),
			},
			{
				Name:        "id",
				Description: "The ID of the build.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Id"),
			},
			{
				Name:        "keep_forever",
				Description: "Indicates whether the build should be skipped by retention policies.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("KeepForever"),
			},
			{
				Name:        "last_changed_by",
				Description: "The identity representing the process or person that last changed the build.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("LastChangedBy"),
			},
			{
				Name:        "last_changed_date",
				Description: "The date the build was last changed.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("LastChangedDate.Time"),
			},
			{
				Name:        "logs",
				Description: "Information about the build logs.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Logs"),
			},
			{
				Name:        "orchestration_plan",
				Description: "The orchestration plan for the build.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("OrchestrationPlan"),
			},
			{
				Name:        "parameters",
				Description: "The parameters for the build.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Parameters"),
			},
			{
				Name:        "plans",
				Description: "Orchestration plans associated with the build (build, cleanup)",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Plans"),
			},
			{
				Name:        "priority",
				Description: "The build's priority.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Priority"),
			},
			{
				Name:        "project_id",
				Description: "The team project.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Project.Id"),
			},
			{
				Name:        "properties",
				Description: "The class represents a property bag as a collection of key-value pairs. Values of all primitive types (any type with a TypeCode != TypeCode.Object) except for DBNull are accepted. Values of type Byte[], Int32, Double, DateType and String preserve their type, other primitives are retuned as a String. Byte[] expected as base64 encoded string.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Properties"),
			},
			{
				Name:        "quality",
				Description: "The quality of the xaml build (good, bad, etc.)",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Quality"),
			},
			{
				Name:        "queue",
				Description: "The queue. This is only set if the definition type is Build.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Queue"),
			},
			{
				Name:        "queue_options",
				Description: "Additional options for queueing the build.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("QueueOptions"),
			},
			{
				Name:        "queue_position",
				Description: "The current position of the build in the queue.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("QueuePosition"),
			},
			{
				Name:        "queue_time",
				Description: "The time that the build was queued.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("QueueTime.Time"),
			},
			{
				Name:        "reason",
				Description: "The reason that the build was created.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Reason"),
			},
			{
				Name:        "repository",
				Description: "The repository.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Repository"),
			},
			{
				Name:        "requested_by",
				Description: "The identity that queued the build.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("RequestedBy"),
			},
			{
				Name:        "requested_for",
				Description: "The identity on whose behalf the build was queued.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("RequestedFor"),
			},
			{
				Name:        "result",
				Description: "The build result.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Result"),
			},
			{
				Name:        "retained_by_release",
				Description: "Indicates whether the build is retained by a release.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("RetainedByRelease"),
			},
			{
				Name:        "source_branch",
				Description: "The source branch.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("SourceBranch"),
			},
			{
				Name:        "source_version",
				Description: "The source version.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("SourceVersion"),
			},
			{
				Name:        "start_time",
				Description: "The time that the build was started.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("StartTime.Time"),
			},
			{
				Name:        "status",
				Description: "The status of the build.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Status"),
			},
			{
				Name:      "tags",
				Type:      proto.ColumnType_JSON,
				Transform: transform.FromField("Status"),
			},
			{
				Name:        "triggered_by_build",
				Description: "The build that triggered this build via a Build completion trigger.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("TriggeredByBuild"),
			},
			{
				Name:        "trigger_info",
				Description: "Sourceprovider-specific information about what triggered the build",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("TriggerInfo"),
			},
			{
				Name:        "uri",
				Description: "The URI of the build.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Uri"),
			},
			{
				Name:        "url",
				Description: "The REST URL of the build.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Url"),
			},
			{
				Name:        "validation_results",
				Description: "Represents the result of validating a build request.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ValidationResults"),
			},
		},
	}
}

func listBuilds(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	projectID := d.KeyColumnQuals["project_id"].GetStringValue()

	connection, err := GetAzureDevOpsConnection(ctx, d)
	if err != nil {
		logger.Error("azuredevops_project.listBuilds", "connection_error", err)
		return nil, err
	}

	client, err := builds.NewClient(ctx, connection)
	if err != nil {
		logger.Error("azuredevops_project.listBuilds", "client_error", err)
		return nil, err
	}

	top := 999
	input := builds.GetBuildsArgs{
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
		response, err := client.GetBuilds(ctx, input)
		if err != nil {
			logger.Error("listBuilds", "list_builds_error", err)
			return nil, err
		}

		for _, build := range (*response).Value {
			d.StreamListItem(ctx, build)

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
