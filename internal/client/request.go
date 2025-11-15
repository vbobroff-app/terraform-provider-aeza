package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Request struct {
	client  *Client
	ctx     context.Context
	method  string
	path    string
	body    interface{}
	headers map[string]string
}

func (r *Request) SetMethod(method string) *Request {
	r.method = strings.ToUpper(method)
	return r
}

func (r *Request) SetPath(path string) *Request {
	r.path = path
	return r
}

func (r *Request) SetBody(body interface{}) *Request {
	r.body = body
	return r
}

func (r *Request) SetHeader(key, value string) *Request {
	if r.headers == nil {
		r.headers = make(map[string]string)
	}
	r.headers[key] = value
	return r
}

func (r *Request) Do(v interface{}) error {
	var bodyReader io.Reader
	if r.body != nil {
		jsonData, err := json.Marshal(r.body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(jsonData)
	}

	req, err := http.NewRequestWithContext(r.ctx, r.method, r.client.baseURL+r.path, bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Set default headers
	req.Header.Set("X-API-Key", r.client.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Set custom headers
	for key, value := range r.headers {
		req.Header.Set(key, value)
	}

	resp, err := r.client.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for API errors
	if resp.StatusCode >= 400 {
		var errorResp ErrorResponse
		if err := json.Unmarshal(body, &errorResp); err == nil && errorResp.Error.Message != "" {
			return &APIError{
				StatusCode: resp.StatusCode,
				Slug:       errorResp.Error.Slug,
				Message:    errorResp.Error.Message,
				Data:       errorResp.Error.Data,
			}
		}
		return fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	// Parse successful response
	if v != nil {
		if err := json.Unmarshal(body, v); err != nil {
			return fmt.Errorf("failed to unmarshal response: %w, body: %s", err, string(body))
		}
	}

	return nil
}
