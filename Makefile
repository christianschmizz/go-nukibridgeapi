
CGO_ENABLED ?= 0
GOOS ?= darwin
GOARCH ?= amd64

.PHONY: build
build:
	@echo Building ...
	@go build -v -a -ldflags '-extldflags "-static"' -o nukibridgectl cmd/nukibridgectl/main.go
