// internal/models/legacy/service_update.go
package legacy

// ServiceUpdateRequest - запрос на обновление услуги через Legacy API
type ServiceUpdateRequest struct {
	Name        *string `json:"name,omitempty"`
	PaymentTerm *string `json:"paymentTerm,omitempty"`
	AutoProlong *bool   `json:"autoProlong,omitempty"`
}
