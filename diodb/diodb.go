package diodb

import (
	"context"

	jsonplugin "github.com/prodigysml/steampipe-plugin-discloseio-diodb/diodb/tables"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "json_plugin",
	}

	// Attach table definitions
	p.TableMap = map[string]*plugin.Table{
		"diodb": jsonplugin.TableJSON(),
	}

	return p
}
