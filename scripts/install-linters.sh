#!/usr/bin/env bash

set -euo pipefail

curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.24.0

go get github.com/fzipp/gocyclo
go get github.com/jingyugao/rowserrcheck
go get github.com/uudashr/gocognit/cmd/gocognit
go get github.com/client9/misspell/cmd/misspell
