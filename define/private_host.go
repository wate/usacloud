package define

import (
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func PrivateHostResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find", "selector"},
			Params:             privateHostListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: privateHostListColumns(),
			UseCustomCommand:   true,
			Category:           "basics",
			Order:              10,
		},
		"create": {
			Type:             schema.CommandCreate,
			Params:           privateHostCreateParam(),
			IncludeFields:    privateHostDetailIncludes(),
			ExcludeFields:    privateHostDetailExcludes(),
			UseCustomCommand: true,
			Category:         "basics",
			Order:            20,
		},
		"read": {
			Type:          schema.CommandRead,
			Params:        privateHostReadParam(),
			IncludeFields: privateHostDetailIncludes(),
			ExcludeFields: privateHostDetailExcludes(),
			Category:      "basics",
			Order:         30,
		},
		"update": {
			Type:          schema.CommandUpdate,
			Params:        privateHostUpdateParam(),
			IncludeFields: privateHostDetailIncludes(),
			ExcludeFields: privateHostDetailExcludes(),
			Category:      "basics",
			Order:         40,
		},
		"delete": {
			Type:          schema.CommandDelete,
			Aliases:       []string{"rm"},
			Params:        privateHostDeleteParam(),
			IncludeFields: privateHostDetailIncludes(),
			ExcludeFields: privateHostDetailExcludes(),
			Category:      "basics",
			Order:         50,
		},
		"server-info": {
			Type:               schema.CommandManipulateSingle,
			Params:             privateHostServerInfoParam(),
			Aliases:            []string{"server-list"},
			TableType:          output.TableSimple,
			TableColumnDefines: serverListColumns(),
			UseCustomCommand:   true,
			NeedlessConfirm:    true,
			Category:           "servers",
			Order:              10,
		},
		"server-add": {
			Type:               schema.CommandManipulateSingle,
			Params:             privateHostServerUpdateParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: serverListColumns(),
			UseCustomCommand:   true,
			Category:           "servers",
			Order:              20,
		},
		"server-delete": {
			Type:               schema.CommandManipulateSingle,
			Params:             privateHostServerUpdateParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: serverListColumns(),
			UseCustomCommand:   true,
			Category:           "servers",
			Order:              30,
		},
	}

	return &schema.Resource{
		Commands:            commands,
		CommandCategories:   privateHostCommandCategories,
		ResourceCategory:    CategoryComputing,
		ListResultFieldName: "PrivateHosts",
	}
}

var privateHostCommandCategories = []schema.Category{
	{
		Key:         "basics",
		DisplayName: "Basics",
		Order:       10,
	},
	{
		Key:         "servers",
		DisplayName: "Server Management",
		Order:       20,
	},
}

func privateHostListParam() map[string]*schema.Schema {
	return mergeParameterMap(CommonListParam, paramTagsCond)
}

func privateHostListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
		{
			Name:    "HostName",
			Sources: []string{"Host.Name"},
		},
		{
			Name:    "Core",
			Sources: []string{"AssignedCore", "TotalCore"},
			Format:  "%s / %s",
		},
		{
			Name:    "Memory",
			Sources: []string{"AssignedMemory", "TotalMemory"},
			Format:  "%s / %s",
		},
	}
}

func privateHostDetailIncludes() []string {
	return []string{}
}

func privateHostDetailExcludes() []string {
	return []string{}
}

func privateHostCreateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name":        paramRequiredName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     paramIconResourceID,
	}
}

func privateHostReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func privateHostUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name":        paramName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     paramIconResourceID,
	}
}

func privateHostDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func privateHostServerInfoParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func privateHostServerUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"server-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set server ID",
			Required:     true,
			ValidateFunc: validateSakuraID(),
			CompleteFunc: completeServerID(),
			Category:     "server",
			Order:        1,
		},
	}
}
