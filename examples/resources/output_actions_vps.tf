output "prolong_transaction" {
  description = "Prolong transaction details"
  value = aeza_service_actions.prolong_service.prolong[0]
}

output "service_actions_id" {
  description = "Service actions resource ID"
  value       = aeza_service_actions.prolong_service.id
}