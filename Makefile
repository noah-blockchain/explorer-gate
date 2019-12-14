APP ?= gate
VERSION ?= $(strip $(shell cat VERSION))
GOOS ?= linux
SRC = ./

COMMIT = $(shell git rev-parse --short HEAD)
BRANCH = $(strip $(shell git rev-parse --abbrev-ref HEAD))
CHANGES = $(shell git rev-list --count ${COMMIT})
BUILDED ?= $(shell date -u '+%Y-%m-%dT%H:%M:%S')
BUILD_FLAGS = "-X main.Version=$(VERSION) -X main.GitCommit=$(COMMIT) -X main.BuildedDate=$(BUILDED)"
BUILD_TAGS?=noah-gate
DOCKER_TAG = latest
SERVER ?= gate.noah.network
PACKAGES=$(shell go list ./... | grep -v '/vendor/')

all: test build

#Run this from CI
create_vendor:
	@rm -rf vendor
	@echo "--> Running go mod vendor"
	@go mod vendor

### Build ###################
build: clean
	GOOS=${GOOS} go build -ldflags $(BUILD_FLAGS) -o ./builds/$(APP)

install:
	GOOS=${GOOS} go install -ldflags $(BUILD_FLAGS)

clean:
	@rm -f $(BINARY)

### Test ####################
test:
	@echo "--> Running tests"
	go test -v ${SRC}

fmt:
	@go fmt ./...

.PHONY: create_vendor build clean fmt test
