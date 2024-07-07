package logging

type Category string
type SubCategory string

const (
	General         Category = "General"
	IO              Category = "IO"
	Internal        Category = "Internal"
	Postgres        Category = "Postgres"
	Redis           Category = "Redis"
	Validation      Category = "Validation"
	RequestResponse Category = "RequestResponse"
	Prometheus      Category = "Prometheus"
)

const (
	Startup         SubCategory = "Startup"
	ExternalService SubCategory = "ExternalService"

	Migration SubCategory = "Migration"
	Select    SubCategory = "Select"
	Rollback  SubCategory = "Rollback"
	Update    SubCategory = "Update"
	Delete    SubCategory = "Delete"
	Insert    SubCategory = "Insert"

	PasswordValidation SubCategory = "PasswordValidation"
)

const (
	Package      string = "Package"
	Method       string = "Method"
	StatusCode   string = "StatusCode"
	RequestBody  string = "RequestBody"
	ResponseBody string = "ResponseBody"
	ErrorMessage string = "ErrorMessage"
)
