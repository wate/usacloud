// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-funcs'; DO NOT EDIT

package command

import (
	"fmt"
)

func SwitchRead(ctx Context, params *ReadSwitchParam) error {

	client := ctx.GetAPIClient()
	api := client.GetSwitchAPI()

	// set params

	// call Read(id)
	res, err := api.Read(params.Id)
	if err != nil {
		return fmt.Errorf("SwitchRead is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}