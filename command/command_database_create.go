package command

import (
	"fmt"
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command/internal"
)

func DatabaseCreate(ctx Context, params *CreateDatabaseParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDatabaseAPI()

	// validation
	sw, err := client.GetSwitchAPI().Read(params.SwitchId)
	if err != nil {
		return fmt.Errorf("Switch(%d) is not found", params.SwitchId)
	}

	// set params
	var p *sacloud.CreateDatabaseValue
	switch params.Database {
	case "postgresql":
		p = sacloud.NewCreatePostgreSQLDatabaseValue()
	case "mariadb":
		p = sacloud.NewCreateMariaDBDatabaseValue()
	}

	p.SwitchID = fmt.Sprintf("%d", sw.ID)
	p.Plan = sacloud.DatabasePlan(params.Plan)
	p.DefaultUser = params.Username
	p.UserPassword = params.Password

	if ctx.IsSet("port") {
		p.ServicePort = fmt.Sprintf("%d", params.Port)
	}
	p.IPAddress1 = params.Ipaddress1
	p.MaskLen = params.NwMaskLen
	p.DefaultRoute = params.DefaultRoute
	p.SourceNetwork = params.SourceNetworks

	// TODO WebUI

	p.BackupTime = params.BackupTime
	p.Name = params.Name
	p.Tags = params.Tags
	p.Description = params.Description
	p.Icon = sacloud.NewResource(params.IconId)

	// call Create(id)
	dbParam := api.New(p)
	res, err := api.Create(dbParam)
	if err != nil {
		return fmt.Errorf("DatabaseCreate is failed: %s", err)
	}

	// wait for boot
	err = internal.ExecWithProgress(
		fmt.Sprintf("Still creating[ID:%d]...", res.ID),
		fmt.Sprintf("Create database[ID:%d]", res.ID),
		GlobalOption.Progress,
		func(compChan chan bool, errChan chan error) {
			// call manipurate functions
			err := api.SleepWhileCopying(res.ID, client.DefaultTimeoutDuration, 20)
			if err != nil {
				errChan <- err
				return
			}
			err = api.SleepUntilDatabaseRunning(res.ID, client.DefaultTimeoutDuration, 30)
			if err != nil {
				errChan <- err
				return
			}
			compChan <- true
		},
	)
	if err != nil {
		return fmt.Errorf("LoadBalancerCreate is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)
}
