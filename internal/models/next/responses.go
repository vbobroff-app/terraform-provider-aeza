// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

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
