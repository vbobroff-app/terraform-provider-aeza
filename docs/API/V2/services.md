# **Услуги (Services) V2**

🎯 Все услуги пользователя в системе Aeza - новая версия API

<img width="1727" height="282" alt="image" src="https://github.com/user-attachments/assets/249104e2-fa02-43a5-b63d-82da11cf7fa2" />

### 📋 **Основная информация**
Услуги V2 (Services) - V2 API для получения сервисов аккаунта пользователя с улучшенной структурой данных и поддержкой новых возможностей.

### 🔧 **API Endpoint**
V2 API

```bash
curl -X 'GET' \
  'https://my.aeza.net/api/v2/services' \
  -H 'accept: application/json' \
  -H 'X-API-KEY: yourAPIkey'
```

### 📊 **Структуры данных**
V2 API Response

```go
type Service struct {
	ID           int                    `json:"id"`
	Name         string                 `json:"name"`
	IP           string                 `json:"ip"`
	Payload      map[string]interface{} `json:"payload"`
	Price        float64                `json:"price"`
	PaymentTerm  string                 `json:"paymentTerm"`
	AutoProlong  bool                   `json:"autoProlong"`
	CreatedAt    string                 `json:"createdAt"`    // ISO 8601: "2025-11-20T13:12:12.733Z"
	ExpiresAt    string                 `json:"expiresAt"`    // ISO 8601: "2025-11-20T13:12:12.733Z"
	Status       string                 `json:"status"`       // "activation_wait", "active", etc.
	TypeSlug     string                 `json:"typeSlug"`     // "vps"
	ProductName  string                 `json:"productName"`  // "SWE-PROMO"
	LocationCode string                 `json:"locationCode"` // "de", "nl", etc.
	CurrentTask  *CurrentTask           `json:"currentTask"`  // nullable объект, а не строка!
	Capabilities []string               `json:"capabilities"` // ["change_password", "rename", ..., "backups"]
}

type CurrentTask struct {
	ID        string `json:"id"`
	Slug      string `json:"slug"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"` // ISO 8601
	Status    string `json:"status"`    // "queued", "running", etc.
}

type ComputedParametersVPS struct {
	CPU      int     `json:"cpu"`
	RAM      int     `json:"ram"`
	ROM      int     `json:"rom"`
	IP       int     `json:"ip"`
	OS       string  `json:"os"`
	Node     string  `json:"node"`
	ISOURL   string  `json:"isoUrl"`
	Recipe   *string `json:"recipe"` // nullable
	Username string  `json:"username"`
}
```

#### Структура **Service**
| Поле | Тип | Описание |
|------|-----|-----------|
| `ID` | int | Уникальный идентификатор услуги |
| `Name` | string | Название услуги |
| `IP` | string | Основной IP-адрес услуги |
| `Payload` | map[string]interface{} | Дополнительные данные услуги |
| `Price` | float64 | Цена услуги |
| `PaymentTerm` | string | Период оплаты (hour, month, year) |
| `AutoProlong` | bool | Флаг автоматического продления |
| `CreatedAt` | string | Дата создания в формате ISO 8601 |
| `ExpiresAt` | string | Дата истечения в формате ISO 8601 |
| `Status` | string | Текущий статус (activation_wait, active и др.) |
| `TypeSlug` | string | Тип услуги (vps, dns, vpn и т.д.) |
| `ProductName` | string | Название продукта |
| `LocationCode` | string | Код локации (de, nl и т.д.) |
| `CurrentTask` | *CurrentTask | Текущая задача (может быть null) |
| `Capabilities` | []string | Доступные возможности услуги |

#### Структура **CurrentTask**
| Поле | Тип | Описание |
|------|-----|-----------|
| `ID` | string | Уникальный идентификатор задачи |
| `Slug` | string | Идентификатор типа задачи |
| `Name` | string | Название задачи |
| `CreatedAt` | string | Дата создания задачи в формате ISO 8601 |
| `Status` | string | Статус задачи (queued, running и др.) |

#### Структура **ComputedParametersVPS**
| Поле | Тип | Описание |
|------|-----|-----------|
| `CPU` | int | Количество ядер процессора |
| `RAM` | int | Объем оперативной памяти (в МБ) |
| `ROM` | int | Объем дискового пространства (в ГБ) |
| `IP` | int | Количество IP-адресов |
| `OS` | string | Операционная система |
| `Node` | string | Нода размещения |
| `ISOURL` | string | URL ISO образа |
| `Recipe` | *string | Рецепт настройки (может быть null) |
| `Username` | string | Имя пользователя по умолчанию |


### 🔄 **Отличия от Legacy API**
- Упрощенная структура - меньше вложенных объектов

- Стандартизированные форматы дат - ISO 8601 вместо Unix timestamp

- Улучшенная обработка задач - отдельный объект для CurrentTask

- Единообразные статусы - стандартизированные значения статусов

- Расширенные возможности - массив capabilities для управления сервисом

Также доступна в [legacy версии api/services](../legacy/services.md)