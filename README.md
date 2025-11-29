# Terraform Provider для Aeza

Провайдер Terraform для управления услугами хостинг-провайдера [Aeza](https://aeza.net/) через API.

## Требования

- [Terraform](https://www.terraform.io/downloads.html) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.19 (при компиляции из исходного кода)

## Установка

### Автоматическая установка

Добавить провайдер в конфигурацию Terraform:

```hcl
terraform {
  required_providers {
    aeza = {
      source = "vbobroff-app/aeza"
      version = "0.1.0"
    }
  }
}

provider "aeza" {
  api_key = var.aeza_api_key
  base_url = var.aeza_base_url
}
```

### Ручная установка
Скачайте бинарный файл провайдера и разместите его в директории плагинов Terraform:

```bash
mkdir -p ~/.terraform.d/plugins/
cp terraform-provider-aeza ~/.terraform.d/plugins/
```

### Настройка провайдера

[Настройка локального провайдера для разработки](/docs/test-init.md)

Провайдер требует аутентификацию через API токен (apiKey)


### Получение API токена
1. Войдите в панель управления Aeza

2. Перейдите в раздел API

3. Сгенерируйте новый API токен

## Доступные ресурсы 

### **Data Sources** 

[`aeza_products`](/docs/Data-sources/aeza_products.md) - Получение списка доступных продуктов

[`aeza_services`](/docs/Data-sources/aeza_services.md) - Получение списка активных услуг

[`aeza_os`](/docs/Data-sources/aeza_os.md) - Получение списка операционных систем

[`aeza_service_groups`](/docs/Data-sources/aeza_service_groups.md) - Получение групп услуг

[`aeza_service_types`](/docs/Data-sources/aeza_service_types.md) - Получение доступных типов услуг



### **Resources**
[`aeza_service`](/docs/Resources/aeza_services.md) - Создание и управление услугой (продуктом)

[`aeza_service_prolong`](/docs/Resources/aeza_service_prolong.md) - Продление и изменение периода оплаты услуги

[`aeza_service_actions`](/docs/Resources/aeza_servise_actions.md) - Управление действиями: включение, выключение, перезагрузка

