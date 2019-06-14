#!/usr/bin/env bash

_dirname=$(cd "$(dirname "$0")"; pwd)

cd $_dirname

CC="$ANDROID_NDK_ROOT/toolchains/llvm/prebuilt/darwin-x86_64/bin/aarch64-linux-android28-clang -pie -fPIE -fPIC" GOOS=android GOARCH=arm64 CGO_ENABLED=1 go build main.go
