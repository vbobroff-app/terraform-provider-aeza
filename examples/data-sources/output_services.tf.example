output "all_services" {
  description = "List of all Aeza services"
  value       = data.aeza_services.all.services
}

output "services_summary" {
  description = "Summary of services by type"
  value = {
    for service in data.aeza_services.all.services :
    service.type_slug => service.name...
  }
}

output "services_by_status" {
  description = "Services grouped by status"
  value = {
    for service in data.aeza_services.all.services :
    service.status => {
      id      = service.id
      name    = service.name
      type    = service.type_slug
      ip      = service.ip
      product = service.product_name
    }...
  }
}

output "active_services" {
  description = "List of active services"
  value = [
    for service in data.aeza_services.all.services :
    service if service.status == "active"
  ]
}

output "services_with_ips" {
  description = "Services with IP addresses"
  value = [
    for service in data.aeza_services.all.services :
    {
      id       = service.id
      name     = service.name
      type     = service.type_slug
      ip       = service.ip
      location = service.location_code
      price    = service.price_display
    } if service.ip != ""
  ]
}

output "services_pricing_summary" {
  description = "Pricing summary of services"
  value = [
    for service in data.aeza_services.all.services :
    {
      id           = service.id
      name         = service.name
      type         = service.type_slug
      price_raw    = service.price_raw
      price_display = service.price_display
      payment_term = service.payment_term
      auto_prolong = service.auto_prolong
      expires_at   = service.expires_at
    }
  ]
}

output "services_by_location" {
  description = "Services grouped by location"
  value = {
    for service in data.aeza_services.all.services :
    service.location_code => {
      id   = service.id
      name = service.name
      type = service.type_slug
    }...
  }
}