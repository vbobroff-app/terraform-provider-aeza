// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package legacy

// ServiceProlongRequest - запрос на продление услуги через Legacy API
type ServiceProlongRequest struct {
	Method string `json:"method"`
	Term   string `json:"term"`
	Count  int64  `json:"count"`
}

// ServiceProlongResponse - ответ на запрос продления услуги
type ServiceProlongResponse struct {
	Data struct {
		Transaction Transaction `json:"transaction"`
	} `json:"data"`
}
