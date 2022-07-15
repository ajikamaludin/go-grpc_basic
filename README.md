# GO GRPC BASIC

### Requirements ( Do It Letter )

### Setup Project 
- create `proto` dir 
- create versioning dir and service dir `health`
- create proto file `health.proto`
- compile / generate proto with `compile-proto.sh` in proto dir

- create `config.yaml`
- create `pkg` dir , create versioning dir and create `configs` dir
- `go get gopkg.in/yaml.v2` , is a lib for parsing yaml config file to struct
- create `config.go` file, implement New and other func
- create `configs` dir on root project , create `configs.go`, this is file that bundle or wrap any services or packages
- `go get github.com/sirupsen/logrus`, is a lib to show log on run
- implement New to `configs.go` file
- create `main.go`, implement to call config and log environtment read is ready