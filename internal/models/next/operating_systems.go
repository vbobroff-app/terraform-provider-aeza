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
