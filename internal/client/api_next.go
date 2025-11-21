// internal/client/api_next.go

package client

import (
	"context"

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
