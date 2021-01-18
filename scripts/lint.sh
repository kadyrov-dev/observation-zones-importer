#!/usr/bin/env bash

set -euo pipefail

$(go env GOPATH)/bin/golangci-lint run --no-config --verbose --print-resources-usage --disable-all \
  --enable=govet --enable=errcheck --enable=staticcheck --enable=unused --enable=gosimple --enable=structcheck \
  --enable=varcheck --enable=ineffassign --enable=deadcode --enable=bodyclose --enable=golint --enable=stylecheck \
  --enable=gosec --enable=interfacer --enable=unconvert --enable=dupl --enable=goconst --enable=gocognit \
  --enable=rowserrcheck --enable=gofmt --enable=goimports --enable=maligned --enable=depguard --enable=misspell \
  --enable=lll --enable=unparam --enable=dogsled --enable=nakedret --enable=prealloc --enable=scopelint \
  --enable=gocritic --enable=gochecknoinits --enable=gochecknoglobals --enable=godox --enable=funlen --enable=wsl \
  --enable=goprintffuncname
