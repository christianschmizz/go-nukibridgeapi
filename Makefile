
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
GOOS ?= darwin
GOARCH ?= amd64

BINARY_NAME = nukibridgectl
BINARY_NAME_LINUX = $(BINARY_NAME)-linux-$(GOARCH)
BINARY_NAME_WINDOWS = $(BINARY_NAME)-windows-$(GOARCH)
BINARY_NAME_DARWIN = $(BINARY_NAME)-darwin-$(GOARCH)

DATE_FMT = +%Y-%m-%d
BUILD_DATE ?= $(shell date "$(DATE_FMT)")

VERSION ?= $(shell git describe --tags 2>/dev/null || git rev-parse --short HEAD)

GO_LDFLAGS := -X github.com/christianschmizz/go-nukibridgeapi/internal/build.Version=$(VERSION) $(GO_LDFLAGS)
GO_LDFLAGS := -X github.com/christianschmizz/go-nukibridgeapi/internal/build.Date=$(BUILD_DATE) $(GO_LDFLAGS)
GO_LDFLAGS_STATIC := '-extldflags "-static"' $(GO_LDFLAGS)

BUILD_FLAGS = -a -trimpath -ldflags "$(GO_LDFLAGS)"
BUILD_FLAGS_STATIC = -a -v -trimpath -ldflags "$(GO_LDFLAGS_STATIC)"

ifeq ($(BUILD_VERBOSE),1)
	BUILD_FLAGS += -v
	BUILD_FLAGS_STATIC += -v
endif

.PHONY: build-static-linux
build-static-linux:
	$(ECHO) building $@
	$(Q)GOOS=linux go build $(BUILD_FLAGS_STATIC) -o "$(BINARY_NAME_LINUX)-static" ./cmd/nukibridgectl/

.PHONY: build-static-windows
build-static-windows:
	$(Q)GOOS=windows go build $(BUILD_FLAGS_STATIC) -o "$(BINARY_NAME_WINDOWS)-static" ./cmd/nukibridgectl/

.PHONY: build-static-darwin
build-static-darwin:
	$(Q)GOOS=darwin go build $(BUILD_FLAGS_STATIC) -o "$(BINARY_NAME_DARWIN)-static" ./cmd/nukibridgectl/

.PHONY: build-static
build-static: build-static-darwin build-static-linux build-static-windows

.PHONY: build
build:
	$(Q)go build $(BUILD_FLAGS) -o "$(BINARY_NAME)-$(GOOS)-$(GOARCH)" ./cmd/nukibridgectl/

.PHONY: test
test:
	$(Q)go test -v ./...