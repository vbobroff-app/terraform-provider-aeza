package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Request представляет HTTP запрос
type Request struct {
	client *Client
	method string
	path   string
	body   interface{}
}

// NewRequest создает новый запрос
func (c *Client) NewRequest(method, path string, body interface{}) *Request {
	return &Request{
		client: c,
		method: method,
		path:   path,
		body:   body,
	}
}

// Do выполняет запрос
func (r *Request) Do(ctx context.Context, result interface{}) error {
	var bodyReader io.Reader
	if r.body != nil {
		jsonData, err := json.Marshal(r.body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequestWithContext(ctx, r.method, r.client.host+r.path, bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+r.client.token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := r.client.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("API error %d: %s", resp.StatusCode, string(body))
	}

	if result != nil {
		if err := json.Unmarshal(body, result); err != nil {
			return fmt.Errorf("failed to unmarshal response: %w", err)
		}
	}

	return nil
}
