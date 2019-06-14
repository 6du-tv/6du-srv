#!/usr/bin/env bash

_dirname=$(cd "$(dirname "$0")"; pwd)

cd $_dirname

protoc -I proto du.proto --dart_out=grpc:dart/proto --go_out=plugins=grpc:go/proto
