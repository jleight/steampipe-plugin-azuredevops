install:
	go build -o ~/.steampipe/plugins/hub.steampipe.io/plugins/jleight/azuredevops@latest/steampipe-plugin-azuredevops.plugin *.go

restart:
	steampipe service restart
