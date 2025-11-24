# **Data Sources**

### **aeza_service_types**

Получает список [доступных типов услуг в Aeza](../API/service-types.md). Используется для определения поддерживаемых категорий сервисов (VPS, выделенные серверы и т.д.) перед созданием ресурсов.

##### **Схема данных**

| Поле | Тип | Обязательный | Описание |
|------|-----|--------------|-----------|
| `slug` | `string` | Только чтение | Уникальный идентификатор типа услуги (vps, dedicated, storage и т.д.) |
| `name` | `string` | Только чтение | Человеко-читаемое название типа услуги |
| `order` | `number` | Только чтение | Порядок отображения в интерфейсе |

**Все readonly поля доступны только для чтения.**

### Пример использования

```hcl
data "aeza_service_types" "all" {}

output "available_service_types" {
  value = data.aeza_service_types.all.service_types
}

output "service_types_map" {
  description = "Service types as slug->name mapping"
  value = {
    for st in data.aeza_service_types.all.service_types :
    st.slug => st.name
  }
}
```
### Пример вывода
```bash
+ service_types_map    = {
      + dedicated = "Выделенный сервер"
      + hicpu     = "Hi-CPU сервер"
      + soft      = "Лицензия ispmanager"
      + storage   = "Storage VPS"
      + vpn       = "VPN"
      + vps       = "Виртуальный сервер"
    }
```