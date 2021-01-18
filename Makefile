GO=go
GOTEST=$(GO) test
GOCOVER=$(GO) tool cover

all: build

clean:
	rm -rf ./build/*

dependencies:
	$(GO) mod download

install_linters:
	$(CURDIR)/scripts/install-linters.sh

lint:
	$(CURDIR)/scripts/lint.sh

test:
	$(GOTEST) -test.failfast -race ./...

ci: dependencies test

cover:
	$(GOTEST) -v -coverprofile=/tmp/coverage.out ./...
	$(GOCOVER) -func=/tmp/coverage.out
	$(GOCOVER) -html=/tmp/coverage.out

build: clean dependencies build-cmd

build-cmd:
	$(CURDIR)/scripts/build.sh
