// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package completion

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func NFSListCompleteArgs(ctx command.Context, params *params.ListNFSParam, cur, prev, commandName string) {

}

func NFSCreateCompleteArgs(ctx command.Context, params *params.CreateNFSParam, cur, prev, commandName string) {

}

func NFSReadCompleteArgs(ctx command.Context, params *params.ReadNFSParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetNFSAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.NFS {
		fmt.Println(res.NFS[i].ID)
		var target interface{} = &res.NFS[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func NFSUpdateCompleteArgs(ctx command.Context, params *params.UpdateNFSParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetNFSAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.NFS {
		fmt.Println(res.NFS[i].ID)
		var target interface{} = &res.NFS[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func NFSDeleteCompleteArgs(ctx command.Context, params *params.DeleteNFSParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetNFSAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.NFS {
		fmt.Println(res.NFS[i].ID)
		var target interface{} = &res.NFS[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func NFSBootCompleteArgs(ctx command.Context, params *params.BootNFSParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetNFSAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.NFS {
		fmt.Println(res.NFS[i].ID)
		var target interface{} = &res.NFS[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func NFSShutdownCompleteArgs(ctx command.Context, params *params.ShutdownNFSParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetNFSAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.NFS {
		fmt.Println(res.NFS[i].ID)
		var target interface{} = &res.NFS[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func NFSShutdownForceCompleteArgs(ctx command.Context, params *params.ShutdownForceNFSParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetNFSAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.NFS {
		fmt.Println(res.NFS[i].ID)
		var target interface{} = &res.NFS[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func NFSResetCompleteArgs(ctx command.Context, params *params.ResetNFSParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetNFSAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.NFS {
		fmt.Println(res.NFS[i].ID)
		var target interface{} = &res.NFS[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func NFSWaitForBootCompleteArgs(ctx command.Context, params *params.WaitForBootNFSParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetNFSAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.NFS {
		fmt.Println(res.NFS[i].ID)
		var target interface{} = &res.NFS[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func NFSWaitForDownCompleteArgs(ctx command.Context, params *params.WaitForDownNFSParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetNFSAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.NFS {
		fmt.Println(res.NFS[i].ID)
		var target interface{} = &res.NFS[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func NFSMonitorNicCompleteArgs(ctx command.Context, params *params.MonitorNicNFSParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetNFSAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.NFS {
		fmt.Println(res.NFS[i].ID)
		var target interface{} = &res.NFS[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func NFSMonitorFreeDiskSizeCompleteArgs(ctx command.Context, params *params.MonitorFreeDiskSizeNFSParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetNFSAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.NFS {
		fmt.Println(res.NFS[i].ID)
		var target interface{} = &res.NFS[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}
