// models/responses/legacy/types.go
package legacy

type ServiceType struct {
	Slug                 string                 `json:"slug"`                 // "vps", "vpn", "domain", etc.
	Order                int                    `json:"order"`                // Порядок отображения
	Names                map[string]string      `json:"names"`                // Локализованные названия
	Payload              map[string]interface{} `json:"payload"`              // Дополнительные данные
	LocaledPayload       map[string]interface{} `json:"localedPayload"`       // Локализованные дополнительные данные
	Name                 string                 `json:"name"`                 // Основное название
	PrettyLocaledPayload map[string]string      `json:"prettyLocaledPayload"` // Форматированные локализованные данные
}
