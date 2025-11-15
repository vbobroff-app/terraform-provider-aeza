package client

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

func NewClient(baseURL, apiKey string) *Client {
	return &Client{
		baseURL: baseURL,
		apiKey:  apiKey,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:       10,
				IdleConnTimeout:    30 * time.Second,
				DisableCompression: false,
			},
		},
	}
}

func (c *Client) WithContext(ctx context.Context) *Request {
	return &Request{
		client:  c,
		ctx:     ctx,
		method:  "GET",
		path:    "",
		headers: make(map[string]string),
	}
}

// Базовые методы
func (c *Client) GetServices(ctx context.Context) (*ServicesResponse, error) {
	req := c.WithContext(ctx).SetPath("/services")
	var resp ServicesResponse
	return &resp, req.Do(&resp)
}

func (c *Client) GetProducts(ctx context.Context) (*ProductsResponse, error) {
	req := c.WithContext(ctx).SetPath("/services/products")
	var resp ProductsResponse
	return &resp, req.Do(&resp)
}

func (c *Client) GetServiceTypes(ctx context.Context) (*ServiceTypesResponse, error) {
	req := c.WithContext(ctx).SetPath("/services/types")
	var resp ServiceTypesResponse
	return &resp, req.Do(&resp)
}

func (c *Client) CreateService(ctx context.Context, data *CreateServiceRequest) (*ServiceResponse, error) {
	req := c.WithContext(ctx).SetMethod("POST").SetPath("/services").SetBody(data)
	var resp ServiceResponse
	return &resp, req.Do(&resp)
}

func (c *Client) DeleteService(ctx context.Context, serviceID string) error {
	req := c.WithContext(ctx).SetMethod("DELETE").SetPath(fmt.Sprintf("/services/%s", serviceID))
	return req.Do(nil)
}
