package main

import (
	"github.com/jleight/steampipe-plugin-azuredevops/azuredevops"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: azuredevops.Plugin,
	})
}
