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
