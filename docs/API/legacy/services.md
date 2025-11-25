# **Услуги (Services)**
🎯 Все услуги пользователя в системе Aeza

<img width="1727" height="282" alt="image" src="https://github.com/user-attachments/assets/249104e2-fa02-43a5-b63d-82da11cf7fa2" />

### 📋 Основная информация
**Услуги (Services)** - это сервисы аккаунта пользователя (VPS, выделенные серверы, домены, VPN и т.д.)

### 🔧 API Endpoint

Legacy API
```bash
curl -X 'GET' \
  'https://my.aeza.net/api/services \
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

type Service struct {
	ID            int     `json:"id"`
	OwnerID       int     `json:"ownerId"`
	ProductID     int     `json:"productId"`
	Name          string  `json:"name"`
	IP            string  `json:"ip"`
	PaymentTerm   string  `json:"paymentTerm"`
	AutoProlong   bool    `json:"autoProlong"`
	Backups       bool    `json:"backups"`
	Status        string  `json:"status"`
	LastStatus    *string `json:"lastStatus"`
	Product       Product `json:"product"`
	LocationCode  string  `json:"locationCode"`
	CurrentStatus string  `json:"currentStatus"`
}


type ServiceVPS struct {
	Service
	// VPS-специфичные  override поля
	Parameters           Parameters                   `json:"parameters"`
	SecureParameters     SecureParameters             `json:"secureParameters"`
	IPs                  []IPAddress                  `json:"ips"`
	IPv6                 []IPv6Address                `json:"ipv6"`
	Payload              map[string]interface{}       `json:"payload"`
	Configuration        map[string]interface{}       `json:"configuration"`
	SpecialPrices        map[string]interface{}       `json:"specialPrices"`
	RelativePrices       map[string]interface{}       `json:"relativePrices"`
	Schedule             *string                      `json:"schedule"`
	Timestamps           ServiceTimestamps            `json:"timestamps"`
	PaymentTermRatio     float64                      `json:"paymentTermRatio"`
	Opportunities        []string                     `json:"opportunities"`
	Capabilities         []string                     `json:"capabilities"`
	RawPrices            map[string]int               `json:"rawPrices"`
	SummaryConfiguration map[string]ConfigurationItem `json:"summaryConfiguration"`
	IndividualPrices     map[string]int               `json:"individualPrices"`
	CurrentTask          *string                      `json:"currentTask"`
}

type Parameters struct {
	OS       string  `json:"os,omitempty"`
	ISOURL   string  `json:"isoUrl,omitempty"`
	Recipe   *string `json:"recipe,omitempty"`
	Username string  `json:"username,omitempty"`
}

type SecureParameters struct {
	Unsecure bool                   `json:"unsecure"`
	Data     map[string]interface{} `json:"data"`
}

type IPAddress struct {
	Key                string `json:"key"`
	Mask               string `json:"mask"`
	Value              string `json:"value"`
	Status             string `json:"status"`
	Gateway            string `json:"gateway"`
	ExtendedProtection bool   `json:"extendedProtection"`
}
type IPv6Address struct {
	IPs     []interface{} `json:"ips"`
	Key     string        `json:"key"`
	Value   string        `json:"value"`
	Prefix  int           `json:"prefix"`
	Gateway string        `json:"gateway"`
}

type ServiceTimestamps struct {
	CreatedAt   int64 `json:"createdAt"`   // Unix timestamp
	ExpiresAt   int64 `json:"expiresAt"`   // Unix timestamp
	PurchasedAt int64 `json:"purchasedAt"` // Unix timestamp
}

type ConfigurationItem struct {
	Prices       map[string]interface{} `json:"prices"`
	Count        int                    `json:"count"`
	Base         int                    `json:"base"`
	Additionally int                    `json:"additionally"`
}

```

#### Структура **ServicesResponse.Data**
| Поле | Тип | Описание |
|------|-----|-----------|
| `SelectorMode` | string | Режим выбора |
| `Filter` | *string | Фильтры (может быть null) |
| `Items` | array | Массив услуг |
| `Total` | int | Общее количество услуг |
| `Edit` | bool | Флаг возможности редактирования |

