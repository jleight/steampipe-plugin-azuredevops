package azuredevops

import (
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/schema"
)

type azureDevOpsConfig struct {
	OrganizationURL     *string `cty:"org_service_url"`
	PersonalAccessToken *string `cty:"personal_access_token"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"org_service_url": {
		Type: schema.TypeString,
	},
	"personal_access_token": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &azureDevOpsConfig{}
}

func GetConfig(connection *plugin.Connection) azureDevOpsConfig {
	if connection == nil || connection.Config == nil {
		return azureDevOpsConfig{}
	}

	config, _ := connection.Config.(azureDevOpsConfig)

	return config
}
