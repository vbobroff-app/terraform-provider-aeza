# **Услуги (Services) - CRUD API**

🎯 Полное управление услугами пользователя в системе Aeza

### 📋 **Обзор /api/services/orders**

**Услуги (Services)** - это сервисы аккаунта пользователя (VPS, выделенные серверы, домены, VPN и т.д.)
___

## **🚀 CRUD Операции**

###  **1. Создание услуги POST **
#### **Legacy API**

```bash
curl -X 'POST' \
  'https://my.aeza.net/api/services/orders' \
  -H 'accept: application/json' \
  -H 'X-API-KEY: yourAPIkey' \
  -H 'Content-Type: application/json' \
  -d '{
    "method": "balance",
    "count": 1,
    "term": "hour",
    "name": "my-vps",
    "productId": 182,
    "parameters": {
      "os": "ubuntu_2404",
      "isoUrl": "",
      "recipe": null
    },
    "autoProlong": false,
    "backups": false
  }'
```

### 📊 **Структуры запроса**

#### Legacy API Request

```go
type ServiceCreateRequest struct {
    Method      string                  `json:"method"`
    Count       int                     `json:"count"`
    Term        string                  `json:"term"`
    Name        string                  `json:"name"`
    ProductID   int64                   `json:"productId"`
    Parameters  ServiceCreateParameters `json:"parameters"`
    AutoProlong bool                    `json:"autoProlong"`
    Backups     bool                    `json:"backups"`
}

type ServiceCreateParameters struct {
    Recipe *string `json:"recipe"`
    OS     string  `json:"os"`
    IsoURL string  `json:"isoUrl"`
}
```

### 📋 **Параметры создания**

| Поле | Тип | Обязательный | Описание | Пример |
|------|-----|-------------|-----------|---------|
| `method` | string | ✅ | Метод оплаты | `"balance"` |
| `count` | int | ✅ | Количество услуг | `1` |
| `term` | string | ✅ | Период оплаты | `"hour"`, `"month"`, `"year"` |
| `name` | string | ✅ | Название услуги | `"my-vps"` |
| `productId` | int64 | ✅ | ID продукта | `182` |
| `parameters.os` | string | ❌ | Операционная система | `"ubuntu_2404"`, `"debian_12"` |
| `parameters.recipe` | *string | ❌ | Рецепт настройки | `"docker"`, `null` |
| `parameters.isoUrl` | string | ❌ | URL ISO образа | `""`, `"https://example.com/image.iso"` |
| `autoProlong` | bool | ✅ | Автопродление | `true`, `false` |
| `backups` | bool | ✅ | Бэкапы | `true`, `false` |

#### ✅  Response
```go
type ServiceCreateResponse struct {
    Data ServiceCreateData `json:"data"`
}

type ServiceCreateData struct {
    Items       []ServiceOrderItem `json:"items"`
    Transaction Transaction        `json:"transaction"`
}

type ServiceOrderItem struct {
    ID                int64   `json:"id"`
    CreatedServiceIds []int64 `json:"createdServiceIds"`
    Status            string  `json:"status"`
    // ... другие поля
}
```
###  **2. Чтение услуги GET /api/services/{id}**
#### Legacy API

```bash
curl -X 'GET' \
  'https://my.aeza.net/api/services/1576759' \
  -H 'accept: application/json' \
  -H 'X-API-KEY: yourAPIkey'
```
#### 📊 Структуры ответа
```go
type ServiceGetResponse struct {
    Data ServiceGetData `json:"data"`
}

type ServiceGetData struct {
    Items []ServiceVPS `json:"items"`
    Total int          `json:"total"`
}

type ServiceVPS struct {
    ID           int64       `json:"id"`
    Name         string      `json:"name"`
    IP           string      `json:"ip"`
    PaymentTerm  string      `json:"paymentTerm"`
    AutoProlong  bool        `json:"autoProlong"`
    Status       string      `json:"status"`
    Parameters   Parameters  `json:"parameters"`
    Product      Product     `json:"product"`
     ... 
}
```
[*Полная структура Service](../legacy/services.md)

### 📋 **Основные поля ответа**

