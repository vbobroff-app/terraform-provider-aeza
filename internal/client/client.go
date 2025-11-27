// internal/client/client.go
package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/vbobroff-app/terraform-provider-aeza/internal/models"
)

type Client struct {
	host       string
	apiKey     string
	httpClient *http.Client
}

func NewClient(baseUrl, apiKey string) (*Client, error) {
	return &Client{
		host:       baseUrl,
		apiKey:     apiKey,
		httpClient: &http.Client{},
	}, nil
}

// Resource methods
// CreateService creates a new service in Aeza
// Parameters:
//   - ctx: context for cancellation and timeouts
//   - req: service creation request with name, product_id, etc.
//
// Returns:
//   - *ServiceCreateResponse: created service details
//   - error: any error that occurred during creation
// func (c *Client) CreateService(ctx context.Context, req models.ServiceCreateRequest) (*models.ServiceCreateResponse, error) {
// 	var response models.ServiceCreateResponse

// 	err := c.NewRequest("POST", "/services", req).Do(ctx, &response)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &response, nil
// }

// func (c *Client) GetService(ctx context.Context, id int64) (*models.ServiceGetResponse, error) {

// 	// ВРЕМЕННО: логируем какой ID запрашивается
// 	log.Printf("[DEBUG] !!!GetService called with ID: %d", id)

// 	var response models.ServiceGetResponse
// 	err := c.NewRequest("GET", fmt.Sprintf("/services/%d", id), nil).Do(ctx, &response)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &response, nil
// }

func (c *Client) UpdateService(ctx context.Context, id int64, req models.ServiceCreateRequest) error {
	return c.NewRequest("PUT", fmt.Sprintf("/services/%d", id), req).Do(ctx, nil)
}

func (r *Request) AddQueryParam(key, value string) *Request {
	r.queryParams[key] = value
	return r
}

func (r *Request) SetQueryParams(params map[string]string) *Request {
	for key, value := range params {
		r.queryParams[key] = value
	}
	return r
}
