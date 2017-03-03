// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package command

import (
	"fmt"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
)

func ServerSshCompleteFlags(ctx Context, params *SshServerParam, flagName string, currentValue string) {
	var comp schema.SchemaCompletionFunc

	switch flagName {
	case "id":
		comp = define.Resources["Server"].Commands["ssh"].Params["id"].CompleteFunc
	case "key", "i":
		comp = define.Resources["Server"].Commands["ssh"].Params["key"].CompleteFunc
	case "password":
		comp = define.Resources["Server"].Commands["ssh"].Params["password"].CompleteFunc
	case "port", "p":
		comp = define.Resources["Server"].Commands["ssh"].Params["port"].CompleteFunc
	case "quiet", "q":
		comp = define.Resources["Server"].Commands["ssh"].Params["quiet"].CompleteFunc
	case "user", "l":
		comp = define.Resources["Server"].Commands["ssh"].Params["user"].CompleteFunc
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}