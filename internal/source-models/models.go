package source_models

// Service представляет услугу в Aeza
type Service struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Status     string `json:"status"`
	LocationID int64  `json:"location_id"`
	ProductID  int64  `json:"product_id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

// Product представляет продукт в Aeza
type Product struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Currency    string  `json:"currency"`
	Category    string  `json:"category"`
}

// ServiceType представляет тип услуги
type ServiceType struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// ListServicesResponse - ответ на запрос списка услуг
type ListServicesResponse struct {
	Services []Service `json:"services"`
	Total    int       `json:"total"`
}

// ListProductsResponse - ответ на запрос списка продуктов
type ListProductsResponse struct {
	Products []Product `json:"products"`
	Total    int       `json:"total"`
}

// ListServiceTypesResponse - ответ на запрос типов услуг
type ListServiceTypesResponse struct {
	ServiceTypes []ServiceType `json:"service_types"`
	Total        int           `json:"total"`
}

// ServiceCreateRequest - запрос на создание услуги
type ServiceCreateRequest struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	LocationID int64  `json:"location_id"`
	ProductID  int64  `json:"product_id"`
}

// ServiceCreateResponse - ответ на создание услуги
type ServiceCreateResponse struct {
	ID     int64   `json:"id"`
	Status string  `json:"status"`
	Error  *string `json:"error,omitempty"`
}

// ServiceGetResponse - ответ на запрос услуги
type ServiceGetResponse struct {
	Service Service `json:"service"`
}
