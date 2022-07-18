#!/bin/bash
WORKDIR=github.com/ajikamaludin/go-grpc_basic

set -ve

for x in $(ls v1)
do
      protoc \
      -I. \
      -I/usr/local/include \
      -I${GOPATH}/src \
      -I${GOPATH}/src/$WORKDIR/proto \
      -I${GOPATH}/src/$WORKDIR/proto/lib \
      --go_out=plugins=grpc:$GOPATH/src \
      --grpc-gateway_out=logtostderr=true:$GOPATH/src \
      v1/$x/$x.proto
done

