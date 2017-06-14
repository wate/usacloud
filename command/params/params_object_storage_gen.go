// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package params

import (
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

// ListObjectStorageParam is input parameters for the sacloud API
type ListObjectStorageParam struct {
	AccessKey         string
	SecretKey         string
	Bucket            string
	ParamTemplate     string
	ParamTemplateFile string
	OutputType        string
	Column            []string
	Quiet             bool
	Format            string
	FormatFile        string
}

// NewListObjectStorageParam return new ListObjectStorageParam
func NewListObjectStorageParam() *ListObjectStorageParam {
	return &ListObjectStorageParam{}
}

// Validate checks current values in model
func (p *ListObjectStorageParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--access-key", p.AccessKey)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--secret-key", p.SecretKey)
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

func (p *ListObjectStorageParam) GetResourceDef() *schema.Resource {
	return define.Resources["ObjectStorage"]
}

func (p *ListObjectStorageParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["list"]
}

func (p *ListObjectStorageParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *ListObjectStorageParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *ListObjectStorageParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *ListObjectStorageParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *ListObjectStorageParam) GetOutputFormat() string {
	return "table"
}

func (p *ListObjectStorageParam) SetAccessKey(v string) {
	p.AccessKey = v
}

func (p *ListObjectStorageParam) GetAccessKey() string {
	return p.AccessKey
}
func (p *ListObjectStorageParam) SetSecretKey(v string) {
	p.SecretKey = v
}

func (p *ListObjectStorageParam) GetSecretKey() string {
	return p.SecretKey
}
func (p *ListObjectStorageParam) SetBucket(v string) {
	p.Bucket = v
}

func (p *ListObjectStorageParam) GetBucket() string {
	return p.Bucket
}
func (p *ListObjectStorageParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ListObjectStorageParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ListObjectStorageParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ListObjectStorageParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ListObjectStorageParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ListObjectStorageParam) GetOutputType() string {
	return p.OutputType
}
func (p *ListObjectStorageParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ListObjectStorageParam) GetColumn() []string {
	return p.Column
}
func (p *ListObjectStorageParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ListObjectStorageParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ListObjectStorageParam) SetFormat(v string) {
	p.Format = v
}

func (p *ListObjectStorageParam) GetFormat() string {
	return p.Format
}
func (p *ListObjectStorageParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ListObjectStorageParam) GetFormatFile() string {
	return p.FormatFile
}

// PutObjectStorageParam is input parameters for the sacloud API
type PutObjectStorageParam struct {
	AccessKey         string
	ContentType       string
	Recursive         bool
	SecretKey         string
	Bucket            string
	Assumeyes         bool
	ParamTemplate     string
	ParamTemplateFile string
}

// NewPutObjectStorageParam return new PutObjectStorageParam
func NewPutObjectStorageParam() *PutObjectStorageParam {
	return &PutObjectStorageParam{

		ContentType: "application/octet-stream",
	}
}

// Validate checks current values in model
func (p *PutObjectStorageParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--access-key", p.AccessKey)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--secret-key", p.SecretKey)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *PutObjectStorageParam) GetResourceDef() *schema.Resource {
	return define.Resources["ObjectStorage"]
}

func (p *PutObjectStorageParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["put"]
}

func (p *PutObjectStorageParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *PutObjectStorageParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *PutObjectStorageParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *PutObjectStorageParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *PutObjectStorageParam) GetOutputFormat() string {
	return "table"
}

func (p *PutObjectStorageParam) SetAccessKey(v string) {
	p.AccessKey = v
}

func (p *PutObjectStorageParam) GetAccessKey() string {
	return p.AccessKey
}
func (p *PutObjectStorageParam) SetContentType(v string) {
	p.ContentType = v
}

func (p *PutObjectStorageParam) GetContentType() string {
	return p.ContentType
}
func (p *PutObjectStorageParam) SetRecursive(v bool) {
	p.Recursive = v
}

func (p *PutObjectStorageParam) GetRecursive() bool {
	return p.Recursive
}
func (p *PutObjectStorageParam) SetSecretKey(v string) {
	p.SecretKey = v
}

func (p *PutObjectStorageParam) GetSecretKey() string {
	return p.SecretKey
}
func (p *PutObjectStorageParam) SetBucket(v string) {
	p.Bucket = v
}

