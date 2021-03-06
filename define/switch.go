package define

import (
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func SwitchResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find", "selector"},
			Params:             switchListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: switchListColumns(),
			Category:           "basics",
			Order:              10,
		},
		"create": {
			Type:          schema.CommandCreate,
			Params:        switchCreateParam(),
			IncludeFields: switchDetailIncludes(),
			ExcludeFields: switchDetailExcludes(),
			Category:      "basics",
			Order:         20,
		},
		"read": {
			Type:          schema.CommandRead,
			Params:        switchReadParam(),
			IncludeFields: switchDetailIncludes(),
			ExcludeFields: switchDetailExcludes(),
			Category:      "basics",
			Order:         30,
		},
		"update": {
			Type:          schema.CommandUpdate,
			Params:        switchUpdateParam(),
			IncludeFields: switchDetailIncludes(),
			ExcludeFields: switchDetailExcludes(),
			Category:      "basics",
			Order:         40,
		},
		"delete": {
			Type:          schema.CommandDelete,
			Aliases:       []string{"rm"},
			Params:        switchDeleteParam(),
			IncludeFields: switchDetailIncludes(),
			ExcludeFields: switchDetailExcludes(),
			Category:      "basics",
			Order:         50,
		},
		"bridge-connect": {
			Type:             schema.CommandManipulateMulti,
			UseCustomCommand: true,
			Params:           switchConnectBridgeParam(),
			IncludeFields:    switchDetailIncludes(),
			ExcludeFields:    switchDetailExcludes(),
			NoOutput:         true,
			Category:         "bridge",
			Order:            10,
		},
		"bridge-disconnect": {
			Type:             schema.CommandManipulateMulti,
			UseCustomCommand: true,
			Params:           switchDisconnectBridgeParam(),
			IncludeFields:    switchDetailIncludes(),
			ExcludeFields:    switchDetailExcludes(),
			NoOutput:         true,
			Category:         "bridge",
			Order:            20,
		},
	}

	return &schema.Resource{
		Commands:            commands,
		CommandCategories:   switchCommandCategories,
		ResourceCategory:    CategoryNetworking,
		ListResultFieldName: "Switches",
	}
}

var switchCommandCategories = []schema.Category{
	{
		Key:         "basics",
		DisplayName: "Basics",
		Order:       10,
	},
	{
		Key:         "bridge",
		DisplayName: "Bridge Connection Management",
		Order:       20,
	},
}

func switchListParam() map[string]*schema.Schema {
	return mergeParameterMap(CommonListParam, paramTagsCond)
}

func switchListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
		{
			Name:    "Server",
			Sources: []string{"ServerCount"},
		},
		{
			Name:    "Appliance",
			Sources: []string{"ApplianceCount"},
		},
		{
			Name:    "Bridge",
			Sources: []string{"Bridge.Name"},
		},
		{
			Name:    "Gateway",
			Sources: []string{"Subnets.0.DefaultRoute"},
		},
		{
			Name: "Network",
			Sources: []string{
				"Subnets.0.NetworkAddress",
				"Subnets.0.NetworkMaskLen",
			},
			Format: "%s/%s",
		},
		{
			Name:    "BandWidth",
			Sources: []string{"Subnets.0.Internet.BandWidthMbps"},
			Format:  "%sMbps",
		},
	}
}

func switchDetailIncludes() []string {
	return []string{}
}

func switchDetailExcludes() []string {
	return []string{
		"Bridge.SwitchInZone",
		"Bridge.Info.Switches",
	}
}

func switchCreateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name":        paramRequiredName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     paramIconResourceID,
	}
}

func switchReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func switchUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name":        paramName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     paramIconResourceID,
	}
}

func switchDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func switchConnectBridgeParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"bridge-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerPathThrough,
			Description:  "set Bridge ID",
			Required:     true,
			ValidateFunc: validateSakuraID(),
			CompleteFunc: completeBridgeID(),
			Category:     "bridge",
			Order:        10,
		},
	}
}

func switchDisconnectBridgeParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}
