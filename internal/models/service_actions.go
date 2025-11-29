// internal/models/service_actions.go
package models

type ServiceProlongRequest struct {
	Method string `json:"method"`
	Term   string `json:"term"`
	Count  int64  `json:"count"`
}

// ServiceProlongResponse - ответ на запрос продления услуги
type ServiceProlongResponse struct {
	Transaction *ProlongedTransaction `json:"transaction,omitempty"`
}

type ProlongedTransaction struct {
	ID        int64  `json:"id"`
	Amount    string `json:"amount"`
	Status    string `json:"status"`
	Type      string `json:"type"`
	CreatedAt int64  `json:"created_at"`
	ServiceID int64  `json:"service_id"`
}
