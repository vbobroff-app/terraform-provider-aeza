// internal/client/models.go
package models

import "github.com/vbobroff-app/terraform-provider-aeza/internal/models/legacy"

type TerraformService struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	IP           string `json:"ip"`
	Price        int    `json:"price"`
	PaymentTerm  string `json:"paymentTerm"`
	AutoProlong  bool   `json:"autoProlong"`
	CreatedAt    string `json:"createdAt"`    // ISO 8601: "2025-11-20T13:12:12.733Z"
	ExpiresAt    string `json:"expiresAt"`    // ISO 8601: "2025-11-20T13:12:12.733Z"
	Status       string `json:"status"`       // "activation_wait", "active", etc.
	TypeSlug     string `json:"typeSlug"`     // "vps"
	ProductName  string `json:"productName"`  // "SWE-PROMO"
	LocationCode string `json:"locationCode"` // "de", "nl", etc.
}

// TerraformServiceDetailed - унифицированная структура для Terraform провайдера
type TerraformServiceDetailed struct {
	// Основная идентификация
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Status       string `json:"status"`
	TypeSlug     string `json:"type_slug"`     // "vps", "vpn", "domain"
	LocationCode string `json:"location_code"` // "nl", "de", "ru"

	// Сетевые настройки
	IP   string   `json:"ip"`
	IPs  []string `json:"ips,omitempty"`  // Все IP адреса
	IPv6 string   `json:"ipv6,omitempty"` // IPv6 подсеть

	// Стоимость (в евро)
	Price        float64 `json:"price"`         // 4.94 (евро)
	PriceDisplay string  `json:"price_display"` // "4,94 € / Месяц"
	PaymentTerm  string  `json:"payment_term"`  // "month", "hour", "year"
	AutoProlong  bool    `json:"auto_prolong"`

	// Конфигурация оборудования (для VPS)
	CPU     int `json:"cpu,omitempty"`      // 1 ядро
	RAM     int `json:"ram,omitempty"`      // 2 GB
	Storage int `json:"storage,omitempty"`  // 30 GB
	IPCount int `json:"ip_count,omitempty"` // Количество IP

	// Программное обеспечение
	ProductName string `json:"product_name"`       // "NLs-1"
	OS          string `json:"os,omitempty"`       // "Ubuntu 24.04"
	Username    string `json:"username,omitempty"` // "root"

	// Временные метки
	CreatedAt   string `json:"created_at"`   // ISO: "2025-11-19T01:04:29Z"
	ExpiresAt   string `json:"expires_at"`   // ISO: "2025-12-19T01:05:32Z"
	CreatedDate string `json:"created_date"` // "19 ноября, 2025 г."
	ExpiresDate string `json:"expires_date"` // "18 декабря, 2025 г."

	// Дополнительные возможности
	Capabilities []string               `json:"capabilities,omitempty"`
	Backups      bool                   `json:"backups"`
	Payload      map[string]interface{} `json:"payload,omitempty"`
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
	ID     int64                   `json:"id"`
	Status string                  `json:"status"`
	Error  *map[string]interface{} `json:"error,omitempty"` // Меняем *string на *map[string]interface{}
}

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

// ServiceGetResponse - ответ на запрос услуги
type ServiceGetResponse struct {
	Service Service `json:"service"`
}

type OperatingSystem struct {
	legacy.OperatingSystem
}
