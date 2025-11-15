package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
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

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("API error: %s, body: %s", resp.Status, string(body))
	}

	return body, nil
}

// Реализуем методы интерфейса DataClient
func (c *Client) ListServices(ctx context.Context) ([]source_models.Service, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.host+"/services", nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var response source_models.ListServicesResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return response.Services, nil
}

func (c *Client) ListProducts(ctx context.Context) ([]source_models.Product, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.host+"/products", nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var response source_models.ListProductsResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return response.Products, nil
}

func (c *Client) ListServiceTypes(ctx context.Context) ([]source_models.ServiceType, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.host+"/service-types", nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var response source_models.ListServiceTypesResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return response.ServiceTypes, nil
}

// Resource methods
func (c *Client) CreateService(ctx context.Context, req source_models.ServiceCreateRequest) (*source_models.ServiceCreateResponse, error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", c.host+"/services", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(httpReq)
	if err != nil {
		return nil, err
	}

	var response source_models.ServiceCreateResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetService(ctx context.Context, id int64) (*source_models.ServiceGetResponse, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/services/%d", c.host, id), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var response source_models.ServiceGetResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) UpdateService(ctx context.Context, id int64, req source_models.ServiceCreateRequest) error {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return err
	}

	httpReq, err := http.NewRequestWithContext(ctx, "PUT", fmt.Sprintf("%s/services/%d", c.host, id), bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	_, err = c.doRequest(httpReq)
	return err
}

func (c *Client) DeleteService(ctx context.Context, id int64) error {
	req, err := http.NewRequestWithContext(ctx, "DELETE", fmt.Sprintf("%s/services/%d", c.host, id), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	return err
}
