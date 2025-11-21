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
