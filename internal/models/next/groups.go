// models/next/groups.go
package next

type ServiceGroup struct {
	ID             int64                  `json:"id"`
	Name           string                 `json:"name"`
	Type           ServiceType            `json:"type"`
	Role           interface{}            `json:"role"`
	ParentID       interface{}            `json:"parentId"`
	Description    string                 `json:"description"`
	Payload        map[string]interface{} `json:"payload"`
	LocaledPayload map[string]interface{} `json:"localedPayload"` // Используем interface{} для гибкости
}

type ServiceType struct {
	Slug           string                 `json:"slug"`
	Name           string                 `json:"name"`
	Parameters     []ServiceParameter     `json:"parameters"`
	ServiceHandler string                 `json:"serviceHandler"`
	Payload        map[string]interface{} `json:"payload"`
	LocaledPayload map[string]interface{} `json:"localedPayload"`
}

type ServiceParameter struct {
	Field string `json:"field"`
	Name  string `json:"name"`
	Type  string `json:"type"`
}
