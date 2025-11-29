// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// models/legacy/groups.go
package legacy

type ServiceGroup struct {
	ID             int64                  `json:"id"`
	Name           string                 `json:"name"`
	Type           string                 `json:"type"`
	Role           interface{}            `json:"role"`
	ParentID       interface{}            `json:"parentId"`
	Description    string                 `json:"description"`
	Payload        map[string]interface{} `json:"payload"`
	LocaledPayload map[string]interface{} `json:"localedPayload"`
	TypeObject     ServiceType            `json:"typeObject"`
	ServiceHandler string                 `json:"serviceHandler"`
}
