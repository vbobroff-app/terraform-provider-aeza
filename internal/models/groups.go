// models/groups.go
package models

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
