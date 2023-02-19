#!/usr/bin/env make

.PHONY: help
help: ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: test
test: ## Run all tests
	go test -cover -race -count=1 ./...

.PHONY: install
install:
	go install h.go

.PHONY: fmt
fmt: ## Formats project
	go fmt ./...
	golangci-lint run --fix

.PHONY: lint
lint: ## Run the linter
	golangci-lint run


