// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package command

import (
	"fmt"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
)

func ServerInterfaceAddForInternetCompleteFlags(ctx Context, params *InterfaceAddForInternetServerParam, flagName string, currentValue string) {
	var comp schema.SchemaCompletionFunc

	switch flagName {
	case "id":
		comp = define.Resources["Server"].Commands["interface-add-for-internet"].Params["id"].CompleteFunc
	case "without-disk-edit":
		comp = define.Resources["Server"].Commands["interface-add-for-internet"].Params["without-disk-edit"].CompleteFunc
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}