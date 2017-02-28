// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package command

import (
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

// ListRegionParam is input parameters for the sacloud API
type ListRegionParam struct {
	From int
	Id   []int64
	Max  int
	Name []string
	Sort []string
}

// NewListRegionParam return new ListRegionParam
func NewListRegionParam() *ListRegionParam {
	return &ListRegionParam{}
}

// Validate checks current values in model
func (p *ListRegionParam) Validate() []error {
	errors := []error{}
	{
		validator := define.Resources["Region"].Commands["list"].Params["id"].ValidateFunc
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

func (p *ListRegionParam) getResourceDef() *schema.Resource {
	return define.Resources["Region"]
}

func (p *ListRegionParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["list"]
}

func (p *ListRegionParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ListRegionParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ListRegionParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ListRegionParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ListRegionParam) SetFrom(v int) {
	p.From = v
}

func (p *ListRegionParam) GetFrom() int {
	return p.From
}
func (p *ListRegionParam) SetId(v []int64) {
	p.Id = v
}

func (p *ListRegionParam) GetId() []int64 {
	return p.Id
}
func (p *ListRegionParam) SetMax(v int) {
	p.Max = v
}

func (p *ListRegionParam) GetMax() int {
	return p.Max
}
func (p *ListRegionParam) SetName(v []string) {
	p.Name = v
}

func (p *ListRegionParam) GetName() []string {
	return p.Name
}
func (p *ListRegionParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListRegionParam) GetSort() []string {
	return p.Sort
}

// ReadRegionParam is input parameters for the sacloud API
type ReadRegionParam struct {
	Id int64
}

// NewReadRegionParam return new ReadRegionParam
func NewReadRegionParam() *ReadRegionParam {
	return &ReadRegionParam{}
}

// Validate checks current values in model
func (p *ReadRegionParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Region"].Commands["read"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ReadRegionParam) getResourceDef() *schema.Resource {
	return define.Resources["Region"]
}

func (p *ReadRegionParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["read"]
}

func (p *ReadRegionParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ReadRegionParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ReadRegionParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ReadRegionParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ReadRegionParam) SetId(v int64) {
	p.Id = v
}

func (p *ReadRegionParam) GetId() int64 {
	return p.Id
}