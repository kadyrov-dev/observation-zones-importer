#!/usr/bin/env bash

set -eou pipefail

BIN="./build/observation-zones-importer"
MODULE="github.com/kadyrov-dev/observation-zones-importer"

CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o $BIN-mac $MODULE
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $BIN-linux $MODULE
