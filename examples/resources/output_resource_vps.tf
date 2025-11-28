output "service_info" {
  value = {
    name               = aeza_service.test_vps.name
    id                 = aeza_service.test_vps.id
    order_id           = aeza_service.test_vps.order_id
    product_type       = aeza_service.test_vps.product_type
    product_name       = aeza_service.test_vps.product_name
    status             = aeza_service.test_vps.status
    ip                 = aeza_service.test_vps.ip
    price              = aeza_service.test_vps.price
    term               = aeza_service.test_vps.payment_term
    transaction_amount = aeza_service.test_vps.transaction_amount
    location_label     = aeza_service.test_vps.location_name
    group_id           = aeza_service.test_vps.group_id
    date               = aeza_service.test_vps.date
    auto_prolong       = aeza_service.test_vps.auto_prolong
  }
}
