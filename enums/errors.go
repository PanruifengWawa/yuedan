package enums

type Errors string

const (
	Success        Errors = ""
	AuthError             = "权限错误"
	ParameterError        = "参数错误"
	DBError               = "数据库错误"
)
