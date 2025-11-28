// internal/utils/service_update_utils.go
package utils

import (
	"github.com/vbobroff-app/terraform-provider-aeza/internal/models"
	"github.com/vbobroff-app/terraform-provider-aeza/internal/models/legacy"
)

// ConvertToLegacy_ServiceUpdateRequest конвертирует Terraform модель в Legacy API запрос
func ConvertToLegacy_ServiceUpdateRequest(req models.ServiceUpdateRequest) legacy.ServiceUpdateRequest {
	return legacy.ServiceUpdateRequest{
		Name:        req.Name,
		PaymentTerm: req.PaymentTerm,
		AutoProlong: req.AutoProlong,
	}
}
