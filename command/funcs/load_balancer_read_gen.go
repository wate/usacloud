// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-funcs'; DO NOT EDIT

package funcs

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func LoadBalancerRead(ctx command.Context, params *params.ReadLoadBalancerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetLoadBalancerAPI()

	// set params

	// call Read(id)
	res, err := api.Read(params.Id)
	if err != nil {
		return fmt.Errorf("LoadBalancerRead is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}
