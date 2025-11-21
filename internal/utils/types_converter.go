// internal/utils/types_converter.go
package utils

import (
	"github.com/vbobroff-app/terraform-provider-aeza/internal/models"
	"github.com/vbobroff-app/terraform-provider-aeza/internal/models/legacy"
)

func ConvertLegacyServiceType(legacyType legacy.ServiceType) models.ServiceType {
	return models.ServiceType{
		Slug:    legacyType.Slug,
		Order:   legacyType.Order,
		Names:   legacyType.Names,
		Payload: legacyType.Payload,
		Name:    legacyType.Name,
		// Поля LocaledPayload и PrettyLocaledPayload игнорируются
		// так как их нет в models.ServiceType
	}
}
