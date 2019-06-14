#!/usr/bin/env bash

_dirname=$(cd "$(dirname "$0")"; pwd)

cd $_dirname

CPU="armv7a-linux-androideabi28-clang"
GOARCH=arm

CC="$ANDROID_NDK_ROOT/toolchains/llvm/prebuilt/darwin-x86_64/bin/$CPU -pie -fPIE -fPIC" GOOS=android GOARCH=$GOARCH CGO_ENABLED=1 go build main.go
