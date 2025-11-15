terraform {
  required_providers {
    aeza = {
      source = "github.com/yourname/aeza"
    }
  }
}

provider "aeza" {
  api_key = var.aeza_api_key
  # base_url = "https://my.aeza.net/api" # optional
}

# Получить информацию о доступных продуктах
data "aeza_products" "europe_vps" {
  type     = "vps"
  location = "NL-SHARED"
}

# Получить типы серверов
data "aeza_service_types" "all" {}

# Создать VPS
resource "aeza_service" "my_server" {
  product_id = data.aeza_products.europe_vps.products[0].id
  name       = "my-terraform-vps"
  type       = "vps"
}

# Вывести информацию о созданном сервере
output "server_info" {
  value = {
    id      = aeza_service.my_server.id
    name    = aeza_service.my_server.name
    status  = aeza_service.my_server.status
    product = aeza_service.my_server.product_name
    location = aeza_service.my_server.location_code
  }
}

# Вывести доступные типы серверов
output "available_types" {
  value = data.aeza_service_types.all.types
}