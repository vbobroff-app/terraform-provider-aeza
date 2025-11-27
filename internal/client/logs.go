// internal/client/logs.go

package client

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

// logRequest логирует HTTP запрос
func (c *Client) logRequest(method, url string, body []byte, headers http.Header) {
	log.Printf("[DEBUG] === HTTP REQUEST ===")
	log.Printf("[DEBUG] Method: %s", method)
	log.Printf("[DEBUG] URL: %s", url)

	// Логируем заголовки
	log.Printf("[DEBUG] Headers:")
	for key, values := range headers {
		if strings.ToLower(key) == "x-api-key" {
			log.Printf("[DEBUG]   %s: [REDACTED]", key)
		} else {
			log.Printf("[DEBUG]   %s: %v", key, values)
		}
	}

	// Логируем тело запроса
	log.Printf("[DEBUG] Request Body:")
	if len(body) > 0 {
		// Декодируем Base64 если это JSON
		bodyStr := string(body)
		if decoded, err := base64.StdEncoding.DecodeString(bodyStr); err == nil {
			// Это Base64 - декодируем и форматируем JSON
			var prettyJSON bytes.Buffer
			if err := json.Indent(&prettyJSON, decoded, "", "  "); err == nil {
				log.Printf("[DEBUG] %s", prettyJSON.String())
			} else {
				log.Printf("[DEBUG] %s", string(decoded))
			}
		} else {
			// Не Base64 - логируем как есть
			if c.isJSONContent(headers) {
				var prettyJSON bytes.Buffer
				if err := json.Indent(&prettyJSON, body, "", "  "); err == nil {
					log.Printf("[DEBUG] %s", prettyJSON.String())
				} else {
					log.Printf("[DEBUG] %s", bodyStr)
				}
			} else {
				log.Printf("[DEBUG] %s", bodyStr)
			}
		}
	} else {
		log.Printf("[DEBUG] [EMPTY]")
	}
	log.Printf("[DEBUG] ====================")
}

// logResponse логирует HTTP ответ
func (c *Client) logResponse(method, url string, status string, body []byte, headers http.Header, duration time.Duration) {
	log.Printf("[DEBUG] === HTTP RESPONSE ===")
	log.Printf("[DEBUG] Method: %s", method)
	log.Printf("[DEBUG] URL: %s", url)
	log.Printf("[DEBUG] Status: %s", status)
	log.Printf("[DEBUG] Duration: %v", duration)

	// Логируем заголовки ответа
	if len(headers) > 0 {
		log.Printf("[DEBUG] Response Headers:")
		for key, values := range headers {
			log.Printf("[DEBUG]   %s: %v", key, values)
		}
	}

	// Логируем тело ответа (если есть) - len() работает для nil slices
	log.Printf("[DEBUG] Response Body:")
	if len(body) > 0 {
		if c.isJSONContent(headers) {
			var prettyJSON bytes.Buffer
			if err := json.Indent(&prettyJSON, body, "", "  "); err == nil {
				log.Printf("[DEBUG] %s", prettyJSON.String())
			} else {
				log.Printf("[DEBUG] %s (invalid JSON)", string(body))
			}
		} else {
			// Обрезаем длинный текст для читаемости
			bodyStr := string(body)
			if len(bodyStr) > 1000 {
				bodyStr = bodyStr[:1000] + "... [TRUNCATED]"
			}
			log.Printf("[DEBUG] %s", bodyStr)
		}
	} else {
		log.Printf("[DEBUG] [EMPTY]")
	}
	log.Printf("[DEBUG] =====================")
}

// isJSONContent проверяет, является ли content-type JSON
func (c *Client) isJSONContent(headers http.Header) bool {
	contentType := headers.Get("Content-Type")
	return strings.Contains(contentType, "application/json") ||
		strings.Contains(contentType, "text/json") ||
		strings.Contains(contentType, "json")
}

// logError логирует ошибки
func (c *Client) logError(method, url string, err error, duration time.Duration) {
	log.Printf("[ERROR] === HTTP ERROR ===")
	log.Printf("[ERROR] Method: %s", method)
	log.Printf("[ERROR] URL: %s", url)
	log.Printf("[ERROR] Error: %v", err)
	log.Printf("[ERROR] Duration: %v", duration)
	log.Printf("[ERROR] ==================")
}
