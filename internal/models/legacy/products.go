// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// models/responses/legacy/product.go
package legacy

type Product struct {
	ID                     int                    `json:"id"`
	Name                   string                 `json:"name"`
	Type                   string                 `json:"type"`
	GroupID                *int64                 `json:"groupId"`
	Order                  int64                  `json:"order"`
	Configuration          []ProductConfig        `json:"configuration"`
	DefaultParameters      map[string]interface{} `json:"defaultParameters"`
	Payload                map[string]interface{} `json:"payload"`
	IsPrivate              bool                   `json:"isPrivate"`
	Prices                 ProductPrices          `json:"prices"`
	RawPrices              ProductPrices          `json:"rawPrices"`
	IndividualPrices       ProductPrices          `json:"individualPrices"`
	InstallPrice           float64                `json:"installPrice"` // меняем int на float64
	FirstPrices            ProductPrices          `json:"firstPrices"`
	IndividualFirstPrices  ProductPrices          `json:"individualFirstPrices"`
	IndividualInstallPrice *float64               `json:"individualInstallPrice"` // меняем *int на *float64
	SummaryConfiguration   map[string]interface{} `json:"summaryConfiguration"`
	LocaledPayload         map[string]interface{} `json:"localedPayload"`
	PrettyLocaledPayload   map[string]interface{} `json:"prettyLocaledPayload"`
	Group                  *ProductGroup          `json:"group"`
	TypeObject             *ServiceType           `json:"typeObject"`
	ServiceHandler         string                 `json:"serviceHandler"`
}

type ProductConfig struct {
	Slug   string                 `json:"slug"`
	Base   float64                `json:"base"` // меняем int на float64
	Max    float64                `json:"max"`  // меняем int на float64
	Type   string                 `json:"type"`
	Count  int                    `json:"count"`
	Prices map[string]interface{} `json:"prices"`
}

type ProductPrices struct {
	Hour        float64 `json:"hour"`         // меняем int на float64
	Month       float64 `json:"month"`        // меняем int на float64
	Year        float64 `json:"year"`         // меняем int на float64
	HalfYear    float64 `json:"half_year"`    // меняем int на float64
	QuarterYear float64 `json:"quarter_year"` // меняем int на float64
}

type ProductGroup struct {
	ID                   int                    `json:"id"`
	Order                int                    `json:"order"`
	Names                map[string]string      `json:"names"`
	Type                 string                 `json:"type"`
	Role                 *string                `json:"role"`
	ParentID             *int                   `json:"parentId"`
	Descriptions         map[string]string      `json:"descriptions"`
	Payload              map[string]interface{} `json:"payload"`
	LocaledPayload       map[string]interface{} `json:"localedPayload"`
	ConfigurationPrices  map[string]interface{} `json:"configurationPrices"`
	HasProducts          bool                   `json:"hasProducts"`
	PrettyLocaledPayload map[string]interface{} `json:"prettyLocaledPayload"`
	Name                 string                 `json:"name"`
	TypeObject           *ServiceType           `json:"typeObject"`
	ServiceHandler       string                 `json:"serviceHandler"`
}
