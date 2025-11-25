# **Операционные системы (Operating Systems) V2**

🎯 Все операционные системы новая V2 версия API

```hcl
[
  "alma_8",                 // Alma Linux 8
  "alma_9",                 // Alma Linux 9  
  "centos_7",               // CentOS 7
  "centos_9",               // CentOS 9 Stream
  "debian_10",              // Debian 10
  "debian_11",              // Debian 11
  "debian_12",              // Debian 12
  "rocky_8",                // Rocky Linux 8
  "rocky_9",                // Rocky Linux 9
  "ubuntu_2004",            // Ubuntu 20.04
  "ubuntu_2204",            // Ubuntu 22.04
  "ubuntu_2404",            // Ubuntu 24.04
  "windows_server_2016",    // Windows Server 2016
  "windows_server_2019",    // Windows Server 2019
  "windows_server_2022"     // Windows Server 2022
]
```
slug - человеко-читаемый уникальный идентификатор

Основной ендпоинт в [V2 API /my.aeza.net/api/v2/](../api.md)

```bash
curl -X 'GET' \
  'https://my.aeza.net/api/v2/services/operating-systems' \
  -H 'accept: application/json' \
  -H 'X-API-KEY: yourAPIkey'
```

Что в ответе - массив операционных систем []OperatingSystem

### **Структура OperatingSystem**

```go
type OperatingSystem struct {
    ID         int            `json:"id"`
    Slug       string         `json:"slug"`
    Name       string         `json:"name"`
    Repository *string        `json:"repository"`
    Group      string         `json:"group"`
    Username   string         `json:"username"`
    Enabled    bool           `json:"enabled"`
    Targets    map[string]int `json:"targets"`
    Order      int            `json:"order"`
}
```

| Поле | Тип | Описание |
|------|-----|-----------|
| `ID` | number | Уникальный числовой идентификатор |
| `Slug` | string | Уникальный строковый идентификатор ОС |
| `Name` | string | Полное название операционной системы |
| `Repository` | string \| null | Репозиторий ОС |
| `Group` | string | Группа ОС (alma, centos, debian, ubuntu, windows_server) |
| `Username` | string | Имя пользователя по умолчанию |
| `Enabled` | boolean | Флаг доступности ОС |
| `Targets` | object | Карта целевых платформ и их идентификаторов |
| `Order` | number | Порядок сортировки |

Также досуступен в [leagacy версии api/os](../legacy/os.md) 