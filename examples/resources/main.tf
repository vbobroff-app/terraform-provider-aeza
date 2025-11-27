# examples/resources/main.tf
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

resource "aeza_service" "test_vps" {
  name          = "test-hourly-vps"
  product_id    = 182
  payment_term  = "hour"
  auto_prolong  = false
  os            = "ubuntu_2404"
  
}

