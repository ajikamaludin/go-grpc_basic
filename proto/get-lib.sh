#!/bin/bash

# debug run
# set -ve

echo "Starting getting googleapi lib"

apidir="lib/google/api"
source="https://raw.githubusercontent.com/googleapis/googleapis/master/google/api"

mkdir -p $apidir

for file in annotations.proto http.proto; do
    curl -Ls "${source}/$file" >${apidir}/$file
done

protobufdir="lib/google/protobuf"
source="https://raw.githubusercontent.com/protocolbuffers/protobuf/master/src/google/protobuf"

mkdir -p $protobufdir

for file in any.proto wrappers.proto empty.proto; do
    curl -Ls "${source}/$file" >${protobufdir}/$file
done

echo "Successfully getting googleapi lib"
