// internal/client/request.go
package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"io"
	"net/http"
	"net/url"
)

// Request представляет HTTP запрос
type Request struct {
	client      *Client
	method      string
	path        string
	body        interface{}
	queryParams map[string]string
}

// NewRequest создает новый запрос
func (c *Client) NewRequest(method, path string, body interface{}) *Request {
	return &Request{
		client:      c,
		method:      method,
		path:        path,
		body:        body,
		queryParams: make(map[string]string),
	}
}

// Do выполняет запрос с полным логированием
func (r *Request) Do(ctx context.Context, result interface{}) error {
	start := time.Now()

	// Подготовка тела запроса
	var bodyReader io.Reader
	var requestBody []byte
	var err error

	if r.body != nil {
		requestBody, err = json.Marshal(r.body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewBuffer(requestBody)
	}

	// Формируем URL с query-параметрами
	fullURL := r.client.host + r.path
	if len(r.queryParams) > 0 {
		query := url.Values{}
		for key, value := range r.queryParams {
			query.Add(key, value)
		}
		fullURL = fullURL + "?" + query.Encode()
	}

	// Создаем запрос
	req, err := http.NewRequestWithContext(ctx, r.method, fullURL, bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Устанавливаем заголовки
	req.Header.Set("X-API-Key", r.client.apiKey)
	if r.body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// Логируем запрос ДО отправки
	r.client.logRequest(req.Method, fullURL, requestBody, req.Header)

	// Выполняем запрос
	resp, err := r.client.httpClient.Do(req)
	if err != nil {
		duration := time.Since(start)
		r.client.logError(req.Method, fullURL, err, duration)
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	// Читаем ответ
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		duration := time.Since(start)
		r.client.logError(req.Method, fullURL, err, duration)
		return fmt.Errorf("failed to read response body: %w", err)
	}

	duration := time.Since(start)

	// Логируем ответ
	r.client.logResponse(req.Method, fullURL, resp.Status, responseBody, resp.Header, duration)

	// Обрабатываем ошибки HTTP
	if resp.StatusCode >= 400 {
		return fmt.Errorf("API error %d: %s", resp.StatusCode, string(responseBody))
	}

	// Парсим ответ если требуется
	if result != nil {
		if err := json.Unmarshal(responseBody, result); err != nil {
			return fmt.Errorf("failed to unmarshal response: %w", err)
		}
	}

	return nil
}
