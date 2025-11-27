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

type ServiceGroupsResponse struct {
	Data struct {
		Items []ServiceGroup `json:"items"`
		Total int            `json:"total"`
	} `json:"data"`
}

// ListProductsResponse - ответ на запрос списка продуктов
type ListProductsResponse struct {
	Data struct {
		Items []Product `json:"items"`
		Total int       `json:"total"`
	} `json:"data"`
}

type OSResponse struct {
	Data struct {
		Items []OperatingSystem `json:"items"`
	} `json:"data"`
}

// ServiceDeleteResponse - ответ на удаление услуги через Legacy API
type ServiceDeleteResponse struct {
	Data string `json:"data"`
}
