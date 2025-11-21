# **Услуги (Services)**
🎯 Все услуги пользователя в системе Aeza

<img width="1727" height="282" alt="image" src="https://github.com/user-attachments/assets/249104e2-fa02-43a5-b63d-82da11cf7fa2" />

### 📋 Основная информация
**Услуги (Services)** - это активные сервисы пользователя (VPS, выделенные серверы, домены, VPN и т.д.)

### 🔧 API Endpoints

Legacy API
```bash
curl -X 'GET' \
  'https://my.aeza.net/api/services \
  -H 'accept: application/json' \
  -H 'X-API-KEY: yourAPIkey'
```

API v2 (не работает)
```bash
curl -X 'GET' \
  'https://my.aeza.net/api/v2/services' \
  -H 'accept: application/json' \
  -H 'X-API-KEY: yourAPIkey'
```


### 📊 **Структуры данных**
Legacy API Response
```go
type ServicesResponse struct {
    Data struct {
        SelectorMode       string    `json:"selectorMode"`
        Filter             *string   `json:"filter"`
        Items              []Service `json:"items"`
        Total              int       `json:"total"`
        Edit               bool      `json:"edit"`
    } `json:"data"`
}
```
API v2 Response
```go
type ServicesResponseV2 struct {
    Items []ServiceV2 `json:"items"`
    Total int         `json:"total"`
}
```


### 🏗️ **Модели услуг**
Legacy Service Model (полная версия)
```go
type Service struct {
    ID               int                      `json:"id"`
    OwnerID          int                      `json:"ownerId"`
    ProductID        int                      `json:"productId"`
    Name             string                   `json:"name"`
    IP               string                   `json:"ip"`
    IPs              []IPAddress              `json:"ips"`
    IPv6             []IPv6Address            `json:"ipv6"`
    Payload          map[string]interface{}   `json:"payload"`
    Configuration    map[string]interface{}   `json:"configuration"`
    PaymentTerm      string                   `json:"paymentTerm"`
    AutoProlong      bool                     `json:"autoProlong"`
    Backups          bool                     `json:"backups"`
    Status           string                   `json:"status"`
    LastStatus       *string                  `json:"lastStatus"`
    Timestamps       ServiceTimestamps        `json:"timestamps"`
    PaymentTermRatio float64                  `json:"paymentTermRatio"`
    Opportunities    []string                 `json:"opportunities"`
    Capabilities     []string                 `json:"capabilities"`
    Product          Product                  `json:"product"`
    Parameters       ServiceParameters        `json:"parameters"`
    SecureParameters SecureParameters         `json:"secureParameters"`
    LocationCode     string                   `json:"locationCode"`
    RawPrices        map[string]int           `json:"rawPrices"`
    CurrentStatus    string                   `json:"currentStatus"`
    CurrentTask      *string                  `json:"currentTask"`
    IndividualPrices map[string]int           `json:"individualPrices"`
}
```

API v2 Service Model (упрощенная версия)
```go
type ServiceV2 struct {
    ID           int                    `json:"id"`
    Name         string                 `json:"name"`
    IP           string                 `json:"ip"`
    Payload      map[string]interface{} `json:"payload"`
    Price        int                    `json:"price"`
    PaymentTerm  string                 `json:"paymentTerm"`
    AutoProlong  bool                   `json:"autoProlong"`
    CreatedAt    string                 `json:"createdAt"`
    ExpiresAt    string                 `json:"expiresAt"`
    Status       string                 `json:"status"`
    TypeSlug     string                 `json:"typeSlug"`
    ProductName  string                 `json:"productName"`
    LocationCode string                 `json:"locationCode"`
    CurrentTask  *string                `json:"currentTask"`
    Capabilities []string               `json:"capabilities"`
}
```

### 🔍 Ключевые поля
ServicesResponse.Data:
- `SelectorMode` - режим выбора (например, "all")

-  `Filter` - примененные фильтры

- `Items` - массив услуг пользователя

- `Total` - общее количество услуг

- `Edit` - флаг возможности редактирования

Основные поля услуги:
- `ID` - уникальный идентификатор услуги

- `Name` - название услуги

- `TypeSlug` - тип услуги (vps, dedicated, domain, etc.)

- `Status` - текущий статус

- `IP` - основной IP-адрес

- `LocationCode` - код локации (nl, de, us, etc.)

- `PaymentTerm` - период оплаты (hour, month, year)

- `Price` - текущая цена

### 📈 **Статусы услуг**

| Статус | Описание |
|--------|----------|
| `active` | Услуга активна и работает |
| `activation_wait` | Ожидание активации |
| `prolong_wait` | Ожидание продления |
| `suspended` | Услуга приостановлена |
| `deleted` | Услуга удалена |
| `blocked` | Услуга заблокирована |
| `expired` | Срок действия услуги истек |
| `pending` | В процессе обработки |
| `failed` | Ошибка активации |
| `cancelled` | Услуга отменена |

### 🌍 **Локации**

| Код | Локация |
|-----|---------|
| `nl` | Нидерланды |
| `de` | Германия |
| `us` | США |
| `fr` | Франция |
| `fi` | Финляндия |
| `ru` | Россия |
| `kz` | Казахстан |
| `ua` | Украина |
| `pl` | Польша |
| `sg` | Сингапур |
| `jp` | Япония |
| `kr` | Южная Корея |

### 💰 **Периоды оплаты**

| Период | Описание |
|--------|----------|
| `hour` | Почасовая оплата |
| `month` | Месячная оплата |
| `year` | Годовая оплата |
| `half_year` | Полугодовая оплата |
| `quarter_year` | Квартальная оплата |
| `two_year` | Двухгодичная оплата |
| `three_year` | Трехгодичная оплата |

Таким образом, этот endpoint используется для получения списка всех услуг пользователя с возможностью фильтрации и пагинации.