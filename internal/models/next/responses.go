// models/next/responses.go
package next

type ListServicesResponse struct {
	Items []Service `json:"items"`
	Total int       `json:"total"`
}

type ComputedParametersResponse struct {
	Data ComputedParametersVPS `json:"data"`
}

type ServiceGroupsResponse struct {
	Items []ServiceGroup `json:"items"`
	Total int            `json:"total"`
}

type OSResponse struct {
	OperatingSystems []OperatingSystem `json:"operating_systems"`
}