#### Структура **Service**
| Поле | Тип | Описание |
|------|-----|-----------|
| `ID` | int | Уникальный идентификатор услуги |
| `OwnerID` | int | ID владельца услуги |
| `ProductID` | int | ID продукта |
| `Name` | string | Название услуги |
| `IP` | string | IP-адрес услуги |
| `PaymentTerm` | string | Период оплаты |
| `AutoProlong` | bool | Флаг автоматического продления |
| `Backups` | bool | Флаг наличия бэкапов |
| `Status` | string | Текущий статус |
| `LastStatus` | *string | Предыдущий статус (может быть null) |
| `Product` | Product | Информация о продукте |
| `LocationCode` | string | Код локации |
| `CurrentStatus` | string | Текущий статус отображения |

#### Структура **ServiceVPS**
| Поле | Тип | Описание |
|------|-----|-----------|
| `Service` | Service | Базовая структура услуги |
| `Parameters` | Parameters | Основные параметры VPS |
| `SecureParameters` | SecureParameters | Защищенные параметры |
| `IPs` | []IPAddress | Список IPv4 адресов |
| `IPv6` | []IPv6Address | Список IPv6 адресов |
| `Payload` | map[string]interface{} | Дополнительные данные |
| `Configuration` | map[string]interface{} | Конфигурация услуги |
| `SpecialPrices` | map[string]interface{} | Специальные цены |
| `RelativePrices` | map[string]interface{} | Относительные цены |
| `Schedule` | *string | Расписание (может быть null) |
| `Timestamps` | ServiceTimestamps | Временные метки |
| `PaymentTermRatio` | float64 | Коэффициент периода оплаты |
| `Opportunities` | []string | Возможности услуги |
| `Capabilities` | []string | Возможности системы |
| `RawPrices` | map[string]int | Цены в числовом формате |
| `SummaryConfiguration` | map[string]ConfigurationItem | Сводная конфигурация |
| `IndividualPrices` | map[string]int | Индивидуальные цены |
| `CurrentTask` | *string | Текущая задача (может быть null) |

#### Структура **Parameters**
| Поле | Тип | Описание |
|------|-----|-----------|
| `OS` | string | Операционная система |
| `ISOURL` | string | URL ISO образа |
| `Recipe` | *string | Рецепт настройки (может быть null) |
| `Username` | string | Имя пользователя |

#### Структура **SecureParameters**
| Поле | Тип | Описание |
|------|-----|-----------|
| `Unsecure` | bool | Флаг безопасности |
| `Data` | map[string]interface{} | Защищенные данные |

#### Структура **IPAddress**
| Поле | Тип | Описание |
|------|-----|-----------|
| `Key` | string | Ключ IP-адреса |
| `Mask` | string | Маска сети |
| `Value` | string | IP-адрес |
| `Status` | string | Статус IP-адреса |
| `Gateway` | string | Шлюз |
| `ExtendedProtection` | bool | Расширенная защита |

#### Структура **IPv6Address**
| Поле | Тип | Описание |
|------|-----|-----------|
| `IPs` | []interface{} | Список IPv6 адресов |
| `Key` | string | Ключ IPv6 |
| `Value` | string | Основной IPv6 адрес |
| `Prefix` | int | Префикс сети |
| `Gateway` | string | IPv6 шлюз |

#### Структура **ServiceTimestamps**
| Поле | Тип | Описание |
|------|-----|-----------|
| `CreatedAt` | int64 | Дата создания (Unix timestamp) |
| `ExpiresAt` | int64 | Дата истечения (Unix timestamp) |
| `PurchasedAt` | int64 | Дата покупки (Unix timestamp) |

#### Структура **ConfigurationItem**
| Поле | Тип | Описание |
|------|-----|-----------|
| `Prices` | map[string]interface{} | Цены конфигурации |
| `Count` | int | Количество |
| `Base` | int | Базовое значение |
| `Additionally` | int | Дополнительное значение |


Также досуступен в новой [V2 версии api/services](../V2/services.md) 