// internal/client/api_next.go

package client

import (
	"context"

	"github.com/vbobroff-app/terraform-provider-aeza/internal/models"
	"github.com/vbobroff-app/terraform-provider-aeza/internal/utils"
)

func (c *Client) ListServices(ctx context.Context) ([]models.TerraformService, error) {
	// Пробуем сначала API v2 (более быстрый и простой)
	nextServices, err := c.ListServices_V2(ctx)
	if err == nil && len(nextServices) > 0 {
		return utils.ConvertNextServices(nextServices), nil
	}

	// Fallback на Legacy API
	legacyServices, err := c.ListServicesVPS_legacy(ctx)
	if err != nil {
		return nil, err
	}

	return utils.ConvertLegacyServices(legacyServices), nil
}

func (c *Client) ListServiceTypes(ctx context.Context) ([]models.ServiceType, error) {

	legacyTypes, err := c.ListServiceTypes_legacy(ctx)
	if err != nil {
		return nil, err
	}

	// Конвертируем legacy.ServiceType в models.ServiceType
	var result []models.ServiceType
	for _, legacyType := range legacyTypes {
		result = append(result, utils.ConvertLegacyServiceType(legacyType))
	}

	return result, nil
}

// ListServiceGroups получает список групп услуг (основной метод)
func (c *Client) ListServiceGroups(ctx context.Context, serviceType string) ([]models.ServiceGroup, error) {
	// Приоритет Legacy API
	legacyGroups, err := c.ListServiceGroups_Legacy(ctx, serviceType)
	if err == nil && len(legacyGroups) > 0 {
		return utils.ConvertLegacyServiceGroups(legacyGroups), nil
	}

	// Fallback на API v2
	nextGroups, err := c.ListServiceGroups_V2(ctx, serviceType)
	if err != nil {
		return nil, err
	}

	// Конвертируем в Terraform модели
	return utils.ConvertNextServiceGroups(nextGroups), nil
}

func (c *Client) ListProducts(ctx context.Context) ([]models.Product, error) {
	legacyProducts, err := c.ListProducts_Legacy(ctx)
	if err != nil {
		return nil, err
	}

	// Конвертируем legacy.Product в models.Product
	var result []models.Product
	for _, legacyProduct := range legacyProducts {
		result = append(result, utils.ConvertLegacyProduct(legacyProduct))
	}

	return result, nil
}

func (c *Client) ListOS(ctx context.Context) ([]models.OperatingSystem, error) {
	// Сначала пытаемся получить данные из V2 API
	nextOS, err := c.ListOS_V2(ctx)
	if err == nil && len(nextOS) > 0 {
		var result []models.OperatingSystem
		for _, os := range nextOS {
			result = append(result, models.OperatingSystem{OperatingSystem: os})
		}
		return result, nil
	}

	// Если V2 не сработал, пробуем Legacy API
	legacyOS, err := c.ListOS_Legacy(ctx)
	if err != nil {
		return nil, err
	}

	var result []models.OperatingSystem
	for _, os := range legacyOS {
		result = append(result, models.OperatingSystem{OperatingSystem: utils.ConvertOsFromLegacy(os)})
	}

	return result, nil
}
