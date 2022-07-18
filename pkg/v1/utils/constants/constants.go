package constants

const (
	EnvProduction = "production"
)

const (
	SuccessCode = `0000`
	SuccesDesc  = `SUCCESS`
)

const (
	Table_Custom_Main = "custom.main"
)

const (
	MethodGET    = "GET"
	MethodPOST   = "POST"
	MethodPUT    = "PUT"
	MethodDELETE = "DELETE"
)

const (
	Host_Reqres = "https://jsonplaceholder.typicode.com"
)

const (
	JSONType       = "application/json"
	XMLType        = "application/xml"
	URLEncodedType = "application/x-www-form-urlencoded"
	STREAMType     = "application/octet-stream"
)

const (
	Jwt_Refresh_Expired_Periode = 3600
	Jwt_Token_Expired_Periode   = 3600
)

const (
	Endpoint_Auth_Register = `/api.gogrpc.v1.auth.AuthService/Register`
	Endpoint_Auth_Login    = `/api.gogrpc.v1.auth.AuthService/Login`
)
