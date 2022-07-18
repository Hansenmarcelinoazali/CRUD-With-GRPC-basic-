#!/bin/bash

set -ex

echo "Starting getting dependency"

# shellcheck disable=SC2068
go install $@ \
  github.com/golang/protobuf/protoc-gen-go \
  github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
  github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

brew install coreutils
brew install protobuf
brew install jq

echo "Successfully getting dependency"
