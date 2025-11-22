// models/responses/legacy/os.go
package legacy

type OperatingSystem struct {
	ID         int            `json:"id"`
	Name       string         `json:"name"`
	Repository string         `json:"repository"`
	Group      string         `json:"group"`
	Enabled    bool           `json:"enabled"`
	Slug       string         `json:"slug"`
	Username   string         `json:"username"`
	Targets    map[string]int `json:"targets"`
}
