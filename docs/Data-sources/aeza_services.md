# **Data Sources**
### **aeza_services**

Получает список всех услуг в аккаунте Aeza. Используется для получения информации о существующих сервисах (VPS, DNS, VPN и т.д.), их статусах, ценах и сроках действия.

### **Схема данных**

| Поле | Тип | Описание |
|------|-----|-----------|
| `id` | number | Уникальный числовой идентификатор услуги |
| `type_slug` | string | Тип услуги (vps, dns, vpn и т.д.) |
| `name` | string | Название услуги |
| `ip` | string | Основной IP-адрес услуги |
| `price_raw` | number | Цена в числовом формате для вычислений (в центах/копейках) |
| `price_display` | string | Отформатированная цена с валютой для отображения |
| `payment_term` | string | Период оплаты (hour, month, year и т.д.) |
| `auto_prolong` | boolean | Флаг автоматического продления услуги |
| `created_at` | string | Дата создания услуги в отформатированном виде |
| `expires_at` | string | Дата истечения срока действия услуги в отформатированном виде |
| `product_name` | string | Название продукта |
| `location_code` | string | Код локации, где развернута услуга (de, nl и т.д.) |
| `status` | string | Текущий статус услуги (activation_wait, active и т.д.) |

Все поля доступны только для чтения. Использует leagacy [API /api/services](../API/legacy/services.md) ендпоинт.

### **Пример использования**
```hcl
# Получение списка всех услуг
data "aeza_services" "all" {}

# Вывод всех услуг
output "all_services" {
  value = data.aeza_services.all.services
}
  
# Фильтрация только активных VPS услуг
locals {
  active_vps_services = [
    for service in data.aeza_services.all.services :
    service if service.type_slug == "vps" && service.status == "active"
  ]
  # Создание карты ID->Name для удобства использования
  services_map = {
    for service in data.aeza_services.all.services :
    tostring(service.id) => service.name
  }
}  

# Использование данных для других ресурсов
resource "some_other_resource" "example" {
  service_id = local.active_vps_services[0].id
  name       = "Backup for ${local.active_vps_services[0].name}"
}
```

### **Пример вывода**
```bash
+ all_services = [
  + {
      + auto_prolong   = true
      + created_at     = "20.11.2025 13:12"
      + expires_at     = "20.12.2025 13:12"
      + id             = 12345
      + ip             = "192.168.1.100"
      + location_code  = "de"
      + name           = "my-vps-server"
      + payment_term   = "month"
      + price_display  = "5.00 €"
      + price_raw      = 500
      + product_name   = "SWE-PROMO"
      + status         = "active"
      + type_slug      = "vps"
    },
  + {
      + auto_prolong   = false
      + created_at     = "15.11.2025 10:30"
      + expires_at     = "15.12.2025 10:30"
      + id             = 12346
      + ip             = "10.0.0.1"
      + location_code  = "nl"
      + name           = "my-dns-service"
      + payment_term   = "year"
      + price_display  = "12.00 €"
      + price_raw      = 1200
      + product_name   = "DNS-PRO"
      + status         = "active"
      + type_slug      = "dns"
    },
  + {
      + auto_prolong   = true
      + created_at     = "18.11.2025 14:45"
      + expires_at     = "18.11.2025 15:45"
      + id             = 12347
      + ip             = "172.16.1.100"
      + location_code  = "us"
      + name           = "test-vpn"
      + payment_term   = "hour"
      + price_display  = "0.50 €"
      + price_raw      = 50
      + product_name   = "VPN-TEST"
      + status         = "activation_wait"
      + type_slug      = "vpn"
    },
]

+ services_map = {
    + "12345" = "my-vps-server"
    + "12346" = "my-dns-service"
    + "12347" = "test-vpn"
  }
```