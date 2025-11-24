# **Data Sources**

### **aeza_os**

Получает список доступных операционных систем в Aeza. Используется для определения поддерживаемых ОС (Linux, Windows и т.д.) перед созданием виртуальных серверов и других ресурсов.

| Поле | Тип | Описание |
|------|-----|-----------|
| `id` | `number` | Уникальный числовой идентификатор операционной системы |
| `name` | `string` | Полное название операционной системы (например, "Ubuntu 22.04") |
| `slug` | `string` | Уникальный строковый идентификатор ОС (например, "ubuntu_2204") |
| `group` | `string` | Группа операционных систем (alma, centos, debian, ubuntu, windows_server и т.д.) |
| `username` | `string` | Имя пользователя по умолчанию для ОС ("root" для Linux, "Administrator" для Windows) |
| `repository` | `string` | Репозиторий ОС (может быть пустым) |
| `enabled` | `boolean` | Флаг доступности ОС для использования |
| `targets` | `map(number)` | Карта целевых платформ и их идентификаторов |
| `order` | `number` | Порядок сортировки для отображения в интерфейсе |

Все поля доступны только для чтения.

### **Пример использования**

```hcl
# Получение списка всех операционных систем
data "aeza_os" "all" {}


output "all_operating_systems" {
  value = data.aeza_os.all.os_list
}

# Фильтрация только Linux систем
locals {
  linux_os = [
    for os in data.aeza_os.all.os_list :
    os if os.group != "windows_server"
  ]
  
  # Создание карты slug->name для удобства использования
  os_map = {
    for os in data.aeza_os.all.os_list :
    os.slug => os.name
  }
  
  # Получение конкретной ОС по slug
  ubuntu_2204 = [
    for os in data.aeza_os.all.os_list :
    os if os.slug == "ubuntu_2204"
  ][0]
}
```

### **Пример вывода**

```bash
+ all_operating_systems = [
  + {
      + enabled    = true
      + group      = "ubuntu"
      + id         = 6
      + name       = "Ubuntu 22.04"
      + order      = 0
      + repository = ""
      + slug       = "ubuntu_2204"
      + targets    = {
          + "vm6"   = 940
          + "vmmgr" = 6
        }
      + username   = "root"
    },
  + {
      + enabled    = true
      + group      = "windows_server"
      + id         = 17
      + name       = "Windows Server 2019"
      + order      = 0
      + repository = ""
      + slug       = "windows_server_2019"
      + targets    = {
          + "vm6"   = 931
          + "vmmgr" = 18
        }
      + username   = "Administrator"
    },
]

+ operating_systems_map = {
    + "alma_8"                 = "Alma Linux 8"
    + "alma_9"                 = "Alma Linux 9"
    + "centos_7"               = "CentOS 7"
    + "centos_9"               = "CentOS 9 Stream"
    + "debian_10"              = "Debian 10"
    + "debian_11"              = "Debian 11"
    + "debian_12"              = "Debian 12"
    + "rocky_8"                = "Rocky Linux 8"
    + "rocky_9"                = "Rocky Linux 9"
    + "ubuntu_2004"            = "Ubuntu 20.04"
    + "ubuntu_2204"            = "Ubuntu 22.04"
    + "ubuntu_2404"            = "Ubuntu 24.04"
    + "windows_server_2016"    = "Windows Server 2016"
    + "windows_server_2019"    = "Windows Server 2019"
    + "windows_server_2022"    = "Windows Server 2022"
  }
```



