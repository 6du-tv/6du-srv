#!/usr/bin/env bash

_dirname=$(cd "$(dirname "$0")"; pwd)

cd $_dirname

export GOPROXY=https://goproxy.io
export GO111MODULE=on

COLOR='\033[0;32m'
NOCOLOR='\033[0m'
fswatch -o  .| xargs -n1 sh -c "echo '\n$COLOR----$NOCOLOR\n' ; go run main.go"
#go run main.go
