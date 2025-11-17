// internal/client/models.go
package source_models

// Service представляет услугу в Aeza
type Service struct {
	ID               int64                  `json:"id"`
	OwnerID          int64                  `json:"ownerId"`
	ProductID        int64                  `json:"productId"`
	Name             string                 `json:"name"`
	IP               string                 `json:"ip"`
	PaymentTerm      string                 `json:"paymentTerm"`
	Parameters       map[string]interface{} `json:"parameters"`
	SecureParameters map[string]interface{} `json:"secureParameters"`
	AutoProlong      bool                   `json:"autoProlong"`
	Backups          bool                   `json:"backups"`
	Status           string                 `json:"status"`
	LastStatus       string                 `json:"lastStatus"`
	Product          Product                `json:"product"`
	LocationCode     string                 `json:"locationCode"`
	Prices           map[string]Price       `json:"prices"`
	CurrentStatus    string                 `json:"currentStatus"`
	CreatedAt        string                 `json:"createdAt"`
	UpdatedAt        string                 `json:"updatedAt"`
}

type Price struct {
	Value    float64 `json:"value"`
	Currency string  `json:"currency"`
}

// Product представляет продукт в Aeza
type Product struct {
	ID                   int64                           `json:"id"`
	Type                 string                          `json:"type"`
	GroupID              int64                           `json:"groupId"`
	Name                 string                          `json:"name"`
	Configuration        []ConfigurationItem             `json:"configuration"`
	Prices               map[string]float64              `json:"prices"`
	SummaryConfiguration map[string]ConfigurationSummary `json:"summaryConfiguration"`
	ServiceHandler       string                          `json:"serviceHandler"`
}

type ConfigurationItem struct {
	Max    float64                `json:"max"`  // Меняем int на float64
	Base   float64                `json:"base"` // Меняем int на float64
	Slug   string                 `json:"slug"`
	Type   string                 `json:"type"`
	Prices map[string]interface{} `json:"prices"`
}

type ConfigurationSummary struct {
	Max  float64 `json:"max"`  // Меняем int на float64
	Base float64 `json:"base"` // Меняем int на float64
	Slug string  `json:"slug"`
	Type string  `json:"type"`
}

// ServiceType представляет тип услуги
type ServiceType struct {
	Slug    string                 `json:"slug"`
	Order   int                    `json:"order"`
	Names   map[string]string      `json:"names"`
	Payload map[string]interface{} `json:"payload"`
	Name    string                 `json:"name"`
}

// ListServicesResponse - ответ на запрос списка услуг
type ListServicesResponse struct {
	Data struct {
		Items []Service `json:"items"`
		Total int       `json:"total"`
	} `json:"data"`
}

// ListProductsResponse - ответ на запрос списка продуктов
type ListProductsResponse struct {
	Data struct {
		Items []Product `json:"items"`
		Total int       `json:"total"`
	} `json:"data"`
}

// ListServiceTypesResponse - ответ на запрос типов услуг
type ListServiceTypesResponse struct {
	Data struct {
		Items []ServiceType `json:"items"`
		Total int           `json:"total"`
	} `json:"data"`
}

// ServiceCreateRequest - запрос на создание услуги
type ServiceCreateRequest struct {
	Name         string `json:"name"`
	ProductID    int64  `json:"productId"`
	LocationCode string `json:"locationCode"`
	PaymentTerm  string `json:"paymentTerm"`
	AutoProlong  bool   `json:"autoProlong"`
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
