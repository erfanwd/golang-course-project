package logging

type Category string
type SubCategory string
type Extra string

const (
	General         Category = "General"
	Internal        Category = "Internal"
	Postgres        Category = "Postgres"
	Redis           Category = "Redis"
	Validation      Category = "Validation"
	RequestResponse Category = "RequestResponse"
)

const (
	StartUp         SubCategory = "StartUp"
	ExternalService SubCategory = "ExternalService"
	//postgres
	Select    SubCategory = "Select"
	Migration SubCategory = "Migration"
	Rollback  SubCategory = "Rollback"
	Update    SubCategory = "Update"
	Delete    SubCategory = "Delete"
	Insert    SubCategory = "Insert"
	//internal
	Api                 SubCategory = "Api"
	HashPassword        SubCategory = "HashPassword"
	DefaultRoleNotFound SubCategory = "DefaultRoleNotFound"
	//validation
	MobileValidation   SubCategory = "MobileValidation"
	PasswordValidation SubCategory = "PasswordValidation"
)

const (
	AppName      Extra = "AppName"
	LoggerName   Extra = "LoggerName"
	ClientIp     Extra = "ClientIp"
	HostIp       Extra = "HostIp"
	Method       Extra = "Method"
	StatusCode   Extra = "StatusCode"
	BodySize     Extra = "BodySize"
	Latency      Extra = "Latency"
	Body         Extra = "Body"
	ErrorMessage Extra = "ErrorMessage"
	Path         Extra = "Path"
	RequestBody  Extra = "RequestBody"
	ResponseBody Extra = "ResponseBody"
	TablesData   Extra = "TablesData"
)
