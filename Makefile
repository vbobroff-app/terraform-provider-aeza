BINARY_NAME=terraform-provider-aeza
VERSION=1.0.0
GOOS?=$(shell go env GOOS)
GOARCH?=$(shell go env GOARCH)

.PHONY: build clean test install

build:
	@echo "Building $(BINARY_NAME)..."
	@go build -o $(BINARY_NAME) ./cmd/terraform-provider-aeza

build-all:
	@echo "Building for all platforms..."
	@GOOS=linux GOARCH=amd64 go build -o bin/linux/amd64/$(BINARY_NAME) ./cmd/terraform-provider-aeza
	@GOOS=darwin GOARCH=amd64 go build -o bin/darwin/amd64/$(BINARY_NAME) ./cmd/terraform-provider-aeza
	@GOOS=windows GOARCH=amd64 go build -o bin/windows/amd64/$(BINARY_NAME).exe ./cmd/terraform-provider-aeza

clean:
	@echo "Cleaning..."
	@rm -f $(BINARY_NAME)
	@rm -rf bin/

test:
	@echo "Running tests..."
	@go test -v ./...

test-race:
	@echo "Running tests with race detector..."
	@go test -race -v ./...

install: build
	@echo "Installing provider..."
	@mkdir -p ~/.terraform.d/plugins/local/yourname/aeza/$(VERSION)/$(GOOS)_$(GOARCH)
	@cp $(BINARY_NAME) ~/.terraform.d/plugins/local/yourname/aeza/$(VERSION)/$(GOOS)_$(GOARCH)/$(BINARY_NAME)_v$(VERSION)

fmt:
	@echo "Formatting code..."
	@gofmt -w -s .

lint:
	@echo "Linting code..."
	@golangci-lint run

.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build      - Build the provider"
	@echo "  build-all  - Build for all platforms"
	@echo "  clean      - Clean build artifacts"
	@echo "  test       - Run tests"
	@echo "  install    - Install provider locally"
	@echo "  fmt        - Format code"
	@echo "  lint       - Run linter"