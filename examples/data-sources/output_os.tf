# Вывод всех операционных систем
output "all_operating_systems" {
  value = data.aeza_os_list.available.os_list
}

# # Вывод только Linux систем
# output "linux_operating_systems" {
#   value = [
#     for os in data.aeza_os_list.available.os_list : os
#     if os.group != "windows_server"
#   ]
# }

# # Вывод только Windows систем
# output "windows_operating_systems" {
#   value = [
#     for os in data.aeza_os_list.available.os_list : os
#     if os.group == "windows_server"
#   ]
# }

# # Вывод списка имен ОС
# output "os_names" {
#   value = [for os in data.aeza_os_list.available.os_list : os.name]
# }

# # Вывод ОС по группам
# output "os_by_groups" {
#   value = {
#     for os in data.aeza_os_list.available.os_list : os.group => os.name...
#   }
# }

# # Поиск конкретной ОС (например, Ubuntu 22.04)
# output "ubuntu_2204" {
#   value = [
#     for os in data.aeza_os_list.available.os_list : os
#     if os.name == "Ubuntu 22.04"
#   ]
# }