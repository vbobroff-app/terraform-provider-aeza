package client

import "time"

// Базовые структуры ответов API
type BaseResponse struct {
	Data struct {
		Items []interface{} `json:"items"`
		Total int           `json:"total"`
	} `json:"data"`
}

type ServicesResponse struct {
	Data struct {
		Items []Service `json:"items"`
		Total int       `json:"total"`
	} `json:"data"`
}

type ProductsResponse struct {
	Data struct {
		Items []Product `json:"items"`
		Total int       `json:"total"`
	} `json:"data"`
}

type ServiceTypesResponse struct {
	Data struct {
		Items []ServiceType `json:"items"`
		Total int           `json:"total"`
	} `json:"data"`
}

type ServiceResponse struct {
	Data Service `json:"data"`
}

// Бизнес-модели
type Service struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Type         string    `json:"type"`
	Status       string    `json:"status"`
	ProductName  string    `json:"productName"`
	LocationCode string    `json:"locationCode"`
	CreatedAt    time.Time `json:"createable"`
}

type Product struct {
	ID            int             `json:"id"`
	Name          string          `json:"name"`
	Type          string          `json:"type"`
	GroupID       int             `json:"groupId"`
	Configuration []ProductConfig `json:"configuration"`
	Prices        ProductPrices   `json:"prices"`
	Group         ProductGroup    `json:"group"`
}

type ProductConfig struct {
	Slug string `json:"slug"`
	Base int    `json:"base"`
	Max  int    `json:"max"`
	Type string `json:"type"`
}

type ProductPrices struct {
	Hour  int `json:"hour"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

type ProductGroup struct {
	ID      int                    `json:"id"`
	Names   map[string]string      `json:"names"`
	Payload map[string]interface{} `json:"payload"`
}

type ServiceType struct {
	Slug    string                 `json:"slug"`
	Name    string                 `json:"name"`
	Order   int                    `json:"order"`
	Payload map[string]interface{} `json:"payload"`
}

type CreateServiceRequest struct {
	ProductID int    `json:"product_id"`
	Name      string `json:"name"`
	Type      string `json:"type,omitempty"`
}
