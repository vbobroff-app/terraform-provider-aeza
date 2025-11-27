// internal/client/api_legacy.go
package client

import (
	"context"
	"fmt"

	"github.com/vbobroff-app/terraform-provider-aeza/internal/models/legacy"
)

// ListServicesBase_legacy - получение списка или конкретной услуги (базовый тип)
func (c *Client) ListServicesBase_legacy(ctx context.Context, serviceID ...int) ([]legacy.ServiceBase, error) {
	var response legacy.ListServicesResponse

	url := "/services"
	if len(serviceID) > 0 {
		url = fmt.Sprintf("/services/%d", serviceID[0])
	}

	err := c.NewRequest("GET", url, nil).Do(ctx, &response)
	if err != nil {
		return nil, err
	}

	return response.Data.Items, nil
}

// ListServicesVPS_legacy - получение списка или конкретной VPS услуги
func (c *Client) ListServicesVPS_legacy(ctx context.Context, serviceID ...int) ([]legacy.ServiceVPS, error) {
	var response legacy.ListServicesVPSResponse

	url := "/services"
	if len(serviceID) > 0 {
		url = fmt.Sprintf("/services/%d", serviceID[0])
	}

	err := c.NewRequest("GET", url, nil).Do(ctx, &response)
	if err != nil {
		return nil, err
	}

	return response.Data.Items, nil
}

func (c *Client) ListServiceTypes_legacy(ctx context.Context) ([]legacy.ServiceType, error) {
	var response legacy.ListServiceTypesResponse
	err := c.NewRequest("GET", "/services/types", nil).Do(ctx, &response)
	if err != nil {
		return nil, err
	}
	return response.Data.Items, nil // Теперь берем из Data.Items
}

// ListServiceGroups_Legacy получает список групп услуг через Legacy API
func (c *Client) ListServiceGroups_Legacy(ctx context.Context, serviceType string) ([]legacy.ServiceGroup, error) {
	var response legacy.ServiceGroupsResponse

	req := c.NewRequest("GET", "/services/groups", nil)
	req.AddQueryParam("extra", "true")

	if serviceType != "" {
		req.AddQueryParam("type", serviceType)
	}

	err := req.Do(ctx, &response)
	if err != nil {
		return nil, err
	}

	return response.Data.Items, nil
}

func (c *Client) ListProducts_Legacy(ctx context.Context) ([]legacy.Product, error) {
	var response legacy.ListProductsResponse
	err := c.NewRequest("GET", "/services/products", nil).Do(ctx, &response)
	if err != nil {
		return nil, err
	}
	return response.Data.Items, nil
}

func (c *Client) ListOS_Legacy(ctx context.Context) ([]legacy.OperatingSystem, error) {
	var response legacy.OSResponse
	err := c.NewRequest("GET", "/os", nil).Do(ctx, &response)
	if err != nil {
		return nil, err
	}
	return response.Data.Items, nil
}

// CreateService_legacy создает услугу через Legacy API (существующий метод)
func (c *Client) CreateService_legacy(ctx context.Context, req legacy.ServiceCreateRequest) (*legacy.ServiceCreateResponse, error) {
	// ✅ ПРАВИЛЬНО: передаем структуру, а не закодированный JSON
	var response legacy.ServiceCreateResponse
	err := c.NewRequest("POST", "/services/orders", req).Do(ctx, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to create service create request: %w", err)
	}

	// Проверяем, что заказ создан успешно
	if len(response.Data.Items) == 0 {
		return nil, fmt.Errorf("no items in legacy response, service creation failed")
	}

	// Проверяем статус созданного заказа
	item := response.Data.Items[0]
	if item.Status != "performed" {
		return nil, fmt.Errorf("legacy service order status is not 'performed': %s", item.Status)
	}

	// Проверяем, что есть ID созданных услуг
	if len(item.CreatedServiceIds) == 0 {
		return nil, fmt.Errorf("no service IDs in legacy response, service creation failed")
	}

	return &response, nil
}
