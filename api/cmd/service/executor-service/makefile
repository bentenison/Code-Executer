# These are the values that will be injected into the BuildInfo struct
VERSION := 1.0.0
GIT_COMMIT := $(shell git rev-parse --short HEAD)
BUILD_TIME := $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')
GO_VERSION := $(shell go version | awk '{print $$3}')
TARGET_OS := $(shell go env GOOS)
TARGET_ARCH := $(shell go env GOARCH)
CONTAINER_ID := $(shell hostname)
CONTAINER_NAME := $(shell cat /proc/1/cpuset | sed 's/\///g')

# Binary output name
BINARY_NAME := myapp

# LDFLAGS to inject the values into the version.BuildInfo struct
LDFLAGS := -X 'version.BuildInfo.Version=$(VERSION)' \
           -X 'version.BuildInfo.GitCommit=$(GIT_COMMIT)' \
           -X 'version.BuildInfo.BuildTime=$(BUILD_TIME)' \
           -X 'version.BuildInfo.GoVersion=$(GO_VERSION)' \
           -X 'version.BuildInfo.TargetOS=$(TARGET_OS)' \
           -X 'version.BuildInfo.TargetArch=$(TARGET_ARCH)' \
           -X 'version.BuildInfo.ContainerId=$(CONTAINER_ID)' \
           -X 'version.BuildInfo.ContainerName=$(CONTAINER_NAME)'

# The default rule to build the binary
all: build

# Rule to build the Go binary
build:
	go build -ldflags "$(LDFLAGS)" -o $(BINARY_NAME) .

# Rule to clean up the binary
clean:
	rm -f $(BINARY_NAME)

.PHONY: all build clean
