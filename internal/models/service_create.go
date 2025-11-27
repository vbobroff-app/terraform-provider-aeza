// internal/models/service_create.go
package models

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
