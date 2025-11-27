// models/legacy/service_create.go
package legacy

// ServiceCreateRequest - запрос на создание услуги через Legacy API
type ServiceCreateRequest struct {
	Method      string                  `json:"method"`
	Count       int                     `json:"count"`
	Term        string                  `json:"term"`
	Name        string                  `json:"name"`
	ProductID   int64                   `json:"productId"`
	Parameters  ServiceCreateParameters `json:"parameters"`
	AutoProlong bool                    `json:"autoProlong"`
	Backups     bool                    `json:"backups"`
}

type ServiceCreateParameters struct {
	Recipe *string `json:"recipe"`
	OS     string  `json:"os"`
	IsoURL string  `json:"isoUrl"`
}

// ServiceCreateResponse - основной ответ на создание услуги через Legacy API
type ServiceCreateResponse struct {
	Data ServiceCreateData `json:"data"`
}

type ServiceCreateData struct {
	Items       []ServiceOrderItem `json:"items"`
	Transaction Transaction        `json:"transaction"`
}

type ServiceOrderItem struct {
	ID                int64                  `json:"id"`
	OwnerID           int64                  `json:"ownerId"`
	ProductID         int64                  `json:"productId"`
	Name              string                 `json:"name"`
	Term              string                 `json:"term"`
	Count             int                    `json:"count"`
	Parameters        ServiceParameters      `json:"parameters"`
	Configuration     map[string]interface{} `json:"configuration"`
	AutoProlong       bool                   `json:"autoProlong"`
	Backups           bool                   `json:"backups"`
	CreatedAt         int64                  `json:"createdAt"`
	Status            string                 `json:"status"`
	CreatedServiceIds []int64                `json:"createdServiceIds"`
	Pincode           bool                   `json:"pincode"`
	Payload           map[string]interface{} `json:"payload"`
	Product           Product                `json:"product"`
	IndividualPrice   float64                `json:"individualPrice"`
}

type ServiceParameters struct {
	OS     string  `json:"os"`
	IsoURL string  `json:"isoUrl"`
	Recipe *string `json:"recipe"`
}

type Transaction struct {
	ID                int64                  `json:"id"`
	OwnerID           int64                  `json:"ownerId"`
	Amount            float64                `json:"amount"`
	BonusAmount       float64                `json:"bonusAmount"`
	Mode              string                 `json:"mode"`
	Status            string                 `json:"status"`
	PerformedAt       interface{}            `json:"performedAt"`
	CreatedAt         int64                  `json:"createdAt"`
	Type              string                 `json:"type"`
	ServiceID         interface{}            `json:"serviceId"`
	ItemSlug          interface{}            `json:"itemSlug"`
	Description       interface{}            `json:"description"`
	RandomString      interface{}            `json:"randomString"`
	Payload           map[string]interface{} `json:"payload"`
	InvoiceID         int64                  `json:"invoiceId"`
	Invoice           interface{}            `json:"invoice"`
	PrettyAmount      PrettyAmount           `json:"prettyAmount"`
	OrderInfo         OrderInfo              `json:"orderInfo"`
	PrettyBonusAmount PrettyAmount           `json:"prettyBonusAmount"`
}

type PrettyAmount struct {
	Value float64 `json:"value"`
}

type OrderInfo struct {
	ID                int64                  `json:"id"`
	OwnerID           int64                  `json:"ownerId"`
	ProductID         int64                  `json:"productId"`
	Name              string                 `json:"name"`
	Term              string                 `json:"term"`
	Count             int                    `json:"count"`
	Parameters        ServiceParameters      `json:"parameters"`
	Configuration     map[string]interface{} `json:"configuration"`
	AutoProlong       bool                   `json:"autoProlong"`
	Backups           bool                   `json:"backups"`
	CreatedAt         int64                  `json:"createdAt"`
	Status            string                 `json:"status"`
	CreatedServiceIds []int64                `json:"createdServiceIds"`
	Pincode           bool                   `json:"pincode"`
	Payload           map[string]interface{} `json:"payload"`
}
