# Базовая архитектура проекта

```text
terraform-provider-aeza/
├── cmd/
│   └── terraform-provider-aeza/
│       └── main.go
├── internal/
│   ├── provider/
│   │   ├── provider.go
│   │   └── config.go
│   ├── resources/
│   │   └── service.go
│   ├── data-sources/
│   │   ├── services.go
│   │   ├── products.go
│   │   └── service_types.go
│   └── client/
|       ├── request.go
│       ├── client.go
│       ├── models.go
│       └── errors.go
├── pkg/
│   └── utils/
|       ├── errors.go
│       ├── validation.go
│       └── converters.go
├── examples/
│   ├── basic/
│   │   └── main.tf
│   └── advanced/
│       └── main.tf
├── docs/
|   ├──variables.tf
│   └── index.md
├── go.mod
└── go.sum
│
└── Makefile
```


A Terraform provider for managing Aeza cloud resources.

## Requirements

- Terraform 1.0+
- Go 1.21+ (if building from source)

## Installation

### Using Terraform Registry

```hcl
terraform {
  required_providers {
    aeza = {
      source = "yourname/aeza"
      version = "1.0.0"
    }
  }
}
```

### Building from source
```bash
git clone https://github.com/yourname/terraform-provider-aeza
cd terraform-provider-aeza
go build -o terraform-provider-aeza
```
Provider Configuration
```hcl
provider "aeza" {
  api_key = "your_api_key_here"
  # base_url = "https://my.aeza.net/api" # optional
}
```

### Data Sources

**aeza_products**

Retrieve information about available products.

```hcl
data "aeza_products" "vps" {
  type = "vps"
}

output "vps_products" {
  value = data.aeza_products.vps.products
}
```

**aeza_services**

Retrieve list of existing services.

```hcl
data "aeza_services" "all" {}

output "my_services" {
  value = data.aeza_services.all.services
}
```
**aeza_service_types**

Retrieve available service types.

```hcl
data "aeza_service_types" "all" {}

output "service_types" {
  value = data.aeza_service_types.all.types
}
```
###Resources

**aeza_service**

Manage Aeza services (VPS, Hi-CPU servers, etc.).

```hcl
resource "aeza_service" "web_server" {
  product_id = 181
  name       = "web-server"
  type       = "vps"
}
```

### Development

**Building the provider**

```bash
go build -o terraform-provider-aeza
```

**Running tests**

```bash
go test ./...
```

**Generating documentation**
```bash
go generate
```

## 11. 🏃‍♂️ Сборка и установка

**Makefile:**
```makefile
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
```
