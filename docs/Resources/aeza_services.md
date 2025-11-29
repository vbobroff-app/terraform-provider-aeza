# **Resources**

### **aeza_service**

Создает и управляет услугами в аккаунте Aeza. Ресурс позволяет развертывать **VPS**, DNS*, VPN* и другие услуги*, а также управлять их базовыми параметрами.
<br>
  *- *в разработке*

### **Схема ресурса**
**Обязательные параметры**

| Поле | Тип | Обязательный | Описание |
|------|-----|-------------|-----------|
| `name` | string | ✅ | Название услуги |
| `product_id` | number | ✅ | ID продукта из каталога услуг |

**Опциональные параметры**

| Поле | Тип | Обязательный | Описание |
|------|-----|-------------|-----------|
| `payment_term` | string | ❌ | Период оплаты (hour, day, month, year). По умолчанию: "month" |
| `auto_prolong` | boolean | ❌ | Автоматическое продление услуги. По умолчанию: true |
| `os` | string | ❌ | Операционная система для VPS (только при создании) |
| `recipe` | string | ❌ | Рецепт установки для VPS (только при создании) |
| `iso_url` | string | ❌ | URL ISO образа для установки (только при создании) |

**Вычисляемые поля (только для чтения)**

| Поле | Тип | Описание |
|------|-----|-----------|
| `id` | number | Уникальный идентификатор услуги |
| `order_id` | number | ID заказа |
| `date` | string | Дата создания заказа |
| `product_type` | string | Тип продукта |
| `group_id` | number | ID группы услуг |
| `product_name` | string | Название продукта |
| `location_name` | string | Название локации |
| `price` | string | Цена услуги |
| `transaction_amount` | string | Сумма транзакции |
| `status` | string | Статус услуги (activation_wait, active и т.д.) |
| `ip` | string | IP-адрес услуги |
| `created_at` | string | Дата создания услуги |
| `updated_at` | string | Дата последнего обновления услуги |
| `expires_at` | string | Дата истечения срока действия услуги |
| `purchased_at` | string | Дата покупки услуги |


**Примечания:**

- Поля os, recipe, iso_url можно установить только при создании услуги (ForceNew)

- Поле payment_term можно установить только при создании, для изменения используйте ресурс aeza_vps_actions

- Поля `name` и `auto_prolong` можно изменять после создания

Использует [Legacy API /api/services](../API/legacy/services_crud.md) ендпоинты.

### **Примеры использования**
#### Создание VPS с почасовой оплатой

```hcl
resource "aeza_service" "test_vps" {
  # Обязательные параметры
  name       = "my-test-vps"
  product_id = 182
  
  # Опциональные параметры
  payment_term = "hour"
  auto_prolong = false
  os          = "ubuntu_2404"
}

# Использование созданной услуги
output "vps_info" {
  value = {
    id      = aeza_service.test_vps.id
    ip      = aeza_service.test_vps.ip
    status  = aeza_service.test_vps.status
    price   = aeza_service.test_vps.price
  }
}
```

#### Создание услуги с ISO установкой
```hcl
resource "aeza_service" "custom_iso_vps" {
  name        = "custom-iso-install"
  product_id  = 183
  payment_term = "month"
  auto_prolong = true
  iso_url     = "https://example.com/custom.iso"
}
```
#### Обновление имени услуги
```hcl
resource "aeza_service" "production_vps" {
  name        = "production-server-updated"  # Можно изменить
  product_id  = 182                          # Нельзя изменить
  payment_term = "month"                     # Нельзя изменить через этот ресурс
  auto_prolong = true                        # Можно изменить
  os          = "debian_12"                  # Нельзя изменить
}
```

#### Использование с data source для поиска продукта
```hcl
# Поиск нужного продукта
data "aeza_products" "all" {}

locals {
  # Находим продукт по имени
  target_product = [
    for product in data.aeza_products.all.products :
    product if contains(product.name, "VPS-PRO")
  ][0]
}

# Создание услуги на основе найденного продукта
resource "aeza_service" "pro_vps" {
  name        = "professional-vps"
  product_id  = local.target_product.id
  payment_term = "month"
  auto_prolong = true
  os          = "centos_9"
}
```
#### Импорт существующих услуг
```bash
terraform import aeza_service.example 12345
```
Где 12345 - ID существующей услуги.

### **Ограничения**
- Изменение product_id, os, recipe, iso_url требует пересоздания услуги

- Изменение payment_term осуществляется через отдельный ресурс aeza_vps_actions

- Удаление услуги через Terraform приведет к фактическому удалению услуги в Aeza

