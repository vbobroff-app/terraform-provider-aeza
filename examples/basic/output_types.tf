output "all_service_types" {
  description = "List of all available Aeza service types"
  value       = data.aeza_service_types.all.service_types
}