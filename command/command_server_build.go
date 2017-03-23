package command

import (
	"fmt"
	"github.com/sacloud/libsacloud/builder"
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/libsacloud/sacloud/ostype"
	"github.com/sacloud/usacloud/command/internal"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func ServerBuild(ctx Context, params *BuildServerParam) error {

	client := ctx.GetAPIClient()

	// validate --- for disk mode params
	errs := validateServerDiskModeParams(ctx, params)
	if len(errs) > 0 {
		return fmt.Errorf("%s", flattenErrors(errs))
	}

	// select builder
	var sb interface{}

	switch params.DiskMode {
	case "create":
		if params.SourceDiskId > 0 {
			sb = builder.ServerFromDisk(client, params.Name, params.SourceDiskId)
		} else if params.SourceArchiveId > 0 {
			sb = builder.ServerFromArchive(client, params.Name, params.SourceArchiveId)
		} else {
			// Windows?
			if isWindows(params.OsType) {
				sb = builder.ServerPublicArchiveWindows(client, strToOSType(params.OsType), params.Name)
			} else {
				sb = builder.ServerPublicArchiveUnix(client, strToOSType(params.OsType), params.Name, params.Password)
			}
		}
	case "connect":
		sb = builder.ServerFromExistsDisk(client, params.Name, params.DiskId)
	case "diskless":
		sb = builder.ServerDiskless(client, params.Name)
	}

	// validate --- for network params
	errs = validateServerNetworkParams(sb, ctx, params)
	if len(errs) > 0 {
		return fmt.Errorf("%s", flattenErrors(errs))
	}

	// set network params
	if sb, ok := sb.(serverNetworkParams); ok {
		switch params.NetworkMode {
		case "shared":
			sb.AddPublicNWConnectedNIC()
		case "switch":
			switch sb := sb.(type) {
			case serverConnectSwitchParam:
				sb.AddExistsSwitchConnectedNIC(fmt.Sprintf("%d", params.SwitchId))
			case serverConnectSwitchParamWithEditableDisk:
				sb.AddExistsSwitchConnectedNIC(
					fmt.Sprintf("%d", params.SwitchId),
					params.Ipaddress,
					params.NwMasklen,
					params.DefaultRoute,
				)
			default:
				panic(fmt.Errorf("This server builder Can't connect to switch : %#v", sb))
			}

		case "disconnect":
			sb.AddDisconnectedNIC()
		case "none":
			// noop
		default:
			panic(fmt.Errorf("Unknown NetworkMode : %s", params.NetworkMode))
		}

		sb.SetUseVirtIONetPCI(params.UseNicVirtio)
		if params.PacketFilterId != sacloud.EmptyID {
			sb.SetPacketFilterIDs([]int64{params.PacketFilterId})
		}
	}

	// validate --- for disk params
	errs = validateServerDiskEditParams(sb, ctx, params)
	if len(errs) > 0 {
		return fmt.Errorf("%s", flattenErrors(errs))
	}

	// set disk edit params
	if sb, ok := sb.(serverEditDiskParam); ok {
		sb.SetHostName(params.Hostname)
		sb.SetPassword(params.Password)
		sb.SetDisablePWAuth(params.DisablePasswordAuth)

		for _, v := range params.StartupScriptIds {
			sb.AddNoteID(v)
		}
		for _, v := range params.StartupScripts {
			sb.AddNote(v)
		}
		sb.SetNotesEphemeral(params.StartupScriptsEphemeral)

		// SSH Key generate params
		switch params.SshKeyMode {
		case "id":
			for _, v := range params.SshKeyIds {
				sb.AddSSHKeyID(v)
			}
		case "generate":
			keyName := params.SshKeyName
			if params.SshKeyEphemeral && keyName == "" {
				keyName = fmt.Sprintf("generated-%d", time.Now().UnixNano())
			}
			sb.SetGenerateSSHKeyName(keyName)
			sb.SetGenerateSSHKeyPassPhrase(params.SshKeyPassPhrase)
			sb.SetGenerateSSHKeyDescription(params.SshKeyDescription)
		case "upload":
			// pubkey(text)
			for _, v := range params.SshKeyPublicKeys {
				sb.AddSSHKey(v)
			}
			// pubkey(from file)
			for _, v := range params.SshKeyPublicKeyFiles {
				b, err := ioutil.ReadFile(v)
				if err != nil {
					return fmt.Errorf("ServerCreate is failed: %s", err)
				}
				sb.AddSSHKey(string(b))

			}
			sb.SetSSHKeysEphemeral(params.SshKeyEphemeral)
		}

	}

	// set disk params
	if sb, ok := sb.(serverDiskParams); ok {
		sb.SetDiskPlan(params.DiskPlan)
		sb.SetDiskConnection(sacloud.EDiskConnection(params.DiskConnection))
		sb.SetDiskSize(params.DiskSize)
		sb.SetDistantFrom(params.DistantFrom)
	}

	// set common params
	var b serverBuilder
	b, ok := sb.(serverBuilder)
	if !ok {
		panic(fmt.Errorf("ServerCreate is failed: %s", "ServerBuilder not implements common property."))
	}

	tags := params.GetTags()

	b.SetCore(params.GetCore())
	b.SetMemory(params.GetMemory())
	b.SetServerName(params.GetName())
	b.SetDescription(params.GetDescription())
	if params.UsKeyboard {
		tags = append(tags, sacloud.TagKeyboardUS)
	}
	b.SetTags(tags)
	b.SetIconID(params.IconId)
	b.SetBootAfterCreate(!params.DisableBootAfterCreate)
	b.SetISOImageID(params.GetIsoImageId())

	// set events
	if diskEventBuilder, ok := sb.(serverDiskEventParam); ok {
		// create disk
		progCreate := internal.NewProgress(
			"Still creating disk...",
			"Create disk",
			GlobalOption.Progress)
		diskEventBuilder.SetDiskEventHandler(builder.DiskBuildOnCreateDiskBefore, func(value *builder.DiskBuildValue, result *builder.DiskBuildResult) {
			progCreate.Start()
		})
		diskEventBuilder.SetDiskEventHandler(builder.DiskBuildOnCreateDiskAfter, func(value *builder.DiskBuildValue, result *builder.DiskBuildResult) {
			progCreate.Stop()
		})

		// edit disk
		progEdit := internal.NewProgress(
			"Still editing disk...",
			"Edit disk",
			GlobalOption.Progress)
		diskEventBuilder.SetDiskEventHandler(builder.DiskBuildOnEditDiskBefore, func(value *builder.DiskBuildValue, result *builder.DiskBuildResult) {
			progEdit.Start()
		})
		diskEventBuilder.SetDiskEventHandler(builder.DiskBuildOnEditDiskAfter, func(value *builder.DiskBuildValue, result *builder.DiskBuildResult) {
			progEdit.Stop()
		})

		// cleanup startup script
		progCleanupNotes := internal.NewProgress(
			"Still cleaning StartupScript...",
			"Cleanup StartupScript",
			GlobalOption.Progress)
		diskEventBuilder.SetDiskEventHandler(builder.DiskBuildOnCleanupNoteBefore, func(value *builder.DiskBuildValue, result *builder.DiskBuildResult) {
			progCleanupNotes.Start()
		})
		diskEventBuilder.SetDiskEventHandler(builder.DiskBuildOnCleanupNoteAfter, func(value *builder.DiskBuildValue, result *builder.DiskBuildResult) {
			progCleanupNotes.Stop()
		})

		// cleanup ssh key script
		progCleanupSSHKey := internal.NewProgress(
			"Still cleaning SSHKey...",
			"Cleanup SSHKey",
			GlobalOption.Progress)
		diskEventBuilder.SetDiskEventHandler(builder.DiskBuildOnCleanupSSHKeyBefore, func(value *builder.DiskBuildValue, result *builder.DiskBuildResult) {
			progCleanupSSHKey.Start()
		})
		diskEventBuilder.SetDiskEventHandler(builder.DiskBuildOnCleanupSSHKeyAfter, func(value *builder.DiskBuildValue, result *builder.DiskBuildResult) {
			progCleanupSSHKey.Stop()
		})
	}

	if serverEventBuilder, ok := sb.(serverEventparam); ok {
		progCreate := internal.NewProgress(
			"Still creating server...",
			"Create server",
			GlobalOption.Progress)

		serverEventBuilder.SetEventHandler(builder.ServerBuildOnCreateServerBefore, func(value *builder.ServerBuildValue, result *builder.ServerBuildResult) {
			progCreate.Start()
		})
		serverEventBuilder.SetEventHandler(builder.ServerBuildOnCreateServerAfter, func(value *builder.ServerBuildValue, result *builder.ServerBuildResult) {
			progCreate.Stop()
		})

		progBoot := internal.NewProgress(
			"Still booting server...",
			"Boot server",
			GlobalOption.Progress)

		serverEventBuilder.SetEventHandler(builder.ServerBuildOnBootBefore, func(value *builder.ServerBuildValue, result *builder.ServerBuildResult) {
			progBoot.Start()
		})
		serverEventBuilder.SetEventHandler(builder.ServerBuildOnBootAfter, func(value *builder.ServerBuildValue, result *builder.ServerBuildResult) {
			progBoot.Stop()
		})

	}

	// call Create(id)
	res, err := b.Build()
	if err != nil {
		return fmt.Errorf("ServerCreate is failed: %s", err)
	}

	if len(res.Disks) > 0 && res.Disks[0].GeneratedSSHKey != nil {
		path := params.SshKeyPrivateKeyOutput
		if path == "" {
			p, err := getSSHPrivateKeyStorePath(res.Server.ID)
			if err != nil {
				return fmt.Errorf("ServerCreate is failed: getting HomeDir is failed:%s", err)
			}
			path = p
		}
		pKey := res.Disks[0].GeneratedSSHKey.PrivateKey
		dir := filepath.Dir(path)
		if err := os.MkdirAll(dir, 0600); err != nil {
			return fmt.Errorf("ServerCreate is failed: creating directory(%s) is failed:%s", dir, err)
		}

		err = ioutil.WriteFile(path, []byte(pKey), os.FileMode(0600))
		if err != nil {
			return fmt.Errorf("ServerCreate is failed: Writing private key to %s is failed:%s", params.SshKeyPrivateKeyOutput, err)
		}
	}

	return ctx.GetOutput().Print(res)
}

func validateServerDiskModeParams(ctx Context, params *BuildServerParam) []error {

	var errs []error
	var appendErrors = func(e []error) {
		errs = append(errs, e...)
	}
	var validateIfCtxIsSet = func(baseParamName string, baseParamValue interface{}, targetParamName string, targetValue interface{}) {
		if ctx.IsSet(targetParamName) {
			appendErrors(validateConflictValues(baseParamName, baseParamValue, map[string]interface{}{
				targetParamName: targetValue,
			}))
		}
	}

	switch params.DiskMode {
	case "create":
		// check required values
		appendErrors(validateRequired("disk-plan", params.DiskPlan))
		appendErrors(validateRequired("disk-connection", params.DiskConnection))
		appendErrors(validateRequired("disk-size", params.DiskSize))

		if params.SourceDiskId == 0 && params.SourceArchiveId == 0 {

			appendErrors(validateRequired("os-type", params.OsType))
			// Windows?
			if !isWindows(params.OsType) {
				appendErrors(validateRequired("password", params.Password))
			}

		} else {
			validateIfCtxIsSet("source-archive-id", params.SourceArchiveId, "os-type", params.OsType)
			validateIfCtxIsSet("source-disk-id", params.SourceArchiveId, "os-type", params.OsType)
		}

		validateIfCtxIsSet("disk-mode", params.DiskMode, "disk-id", params.DiskId)

	case "connect":
		appendErrors(validateRequired("disk-id", params.DiskId))
		validateIfCtxIsSet("disk-mode", params.DiskMode, "disk-plan", params.DiskPlan)
		validateIfCtxIsSet("disk-mode", params.DiskMode, "disk-connection", params.DiskConnection)
		validateIfCtxIsSet("disk-size", params.DiskMode, "disk-size", params.DiskSize)
		validateIfCtxIsSet("disk-size", params.DiskMode, "os-type", params.OsType)

	case "diskless":
		validateIfCtxIsSet("disk-mode", params.DiskMode, "disk-id", params.DiskId)
		validateIfCtxIsSet("disk-mode", params.DiskMode, "disk-plan", params.DiskPlan)
		validateIfCtxIsSet("disk-mode", params.DiskMode, "disk-connection", params.DiskConnection)
		validateIfCtxIsSet("disk-size", params.DiskMode, "disk-size", params.DiskSize)
		validateIfCtxIsSet("disk-size", params.DiskMode, "os-type", params.OsType)
	}

	return errs
}

func validateServerNetworkParams(sb interface{}, ctx Context, params *BuildServerParam) []error {
	var errs []error
	var appendErrors = func(e []error) {
		errs = append(errs, e...)
	}
	var validateIfCtxIsSet = func(baseParamName string, baseParamValue interface{}, targetParamName string, targetValue interface{}) {
		if ctx.IsSet(targetParamName) {
			appendErrors(validateConflictValues(baseParamName, baseParamValue, map[string]interface{}{
				targetParamName: targetValue,
			}))
		}
	}
	var validateProhibitedIfCtxIsSet = func(paramName string, paramValue interface{}) {
		if ctx.IsSet(paramName) {
			appendErrors(validateSetProhibited(paramName, paramValue))
		}
	}

	if sb, ok := sb.(serverNetworkParams); ok {
		switch params.NetworkMode {
		case "shared", "disconnect", "none":
			validateIfCtxIsSet("network-mode", params.NetworkMode, "switch-id", params.SwitchId)
			validateIfCtxIsSet("network-mode", params.NetworkMode, "ipaddress", params.Ipaddress)
			validateIfCtxIsSet("network-mode", params.NetworkMode, "nw-masklen", params.NwMasklen)
			validateIfCtxIsSet("network-mode", params.NetworkMode, "default-route", params.DefaultRoute)

			if params.NetworkMode == "none" {
				validateIfCtxIsSet("network-mode", params.NetworkMode, "use-nic-virtio", params.UseNicVirtio)
				validateIfCtxIsSet("network-mode", params.NetworkMode, "packet-filter-id", params.PacketFilterId)
			}

		case "switch":
			switch sb.(type) {
			case serverConnectSwitchParam:
				appendErrors(validateRequired("switch-id", params.SwitchId))

				validateProhibitedIfCtxIsSet("ipaddress", params.Ipaddress)
				validateProhibitedIfCtxIsSet("nw-masklen", params.NwMasklen)
				validateProhibitedIfCtxIsSet("default-route", params.DefaultRoute)

			case serverConnectSwitchParamWithEditableDisk:

				appendErrors(validateRequired("switch-id", params.SwitchId))
			}
		}

	} else {
		validateProhibitedIfCtxIsSet("network-mode", params.NetworkMode)
		validateProhibitedIfCtxIsSet("switch-id", params.SwitchId)
		validateProhibitedIfCtxIsSet("ipaddress", params.Ipaddress)
		validateProhibitedIfCtxIsSet("nw-masklen", params.NwMasklen)
		validateProhibitedIfCtxIsSet("default-route", params.DefaultRoute)
		validateProhibitedIfCtxIsSet("use-nic-virtio", params.UseNicVirtio)
		validateProhibitedIfCtxIsSet("packet-filter-id", params.PacketFilterId)
	}

	return errs
}

func validateServerDiskEditParams(sb interface{}, ctx Context, params *BuildServerParam) []error {
	var errs []error
	var appendErrors = func(e []error) {
		errs = append(errs, e...)
	}
	var validateIfCtxIsSet = func(baseParamName string, baseParamValue interface{}, targetParamName string, targetValue interface{}) {
		if ctx.IsSet(targetParamName) {
			appendErrors(validateConflictValues(baseParamName, baseParamValue, map[string]interface{}{
				targetParamName: targetValue,
			}))
		}
	}
	var validateProhibitedIfCtxIsSet = func(paramName string, paramValue interface{}) {
		if ctx.IsSet(paramName) {
			appendErrors(validateSetProhibited(paramName, paramValue))
		}
	}

	if sb, ok := sb.(serverEditDiskParam); ok {

		// SSH Key generate params
		switch params.SshKeyMode {
		case "id":
			for _, v := range params.SshKeyIds {
				sb.AddSSHKeyID(v)
			}
			validateIfCtxIsSet("ssh-key-mode", params.SshKeyMode, "ssh-key-name", params.SshKeyName)
			validateIfCtxIsSet("ssh-key-mode", params.SshKeyMode, "ssh-key-pass-phrase", params.SshKeyPassPhrase)
			validateIfCtxIsSet("ssh-key-mode", params.SshKeyMode, "ssh-key-description", params.SshKeyDescription)
			validateIfCtxIsSet("ssh-key-mode", params.SshKeyMode, "ssh-key-private-key-output", params.SshKeyPrivateKeyOutput)
			validateIfCtxIsSet("ssh-key-mode", params.SshKeyMode, "ssh-key-public-keys", params.SshKeyPublicKeys)
			validateIfCtxIsSet("ssh-key-mode", params.SshKeyMode, "ssh-key-public-key-files", params.SshKeyPublicKeyFiles)
			validateIfCtxIsSet("ssh-key-mode", params.SshKeyMode, "ssh-key-ephemeral", params.SshKeyEphemeral)
		case "generate":
			if !params.SshKeyEphemeral {
				appendErrors(validateRequired("ssh-key-name", params.SshKeyName))
			}
			validateIfCtxIsSet("ssh-key-mode", params.SshKeyMode, "ssh-key-ids", params.SshKeyIds)
			validateIfCtxIsSet("ssh-key-mode", params.SshKeyMode, "ssh-key-private-key-output", params.SshKeyPrivateKeyOutput)
			validateIfCtxIsSet("ssh-key-mode", params.SshKeyMode, "ssh-key-public-keys", params.SshKeyPublicKeys)
			validateIfCtxIsSet("ssh-key-mode", params.SshKeyMode, "ssh-key-public-key-files", params.SshKeyPublicKeyFiles)
			validateIfCtxIsSet("ssh-key-mode", params.SshKeyMode, "ssh-key-ephemeral", params.SshKeyEphemeral)
		case "upload":

			if len(params.SshKeyPublicKeys) == 0 && len(params.SshKeyPublicKeyFiles) == 0 {
				errs = append(errs,
					fmt.Errorf("%q or %q is required when %q is %q",
						"ssh-key-public-keys",
						"ssh-key-public-key-files",
						"ssh-key-mode",
						"upload",
					))
			}
			validateIfCtxIsSet("ssh-key-mode", params.SshKeyMode, "ssh-key-ids", params.SshKeyIds)
			validateIfCtxIsSet("ssh-key-mode", params.SshKeyMode, "ssh-key-name", params.SshKeyName)
			validateIfCtxIsSet("ssh-key-mode", params.SshKeyMode, "ssh-key-pass-phrase", params.SshKeyPassPhrase)
			validateIfCtxIsSet("ssh-key-mode", params.SshKeyMode, "ssh-key-description", params.SshKeyDescription)
			validateIfCtxIsSet("ssh-key-mode", params.SshKeyMode, "ssh-key-private-key-output", params.SshKeyPrivateKeyOutput)
		case "none":
			validateProhibitedIfCtxIsSet("ssh-key-mode", params.SshKeyMode)
			validateProhibitedIfCtxIsSet("ssh-key-ids", params.SshKeyIds)
			validateProhibitedIfCtxIsSet("ssh-key-name", params.SshKeyName)
			validateProhibitedIfCtxIsSet("ssh-key-pass-phrase", params.SshKeyPassPhrase)
			validateProhibitedIfCtxIsSet("ssh-key-description", params.SshKeyDescription)
			validateProhibitedIfCtxIsSet("ssh-key-private-key-output", params.SshKeyPrivateKeyOutput)
			validateProhibitedIfCtxIsSet("ssh-key-public-keys", params.SshKeyPublicKeys)
			validateProhibitedIfCtxIsSet("ssh-key-public-key-files", params.SshKeyPublicKeyFiles)
			validateProhibitedIfCtxIsSet("ssh-key-ephemeral", params.SshKeyEphemeral)
		}

	} else {
		validateProhibitedIfCtxIsSet("hostname", params.Hostname)
		validateProhibitedIfCtxIsSet("password", params.Password)
		validateProhibitedIfCtxIsSet("disable-password-auth", params.DisablePasswordAuth)
		validateProhibitedIfCtxIsSet("startup-script-ids", params.StartupScriptIds)
		validateProhibitedIfCtxIsSet("startup-scripts", params.StartupScripts)
		validateProhibitedIfCtxIsSet("startup-scripts-ephemeral", params.StartupScriptsEphemeral)
		validateProhibitedIfCtxIsSet("ssh-key-mode", params.SshKeyMode)
		validateProhibitedIfCtxIsSet("ssh-key-ids", params.SshKeyIds)
		validateProhibitedIfCtxIsSet("ssh-key-name", params.SshKeyName)
		validateProhibitedIfCtxIsSet("ssh-key-pass-phrase", params.SshKeyPassPhrase)
		validateProhibitedIfCtxIsSet("ssh-key-description", params.SshKeyDescription)
		validateProhibitedIfCtxIsSet("ssh-key-private-key-output", params.SshKeyPrivateKeyOutput)
		validateProhibitedIfCtxIsSet("ssh-key-public-keys", params.SshKeyPublicKeys)
		validateProhibitedIfCtxIsSet("ssh-key-public-key-files", params.SshKeyPublicKeyFiles)
		validateProhibitedIfCtxIsSet("ssh-key-ephemeral", params.SshKeyEphemeral)
	}

	return errs
}

func isWindows(osType string) bool {
	windowsTypes := []string{
		"windows2008", "windows2008-rds", "windows2008-rds-office",
		"windows2012", "windows2012-rds", "windows2012-rds-office",
		"windows2016", "windows2016-rds", "windows2016-rds-office",
		"windows2016-sql-web", "windows2016-sql-standard",
	}
	for _, v := range windowsTypes {
		if v == osType {
			return true
		}
	}
	return false
}

func strToOSType(osType string) ostype.ArchiveOSTypes {
	switch osType {
	case "centos":
		return ostype.CentOS
	case "ubuntu":
		return ostype.Ubuntu
	case "debian":
		return ostype.Debian
	case "vyos":
		return ostype.VyOS
	case "coreos":
		return ostype.CoreOS
	case "kusanagi":
		return ostype.Kusanagi
	case "site-guard":
		return ostype.SiteGuard
	case "freebsd":
		return ostype.FreeBSD
	case "windows2008":
		return ostype.Windows2008
	case "windows2008-rds":
		return ostype.Windows2008RDS
	case "windows2008-rds-office":
		return ostype.Windows2008RDSOffice
	case "windows2012":
		return ostype.Windows2012
	case "windows2012-rds":
		return ostype.Windows2012RDS
	case "windows2012-rds-office":
		return ostype.Windows2012RDSOffice
	case "windows2016":
		return ostype.Windows2016
	case "windows2016-rds":
		return ostype.Windows2016RDS
	case "windows2016-rds-office":
		return ostype.Windows2016RDSOffice
	case "windows2016-sql-web":
		return ostype.Windows2016SQLServerWeb
	case "windows2016-sql-standard":
		return ostype.Windows2016SQLServerStandard
	default:
		return ostype.Custom
	}
}

type serverBuilder interface {
	SetCore(int)
	SetMemory(int)
	SetServerName(string)
	SetDescription(string)
	SetTags([]string)
	SetIconID(int64)
	SetBootAfterCreate(bool)
	SetISOImageID(int64)

	Build() (*builder.ServerBuildResult, error)
}

type serverDiskParams interface {
	SetDiskPlan(string)
	SetDiskConnection(sacloud.EDiskConnection)
	SetDiskSize(int)
	SetDistantFrom([]int64)
}

type serverNetworkParams interface {
	SetUseVirtIONetPCI(bool)
	SetPacketFilterIDs([]int64)
	AddPublicNWConnectedNIC()
	AddDisconnectedNIC()
}

type serverConnectSwitchParamWithEditableDisk interface {
	AddExistsSwitchConnectedNIC(id string, ipaddress string, maskLen int, defRoute string)
}

type serverConnectSwitchParam interface {
	AddExistsSwitchConnectedNIC(id string)
}

type serverEditDiskParam interface {
	SetHostName(string)
	SetPassword(string)
	SetDisablePWAuth(bool)
	AddNote(string)
	AddNoteID(int64)
	SetNotesEphemeral(bool)
	AddSSHKey(string)
	AddSSHKeyID(int64)
	SetSSHKeysEphemeral(bool)
	SetGenerateSSHKeyName(string)
	SetGenerateSSHKeyPassPhrase(string)
	SetGenerateSSHKeyDescription(string)
}

type serverDiskEventParam interface {
	SetDiskEventHandler(event builder.DiskBuildEvents, handler builder.DiskBuildEventHandler)
}

type serverEventparam interface {
	SetEventHandler(event builder.ServerBuildEvents, handler builder.ServerBuildEventHandler)
}
