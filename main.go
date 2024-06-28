package main

import (
	"github.com/prodigysml/steampipe-plugin-discloseio-diodb/diodb"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: diodb.Plugin,
	})
}
