// models/responses/legacy/product.go
package legacy

type Product struct {
	ID                int                    `json:"id"`
	Name              string                 `json:"name"`
	Type              string                 `json:"type"`
	GroupID           int                    `json:"groupId"`
	Order             int                    `json:"order"`
	Configuration     []ProductConfig        `json:"configuration"`
	DefaultParameters map[string]interface{} `json:"defaultParameters"`
	Payload           map[string]interface{} `json:"payload"`
	IsPrivate         bool                   `json:"isPrivate"`
	Prices            ProductPrices          `json:"prices"`
	RawPrices         map[string]int         `json:"rawPrices"`
	IndividualPrices  map[string]int         `json:"individualPrices"`
	Group             ProductGroup           `json:"group"`
	TypeObject        ServiceType            `json:"typeObject"`
	ServiceHandler    string                 `json:"serviceHandler"`
}

type ProductConfig struct {
	Slug string `json:"slug"`
	Base int    `json:"base"`
	Max  int    `json:"max"`
	Type string `json:"type"`
}

type ProductPrices struct {
	Hour        int `json:"hour"`
	Month       int `json:"month"`
	Year        int `json:"year"`
	HalfYear    int `json:"half_year"`
	QuarterYear int `json:"quarter_year"`
}

type ProductGroup struct {
	ID      int                    `json:"id"`
	Names   map[string]string      `json:"names"`
	Payload map[string]interface{} `json:"payload"`
}

type Price struct {
	Value    float64 `json:"value"`
	Currency string  `json:"currency"`
}
