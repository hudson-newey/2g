#!/usr/bin/env bash
set -euo pipefail

declare BUILD_DIR="./build"

# build the user facing CLI
go build -o "${BUILD_DIR}/2g" ./src/main.go

# build the daemon
go build -o "${BUILD_DIR}/2g-daemon" ./daemon/main.go
