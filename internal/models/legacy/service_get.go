// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/models/legacy/service_get.go
package legacy

type ServiceGetResponse struct {
	Data ServiceGetData `json:"data"`
}

type ServiceGetData struct {
	Items []ServiceGet `json:"items"`
	Total int          `json:"total"`
}

type ServiceGet struct {
	ServiceVPS
	// В будущем можно добавить другие embedded структуры для разных типов услуг
	// ServiceDedicated для выделенных серверов
}
