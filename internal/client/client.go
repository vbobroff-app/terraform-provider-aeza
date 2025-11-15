package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/vbobroff-app/terraform-provider-aeza/internal/source-models"
)

type Client struct {
	host       string
	token      string
	httpClient *http.Client
}

func NewClient(host, token string) (*Client, error) {
	return &Client{
		host:       host,
		token:      token,
		httpClient: &http.Client{},
	}, nil
}

// DataClient methods
func (c *Client) ListServices(ctx context.Context) ([]source_models.Service, error) {
	var response source_models.ListServicesResponse
	err := c.NewRequest("GET", "/services", nil).Do(ctx, &response)
	if err != nil {
		return nil, err
	}
	return response.Services, nil
}

func (c *Client) ListProducts(ctx context.Context) ([]source_models.Product, error) {
	var response source_models.ListProductsResponse
	err := c.NewRequest("GET", "/products", nil).Do(ctx, &response)
	if err != nil {
		return nil, err
	}
	return response.Products, nil
}

func (c *Client) ListServiceTypes(ctx context.Context) ([]source_models.ServiceType, error) {
	var response source_models.ListServiceTypesResponse
	err := c.NewRequest("GET", "/service-types", nil).Do(ctx, &response)
	if err != nil {
		return nil, err
	}
	return response.ServiceTypes, nil
}

// Resource methods
func (c *Client) CreateService(ctx context.Context, req source_models.ServiceCreateRequest) (*source_models.ServiceCreateResponse, error) {
	var response source_models.ServiceCreateResponse
	err := c.NewRequest("POST", "/services", req).Do(ctx, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetService(ctx context.Context, id int64) (*source_models.ServiceGetResponse, error) {
	var response source_models.ServiceGetResponse
	err := c.NewRequest("GET", fmt.Sprintf("/services/%d", id), nil).Do(ctx, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) UpdateService(ctx context.Context, id int64, req source_models.ServiceCreateRequest) error {
	return c.NewRequest("PUT", fmt.Sprintf("/services/%d", id), req).Do(ctx, nil)
}

func (c *Client) DeleteService(ctx context.Context, id int64) error {
	return c.NewRequest("DELETE", fmt.Sprintf("/services/%d", id), nil).Do(ctx, nil)
}
