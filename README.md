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
Провайдер требует аутентификацию через API токен (apiKey)


### Получение API токена
1. Войдите в панель управления Aeza

2. Перейдите в раздел API

3. Сгенерируйте новый API токен

## Доступные ресурсы 

### **Data Sources** 

`aeza_products` - Получение списка доступных продуктов

`aeza_services` - Получение списка активных [услуг](/docs/services.md)

`aeza_service_groups` - Получение [групп услуг](/docs/groups.md)

`aeza_service_types` - Получение доступных [типов услуг](/docs/service-types.md) 

[Спецификации здесь](/docs/data-sources.md)

### **Resources**
В разработке

