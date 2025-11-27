// models/legacy/services.go
package legacy

type Service struct {
	ID            int     `json:"id"`
	OwnerID       int     `json:"ownerId"`
	ProductID     int     `json:"productId"`
	Name          string  `json:"name"`
	IP            string  `json:"ip"`
	PaymentTerm   string  `json:"paymentTerm"`
	AutoProlong   bool    `json:"autoProlong"`
	Backups       bool    `json:"backups"`
	Status        string  `json:"status"`
	LastStatus    *string `json:"lastStatus"`
	Product       Product `json:"product"`
	LocationCode  string  `json:"locationCode"`
	CurrentStatus string  `json:"currentStatus"`
}

type ServiceBase struct {
	Service
	Parameters       map[string]interface{} `json:"parameters"`
	SecureParameters map[string]interface{} `json:"secureParameters"`

	Payload          map[string]interface{} `json:"payload"`
	Configuration    map[string]interface{} `json:"configuration"`
	SpecialPrices    map[string]interface{} `json:"specialPrices"`
	RelativePrices   map[string]interface{} `json:"relativePrices"`
	Schedule         *string                `json:"schedule"`
	Timestamps       ServiceTimestamps      `json:"timestamps"`
	PaymentTermRatio float64                `json:"paymentTermRatio"`
	Opportunities    []string               `json:"opportunities"`
	Capabilities     []string               `json:"capabilities"`
	RawPrices        map[string]int         `json:"rawPrices"`
	IndividualPrices map[string]int         `json:"individualPrices"`
	CurrentTask      *string                `json:"currentTask"`
}

// ServiceVPS embedded все поля Service и добавляет VPS-специфичные поля
type ServiceVPS struct {
	Service
	// VPS-специфичные  override поля
	Parameters           Parameters                   `json:"parameters"`
	SecureParameters     SecureParameters             `json:"secureParameters"`
	IPs                  []IPAddress                  `json:"ips"`
	IPv6                 []IPv6Address                `json:"ipv6"`
	Payload              map[string]interface{}       `json:"payload"`
	Configuration        map[string]interface{}       `json:"configuration"`
	SpecialPrices        map[string]interface{}       `json:"specialPrices"`
	RelativePrices       map[string]interface{}       `json:"relativePrices"`
	Schedule             *string                      `json:"schedule"`
	Timestamps           ServiceTimestamps            `json:"timestamps"`
	PaymentTermRatio     float64                      `json:"paymentTermRatio"`
	Opportunities        []string                     `json:"opportunities"`
	Capabilities         []string                     `json:"capabilities"`
	RawPrices            map[string]int               `json:"rawPrices"`
	SummaryConfiguration map[string]ConfigurationItem `json:"summaryConfiguration"`
	IndividualPrices     map[string]int               `json:"individualPrices"`
	CurrentTask          interface{}                  `json:"currentTask"`
}

type Parameters struct {
	OS       string  `json:"os,omitempty"`
	ISOURL   string  `json:"isoUrl,omitempty"`
	Recipe   *string `json:"recipe,omitempty"`
	Username string  `json:"username,omitempty"`
}

type SecureParameters struct {
	Unsecure bool                   `json:"unsecure"`
	Data     map[string]interface{} `json:"data"`
}

type IPAddress struct {
	Key                string `json:"key"`
	Mask               string `json:"mask"`
	Value              string `json:"value"`
	Status             string `json:"status"`
	Gateway            string `json:"gateway"`
	ExtendedProtection bool   `json:"extendedProtection"`
}
type IPv6Address struct {
	IPs     []interface{} `json:"ips"`
	Key     string        `json:"key"`
	Value   string        `json:"value"`
	Prefix  int           `json:"prefix"`
	Gateway string        `json:"gateway"`
}

type ServiceTimestamps struct {
	CreatedAt   int64 `json:"createdAt"`   // Unix timestamp
	ExpiresAt   int64 `json:"expiresAt"`   // Unix timestamp
	PurchasedAt int64 `json:"purchasedAt"` // Unix timestamp
}

type ConfigurationItem struct {
	Prices       map[string]interface{} `json:"prices"`
	Count        int                    `json:"count"`
	Base         int                    `json:"base"`
	Additionally int                    `json:"additionally"`
}
