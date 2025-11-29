// internal/client/api.go

package client

import (
	"context"
	"fmt"

	"github.com/vbobroff-app/terraform-provider-aeza/internal/models"
	"github.com/vbobroff-app/terraform-provider-aeza/internal/utils"
)

func (c *Client) ListServices(ctx context.Context) ([]models.TerraformService, error) {
	// Пробуем сначала Legacy API
	legacyServices, err := c.ListServicesVPS_legacy(ctx)
	if err == nil && len(legacyServices) > 0 {
		return utils.ConvertLegacyServices(legacyServices), nil
	}

	// Fallback на API v2
	nextServices, err := c.ListServices_V2(ctx)
	if err != nil {
		return nil, err
	}

	return utils.ConvertNextServices(nextServices), nil
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

// CreateService создает услугу через Legacy API с конвертацией моделей
func (c *Client) CreateService(ctx context.Context, req models.ServiceCreateRequest) (*models.ServiceCreateResponse, error) {
	// Конвертируем Terraform модель в Legacy API запрос
	legacyReq := utils.ConvertToLegacy_ServiceCreateRequest(req)

	// Вызываем Legacy метод
	legacyResp, err := c.CreateService_legacy(ctx, legacyReq)
	if err != nil {
		return nil, fmt.Errorf("failed to create service via legacy API: %w", err)
	}

	// Конвертируем Legacy API ответ в Terraform модель
	terraformResp := utils.ConvertFromLegacy_ServiceCreateResponse(*legacyResp)

	return &terraformResp, nil
}

func (c *Client) DeleteService(ctx context.Context, id int64) error {
	return c.DeleteService_legacy(ctx, id)
}

func (c *Client) GetService(ctx context.Context, id int64) (*models.Service, error) {

	// Вызываем legacy метод
	legacyResp, err := c.GetService_legacy(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get service: %w", err)
	}

	// Проверяем что есть items
	if len(legacyResp.Data.Items) == 0 {
		return nil, fmt.Errorf("service with ID %d not found", id)
	}

	legacyService := legacyResp.Data.Items[0]
	terraformService := utils.ConvertLegacyServiceGetToTerraform(legacyService)

	return &terraformService, nil
}

func (c *Client) UpdateService(ctx context.Context, id int64, req models.ServiceUpdateRequest) error {
	legacyReq := utils.ConvertToLegacy_ServiceUpdateRequest(req)
	return c.UpdateService_legacy(ctx, id, legacyReq)
}

// ProlongService продлевает услугу
func (c *Client) ProlongService(ctx context.Context, serviceID int64, req models.ServiceProlongRequest) (*models.ServiceProlongResponse, error) {

	legacyReq := utils.ConvertToLegacy_ServiceProlongRequest(req)

	// 2. Вызываем legacy метод (в будущем может быть V2 метод)
	legacyResp, err := c.ProlongService_Legacy(ctx, serviceID, legacyReq)
	if err != nil {
		return nil, err
	}

	resp := utils.ConvertFromLegacy_ServiceProlongResponse(legacyResp)

	return resp, nil
}

func (c *Client) ControlService(ctx context.Context, serviceID int64, action string) error {

	return c.ControlService_V2(ctx, serviceID, action)
}
