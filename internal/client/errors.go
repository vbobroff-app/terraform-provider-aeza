package client

import "fmt"

type APIError struct {
	StatusCode int
	Slug       string
	Message    string
	Data       interface{}
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API Error %d: %s (%s)", e.StatusCode, e.Message, e.Slug)
}

type ErrorResponse struct {
	Error struct {
		Slug    string      `json:"slug"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	} `json:"error"`
}
