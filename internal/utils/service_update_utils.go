// internal/utils/service_update_utils.go
package utils

import (
	"github.com/vbobroff-app/terraform-provider-aeza/internal/models"
	"github.com/vbobroff-app/terraform-provider-aeza/internal/models/legacy"
)

func ConvertToLegacy_ServiceUpdateRequest(req models.ServiceUpdateRequest) legacy.ServiceUpdateRequest {
	return legacy.ServiceUpdateRequest{
		Name:        req.Name,
		AutoProlong: req.AutoProlong,
	}
}
