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
