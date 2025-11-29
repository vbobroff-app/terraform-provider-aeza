// internal/client/api_next.go
package client

import (
	"context"
	"fmt"

	"github.com/vbobroff-app/terraform-provider-aeza/internal/models/next"
)

func (c *Client) ListServices_V2(ctx context.Context) ([]next.Service, error) {
	var response next.ListServicesResponse
	err := c.NewRequest("GET", "/v2/services", nil).Do(ctx, &response)
	if err != nil {
		return nil, err
	}

	return response.Items, nil
}

// ListServiceGroups_V2 получает список групп услуг через API v2
func (c *Client) ListServiceGroups_V2(ctx context.Context, serviceType string) ([]next.ServiceGroup, error) {
	var response next.ServiceGroupsResponse

	req := c.NewRequest("GET", "/v2/services/groups", nil)
	if serviceType != "" {
		req.AddQueryParam("type", serviceType)
	}

	err := req.Do(ctx, &response)
	if err != nil {
		return nil, err
	}

	return response.Items, nil
}

func (c *Client) ListOS_V2(ctx context.Context) ([]next.OperatingSystem, error) {
	var response []next.OperatingSystem
	err := c.NewRequest("GET", "/v2/services/operating-systems", nil).Do(ctx, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ControlService_V2(ctx context.Context, serviceID int64, action string) error {
	// Проверяем допустимые действия
	allowedActions := map[string]bool{
		"suspend": true,
		"resume":  true,
		"restart": true,
	}

	if !allowedActions[action] {
		return fmt.Errorf("invalid service control action: %s", action)
	}

	url := fmt.Sprintf("/v2/services/%d/ctl/%s", serviceID, action)

	err := c.NewRequest("POST", url, nil).Do(ctx, nil)
	if err != nil {
		return fmt.Errorf("service control request failed for action %s: %w", action, err)
	}

	return nil
}
