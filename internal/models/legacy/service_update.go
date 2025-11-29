// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/models/legacy/service_update.go
package legacy

// ServiceUpdateRequest - запрос на обновление услуги через Legacy API
type ServiceUpdateRequest struct {
	Name        *string `json:"name,omitempty"`
	AutoProlong *bool   `json:"autoProlong,omitempty"`
}
