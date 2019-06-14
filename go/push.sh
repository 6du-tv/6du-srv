#!/usr/bin/env bash

_dirname=$(cd "$(dirname "$0")"; pwd)

cd $_dirname

adb push main /data/local/tmp
adb shell "cd /data/local/tmp;chmod +x main"
