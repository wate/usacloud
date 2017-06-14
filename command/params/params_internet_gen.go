// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package params

import (
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

// ListInternetParam is input parameters for the sacloud API
type ListInternetParam struct {
	Name              []string
	Id                []int64
	Tags              []string
	From              int
	Max               int
	Sort              []string
	ParamTemplate     string
	ParamTemplateFile string
	OutputType        string
	Column            []string
	Quiet             bool
	Format            string
	FormatFile        string
}

// NewListInternetParam return new ListInternetParam
func NewListInternetParam() *ListInternetParam {
	return &ListInternetParam{}
}

// Validate checks current values in model
func (p *ListInternetParam) Validate() []error {
	errors := []error{}
	{
		errs := validateConflicts("--name", p.Name, map[string]interface{}{

			"--id": p.Id,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Internet"].Commands["list"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateConflicts("--id", p.Id, map[string]interface{}{

			"--name": p.Name,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Internet"].Commands["list"].Params["tags"].ValidateFunc
		errs := validator("--tags", p.Tags)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues("json", "csv", "tsv")
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateInputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateOutputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ListInternetParam) GetResourceDef() *schema.Resource {
	return define.Resources["Internet"]
}

func (p *ListInternetParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["list"]
}

func (p *ListInternetParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *ListInternetParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *ListInternetParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *ListInternetParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *ListInternetParam) GetOutputFormat() string {
	return "table"
}

func (p *ListInternetParam) SetName(v []string) {
	p.Name = v
}

func (p *ListInternetParam) GetName() []string {
	return p.Name
}
func (p *ListInternetParam) SetId(v []int64) {
	p.Id = v
}

func (p *ListInternetParam) GetId() []int64 {
	return p.Id
}
func (p *ListInternetParam) SetTags(v []string) {
	p.Tags = v
}

func (p *ListInternetParam) GetTags() []string {
	return p.Tags
}
func (p *ListInternetParam) SetFrom(v int) {
	p.From = v
}

func (p *ListInternetParam) GetFrom() int {
	return p.From
}
func (p *ListInternetParam) SetMax(v int) {
	p.Max = v
}

func (p *ListInternetParam) GetMax() int {
	return p.Max
}
func (p *ListInternetParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListInternetParam) GetSort() []string {
	return p.Sort
}
func (p *ListInternetParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ListInternetParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ListInternetParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ListInternetParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ListInternetParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ListInternetParam) GetOutputType() string {
	return p.OutputType
}
func (p *ListInternetParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ListInternetParam) GetColumn() []string {
	return p.Column
}
func (p *ListInternetParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ListInternetParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ListInternetParam) SetFormat(v string) {
	p.Format = v
}

func (p *ListInternetParam) GetFormat() string {
	return p.Format
}
func (p *ListInternetParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ListInternetParam) GetFormatFile() string {
	return p.FormatFile
}

// MonitorInternetParam is input parameters for the sacloud API
type MonitorInternetParam struct {
	Start             string
	End               string
	KeyFormat         string
	ParamTemplate     string
	ParamTemplateFile string
	OutputType        string
	Column            []string
	Quiet             bool
	Format            string
	FormatFile        string
	Id                int64
}

// NewMonitorInternetParam return new MonitorInternetParam
func NewMonitorInternetParam() *MonitorInternetParam {
	return &MonitorInternetParam{

		KeyFormat: "sakuracloud.{{.ID}}.internet",
	}
}

// Validate checks current values in model
func (p *MonitorInternetParam) Validate() []error {
	errors := []error{}
	{
		validator := define.Resources["Internet"].Commands["monitor"].Params["start"].ValidateFunc
		errs := validator("--start", p.Start)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Internet"].Commands["monitor"].Params["end"].ValidateFunc
		errs := validator("--end", p.End)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--key-format", p.KeyFormat)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateSakuraID
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues("json", "csv", "tsv")
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateInputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateOutputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *MonitorInternetParam) GetResourceDef() *schema.Resource {
	return define.Resources["Internet"]
}

func (p *MonitorInternetParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["monitor"]
}

func (p *MonitorInternetParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *MonitorInternetParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *MonitorInternetParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *MonitorInternetParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *MonitorInternetParam) GetOutputFormat() string {
	return "table"
}

func (p *MonitorInternetParam) SetStart(v string) {
	p.Start = v
}

func (p *MonitorInternetParam) GetStart() string {
	return p.Start
}
func (p *MonitorInternetParam) SetEnd(v string) {
	p.End = v
}

func (p *MonitorInternetParam) GetEnd() string {
	return p.End
}
func (p *MonitorInternetParam) SetKeyFormat(v string) {
	p.KeyFormat = v
}

func (p *MonitorInternetParam) GetKeyFormat() string {
	return p.KeyFormat
}
func (p *MonitorInternetParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *MonitorInternetParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *MonitorInternetParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *MonitorInternetParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *MonitorInternetParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *MonitorInternetParam) GetOutputType() string {
	return p.OutputType
}
func (p *MonitorInternetParam) SetColumn(v []string) {
	p.Column = v
}

func (p *MonitorInternetParam) GetColumn() []string {
	return p.Column
}
func (p *MonitorInternetParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *MonitorInternetParam) GetQuiet() bool {
	return p.Quiet
}
func (p *MonitorInternetParam) SetFormat(v string) {
	p.Format = v
}

func (p *MonitorInternetParam) GetFormat() string {
	return p.Format
}
func (p *MonitorInternetParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *MonitorInternetParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *MonitorInternetParam) SetId(v int64) {
	p.Id = v
}

func (p *MonitorInternetParam) GetId() int64 {
	return p.Id
}

// UpdateBandwidthInternetParam is input parameters for the sacloud API
type UpdateBandwidthInternetParam struct {
	BandWidth         int
	Assumeyes         bool
	ParamTemplate     string
	ParamTemplateFile string
	OutputType        string
	Column            []string
	Quiet             bool
	Format            string
	FormatFile        string
	Id                int64
}

// NewUpdateBandwidthInternetParam return new UpdateBandwidthInternetParam
func NewUpdateBandwidthInternetParam() *UpdateBandwidthInternetParam {
	return &UpdateBandwidthInternetParam{

		BandWidth: 100,
	}
}

// Validate checks current values in model
func (p *UpdateBandwidthInternetParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--band-width", p.BandWidth)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Internet"].Commands["update-bandwidth"].Params["band-width"].ValidateFunc
		errs := validator("--band-width", p.BandWidth)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateSakuraID
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues("json", "csv", "tsv")
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateInputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateOutputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *UpdateBandwidthInternetParam) GetResourceDef() *schema.Resource {
	return define.Resources["Internet"]
}

func (p *UpdateBandwidthInternetParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["update-bandwidth"]
}

func (p *UpdateBandwidthInternetParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *UpdateBandwidthInternetParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *UpdateBandwidthInternetParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *UpdateBandwidthInternetParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *UpdateBandwidthInternetParam) GetOutputFormat() string {
	return "table"
}

func (p *UpdateBandwidthInternetParam) SetBandWidth(v int) {
	p.BandWidth = v
}

func (p *UpdateBandwidthInternetParam) GetBandWidth() int {
	return p.BandWidth
}
func (p *UpdateBandwidthInternetParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *UpdateBandwidthInternetParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *UpdateBandwidthInternetParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *UpdateBandwidthInternetParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *UpdateBandwidthInternetParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *UpdateBandwidthInternetParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *UpdateBandwidthInternetParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *UpdateBandwidthInternetParam) GetOutputType() string {
	return p.OutputType
}
func (p *UpdateBandwidthInternetParam) SetColumn(v []string) {
	p.Column = v
}

func (p *UpdateBandwidthInternetParam) GetColumn() []string {
	return p.Column
}
func (p *UpdateBandwidthInternetParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *UpdateBandwidthInternetParam) GetQuiet() bool {
	return p.Quiet
}
func (p *UpdateBandwidthInternetParam) SetFormat(v string) {
	p.Format = v
}

func (p *UpdateBandwidthInternetParam) GetFormat() string {
	return p.Format
}
func (p *UpdateBandwidthInternetParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *UpdateBandwidthInternetParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *UpdateBandwidthInternetParam) SetId(v int64) {
	p.Id = v
}

func (p *UpdateBandwidthInternetParam) GetId() int64 {
	return p.Id
}

// CreateInternetParam is input parameters for the sacloud API
type CreateInternetParam struct {
	NwMasklen         int
	BandWidth         int
	Name              string
	Description       string
	Tags              []string
	IconId            int64
	Assumeyes         bool
	ParamTemplate     string
	ParamTemplateFile string
	OutputType        string
	Column            []string
	Quiet             bool
	Format            string
	FormatFile        string
}

// NewCreateInternetParam return new CreateInternetParam
func NewCreateInternetParam() *CreateInternetParam {
	return &CreateInternetParam{

		NwMasklen: 28,

		BandWidth: 100,
	}
}

// Validate checks current values in model
func (p *CreateInternetParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--nw-masklen", p.NwMasklen)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Internet"].Commands["create"].Params["nw-masklen"].ValidateFunc
		errs := validator("--nw-masklen", p.NwMasklen)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--band-width", p.BandWidth)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Internet"].Commands["create"].Params["band-width"].ValidateFunc
		errs := validator("--band-width", p.BandWidth)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Internet"].Commands["create"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Internet"].Commands["create"].Params["description"].ValidateFunc
		errs := validator("--description", p.Description)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Internet"].Commands["create"].Params["tags"].ValidateFunc
		errs := validator("--tags", p.Tags)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Internet"].Commands["create"].Params["icon-id"].ValidateFunc
		errs := validator("--icon-id", p.IconId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues("json", "csv", "tsv")
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateInputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateOutputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *CreateInternetParam) GetResourceDef() *schema.Resource {
	return define.Resources["Internet"]
}

func (p *CreateInternetParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["create"]
}

func (p *CreateInternetParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *CreateInternetParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *CreateInternetParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *CreateInternetParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *CreateInternetParam) GetOutputFormat() string {
	return "table"
}

func (p *CreateInternetParam) SetNwMasklen(v int) {
	p.NwMasklen = v
}

func (p *CreateInternetParam) GetNwMasklen() int {
	return p.NwMasklen
}
func (p *CreateInternetParam) SetBandWidth(v int) {
	p.BandWidth = v
}

func (p *CreateInternetParam) GetBandWidth() int {
	return p.BandWidth
}
func (p *CreateInternetParam) SetName(v string) {
	p.Name = v
}

func (p *CreateInternetParam) GetName() string {
	return p.Name
}
func (p *CreateInternetParam) SetDescription(v string) {
	p.Description = v
}

func (p *CreateInternetParam) GetDescription() string {
	return p.Description
}
func (p *CreateInternetParam) SetTags(v []string) {
	p.Tags = v
}

func (p *CreateInternetParam) GetTags() []string {
	return p.Tags
}
func (p *CreateInternetParam) SetIconId(v int64) {
	p.IconId = v
}

func (p *CreateInternetParam) GetIconId() int64 {
	return p.IconId
}
func (p *CreateInternetParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *CreateInternetParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *CreateInternetParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *CreateInternetParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *CreateInternetParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *CreateInternetParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *CreateInternetParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *CreateInternetParam) GetOutputType() string {
	return p.OutputType
}
func (p *CreateInternetParam) SetColumn(v []string) {
	p.Column = v
}

func (p *CreateInternetParam) GetColumn() []string {
	return p.Column
}
func (p *CreateInternetParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *CreateInternetParam) GetQuiet() bool {
	return p.Quiet
}
func (p *CreateInternetParam) SetFormat(v string) {
	p.Format = v
}

func (p *CreateInternetParam) GetFormat() string {
	return p.Format
}
func (p *CreateInternetParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *CreateInternetParam) GetFormatFile() string {
	return p.FormatFile
}

// ReadInternetParam is input parameters for the sacloud API
type ReadInternetParam struct {
	ParamTemplate     string
	ParamTemplateFile string
	OutputType        string
	Column            []string
	Quiet             bool
	Format            string
	FormatFile        string
	Id                int64
}

// NewReadInternetParam return new ReadInternetParam
func NewReadInternetParam() *ReadInternetParam {
	return &ReadInternetParam{}
}

// Validate checks current values in model
func (p *ReadInternetParam) Validate() []error {
	errors := []error{}
	{
		validator := validateSakuraID
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues("json", "csv", "tsv")
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateInputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateOutputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ReadInternetParam) GetResourceDef() *schema.Resource {
	return define.Resources["Internet"]
}

func (p *ReadInternetParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["read"]
}

func (p *ReadInternetParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *ReadInternetParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *ReadInternetParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *ReadInternetParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *ReadInternetParam) GetOutputFormat() string {
	return "table"
}

func (p *ReadInternetParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ReadInternetParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ReadInternetParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ReadInternetParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ReadInternetParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ReadInternetParam) GetOutputType() string {
	return p.OutputType
}
func (p *ReadInternetParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ReadInternetParam) GetColumn() []string {
	return p.Column
}
func (p *ReadInternetParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ReadInternetParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ReadInternetParam) SetFormat(v string) {
	p.Format = v
}

func (p *ReadInternetParam) GetFormat() string {
	return p.Format
}
func (p *ReadInternetParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ReadInternetParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ReadInternetParam) SetId(v int64) {
	p.Id = v
}

func (p *ReadInternetParam) GetId() int64 {
	return p.Id
}

// UpdateInternetParam is input parameters for the sacloud API
type UpdateInternetParam struct {
	BandWidth         int
	Name              string
	Description       string
	Tags              []string
	IconId            int64
	Assumeyes         bool
	ParamTemplate     string
	ParamTemplateFile string
	OutputType        string
	Column            []string
	Quiet             bool
	Format            string
	FormatFile        string
	Id                int64
}

// NewUpdateInternetParam return new UpdateInternetParam
func NewUpdateInternetParam() *UpdateInternetParam {
	return &UpdateInternetParam{}
}

// Validate checks current values in model
func (p *UpdateInternetParam) Validate() []error {
	errors := []error{}
	{
		validator := define.Resources["Internet"].Commands["update"].Params["band-width"].ValidateFunc
		errs := validator("--band-width", p.BandWidth)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Internet"].Commands["update"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Internet"].Commands["update"].Params["description"].ValidateFunc
		errs := validator("--description", p.Description)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Internet"].Commands["update"].Params["tags"].ValidateFunc
		errs := validator("--tags", p.Tags)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Internet"].Commands["update"].Params["icon-id"].ValidateFunc
		errs := validator("--icon-id", p.IconId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateSakuraID
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues("json", "csv", "tsv")
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateInputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateOutputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *UpdateInternetParam) GetResourceDef() *schema.Resource {
	return define.Resources["Internet"]
}

func (p *UpdateInternetParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["update"]
}

func (p *UpdateInternetParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *UpdateInternetParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *UpdateInternetParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *UpdateInternetParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *UpdateInternetParam) GetOutputFormat() string {
	return "table"
}

func (p *UpdateInternetParam) SetBandWidth(v int) {
	p.BandWidth = v
}

func (p *UpdateInternetParam) GetBandWidth() int {
	return p.BandWidth
}
func (p *UpdateInternetParam) SetName(v string) {
	p.Name = v
}

func (p *UpdateInternetParam) GetName() string {
	return p.Name
}
func (p *UpdateInternetParam) SetDescription(v string) {
	p.Description = v
}

func (p *UpdateInternetParam) GetDescription() string {
	return p.Description
}
func (p *UpdateInternetParam) SetTags(v []string) {
	p.Tags = v
}

func (p *UpdateInternetParam) GetTags() []string {
	return p.Tags
}
func (p *UpdateInternetParam) SetIconId(v int64) {
	p.IconId = v
}

func (p *UpdateInternetParam) GetIconId() int64 {
	return p.IconId
}
func (p *UpdateInternetParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *UpdateInternetParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *UpdateInternetParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *UpdateInternetParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *UpdateInternetParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *UpdateInternetParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *UpdateInternetParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *UpdateInternetParam) GetOutputType() string {
	return p.OutputType
}
func (p *UpdateInternetParam) SetColumn(v []string) {
	p.Column = v
}

func (p *UpdateInternetParam) GetColumn() []string {
	return p.Column
}
func (p *UpdateInternetParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *UpdateInternetParam) GetQuiet() bool {
	return p.Quiet
}
func (p *UpdateInternetParam) SetFormat(v string) {
	p.Format = v
}

func (p *UpdateInternetParam) GetFormat() string {
	return p.Format
}
func (p *UpdateInternetParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *UpdateInternetParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *UpdateInternetParam) SetId(v int64) {
	p.Id = v
}

func (p *UpdateInternetParam) GetId() int64 {
	return p.Id
}

// DeleteInternetParam is input parameters for the sacloud API
type DeleteInternetParam struct {
	Assumeyes         bool
	ParamTemplate     string
	ParamTemplateFile string
	OutputType        string
	Column            []string
	Quiet             bool
	Format            string
	FormatFile        string
	Id                int64
}

// NewDeleteInternetParam return new DeleteInternetParam
func NewDeleteInternetParam() *DeleteInternetParam {
	return &DeleteInternetParam{}
}

// Validate checks current values in model
func (p *DeleteInternetParam) Validate() []error {
	errors := []error{}
	{
		validator := validateSakuraID
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues("json", "csv", "tsv")
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateInputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateOutputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *DeleteInternetParam) GetResourceDef() *schema.Resource {
	return define.Resources["Internet"]
}

func (p *DeleteInternetParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["delete"]
}

func (p *DeleteInternetParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *DeleteInternetParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *DeleteInternetParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *DeleteInternetParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *DeleteInternetParam) GetOutputFormat() string {
	return "table"
}

func (p *DeleteInternetParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *DeleteInternetParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *DeleteInternetParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *DeleteInternetParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *DeleteInternetParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *DeleteInternetParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *DeleteInternetParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *DeleteInternetParam) GetOutputType() string {
	return p.OutputType
}
func (p *DeleteInternetParam) SetColumn(v []string) {
	p.Column = v
}

func (p *DeleteInternetParam) GetColumn() []string {
	return p.Column
}
func (p *DeleteInternetParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *DeleteInternetParam) GetQuiet() bool {
	return p.Quiet
}
func (p *DeleteInternetParam) SetFormat(v string) {
	p.Format = v
}

func (p *DeleteInternetParam) GetFormat() string {
	return p.Format
}
func (p *DeleteInternetParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *DeleteInternetParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *DeleteInternetParam) SetId(v int64) {
	p.Id = v
}

func (p *DeleteInternetParam) GetId() int64 {
	return p.Id
}