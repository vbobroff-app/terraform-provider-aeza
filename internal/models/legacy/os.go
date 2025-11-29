// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// models/responses/legacy/os.go
package legacy

type OperatingSystem struct {
	ID         int            `json:"id"`
	Name       string         `json:"name"`
	Repository *string        `json:"repository"`
	Group      string         `json:"group"`
	Enabled    bool           `json:"enabled"`
	Slug       string         `json:"slug"`
	Username   string         `json:"username"`
	Targets    map[string]int `json:"targets"`
}
