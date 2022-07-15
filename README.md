# GO GRPC BASIC

### Requirements ( Do It Letter )
-
-

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
- `go get github.com/soheilhy/cmux`, is for ?
- create `router.go` in router dir and implement IgnoreErr, this is for ignore error so can be safely ignore
- `go get golang.org/x/sync/errgroup`, is for ?
- implement `main.go` to create grpc server from grpc.go with errgroup handler
- `go run .`, run server grpc
- `go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest`
- test `grpcurl -plaintext localhost:5566 list`, to show list name of services
- test `grpcurl -plaintext localhost:5566 list api.gogrpc.v1.health.HealthService`, to show list name of service methods
- test `grpcurl -plaintext localhost:5566 api.gogrpc.v1.health.HealthService.Status`, to test method call in grpc 
- result 
```
{
    "success": true,
    "code": "0000",
    "desc": "SUCCESS"
}
``` 
- or test via postman , new -> grpc request -> Enter Server Url : localhost:5566 -> Import proto file / Select Method : Status -> Click Invoke

### Implment gRPC-Gateway
ref https://github.com/grpc-ecosystem/grpc-gateway
- implement `import "google/api/annotations.proto";` in proto file 
- changes line below in all service methods for rest compile to rest  
```
service HealthService {
    rpc Status(google.protobuf.Empty) returns (Response) {
        + option (google.api.http) = {
        +    get: "/api/gogrpc/v1/health/status"
        + };
    }
}
```
- re - compile / re - generate proto with `compile-proto.sh` in proto dir
- `go mod tidy`
- `go get "github.com/gorilla/handlers"`
- create `http.go` in router dir and implement NewHTTPServer and register health api service 
- register httpserver to `main.go`
```
g.Go(func() error { return router.NewHTTPServer(configs, logger) })
```
- `go run .` run grpc and http server
- test `curl localhost:8080/api/v1/health/status`, reponse
```
{"success":true,"code":"0000","desc":"SUCCESS"}
```