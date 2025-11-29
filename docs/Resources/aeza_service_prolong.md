# **Resources**


### **aeza_service_prolong**
Управляет продлением услуг в аккаунте Aeza. Ресурс позволяет продлевать VPS, DNS, VPN и другие услуги* на указанный период с созданием транзакции.
<br>*- *в разработке*

### **Схема ресурса**
**Обязательные параметры**

| Поле | Тип | Обязательный | Описание |
|------|-----|-------------|-----------|
| `service_id` | number | ✅ | ID услуги для продления |
| `term` | string | ✅ | Период продления (hour, day, month, year) |
| `term_count` | number | ✅ | Количество периодов для продления |

**Опциональные параметры**

| Поле | Тип | Обязательный | Описание |
|------|-----|-------------|-----------|
| `method` | string | ❌ | Способ оплаты (balance, sms и т.д.). По умолчанию: "balance" |
| `force` | boolean | ❌ | Принудительное продление при каждом apply. По умолчанию: false |

**Вычисляемые поля (только для чтения)**

| Поле | Тип | Описание |
|------|-----|-----------|
| `id` | number | ID услуги (совпадает с service_id) |
| `transaction_id` | number | Уникальный идентификатор созданной транзакции |
| `transaction_amount` | string | Сумма транзакции в отформатированном виде |
| `transaction_status` | string | Статус транзакции (created, pending, completed) |
| `transaction_type` | string | Тип транзакции (prolong) |
| `transaction_created_at` | number | Timestamp создания транзакции |

#### **Примечания:**

- Ресурс создает финансовую транзакцию при каждом создании и изменении

- Поле force = true для принудительного продления услуг

- Все поля транзакции заполняются после успешного выполнения операции продления

Использует [Legacy API /api/services/{id}/prolong](../API/legacy/prolong.md) ендпоинт.

### **Примеры использования**
#### Базовое продление услуги на 1 час

```hcl
# Создание VPS
resource "aeza_service" "test_vps" {
  name        = "my-vps"
  product_id  = 182
  payment_term = "hour"
}

# Продление на 1 час через баланс
resource "aeza_service_prolong" "extend_hourly" {
  service_id = aeza_service.test_vps.id
  term       = "hour"
  term_count     = 1
  # method = "balance" (по умолчанию)
  # force  = false (по умолчанию)
}

# Информация о транзакции
output "prolong_info" {
  value = {
    transaction_id     = aeza_service_prolong.extend_hourly.transaction_id
    transaction_amount = aeza_service_prolong.extend_hourly.transaction_amount
    transaction_status = aeza_service_prolong.extend_hourly.transaction_status
  }
}
```

Продление на месяц с принудительным выполнением
```hcl
resource "aeza_service_prolong" "extend_monthly" {
  service_id = aeza_service.production.id
  term       = "month"
  term_count      = 1
  method     = "balance"
  force      = true  # Всегда выполнять продление при apply
}
```

#### Продление на несколько периодов
```hcl
resource "aeza_service_prolong" "extend_multiple" {
  service_id = aeza_service.vps.id
  term       = "hour"
  term_count      = 4     # Продлить на 4 часа
}
```
#### Продление со сменой периода
```hcl
# Продление на 2 дня
resource "aeza_service_prolong" "daily" {
  service_id = aeza_service.test.id
  term       = "day"
  term_count      = 2
}
```

### **Ограничения**
- Ресурс не может быть импортирован (управляет операциями, а не состоянием)

- Изменение параметров продления приводит к созданию новой транзакции

- Удаление ресурса из конфигурации не отменяет выполненные транзакции

- Для отмены автопродления используйте ресурс aeza_service с auto_prolong = false