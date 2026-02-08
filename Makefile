.PHONY: help build run test clean fmt lint vet wire

help: ## Display this help message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

wire: ## Generate Wire dependencies
	go run github.com/google/wire/cmd/wire@latest ./cmd/main

build: wire ## Build the application
	go build -o bin/app ./cmd/main

run: build ## Build and run the application
	./bin/app

clean: ## Clean build artifacts
	rm -rf bin/ dist/ coverage.out coverage.html

fmt: ## Format code
	go fmt ./...

vet: ## Run go vet
	go vet ./...

lint: ## Run golangci-lint (requires installation)
	golangci-lint run ./...

deps: ## Download dependencies
	go mod download
	go mod tidy

deps-update: ## Update all dependencies
	go get -u ./...
	go mod tidy
