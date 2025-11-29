# **Продление услуги (Service Prolong)**
🎯 Продление срока действия услуги с созданием финансовой транзакции

### 📋 **Основная информация**
Продление услуги - операция увеличения срока действия услуги на указанный период с списанием средств выбранным способом оплаты.

### 🔧 **API Endpoint**
#### Legacy API

```bash
curl -X 'POST' \
  'https://my.aeza.net/api/services/{id}/prolong' \
  -H 'accept: application/json' \
  -H 'X-API-KEY: yourAPIkey' \
  -H 'Content-Type: application/json' \
  -d '{
    "method": "balance",
    "term": "hour", 
    "count": 2
  }'
```
### 📥 **Request Structure**
```go
type ServiceProlongRequest struct {
    Method string `json:"method"`  // Способ оплаты
    Term   string `json:"term"`    // Период продления
    Count  int64  `json:"count"`   // Количество периодов
}
```

#### Параметры запроса
| Поле | Тип | Обязательный | Описание | Допустимые значения |
|------|-----|-------------|-----------|-------------------|
| `method` | string | ✅ | Способ оплаты | `"balance"`, `"sms"`, и другие доступные методы |
| `term` | string | ✅ | Период продления | `"hour"`, `"day"`, `"month"`, `"year"` |
| `count` | int64 | ✅ | Количество периодов | Положительное целое число |

###📤 **Response Structure**
```go
type ServiceProlongResponse struct {
    Data struct {
        Transaction Transaction `json:"transaction"`
    } `json:"data"`
}

type Transaction struct {
    ID                int64                  `json:"id"`
    OwnerID           int64                  `json:"ownerId"`
    Amount            float64                `json:"amount"`
    BonusAmount       float64                `json:"bonusAmount"`
    Mode              string                 `json:"mode"`
    Status            string                 `json:"status"`
    PerformedAt       interface{}            `json:"performedAt"`
    CreatedAt         int64                  `json:"createdAt"`
    Type              string                 `json:"type"`
    ServiceID         interface{}            `json:"serviceId"`
    ItemSlug          interface{}            `json:"itemSlug"`
    Description       interface{}            `json:"description"`
    RandomString      interface{}            `json:"randomString"`
    Payload           map[string]interface{} `json:"payload"`
    InvoiceID         int64                  `json:"invoiceId"`
    Invoice           interface{}            `json:"invoice"`
    PrettyAmount      PrettyAmount           `json:"prettyAmount"`
    OrderInfo         interface{}            `json:"orderInfo"`
    PrettyBonusAmount PrettyAmount           `json:"prettyBonusAmount"`
}

type PrettyAmount struct {
    Value float64 `json:"value"`
}
```

#### Структура **ServiceProlongResponse.Data**
| Поле | Тип | Описание |
|------|-----|-----------|
| `Transaction` | Transaction | Созданная транзакция |

#### Структура **Transaction**
| Поле | Тип | Описание |
|------|-----|-----------|
| `ID` | int64 | Уникальный идентификатор транзакции |
| `OwnerID` | int64 | ID владельца транзакции |
| `Amount` | float64 | Сумма транзакции (может быть отрицательной для списаний) |
| `BonusAmount` | float64 | Бонусная сумма |
| `Mode` | string | Режим транзакции |
| `Status` | string | Статус транзакции |
| `PerformedAt` | interface{} | Время выполнения (может быть null) |
| `CreatedAt` | int64 | Время создания (Unix timestamp) |
| `Type` | string | Тип транзакции (`"prolong"`) |
| `ServiceID` | interface{} | ID услуги (может быть int64 или null) |
| `ItemSlug` | interface{} | Слаг товара (может быть null) |
| `Description` | interface{} | Описание (может быть null) |
| `RandomString` | interface{} | Случайная строка (может быть null) |
| `Payload` | map[string]interface{} | Дополнительные данные транзакции |
| `InvoiceID` | int64 | ID инвойса |
| `Invoice` | interface{} | Инвойс (может быть null) |
| `PrettyAmount` | PrettyAmount | Отформатированная сумма |
| `OrderInfo` | interface{} | Информация о заказе (может быть null) |
| `PrettyBonusAmount` | PrettyAmount | Отформатированная бонусная сумма |
___

#### Структура **PrettyAmount**
| Поле | Тип | Описание |
|------|-----|-----------|
|Value	|float64	|Числовое значение суммы|
___

### 💡 **Payload структура**
#### Содержимое поля Payload зависит от выбранного периода (term):

 Для term = "hour"
```json
{
  "hours": 4,
  "term": "hour"
}
```
 Для term = "month"
```json
{
  "months": 1,
  "term": "month"
}
```
#### 🚨 **Коды ответа**
| Код | Описание |
|-----|-----------|
| `200` | Успешное выполнение |
| `400` | Неверные параметры запроса |
| `401` | Ошибка аутентификации |
| `403` | Недостаточно средств |
| `404` | Услуга не найдена |
| `409` | Конфликт состояния услуги |

### 📝 **Примеры использования**
#### Продление на 2 часа через баланс
```bash
curl -X 'POST' \
  'https://my.aeza.net/api/services/12345/prolong' \
  -H 'X-API-KEY: yourAPIkey' \
  -H 'Content-Type: application/json' \
  -d '{
    "method": "balance",
    "term": "hour",
    "count": 2
  }'
```

Продление на 2 месяца
```bash
curl -X 'POST' \
  'https://my.aeza.net/api/services/12345/prolong' \
  -H 'X-API-KEY: yourAPIkey' \
  -H 'Content-Type: application/json' \
  -d '{
    "method": "balance",
    "term": "month",
    "count": 2
  }'
```

#### Успешный ответ
```json
{
  "data": {
    "transaction": {
      "id": 24028197,
      "ownerId": 770366,
      "amount": -2,
      "bonusAmount": 0,
      "mode": "pending",
      "status": "created",
      "performedAt": null,
      "createdAt": 1764394534,
      "type": "prolong",
      "serviceId": 1577409,
      "itemSlug": null,
      "description": null,
      "randomString": null,
      "payload": {
        "hours": 1,
        "term": "hour"
      },
      "invoiceId": 0,
      "invoice": null,
      "prettyAmount": {
        "value": -2
      },
      "orderInfo": null,
      "prettyBonusAmount": {
        "value": 0
      }
    }
  }
}
```

###⚠️ **Особенности**
- Создает финансовую транзакцию с типом "prolong"

- Сумма может быть отрицательная (списание средств)

- Поле payload содержит детали продления в зависимости от периода

- Статус транзакции обычно "created" или "pending"