func (p *PutObjectStorageParam) GetBucket() string {
	return p.Bucket
}
func (p *PutObjectStorageParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *PutObjectStorageParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *PutObjectStorageParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *PutObjectStorageParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *PutObjectStorageParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *PutObjectStorageParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}

// GetObjectStorageParam is input parameters for the sacloud API
type GetObjectStorageParam struct {
	AccessKey         string
	Recursive         bool
	SecretKey         string
	Bucket            string
	ParamTemplate     string
	ParamTemplateFile string
}

// NewGetObjectStorageParam return new GetObjectStorageParam
func NewGetObjectStorageParam() *GetObjectStorageParam {
	return &GetObjectStorageParam{}
}

// Validate checks current values in model
func (p *GetObjectStorageParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--access-key", p.AccessKey)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--secret-key", p.SecretKey)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *GetObjectStorageParam) GetResourceDef() *schema.Resource {
	return define.Resources["ObjectStorage"]
}

func (p *GetObjectStorageParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["get"]
}

func (p *GetObjectStorageParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *GetObjectStorageParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *GetObjectStorageParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *GetObjectStorageParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *GetObjectStorageParam) GetOutputFormat() string {
	return "table"
}

func (p *GetObjectStorageParam) SetAccessKey(v string) {
	p.AccessKey = v
}

func (p *GetObjectStorageParam) GetAccessKey() string {
	return p.AccessKey
}
func (p *GetObjectStorageParam) SetRecursive(v bool) {
	p.Recursive = v
}

func (p *GetObjectStorageParam) GetRecursive() bool {
	return p.Recursive
}
func (p *GetObjectStorageParam) SetSecretKey(v string) {
	p.SecretKey = v
}

func (p *GetObjectStorageParam) GetSecretKey() string {
	return p.SecretKey
}
func (p *GetObjectStorageParam) SetBucket(v string) {
	p.Bucket = v
}

func (p *GetObjectStorageParam) GetBucket() string {
	return p.Bucket
}
func (p *GetObjectStorageParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *GetObjectStorageParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *GetObjectStorageParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *GetObjectStorageParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}

// DeleteObjectStorageParam is input parameters for the sacloud API
type DeleteObjectStorageParam struct {
	AccessKey         string
	Recursive         bool
	SecretKey         string
	Bucket            string
	Assumeyes         bool
	ParamTemplate     string
	ParamTemplateFile string
}

// NewDeleteObjectStorageParam return new DeleteObjectStorageParam
func NewDeleteObjectStorageParam() *DeleteObjectStorageParam {
	return &DeleteObjectStorageParam{}
}

// Validate checks current values in model
func (p *DeleteObjectStorageParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--access-key", p.AccessKey)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--secret-key", p.SecretKey)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *DeleteObjectStorageParam) GetResourceDef() *schema.Resource {
	return define.Resources["ObjectStorage"]
}

func (p *DeleteObjectStorageParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["delete"]
}

func (p *DeleteObjectStorageParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *DeleteObjectStorageParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *DeleteObjectStorageParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *DeleteObjectStorageParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *DeleteObjectStorageParam) GetOutputFormat() string {
	return "table"
}

func (p *DeleteObjectStorageParam) SetAccessKey(v string) {
	p.AccessKey = v
}

func (p *DeleteObjectStorageParam) GetAccessKey() string {
	return p.AccessKey
}
func (p *DeleteObjectStorageParam) SetRecursive(v bool) {
	p.Recursive = v
}

func (p *DeleteObjectStorageParam) GetRecursive() bool {
	return p.Recursive
}
func (p *DeleteObjectStorageParam) SetSecretKey(v string) {
	p.SecretKey = v
}

func (p *DeleteObjectStorageParam) GetSecretKey() string {
	return p.SecretKey
}
func (p *DeleteObjectStorageParam) SetBucket(v string) {
	p.Bucket = v
}

func (p *DeleteObjectStorageParam) GetBucket() string {
	return p.Bucket
}
func (p *DeleteObjectStorageParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *DeleteObjectStorageParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *DeleteObjectStorageParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *DeleteObjectStorageParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *DeleteObjectStorageParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *DeleteObjectStorageParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}