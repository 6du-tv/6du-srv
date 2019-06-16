#!/usr/bin/env bash

_dirname=$(cd "$(dirname "$0")"; pwd)

cd $_dirname

wget -O- https://raw.githubusercontent.com/ngosang/trackerslist/master/trackers_all_udp.txt | sort | sed '/^[[:space:]]*$/d' > udp.txt
