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
# неизменные параметры
  product_id    = 182
  os            = "ubuntu_2404"
  payment_term  = "hour"
# изменяемые параметры
  name          = "test-updated-vps"                           
  auto_prolong  = false
}

resource "aeza_service_actions" "prolong_service" {
  service_id = aeza_service.test_vps.id
  
  prolong {
    term = "hour"
    # method = "balance" (по умолчанию)
    count = 1
  }
}