package azuredevops

import (
	"context"
	"os"

	ado "github.com/microsoft/azure-devops-go-api/azuredevops/v6"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func GetAzureDevOpsConnection(_ context.Context, d *plugin.QueryData) (*ado.Connection, error) {
	var organizationURL, personalAccessToken string

	azureDevOpsConfig := GetConfig(d.Connection)

	if azureDevOpsConfig.OrganizationURL != nil {
		organizationURL = *azureDevOpsConfig.OrganizationURL
	} else {
		organizationURL = os.Getenv("AZDO_ORG_SERVICE_URL")
	}

	if azureDevOpsConfig.PersonalAccessToken != nil {
		personalAccessToken = *azureDevOpsConfig.PersonalAccessToken
	} else {
		personalAccessToken = os.Getenv("AZDO_PERSONAL_ACCESS_TOKEN")
	}

	connection := ado.NewPatConnection(organizationURL, personalAccessToken)
	return connection, nil
}
