// models/groups.go
package models

// ProductGroupsResponse - ответ на запрос групп продуктов
type ProductGroupsResponse struct {
	Items []ProductGroup `json:"items"`
	Total int            `json:"total"`
}

// ProductGroup - группа продуктов
type ProductGroup struct {
	ID             int                    `json:"id"`
	Name           string                 `json:"name"`
	Type           GroupType              `json:"type"`
	Description    *string                `json:"description"`
	Payload        map[string]interface{} `json:"payload"`
	LocaledPayload map[string]interface{} `json:"localedPayload"`
}

// GroupType - тип группы
type GroupType struct {
	Slug           string                 `json:"slug"`
	Name           string                 `json:"name"`
	Parameters     []Parameter            `json:"parameters"`
	ServiceHandler string                 `json:"serviceHandler"`
	Payload        map[string]interface{} `json:"payload"`
	LocaledPayload map[string]interface{} `json:"localedPayload"`
}

// Parameter - параметр группы
type Parameter struct {
	Field string `json:"field"`
	Name  string `json:"name"`
	Type  string `json:"type"`
}
