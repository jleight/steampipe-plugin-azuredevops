package azuredevops

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

const pluginName = "steampipe-plugin-azuredevops"

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: pluginName,

		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			Schema:      ConfigSchema,
			NewInstance: ConfigInstance,
		},

		DefaultTransform: transform.FromCamel(),

		TableMap: map[string]*plugin.Table{
			"azuredevops_build":            tableAzureDevOpsBuild(ctx),
			"azuredevops_build_definition": tableAzureDevOpsBuildDefinition(ctx),
			"azuredevops_git_repository":   tableAzureDevOpsGetRepository(ctx),
			"azuredevops_pipeline":         tableAzureDevOpsPipeline(ctx),
			"azuredevops_project":          tableAzureDevOpsProject(ctx),
		},
	}

	return p
}
