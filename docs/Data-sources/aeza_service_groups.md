# **Data Sources**
### **aeza_service_groups**
Получает [список групп услуг Aeza](../API/legacy/groups.md). Используется для получения информации о доступных локациях, типах серверов и их характеристиках перед созданием ресурсов.

##### **Схема данных**

| Поле | Тип | Описание |
|------|-----|-----------|
| `id` | number | Уникальный идентификатор группы |
| `group_type` | string | Тип группы: server, location, geography, special, unknown |
| `name` | string | Название группы |
| `type` | string | Тип услуги (vps, dedicated, hicpu и т.д.) |
| `location` | string | Локация (например: "NL-SHARED", "US-DEDICATED") |
| `country_code` | string | Код страны (например: "nl", "de", "fr") |
| `server_type` | string | Тип сервера (shared, dedicated) |
| `is_disabled` | boolean | Флаг отключения группы |
| `description` | string | Описание группы |
| `features` | string | Локализованное описание возможностей |
| `network_speed` | string | Скорость сети |
| `ipv4_count` | number | Количество IPv4 адресов |
| `ipv6_subnet` | string | IPv6 подсеть |
| `cpu_model` | string | Модель процессора |
| `cpu_frequency` | string | Частота процессора |
| `service_handler` | string | Обработчик сервиса |


### **Пример использования**
```hcl
# Получение всех групп услуг
data "aeza_service_groups" "all" {}

# Фильтрация по типу услуги
data "aeza_service_groups" "vps" {
  service_type = "vps"
}

# Вывод общей информации
output "groups_count" {
  value = length(data.aeza_service_groups.all.groups)
}

# Группировка по типам групп
output "groups_by_type" {
  value = {
    for group in data.aeza_service_groups.all.groups :
    group.group_type => group.name...
  }
}

# Поиск серверных групп
output "server_groups" {
  value = [for group in data.aeza_service_groups.all.groups : 
    group if group.group_type == "server"
  ]
}

# Локации с выделенными серверами
output "dedicated_locations" {
  value = [for group in data.aeza_service_groups.all.groups : 
    group if group.server_type == "dedicated" && group.group_type == "location"
  ]
}

# Создание карты для удобного доступа
locals {
  groups_map = {
    for group in data.aeza_service_groups.all.groups :
    tostring(group.id) => group.name
  }
  
  # Географические группы
  geography_groups = [
    for group in data.aeza_service_groups.all.groups :
    group if group.group_type == "geography"
  ]
  
  # Специальные сервисы
  special_services = [
    for group in data.aeza_service_groups.all.groups :
    group if group.group_type == "special"
  ]
}
```

### **Пример вывода**
```bash
+ groups_count = 25
+ groups_by_type = {
    + location = [
        + "NL-SHARED",
        + "DE-DEDICATED",
        + "US-DEDICATED"
      ]
    + server   = [
        + "AMD Ryzen Cluster",
        + "Intel Xeon Pool"
      ]
    + geography = [
        + "Europe",
        + "North America"
      ]
    + special = [
        + "Backup Services",
        + "CDN Network"
      ]
    }
+ server_groups = [
    + {
        + country_code    = "nl"
        + cpu_frequency   = "3.4 GHz"
        + cpu_model       = "AMD EPYC"
        + description     = "High performance AMD cluster"
        + features        = "NVMe storage, DDoS protection"
        + group_type      = "server"
        + id              = 2
        + ipv4_count      = 1
        + ipv6_subnet     = "/64"
        + is_disabled     = false
        + location        = "NL-AMD"
        + name            = "AMD Ryzen Cluster"
        + network_speed   = "10 Gbps"
        + server_type     = "shared"
        + service_handler = "vm6"
        + type            = "vps"
      }
    ]
+ dedicated_locations = [
    + {
        + country_code    = "de"
        + cpu_frequency   = "3.4 GHz"
        + cpu_model       = "Intel Xeon E5"
        + description     = "German dedicated servers"
        + features        = "DDoS protection, High performance"
        + group_type      = "location"
        + id              = 1
        + ipv4_count      = 1
        + ipv6_subnet     = "/64"
        + is_disabled     = false
        + location        = "DE-DEDICATED"
        + name            = "Germany Dedicated"
        + network_speed   = "1 Gbps"
        + server_type     = "dedicated"
        + service_handler = "manual"
        + type            = "dedicated"
      }
    ]
```
### **Примечания**
##### **Типы групп:**

- `location` - чистые локации (имеют роль локации, родительские элементы)

- `server` - группы серверов (VPS, Dedicated, ..., HiCPU)

- `geography` - географические регионы (для локации)

- `special` - специальные сервисы (домены, VPN, CDN и т.д.)

- `unknown` - неопределенный тип

##### **Типы серверов:**

- `shared` - виртуальные серверы (VPS)

- `dedicated` - выделенные серверы

##### **Service Handler** определяет систему управления:

- `vm6` - облачные VPS серверы

- `manual` - выделенные серверы

Другие обработчики в зависимости от типа услуги