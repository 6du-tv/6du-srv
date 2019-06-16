#!/usr/bin/env bash

_dirname=$(cd "$(dirname "$0")"; pwd)

cd $_dirname

export GOPROXY=https://goproxy.io
export GO111MODULE=on

go run main.go
