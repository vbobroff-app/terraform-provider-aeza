// models/legacy/groups.go
package legacy

type ServiceGroup struct {
	ID             int64                  `json:"id"`
	Name           string                 `json:"name"`
	Type           string                 `json:"type"`
	Role           interface{}            `json:"role"`
	ParentID       interface{}            `json:"parentId"`
	Description    string                 `json:"description"`
	Payload        map[string]interface{} `json:"payload"`
	LocaledPayload map[string]interface{} `json:"localedPayload"`
	TypeObject     ServiceType            `json:"typeObject"`
	ServiceHandler string                 `json:"serviceHandler"`
}
