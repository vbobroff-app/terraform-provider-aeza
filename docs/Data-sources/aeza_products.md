# **Data Sources**

### **aeza_products**

Получает список [доступных продуктов Aeza](../API/legacy/products.md). Используется для получения информации о конфигурациях, ценах и типах сервисов перед созданием ресурсов.

##### **Схема данных**

| Поле | Тип | Обязательный | Описание |
|------|-----|--------------|-----------|
| `id` | `number` | Только чтение | Уникальный идентификатор продукта |
| `name` | `string` | Только чтение | Название продукта (например: "MSK-1", "AMD Ryzen 9 5950X – 10 Gbps") |
| `type` | `string` | Только чтение | Тип продукта (vps, dedicated, domain, storage, hicpu и т.д.) |
| `group_id` | `number` | Только чтение | Идентификатор группы продуктов |
| `service_handler` | `string` | Только чтение | Обработчик сервиса (vm6, manual, feru, s3, ispmgr и т.д.) |
| `prices` | `map` | Только чтение | Цены продукта в различных периодах (hour, month, year, half_year, quarter_year) |

Все поля доступны только для чтения.

### **Пример использования**

```hcl
# Получение всех продуктов
data "aeza_products" "all" {}

# Вывод общей информации
output "products_count" {
  value = length(data.aeza_products.all.products)
}

# Группировка продуктов по типам
output "products_by_type" {
  value = {
    for product in data.aeza_products.all.products :
    product.type => product.name...
  }
}

# Фильтрация VPS продуктов
output "vps_products" {
  value = [for product in data.aeza_products.all.products : 
    product if product.type == "vps"
  ]
}

# Анализ ценовой политики
output "products_with_monthly_pricing" {
  value = [for product in data.aeza_products.all.products : 
    product if try(product.prices["month"] != null, false)
  ]
}
```

### **Пример вывода**

```bash
+ products_count = 150
+ products_by_type = {
    + dedicated = [
        + "AMD Ryzen 9 5950X – 10 Gbps",
        + "Intel Core i7-6700",
        + "AMD EPYC 7702P – 100 Gbps"
      ]
    + domain    = [
        + "ru",
        + "com", 
        + "net",
        + "org"
      ]
    + vps       = [
        + "MSK-1",
        + "MSK-2",
        + "SPB-1"
      ]
    }
+ vps_products = [
    + {
        + group_id        = 1
        + id              = 148
        + name            = "MSK-1"
        + prices          = {
            + hour     = 3
            + month    = 707
            + year     = 7466
          }
        + service_handler = "vm6"
        + type            = "vps"
      },
    + {
        + group_id        = 1
        + id              = 149
        + name            = "MSK-2" 
        + prices          = {
            + hour     = 5
            + month    = 1415
            + year     = 14943
          }
        + service_handler = "vm6"
        + type            = "vps"
      }
    ]
```

### **Примечания**
**Типы продуктов** соответствуют типам сервисов из [aeza_service_types](aeza_service_types.md)

**Service Handler** определяет систему управления сервисом:

- `vm6` - облачные серверы

- `manual` - выделенные серверы

- `feru` - домены

- `s3` - объектное хранилище

- `ispmgr` - лицензии ISPManager

Цены указаны в минимальных единицах валюты (копейки, центы и т.д.)