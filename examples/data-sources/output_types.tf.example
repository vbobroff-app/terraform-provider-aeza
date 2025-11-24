output "all_service_types" {
  description = "List of all available Aeza service types"
  value       = data.aeza_service_types.all.service_types
}

output "service_types_sorted" {
  description = "Service types sorted by order"
  value = [for st in data.aeza_service_types.all.service_types : {
    slug  = st.slug
    name  = st.name
    order = st.order
  }]
}

output "service_type_slugs" {
  description = "List of service type slugs"
  value = [
    for st in data.aeza_service_types.all.service_types : st.slug
  ]
}

output "service_type_names" {
  description = "List of service type names"
  value = [
    for st in data.aeza_service_types.all.service_types : st.name
  ]
}

output "service_types_map" {
  description = "Service types as slug->name mapping"
  value = {
    for st in data.aeza_service_types.all.service_types :
    st.slug => st.name
  }
}