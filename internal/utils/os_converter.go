// internal/utils/os_converter.go
package utils

import (
	"github.com/vbobroff-app/terraform-provider-aeza/internal/models/legacy"
	"github.com/vbobroff-app/terraform-provider-aeza/internal/models/next"
)

func ConvertOsFromLegacy(legacy legacy.OperatingSystem) next.OperatingSystem {
	return next.OperatingSystem{
		ID:         legacy.ID,
		Slug:       legacy.Slug,
		Name:       legacy.Name,
		Repository: legacy.Repository,
		Group:      legacy.Group,
		Username:   legacy.Username,
		Enabled:    legacy.Enabled,
		Targets:    legacy.Targets,
		Order:      0, // значение по умолчанию
	}
}
