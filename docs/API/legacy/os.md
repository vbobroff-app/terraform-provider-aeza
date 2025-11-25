# **Операционные системы**

🎯 Все операционные системы

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

Legacy ендпоинт [Leagacy API /my.aeza.net/api/](../api.md)

```bash
curl -X 'GET' \
  'https://my.aeza.net/api/v2/services/operating-systems' \
  -H 'accept: application/json' \
  -H 'X-API-KEY: yourAPIkey'
```

Что в ответе - объект с данными Data c массивом операционных систем []OperatingSystem

### **Структура OperatingSystem**

```go
type OSResponse struct {
    Data struct {
        Items []OperatingSystem `json:"items"`
    } `json:"data"`
}

type OperatingSystem struct {
    ID         int            `json:"id"`
    Name       string         `json:"name"`
    Repository *string         `json:"repository"`
    Group      string         `json:"group"`
    Enabled    bool           `json:"enabled"`
    Slug       string         `json:"slug"`
    Username   string         `json:"username"`
    Targets    map[string]int `json:"targets"`
}
```

| Поле | Тип | Описание |
|------|-----|-----------|
| `ID` | number | Уникальный числовой идентификатор |
| `Name` | string | Полное название операционной системы |
| `Repository` | string \| null | Репозиторий ОС |
| `Group` | string | Группа ОС (alma, centos, debian, ubuntu, windows_server) |
| `Username` | string | Имя пользователя по умолчанию |
| `Enabled` | boolean | Флаг доступности ОС |
| `Slug` | string | Уникальный строковый идентификатор ОС |
| `Targets` | object | Карта целевых платформ и их идентификаторов |


Основоной ендпоинт в [V2 версии api/v2/operating-systems](../V2/operating-sistem.md) 