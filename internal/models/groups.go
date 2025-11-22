// models/groups.go
package models

// // ProductGroupsResponse - ответ на запрос групп продуктов
// type ProductGroupsResponse struct {
// 	Items []ProductGroup `json:"items"`
// 	Total int            `json:"total"`
// }

// // ProductGroup - группа продуктов
// type ProductGroup struct {
// 	ID             int                    `json:"id"`
// 	Name           string                 `json:"name"`
// 	Type           GroupType              `json:"type"`
// 	Description    *string                `json:"description"`
// 	Payload        map[string]interface{} `json:"payload"`
// 	LocaledPayload map[string]interface{} `json:"localedPayload"`
// }

// // GroupType - тип группы
// type GroupType struct {
// 	Slug           string                 `json:"slug"`
// 	Name           string                 `json:"name"`
// 	Parameters     []Parameter            `json:"parameters"`
// 	ServiceHandler string                 `json:"serviceHandler"`
// 	Payload        map[string]interface{} `json:"payload"`
// 	LocaledPayload map[string]interface{} `json:"localedPayload"`
// }

// // Parameter - параметр группы
// type Parameter struct {
// 	Field string `json:"field"`
// 	Name  string `json:"name"`
// 	Type  string `json:"type"`
// }

type ServiceGroup struct {
	ID             int64  `json:"id" tfsdk:"id"`
	GroupType      string `json:"group_type" tfsdk:"group_type"` // тип группы: server, location, geography
	Name           string `json:"name" tfsdk:"name"`
	Type           string `json:"type" tfsdk:"type"`
	Location       string `json:"location" tfsdk:"location"`
	CountryCode    string `json:"country_code" tfsdk:"country_code"`
	ServerType     string `json:"server_type" tfsdk:"server_type"` // "shared", "dedicated"
	IsDisabled     bool   `json:"is_disabled" tfsdk:"is_disabled"`
	Description    string `json:"description" tfsdk:"description"`
	Features       string `json:"features" tfsdk:"features"`
	NetworkSpeed   string `json:"network_speed" tfsdk:"network_speed"`
	IPv4Count      int    `json:"ipv4_count" tfsdk:"ipv4_count"`
	IPv6Subnet     string `json:"ipv6_subnet" tfsdk:"ipv6_subnet"`
	CPUModel       string `json:"cpu_model" tfsdk:"cpu_model"`
	CPUFrequency   string `json:"cpu_frequency" tfsdk:"cpu_frequency"`
	ServiceHandler string `json:"service_handler" tfsdk:"service_handler"`
}
