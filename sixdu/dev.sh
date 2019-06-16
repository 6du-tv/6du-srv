#!/usr/bin/env bash

_dirname=$(cd "$(dirname "$0")"; pwd)

cd $_dirname

export GOPROXY=https://goproxy.io

go run main.go
