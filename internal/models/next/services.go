// models/responses/legacy/legacy_models.go
package next

type Service struct {
	ID           int                    `json:"id"`
	Name         string                 `json:"name"`
	IP           string                 `json:"ip"`
	Payload      map[string]interface{} `json:"payload"`
	Price        int                    `json:"price"`
	PaymentTerm  string                 `json:"paymentTerm"`
	AutoProlong  bool                   `json:"autoProlong"`
	CreatedAt    string                 `json:"createdAt"`    // ISO 8601: "2025-11-20T13:12:12.733Z"
	ExpiresAt    string                 `json:"expiresAt"`    // ISO 8601: "2025-11-20T13:12:12.733Z"
	Status       string                 `json:"status"`       // "activation_wait", "active", etc.
	TypeSlug     string                 `json:"typeSlug"`     // "vps"
	ProductName  string                 `json:"productName"`  // "SWE-PROMO"
	LocationCode string                 `json:"locationCode"` // "de", "nl", etc.
	CurrentTask  *CurrentTask           `json:"currentTask"`  // nullable объект, а не строка!
	Capabilities []string               `json:"capabilities"` // ["change_password", "rename", ..., "backups"]
}

type CurrentTask struct {
	ID        string `json:"id"`
	Slug      string `json:"slug"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"` // ISO 8601
	Status    string `json:"status"`    // "queued", "running", etc.
}

type ComputedParametersVPS struct {
	CPU      int     `json:"cpu"`
	RAM      int     `json:"ram"`
	ROM      int     `json:"rom"`
	IP       int     `json:"ip"`
	OS       string  `json:"os"`
	Node     string  `json:"node"`
	ISOURL   string  `json:"isoUrl"`
	Recipe   *string `json:"recipe"` // nullable
	Username string  `json:"username"`
}
