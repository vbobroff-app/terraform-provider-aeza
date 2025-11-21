// models/legacy/responses.go
package legacy

// Универсальный response для списка услуг
type ListServicesResponse struct {
	Data struct {
		SelectorMode string        `json:"selectorMode"`
		Filter       *string       `json:"filter"`
		Items        []ServiceBase `json:"items"` // Базовый тип для списка
		Total        int           `json:"total"`
		Edit         bool          `json:"edit"`
	} `json:"data"`
}

type ListServicesVPSResponse struct {
	Data struct {
		SelectorMode string       `json:"selectorMode"`
		Filter       *string      `json:"filter"`
		Items        []ServiceVPS `json:"items"` // Базовый тип для списка
		Total        int          `json:"total"`
		Edit         bool         `json:"edit"`
	} `json:"data"`
}

// ListServiceTypesResponse - ответ на запрос типов услуг
type ListServiceTypesResponse struct {
	Data struct {
		Items []ServiceType `json:"items"`
		Total int           `json:"total"`
	} `json:"data"`
}
