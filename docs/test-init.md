# **Настройка локального провайдера для разработки**

### **Шаг 1: Создаём тестовую директорию**

```bash
mkdir test-terraform
cd test-terraform
```

### **Шаг 2: Создаём структуру для локального провайдера**
```bash
mkdir -p .terraform/plugins/registry.terraform.io/vbobroff-app/aeza/0.1.0/linux_amd64
```
⚠️ Warning

Это стандартный путь для файловых зеркал Terraform, формат:

 `registry.terraform.io/<NAMESPACE>/<NAME>/<VERSION>/<ARCHITECTURE>`

Где:

- `registry.terraform.io` - хост registry по умолчанию

- `vbobroff-app` - namespace (владелец GitHub репозитория)

- `aeza` - имя провайдера (берётся из названия репозитория terraform-provider-aeza без префикса)

- `0.1.0` - версия провайдера

- `linux_amd64` - архитектура ОС

Terraform ищет провайдеры именно по такой структуре путей при использовании filesystem_mirror.



### **Шаг 3: Копируем собранный бинарник провайдера**
```bash
cp ../terraform-provider-aeza .terraform/plugins/registry.terraform.io/vbobroff-app/aeza/0.1.0/linux_amd64/
```
### **Шаг 4: Создаём конфигурацию Terraform CLI**
```bash
cat > ~/.terraformrc << EOF
provider_installation {
  filesystem_mirror {
    path = "/mnt/c/sources/terraform-provider-aeza/test-terraform/.terraform/plugins"
  }
  direct {
    exclude = ["vbobroff-app/aeza"]
  }
}
EOF
```

**Проверить создание файла:**

```bash
ls -la ~/.terraformrc
cat ~/.terraformrc
```
Файл скрытый (начинается с точки), находится в домашней директории. После создания он будет применяться ко всем Terraform проектам на этой машине.


### **Шаг 5: Создаём main.tf**
```hcl
# test-terraform/main.tf
terraform {
  required_providers {
    aeza = {
      source  = "vbobroff-app/aeza"
      version = "0.1.0"
    }
  }
}

provider "aeza" {
  api_key = "test_key_123"
}

data "aeza_services" "test" {}

output "test_output" {
  value = "Configuration syntax is valid"
}

```
>[!WARNING]  
>Важное уточнение: Имя vbobroff-app/aeza соответствует формату Terraform Registry:
> <namespace>/<name>
>Где:
> - namespace = владелец GitHub репозитория (vbobroff-app)
> - name = название провайдера (берётся из имени репозитория terraform-provider-aeza → aeza)
>
> Для нашего репозитория https://github.com/vbobroff-app/terraform-provider-aeza:
> - Владелец: vbobroff-app
> - Имя провайдера: aeza (terraform-provider-aeza)
>
> **Правило:** namespace всегда соответствует владельцу GitHub репозитория, name - части имени репозитория после terraform-provider-.
>
>Выполняем с vbobroff-app/aeza.

### **Шаг 6: Инициализируем Terraform**
```bash
terraform init
```
Результат: Успешное выполнение terraform init


*в случае отладки
```bash
rm -rf .terraform .terraform.lock.hcl
terraform init
```