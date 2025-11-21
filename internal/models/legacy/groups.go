// models/responses/legacy/groups.go
package legacy

type ServicesGroup struct {
	ID                   int64                             `json:"id"`
	Order                int                               `json:"order"`
	Names                map[string]string                 `json:"names"`
	Type                 string                            `json:"type"`
	Role                 *string                           `json:"role"`
	ParentID             *int64                            `json:"parentId"`
	Descriptions         map[string]string                 `json:"descriptions"`
	Payload              ServicesGroupPayload              `json:"payload"`
	LocaledPayload       ServicesGroupLocaledPayload       `json:"localedPayload"`
	ConfigurationPrices  map[string]interface{}            `json:"configurationPrices"`
	Description          string                            `json:"description"`
	HasProducts          bool                              `json:"hasProducts"`
	PrettyLocaledPayload ServicesGroupPrettyLocaledPayload `json:"prettyLocaledPayload"`
	Name                 string                            `json:"name"`
	TypeObject           ServicesGroupTypeObject           `json:"typeObject"`
	ServiceHandler       string                            `json:"serviceHandler"`
}

// ServicesGroupPayload - структура для payload в ServicesGroup
type ServicesGroupPayload struct {
	Mode       string      `json:"mode,omitempty"`  // "dedicated", "shared", "amd", "intel"
	Label      string      `json:"label,omitempty"` // "FI-DEDICATED", "US-DEDICATED", etc
	IsDisabled bool        `json:"isDisabled,omitempty"`
	Mgr        string      `json:"mgr,omitempty"`     // "usa", "hel"
	Network    interface{} `json:"network,omitempty"` // может быть string "1" или number 1
}

type ServicesGroupLocaledPayload struct {
	Description map[string]string `json:"description,omitempty"`
	Features    map[string]string `json:"features,omitempty"`
}

type ServicesGroupPrettyLocaledPayload struct {
	Description string `json:"description,omitempty"`
	Features    string `json:"features,omitempty"`
}

type ServicesGroupTypeObject struct {
	Slug                 string                 `json:"slug"`
	Order                int                    `json:"order"`
	Names                map[string]string      `json:"names"`
	Payload              ServiceTypePayload     `json:"payload"`
	LocaledPayload       map[string]interface{} `json:"localedPayload"`
	Name                 string                 `json:"name"`
	PrettyLocaledPayload map[string]interface{} `json:"prettyLocaledPayload"`
}

// ServiceTypePayload - структура для payload в ServiceType
type ServiceTypePayload struct {
	Badge          string `json:"badge,omitempty"`          // "-50%"
	BadgeColor     string `json:"badgeColor,omitempty"`     // "#ff003d"
	BadgeTextColor string `json:"badgeTextColor,omitempty"` // "#ffffff"
}
