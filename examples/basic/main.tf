# examples/basic/main.tf
terraform {
  required_providers {
    aeza = {
      source  = "vbobroff-app/aeza"
      version = "0.1.0"
    }
  }
}

provider "aeza" {
  api_key = var.aeza_api_key
  base_url = var.aeza_base_url
}

# Data source для типов услуг
data "aeza_service_types" "all" {}