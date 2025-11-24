# **Продукты Aeza**

🎯 Предоставляет предопределенные сформированные продукты, распределенные по типам. Например ниже на изображенни для типа VPS:

<img width="1737" height="1004" alt="image" src="https://github.com/user-attachments/assets/93be1f4e-a30f-4dab-b63b-26a4448c7ffb" />

### **Основные типы продуктов**

| Тип | Название | Примеры |
|-----|----------|----------|
| `vps` | Виртуальный сервер | MSK-1, SPB-2, CLT-4 |
| `dedicated` | Выделенный сервер | AMD Ryzen 9 5950X, Intel Core i7-6700 |
| `domain` | Домен | .ru, .com, .net, .org |
| `hicpu` | Hi-CPU сервер | EFH-5, THR-5 |
| `storage` | Storage VPS | NAS-2 |
| `soft` | Лицензия | ISPManager Lite, Pro, Host |
| `s3` | Объектное хранилище | 10GB, 100GB, 500GB |

Доступно в основной (legacy) версии /my.aeza.net/api/

```bash
curl -X 'GET' \
  'https://my.aeza.net/api/services/products' \
  -H 'accept: application/json' \
  -H 'X-API-KEY: yourAPIkey'
```

```go
type ProductsResponse struct {
    Data struct {
        SelectorMode string    `json:"selectorMode"`
        Filter       *string   `json:"filter"`
        Items        []Product `json:"items"`
    } `json:"data"`
}

type Product struct {
    ID                     int                    `json:"id"`
    Name                   string                 `json:"name"`
    Type                   string                 `json:"type"`
    GroupID                *int                   `json:"groupId"`
    Order                  int                    `json:"order"`
    Configuration          []ProductConfig        `json:"configuration"`
    DefaultParameters      map[string]interface{} `json:"defaultParameters"`
    Payload                map[string]interface{} `json:"payload"`
    IsPrivate              bool                   `json:"isPrivate"`
    Prices                 ProductPrices          `json:"prices"`
    RawPrices              ProductPrices          `json:"rawPrices"`
    IndividualPrices       ProductPrices          `json:"individualPrices"`
    InstallPrice           float64                `json:"installPrice"`
    FirstPrices            ProductPrices          `json:"firstPrices"`
    IndividualFirstPrices  ProductPrices          `json:"individualFirstPrices"`
    IndividualInstallPrice *float64               `json:"individualInstallPrice"`
    SummaryConfiguration   map[string]interface{} `json:"summaryConfiguration"`
    LocaledPayload         map[string]interface{} `json:"localedPayload"`
    PrettyLocaledPayload   map[string]interface{} `json:"prettyLocaledPayload"`
    Group                  *ProductGroup          `json:"group"`
    TypeObject             *ServiceType           `json:"typeObject"`
    ServiceHandler         string                 `json:"serviceHandler"`
}

type ProductPrices struct {
    Hour        float64 `json:"hour"`
    Month       float64 `json:"month"`
    Year        float64 `json:"year"`
    HalfYear    float64 `json:"half_year"`
    QuarterYear float64 `json:"quarter_year"`
}

type ProductConfig struct {
    Slug   string                 `json:"slug"`
    Base   float64                `json:"base"`
    Max    float64                `json:"max"`
    Type   string                 `json:"type"`
    Count  int                    `json:"count"`
    Prices map[string]interface{} `json:"prices"`
}

type ProductGroup struct {
    ID                   int                    `json:"id"`
    Order                int                    `json:"order"`
    Names                map[string]string      `json:"names"`
    Type                 string                 `json:"type"`
    Role                 *string                `json:"role"`
    ParentID             *int                   `json:"parentId"`
    Descriptions         map[string]string      `json:"descriptions"`
    Payload              map[string]interface{} `json:"payload"`
    LocaledPayload       map[string]interface{} `json:"localedPayload"`
    ConfigurationPrices  map[string]interface{} `json:"configurationPrices"`
    HasProducts          bool                   `json:"hasProducts"`
    PrettyLocaledPayload map[string]interface{} `json:"prettyLocaledPayload"`
    Name                 string                 `json:"name"`
    TypeObject           *ServiceType           `json:"typeObject"`
    ServiceHandler       string                 `json:"serviceHandler"`
}
```

### **ProductsResponse.Data**

| Поле | Тип | Описание |
|------|-----|----------|
|`SelectorMode` | string | Режим выбора (например, "all") |
|`Filter` | *string | Фильтры (может быть null) |
|`Items` | []Product | Массив продуктов |

### **Product - основные поля**

| Поле | Тип | Описание | Пример |
|------|-----|----------|---------|
|`ID` | int | Уникальный идентификатор продукта | `616`, `148` |
|`Name` | string | Название продукта | "MSK-1", "AMD Ryzen 9 5950X" |
|`Type` | string | Тип продукта | "vps", "dedicated", "domain" |
|`GroupID` | *int | Идентификатор группы продуктов | `1`, `48` |
|`ServiceHandler` | string | Система управления сервисом | "vm6", "manual", "feru" |
|`InstallPrice` | float64 | Цена установки/активации | `5227.0`, `37133.0` |
|`IndividualInstallPrice` | *float64 | Индивидуальная цена установки | `51578.0`, `null` |
|`Prices` | ProductPrices | Цены в различных периодах | - |

### **ProductConfig - конфигурация**

| Поле | Тип | Описание | Пример |
|------|-----|----------|---------|
|`Slug` | string | Идентификатор параметра | "cpu", "ram", "rom", "quota" |
|`Base` | float64 | Базовое значение | `2.0`, `3.5`, `60.0` |
|`Max` | float64 | Максимальное значение | `16.0`, `128.0`, `2000.0` |
|`Type` | string | Тип ресурса | "nvme", "ssd", "ddr4" |
|`Count` | int | Количество дисков/устройств | `1`, `2` |
|`Prices` | map[string]interface{} | Дополнительные цены | - |

### **ProductPrices - ценовые периоды**

| Поле | Тип | Описание | Пример |
|------|-----|----------|---------|
|`Hour` | float64 | Почасовая оплата | `3.0`, `4.5`, `24.3` |
|`Month` | float64 | Помесячная оплата | `707.8`, `1415.0` |
|`Year` | float64 | Годовая оплата | `7466.0`, `14943.0` |
|`HalfYear` | float64 | Полугодовая оплата | `3862.0`, `7723.0`  |
|`QuarterYear` | float64 | Квартальная оплата | `2016.0`, `4032.0` |

### **Дополнительные ценовые поля**

| Поле | Тип | Описание |
|------|-----|----------|
|`RawPrices` | ProductPrices | Исходные цены без скидок |
|`IndividualPrices` | ProductPrices | Индивидуальные цены |
|`FirstPrices` | ProductPrices | Цены первого периода (акционные) |
|`IndividualFirstPrices` | ProductPrices | Индивидуальные цены первого периода |

### **Особенности использования**
- Все числовые поля теперь используют float64 для поддержки дробных значений

- InstallPrice может содержать дробные значения для точного расчета стоимости

- GroupID и IndividualInstallPrice могут быть null

- Configuration параметры поддерживают дробные значения для гибкой настройки ресурсов

Этот endpoint предоставляет полную информацию о всех доступных продуктах, их конфигурациях и ценах для интеграции и автоматизации.