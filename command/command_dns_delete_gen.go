// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-funcs'; DO NOT EDIT

package command

import (
	"fmt"
)

func DNSDelete(ctx Context, params *DeleteDNSParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDNSAPI()

	// set params

	// call Delete(id)
	res, err := api.Delete(params.Id)
	if err != nil {
		return fmt.Errorf("DNSDelete is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}