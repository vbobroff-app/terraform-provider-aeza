// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// models/responses/next/operating_systems.go
package next

type OperatingSystem struct {
	ID         int            `json:"id"`
	Slug       string         `json:"slug"`
	Name       string         `json:"name"`
	Repository *string        `json:"repository"`
	Group      string         `json:"group"`
	Username   string         `json:"username"`
	Enabled    bool           `json:"enabled"`
	Targets    map[string]int `json:"targets"`
	Order      int            `json:"order"`
}
