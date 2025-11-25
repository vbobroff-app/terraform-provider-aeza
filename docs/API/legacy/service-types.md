# **Типы услуг**

🎯 Все типы услуг

<img width="1130" height="650" alt="image" src="https://github.com/user-attachments/assets/b655a38e-3562-4dbd-8f12-70f352530442" />


```hcl
[
  "hicpu",      // Hi-CPU сервер
  "domain",     // Домен
  "vpn",        // VPN
  "soft",       // Лицензия ispmanager
  "dedicated",  // Выделенный сервер
  "vps"         // Виртуальный сервер
]
```

slug - человеко-читаемый уникальный идентификатор


| Slug | Название | Описание |
|------|----------|----------|
| `hicpu` | Hi-CPU сервер | Серверы с высокопроизводительными CPU |
| `domain` | Домен | Доменные имена |
| `vpn` | VPN | VPN сервисы |
| `soft` | Лицензия ispmanager | Лицензии панели управления |
| `dedicated` | Выделенный сервер | Выделенные физические серверы |
| `vps` | Виртуальный сервер | Виртуальные приватные серверы |


Доступно в [основной (legacy) версии](./api.md)  /my.aeza.net/api/

```bash
 curl -X 'GET'/
'https://my.aeza.net/api/services/types?offset=0&limit=9'/ 
-H 'accept: application/json'/
-H 'X-API-KEY: yourAPIkey'
```

Что в ответе - legacy интерфейс с Data, SelectedMode и пр.
```go
type ResponseServiceType struct {
    Data struct {
        SelectorMode       string        `json:"selectorMode"`
        Filter             interface{}   `json:"filter"`
        Items              []ServiceType `json:"items"`
        Total              int           `json:"total"`
        Edit               bool          `json:"edit"`
    } `json:"data"`
}

type ServiceType struct {
    Slug                string                 `json:"slug"`
    Order               int                    `json:"order"`
    Names               map[string]string      `json:"names"`
    Payload             map[string]interface{} `json:"payload"`
    LocaledPayload      map[string]interface{} `json:"localedPayload"`
    Name                string                 `json:"name"`
    PrettyLocaledPayload map[string]interface{} `json:"prettyLocaledPayload"`
}
```
</br>

**ResponseServiceType.Data:**
| Поле | Тип | Описание |
|------|-----|-----------|
| `SelectorMode` | string | Режим выбора (например, "all") |
|` Filter` | mixed | Фильтры (может быть null, string или object) |
| `Items` | array | Массив типов услуг |
| `Total` | number | Общее количество типов |
| `Edit` | boolean | Флаг возможности редактирования |

</br>

**ServiceType:**
| Поле | Тип | Описание |
|------|-----|-----------|
| `Slug` | string | Уникальный идентификатор типа (vps, dedicated, domain и т.д.) |
| `Order` | number | Порядок отображения |
| `Names` | object | Многоязычные названия (map с кодами языков: en, ru, zh и т.д.) |
| `Payload` | object | Дополнительные данные (бейджи, цвета, специфичные параметры) |
| `LocaledPayload` | object | Локализованные дополнительные данные |
| `Name` | string | Основное название (вероятно на языке пользователя) |
| `PrettyLocaledPayload` | object | Форматированные локализованные данные |