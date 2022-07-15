# GO GRPC BASIC

## Requirements
check other branch to step by step create project 
### Validate Go Installation
```bash
$ go version
go version go1.18.3 linux/amd64
```

### Validate GOPATH
```bash
echo $GOPATH
```

### Start Project in Go
```bash
cd $GOPATH/src
mkdir -p github.com/ajikamaludin/go-grpc_basic 
cd github.com/ajikamaludin/go-grpc_basic 
go mod init github.com/ajikamaludin/go-grpc_basic
go mod tidy
```

### Install Protoc
Linux : Download zip file, extract zip file to ~/.local/, add PATH ~/.local/bin
please read documentation from this link for more detail information
```
https://grpc.io/docs/protoc-installation/
```

### Validate Protoc Installation
```bash
$ protoc --version
libprotoc 3.21.2
```

### Install Protoc Dependecy for Golang
```bash
# protoc-gen-go
go install github.com/golang/protobuf/protoc-gen-go@latest
# proto (optional, execute in project dir)
go get google.golang.org/protobuf/proto@latest
# protoc-gen-grpc-gateway
go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@latest
# protoc-gen-swagger
go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@latest
```

### Add PATH
add to your ~/.bashrc or ~/.zshrc file
```bash
export PATH="$PATH:$GOPATH/bin"
```
### Validate Protoc Dependency Golang Installation
```bash
~ ❯ ls $GOPATH/bin                                             
dlv         gopls    protoc-gen-go       protoc-gen-grpc-gateway  revel
go-outline  grpcurl  protoc-gen-go-grpc  protoc-gen-swagger
```

## Start Project
### Setup Project 
- create `proto` dir 
- create versioning dir and service dir `health`
- create proto file `health.proto`
- compile / generate proto with `compile-proto.sh` in proto dir

### Setup config file
- create `config.yaml`
- create `pkg` dir , create versioning dir and create `configs` dir
- `go get gopkg.in/yaml.v2` , is a lib for parsing yaml config file to struct
- create `config.go` file, implement New and other func
- create `configs` dir on root project , create `configs.go`, this is file that bundle or wrap any services or packages
- `go get github.com/sirupsen/logrus`, is a lib to show log on run
- implement New to `configs.go` file
- create `main.go`, implement to call config and log environtment read is ready
- test `go run .`

### Implement Server GRPC 
- create `utils/constants` dir in `pkg/v1`, to create global constants, implement EnvProduction, Successcode, SuccessDesc
- implement grpc service create `api/v1/health` service, create `health.go` as server service
- create `api/v1/health/status.go` as method implment from protobuf / pb file 
- create `router` dir in root project
- create `grpc.go` in router dir and implement NewGRPCServer and register health api service 
- `go get github.com/soheilhy/cmux`, TODO: what is for ?
- create `router.go` in router dir and implement IgnoreErr, this is for ignore error so can be safely ignore
- `go get golang.org/x/sync/errgroup`, TODO: what is for ?
- implement `main.go` to create grpc server from grpc.go with errgroup handler
- `go run .`, run server grpc
- `go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest`, this tool is 
- test `grpcurl -plaintext localhost:5566 list`, to show list name of services
- test `grpcurl -plaintext localhost:5566 list api.gogrpc.v1.health.HealthService`, to show list name of service methods
- test `grpcurl -plaintext localhost:5566 api.gogrpc.v1.health.HealthService.Status`, to test method call in grpc 
- result 
```json
{
    "success": true,
    "code": "0000",
    "desc": "SUCCESS"
}
``` 
- or test via postman , new -> grpc request -> Enter Server Url : localhost:5566 -> Import proto file / Select Method : Status -> Click Invoke

### Implement gRPC-Gateway
ref https://github.com/grpc-ecosystem/grpc-gateway
- implement `import "google/api/annotations.proto";` in proto file 
- changes line below in all service methods for rest compile to rest  
```proto
service HealthService {
    rpc Status(google.protobuf.Empty) returns (Response) {
        option (google.api.http) = {
           get: "/api/gogrpc/v1/health/status"
        };
    }
}
```
- re - compile / re - generate proto with `compile-proto.sh` in proto dir
- `go mod tidy`
- `go get "github.com/gorilla/handlers"` TODO: what is for ?
- create `http.go` in router dir and implement NewHTTPServer and register health api service 
- register httpserver to `main.go`
```go
g.Go(func() error { return router.NewHTTPServer(configs, logger) })
```
- `go run .` run grpc and http server
- test `curl localhost:8080/api/v1/health/status`, reponse
```json
{
    "success":true,
    "code":"0000",
    "desc":"SUCCESS"
}
```
- GENERATE API DOCS: 
- `mkdir swagger`
- `cd proto`
- `./gen-apidoc.sh`, will be generated in `swagger/docs.json`
- register apidoc to http server in `http.go` implement 
```go
/////////

if configs.Config.Env != constants.EnvProduction {
	CreateSwagger(mux)
}

/////////

func CreateSwagger(gwmux *http.ServeMux) {
	gwmux.HandleFunc("/api/health/docs.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "swagger/docs.json")
	})
}
```

### Implement DB_Connection (With PostgreSQL)
- Execute `example.sql` in db, you know how to do it
- add `config.yaml` with line below 
```yaml
postgres:
  host: 127.0.0.1
  port: 5432
  dbname: test
  username: postgres
  password: mysecretpassword
```
- changes `pkg/v1/config/config.go` to add new config line in config struct and and validateConfigData rule
- `go get github.com/lib/pq`, this is a lib for standart go lib `database/sql` to connect postgresql, usage in call
```go
import (
    _ "github.com/lib/pq"
)
// ref : https://github.com/lib/pq
```
- create `pkg/v1/postgres` dir, create `postgres.go` file in there implement string conn and test connection
- create `pkg/v1/utils/converter` dir, create `converter.go` file in there, to convert camelcase to snake_case
- create `pkg/v1/postgres/custom.main.go` to implement all query to database table custom
- changes `configs/configs.go` to bundle pg connection
- how to use custom.main.go call function from api calldb, check `api/v1/health/calldb.go`

### Example of call Other Rest API
- add new environtment to `config.yaml` add cert path
- changes `pkg/v1/config/config.go` to validate and add  new environtment 
- changes `pkg/v1/utils/constants` append method POST/GET/PUT/DELETE constant, JsonType Header, Hostname endpoint
- create `pkg/v1/cert` dir, create `cert.go`, this file is for handle Insecure ssl connection (self gen cert)
- create `pkg/v1/requestapi` dir, create `requestapi.go`, implement http call 
- create `services/v1/jsonplaceholder` dir, create service name `jsonplaceholder.go`, implement any of endpoint http call from requestapi.go
- create new method from service health and implement call jsonplaceholder service from there, check `api/v1/health/callapi.go`

### HTTP Custom Error Reponse 
- create `pgk/v1/utils/errors/errors.go`, implement error handler for custom http
- register runtime Http error in `http.go`
```go
runtime.HTTPError = errors.CustomHTTPError
```

### File Logger
- goal : create log file for every action like , request / response from / to client, request / response from / to other api, 
- `go get github.com/lestrrat/go-file-rotatelogs`, this lib for create rotate log file automated
- create `pkg/v1/utils/logger/logger.go`, implement logger function to write log to file
- implement logger for `pkg/v1/requestapi/requestapi.go` to log request and response api from other service
- implement logger for `pkg/v1/utils/errors/errors.go` to log any of error from our service
- implement logger for `router/http.go` to log response grpc rest api service 
- `go run .`
- test call method callapi will generate hif log, custom-error generate error log, call rest will generate rest log 