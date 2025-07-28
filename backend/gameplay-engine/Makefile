APP_NAME = gameplay-engine
GO      ?= go

.PHONY: lint test build run

lint:
@golangci-lint run ./...

test:
@$(GO) test ./... -race -coverprofile=coverage.out

build:
@$(GO) build -o bin/$(APP_NAME) ./cmd/engine

run:
@$(GO) run ./cmd/engine
