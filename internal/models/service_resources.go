// internal/models/service_resources.go
package models

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
	ExpiresAt        string                 `json:"expiresAt"`
	PurchasedAt      string                 `json:"purchasedAt"`
}

type Price struct {
	Value    float64 `json:"value"`
	Currency string  `json:"currency"`
}

// ServiceCreateRequest - Terraform модель для создания услуги
type ServiceCreateRequest struct {
	Name        string `json:"name"`
	ProductID   int64  `json:"product_id"`
	PaymentTerm string `json:"payment_term"`
	AutoProlong bool   `json:"auto_prolong"`
	OS          string `json:"os,omitempty"`
	Recipe      string `json:"recipe,omitempty"`
	IsoURL      string `json:"iso_url,omitempty"`
}

// ServiceCreateResponse - Terraform модель ответа создания услуги
type ServiceCreateResponse struct {
	ID      int64  `json:"id"`
	OrderID int64  `json:"order_id"`
	Status  string `json:"status"`
	Date    string `json:"date"`

	// Информация о продукте
	ProductId   int64  `json:"product_id"`
	ProductType string `json:"product_type"`
	GroupId     *int64 `json:"group_id"`
	ProductName string `json:"product_name"`

	// Информация о локации
	LocationName string `json:"location_name"`

	// Информация о цене
	Term              string `json:"term"`
	Price             string `json:"price"`
	TransactionAmount string `json:"transaction_amount"`
}

type ServiceUpdateRequest struct {
	Name        *string `json:"name,omitempty"`
	AutoProlong *bool   `json:"auto_prolong,omitempty"`
}

type ServiceProlongRequest struct {
	Method string `json:"method"`
	Term   string `json:"term"`
	Count  int64  `json:"count"`
}
