
ifeq ("$(origin V)", "command line")
  BUILD_VERBOSE = $(V)
endif
ifndef BUILD_VERBOSE
  BUILD_VERBOSE = 0
endif

ifeq ($(BUILD_VERBOSE),1)
  Q =
else
  Q = @
endif
ECHO := @echo

CGO_ENABLED ?= 0
GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)

BINARY_NAME := nukibridgectl
BINARY_NAME_LINUX = $(BINARY_NAME)-linux-$(GOARCH)
BINARY_NAME_WINDOWS = $(BINARY_NAME)-windows-$(GOARCH)
BINARY_NAME_DARWIN = $(BINARY_NAME)-darwin-$(GOARCH)

COMMAND = ./cmd/$(BINARY_NAME)/

DATE_FMT = +%Y-%m-%d
BUILD_DATE ?= $(shell date "$(DATE_FMT)")

VERSION ?= $(shell git describe --tags 2>/dev/null || git rev-parse --short HEAD)

GO_LDFLAGS := -X github.com/christianschmizz/go-nukibridgeapi/internal/build.Version=$(VERSION) $(GO_LDFLAGS)
GO_LDFLAGS := -X github.com/christianschmizz/go-nukibridgeapi/internal/build.Date=$(BUILD_DATE) $(GO_LDFLAGS)
GO_LDFLAGS_STATIC := '-extldflags "-static"' $(GO_LDFLAGS)

BUILD_FLAGS = -a -trimpath -ldflags "$(GO_LDFLAGS)"
BUILD_FLAGS_STATIC = -a -trimpath -ldflags "$(GO_LDFLAGS_STATIC)"

ifeq ($(BUILD_VERBOSE),1)
	BUILD_FLAGS += -v
	BUILD_FLAGS_STATIC += -v
endif

.PHONY: build-static-linux
build-static-linux:
	$(ECHO) building $@
	$(Q)GOOS=linux go build $(BUILD_FLAGS_STATIC) -o "$(BINARY_NAME_LINUX)-static" $(COMMAND)

.PHONY: build-static-windows
build-static-windows:
	$(Q)GOOS=windows go build $(BUILD_FLAGS_STATIC) -o "$(BINARY_NAME_WINDOWS)-static" $(COMMAND)

.PHONY: build-static-darwin
build-static-darwin:
	$(Q)GOOS=darwin go build $(BUILD_FLAGS_STATIC) -o "$(BINARY_NAME_DARWIN)-static" $(COMMAND)

.PHONY: build-cross-static
build-cross-static: build-static-darwin build-static-linux build-static-windows

.PHONY: build-static
build-static:
	$(Q)go build $(BUILD_FLAGS_STATIC) -o "$(BINARY_NAME_DARWIN)-static" $(COMMAND)

.PHONY: build
build: | checkdeps lint test
	@echo "Building $(BINARY_NAME)"
	$(Q)go build $(BUILD_FLAGS) -o "$(BINARY_NAME)-$(GOOS)-$(GOARCH)" $(COMMAND)

.PHONY: install
install:
	$(ECHO) "Installing $(BINARY_NAME) binary to '$(GOPATH)/bin/$(BINARY_NAME)'"
	$(Q)go install $(BUILD_FLAGS) $(COMMAND)
	$(ECHO) "Installation successful. To learn more, try \"$(BINARY_NAME) --help\"."

.PHONY: test
test:
	$(Q)go test -v ./...

.PHONY: integration-test
integration-test:
	$(Q)go test -v ./... -failfast -tags=integration -args -host $(BRIDGE_HOST) -token $(BRIDGE_TOKEN)

.PHONY:
checkdeps:
	@echo "Checking dependencies"
ifeq ($(shell which golangci-lint),)
	$(ECHO) Installing golangci-lint
ifeq ($(shell which brew),)
	$(Q)brew install golangci/tap/golangci-lint
else
	$(Q)mkdir -p ${GOPATH}/bin
	$(Q)curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.31.0
endif
endif

.PHONY: lint
lint:
# see https://github.com/golangci/golangci-lint/issues/1040
	$(ECHO) "Running $@ check"
	@GO111MODULE=on ${GOPATH}/bin/golangci-lint cache clean
	@GO111MODULE=on ${GOPATH}/bin/golangci-lint run --timeout=5m --config ./.golangci.yml

.PHONY: version
version:
	$(ECHO) $(VERSION)