| Поле | Тип | Описание | Пример |
|------|-----|-----------|---------|
| `id` | int64 | Уникальный ID услуги | `1576759` |
| `name` | string | Название услуги | `"my-vps"` |
| `ip` | string | Основной IP-адрес | `"85.192.29.68"` |
| `paymentTerm` | string | Период оплаты | `"hour"`, `"month"`, `"year"` |
| `autoProlong` | bool | Автопродление | `true`, `false` |
| `status` | string | Статус услуги | `"active"`, `"prolong_wait"`, `"activation_wait"` |
| `parameters.os` | string | Операционная система | `"ubuntu_2404"`, `"debian_12"` |
| `parameters.isoUrl` | string | URL ISO образа | `""` |
| `parameters.recipe` | *string | Рецепт настройки | `"docker"`, `null` |
| `product.id` | int64 | ID продукта | `182` |
| `product.name` | string | Название продукта | `"NLs-1"` |
| `product.type` | string | Тип продукта | `"vps"`, `"kubernetes"` |
| `locationCode` | string | Код локации | `"nl"`, `"de"` |
| `currentStatus` | string | Текущий статус отображения | `"prolong_wait"` |
| `lastStatus` | *string | Предыдущий статус | `"active"`, `null` |
| `createdAt` | int64 | Дата создания (Unix timestamp) | `1764245933` |

### **3. Обновление услуги PUT /api/services/{id}**
#### Legacy API

```bash
curl -X 'PUT' \
  'https://my.aeza.net/api/services/1576759' \
  -H 'accept: application/json' \
  -H 'X-API-KEY: yourAPIkey' \
  -H 'Content-Type: application/json' \
  -d '{
    "name": "updated-name",
    "autoProlong": true
  }'
```
#### 📊 Структуры запроса
```go
type ServiceUpdateRequest struct {
    Name        string `json:"name"`
    AutoProlong bool   `json:"autoProlong"`
```
### 📋 Обновляемые поля
| Поле | Тип | Описание | Пример |
|------|-----|-----------|---------|
| `name` | string | Новое название услуги | `"updated-vps-name"` |
| `autoProlong` | bool | Новый статус автопродления | `true`, `false` |

⚠️ Не обновляемые поля
- productId - требует пересоздания услуги

- os, recipe, isoUrl - параметры установки

- Все вычисляемые поля (status, ip, createdAt, etc.)

### **4. Удаление услуги DELETE /api/services/{id}**
#### Legacy API

```bash
curl -X 'DELETE' \
  'https://my.aeza.net/api/services/1576759' \
  -H 'accept: application/json' \
  -H 'X-API-KEY: yourAPIkey'
```
#### ✅ Response
```json
{
  "data": "ok"
}
```
### 🎯 **Статусы услуг**
#### 📊 Основные статусы

| Статус | Описание |
|--------|-----------|
| `activation_wait` | Ожидание активации услуги |
| `active` | Активная и работающая услуга |
| `prolong_wait` | Ожидание продления услуги |
| `performed` | Заказ выполнен (статус после создания) |
| `suspended` | Услуга приостановлена |
| `deleted` | Услуга удалена |

##### 🔄 Жизненный цикл
```text
creation → performed → activation_wait → active → prolong_wait → (renew) active
```

### ⚠️ **Ограничения и особенности**
#### 🔒 Безопасность
- Все запросы требуют X-API-KEY

- Доступ только к услугам текущего пользователя

- SecureParameters содержат чувствительные данные (пароли, ключи)

#### ⏱️ Таймауты
- Создание: 30-60 секунд (зависит от типа услуги)

- Активация: 1-5 минут для VPS

- Удаление: мгновенное

#### 🔄 Идемпотентность
- GET, PUT, DELETE - идемпотентные

- POST - не идемпотентный (создает новую услугу)

### 🛠️ **Terraform Провайдер**
#### 📝 Пример использования
```hcl
resource "aeza_service" "my_vps" {
  name          = "production-vps"
  product_id    = 182
  payment_term  = "month"
  auto_prolong  = true
  os            = "ubuntu_2404"
}

output "service_info" {
  value = {
    id        = aeza_service.my_vps.id
    name      = aeza_service.my_vps.name
    ip        = aeza_service.my_vps.ip
    status    = aeza_service.my_vps.status
    product   = aeza_service.my_vps.product_name
  }
}
```

### 🔗 **Связанные API**
- Продукты - Список доступных продуктов

- Группы услуг - Категории услуг

- Типы ОС - Доступные операционные системы

- VPS Действия - Управление VPS (старт/стоп/ребут)

Также доступен в новой V2 версии api/services*

___
<br>

**Примечание:** Legacy API является основным для управления услугами, *V2 API используется для дополнительных операций и мониторинга.
