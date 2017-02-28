// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package command

import (
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

// ListProductInternetParam is input parameters for the sacloud API
type ListProductInternetParam struct {
	From int
	Id   []int64
	Max  int
	Name []string
	Sort []string
}

// NewListProductInternetParam return new ListProductInternetParam
func NewListProductInternetParam() *ListProductInternetParam {
	return &ListProductInternetParam{}
}

// Validate checks current values in model
func (p *ListProductInternetParam) Validate() []error {
	errors := []error{}
	{
		validator := define.Resources["ProductInternet"].Commands["list"].Params["id"].ValidateFunc
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
		errs := validateConflicts("--name", p.Name, map[string]interface{}{

			"--id": p.Id,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ListProductInternetParam) getResourceDef() *schema.Resource {
	return define.Resources["ProductInternet"]
}

func (p *ListProductInternetParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["list"]
}

func (p *ListProductInternetParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ListProductInternetParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ListProductInternetParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ListProductInternetParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ListProductInternetParam) SetFrom(v int) {
	p.From = v
}

func (p *ListProductInternetParam) GetFrom() int {
	return p.From
}
func (p *ListProductInternetParam) SetId(v []int64) {
	p.Id = v
}

func (p *ListProductInternetParam) GetId() []int64 {
	return p.Id
}
func (p *ListProductInternetParam) SetMax(v int) {
	p.Max = v
}

func (p *ListProductInternetParam) GetMax() int {
	return p.Max
}
func (p *ListProductInternetParam) SetName(v []string) {
	p.Name = v
}

func (p *ListProductInternetParam) GetName() []string {
	return p.Name
}
func (p *ListProductInternetParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListProductInternetParam) GetSort() []string {
	return p.Sort
}

// ReadProductInternetParam is input parameters for the sacloud API
type ReadProductInternetParam struct {
	Id int64
}

// NewReadProductInternetParam return new ReadProductInternetParam
func NewReadProductInternetParam() *ReadProductInternetParam {
	return &ReadProductInternetParam{}
}

// Validate checks current values in model
func (p *ReadProductInternetParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["ProductInternet"].Commands["read"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ReadProductInternetParam) getResourceDef() *schema.Resource {
	return define.Resources["ProductInternet"]
}

func (p *ReadProductInternetParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["read"]
}

func (p *ReadProductInternetParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ReadProductInternetParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ReadProductInternetParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ReadProductInternetParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ReadProductInternetParam) SetId(v int64) {
	p.Id = v
}

func (p *ReadProductInternetParam) GetId() int64 {
	return p.Id
}