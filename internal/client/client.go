// internal/client/client.go
package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/vbobroff-app/terraform-provider-aeza/internal/models"
	"github.com/vbobroff-app/terraform-provider-aeza/internal/utils"
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

// func (c *Client) ListServices(ctx context.Context) ([]models.Service, error) {
// 	var response models.ListServicesResponse
// 	err := c.NewRequest("GET", "/services", nil).Do(ctx, &response)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return response.Data.Items, nil
// }

func (c *Client) ListProducts(ctx context.Context) ([]models.Product, error) {
	var response models.ListProductsResponse
	err := c.NewRequest("GET", "/services/products", nil).Do(ctx, &response)
	if err != nil {
		return nil, err
	}
	return response.Data.Items, nil // Теперь берем из Data.Items
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
func (c *Client) CreateService(ctx context.Context, req models.ServiceCreateRequest) (*models.ServiceCreateResponse, error) {
	var response models.ServiceCreateResponse

	// Добавляем логирование запроса
	fmt.Printf("DEBUG: API CreateService request: %+v\n", req)

	// ... существующий код ...

	err := c.NewRequest("POST", "/services", req).Do(ctx, &response)
	if err != nil {
		return nil, err
	}

	// Добавляем логирование ответа
	fmt.Printf("DEBUG: API CreateService response: %+v\n", response)

	return &response, nil
}

func (c *Client) GetService(ctx context.Context, id int64) (*models.ServiceGetResponse, error) {
	var response models.ServiceGetResponse
	err := c.NewRequest("GET", fmt.Sprintf("/services/%d", id), nil).Do(ctx, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) UpdateService(ctx context.Context, id int64, req models.ServiceCreateRequest) error {
	return c.NewRequest("PUT", fmt.Sprintf("/services/%d", id), req).Do(ctx, nil)
}

func (c *Client) DeleteService(ctx context.Context, id int64) error {
	return c.NewRequest("DELETE", fmt.Sprintf("/services/%d", id), nil).Do(ctx, nil)
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

// ListServiceGroups получает список групп услуг (основной метод)
func (c *Client) ListServiceGroups(ctx context.Context, serviceType string) ([]models.ServiceGroup, error) {
	// Используем API v2
	nextGroups, err := c.ListServiceGroups_V2(ctx, serviceType)
	if err != nil {
		return nil, err
	}

	// Конвертируем в Terraform модели
	return utils.ConvertNextServiceGroups(nextGroups), nil
}
