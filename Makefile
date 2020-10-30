
CGO_ENABLED ?= 0
GOOS ?= darwin
GOARCH ?= amd64

BINARY_NAME = nukibridgectl
BINARY_NAME_LINUX = $(BINARY_NAME)-linux-$(GOARCH)
BINARY_NAME_WINDOWS = $(BINARY_NAME)-windows-$(GOARCH)
BINARY_NAME_DARWIN = $(BINARY_NAME)-darwin-$(GOARCH)

BUILD_FLAGS = -v
BUILD_FLAGS_STATIC = $(BUILD_FLAGS) -a -ldflags '-extldflags "-static"'

.PHONY: build-static-linux
build-static-linux:
	@GOOS=linux go build $(BUILD_FLAGS_STATIC) -o "$(BINARY_NAME_LINUX)-static" cmd/nukibridgectl/main.go

.PHONY: build-static-windows
build-static-windows:
	@GOOS=windows go build $(BUILD_FLAGS_STATIC) -o "$(BINARY_NAME_WINDOWS)-static" cmd/nukibridgectl/main.go

.PHONY: build-static-darwin
build-static-darwin:
	@GOOS=darwin go build $(BUILD_FLAGS_STATIC) -o "$(BINARY_NAME_DARWIN)-static" cmd/nukibridgectl/main.go

.PHONY: build-static
build-static: build-static-darwin build-static-linux build-static-windows

.PHONY: build
build:
	@go build $(BUILD_FLAGS) -o "nukibridgectl-$(GOOS)-$(GOARCH)" cmd/nukibridgectl/main.go

.PHONY: test
test:
	@go test -v ./...