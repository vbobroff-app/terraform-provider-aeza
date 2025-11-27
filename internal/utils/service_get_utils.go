// internal/utils/service_get_utils.go
package utils

import (
	"github.com/vbobroff-app/terraform-provider-aeza/internal/models"
	"github.com/vbobroff-app/terraform-provider-aeza/internal/models/legacy"
)

func ConvertLegacyServiceGetToTerraform(legacyService legacy.ServiceGet) models.Service {
	ip := legacyService.IP
	if ip == "" && len(legacyService.IPs) > 0 {
		ip = legacyService.IPs[0].Value
	}

	// Создаем map для Parameters с учетом разных типов услуг
	parameters := make(map[string]interface{})

	// Для VPS услуг добавляем стандартные поля
	parameters["os"] = legacyService.Parameters.OS
	parameters["isoUrl"] = legacyService.Parameters.ISOURL
	if legacyService.Parameters.Recipe != nil {
		parameters["recipe"] = *legacyService.Parameters.Recipe
	}

	// ✅ Можно добавить другие поля для разных типов услуг в будущем
	// if legacyService.Product.Type == "kubernetes" {
	//     parameters["kubernetes_version"] = legacyService.KubernetesVersion
	// }

	createdAt := FormatDateFromUnix(legacyService.Timestamps.CreatedAt)

	return models.Service{
		ID:            int64(legacyService.ID),
		OwnerID:       int64(legacyService.OwnerID),
		ProductID:     int64(legacyService.ProductID),
		Name:          legacyService.Name,
		IP:            ip,
		PaymentTerm:   legacyService.PaymentTerm,
		Parameters:    parameters, // Гибкий map для разных типов услуг
		AutoProlong:   legacyService.AutoProlong,
		Backups:       legacyService.Backups,
		Status:        legacyService.Status,
		LastStatus:    stringToEmpty(legacyService.LastStatus),
		Product:       ConvertLegacyProduct(legacyService.Product),
		LocationCode:  legacyService.LocationCode,
		CurrentStatus: legacyService.CurrentStatus,
		CreatedAt:     createdAt,
		UpdatedAt:     "",
	}
}

func stringToEmpty(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
