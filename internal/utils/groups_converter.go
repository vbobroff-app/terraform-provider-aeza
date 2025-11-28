// internal/utils/legacy_converter.go
package utils

import (
	"strings"

	"github.com/vbobroff-app/terraform-provider-aeza/internal/models"
	"github.com/vbobroff-app/terraform-provider-aeza/internal/models/legacy"
)

func ConvertLegacyServiceGroup(legacyGroup legacy.ServiceGroup) models.ServiceGroup {
	groupType := getLegacyGroupType(legacyGroup)
	serverType := getStringFromPayload(legacyGroup.Payload, "mode")

	// Для серверных групп уточняем server_type
	if groupType == "server" && serverType == "" {
		serverType = getLegacyServerSubtype(legacyGroup)
	}

	group := models.ServiceGroup{
		ID:             legacyGroup.ID,
		Name:           legacyGroup.Name,
		Type:           legacyGroup.TypeObject.Slug,
		GroupType:      groupType,
		ServiceHandler: legacyGroup.ServiceHandler,
		Description:    legacyGroup.Description,
		Location:       getStringFromPayload(legacyGroup.Payload, "label"),
		CountryCode:    getStringFromPayload(legacyGroup.Payload, "code"),
		ServerType:     serverType,
		IsDisabled:     getLegacyBoolFromPayload(legacyGroup.Payload, "isDisabled"),
		Features:       getLegacyFeatures(legacyGroup.LocaledPayload),
	}

	// Заполняем дополнительные поля только для серверных групп
	if groupType == "server" {
		features := getLegacyFeatures(legacyGroup.LocaledPayload)
		group.CPUModel = extractCPUModel(features)
		group.CPUFrequency = extractCPUFrequency(features)
		group.NetworkSpeed = extractNetworkSpeed(features)
		group.IPv4Count = extractIPv4Count(features)
		group.IPv6Subnet = extractIPv6Subnet(features)
	}

	return group
}

func getLegacyGroupType(legacyGroup legacy.ServiceGroup) string {
	mode := getStringFromPayload(legacyGroup.Payload, "mode")
	code := getStringFromPayload(legacyGroup.Payload, "code")
	role := getLegacyRole(legacyGroup.Role)

	// 1. Location - есть code (чистые локации)
	if code != "" {
		return "location"
	}

	// 2. Server groups - есть mode
	if mode != "" {
		return "server"
	}

	// 3. Special services
	if isSpecialService(legacyGroup.TypeObject.Slug) {
		return "special"
	}

	// 4. География (если есть role но нет code)
	if role == "location" {
		return "geography"
	}

	return "unknown"
}

// Получаем role из interface{}
func getLegacyRole(role interface{}) string {
	if role == nil {
		return ""
	}
	if str, ok := role.(string); ok {
		return str
	}
	return ""
}

// Универсальная функция для извлечения features из legacy API
func getLegacyFeatures(localed map[string]interface{}) string {
	if localed == nil {
		return ""
	}

	// Пробуем разные форматы features в legacy API

	// Формат 1: features как map[string]interface{}
	if features, exists := localed["features"]; exists {
		switch v := features.(type) {
		case map[string]interface{}:
			// Берем русскую версию
			if ru, exists := v["ru"]; exists {
				if str, ok := ru.(string); ok {
					return str
				}
			}
			// Или английскую
			if en, exists := v["en"]; exists {
				if str, ok := en.(string); ok {
					return str
				}
			}
		case string:
			// Формат 2: features как строка
			return v
		}
	}

	// Формат 3: prettyLocaledPayload.features
	if pretty, exists := localed["prettyLocaledPayload"]; exists {
		if prettyMap, ok := pretty.(map[string]interface{}); ok {
			if features, exists := prettyMap["features"]; exists {
				if str, ok := features.(string); ok {
					return str
				}
			}
		}
	}

	return ""
}

func getLegacyServerSubtype(legacyGroup legacy.ServiceGroup) string {
	mode := getStringFromPayload(legacyGroup.Payload, "mode")
	label := getStringFromPayload(legacyGroup.Payload, "label")

	if mode == "shared" {
		return "shared"
	}
	if mode == "dedicated" {
		return "dedicated"
	}

	// Определяем по label если mode нет
	if strings.Contains(label, "SHARED") {
		return "shared"
	}
	if strings.Contains(label, "DEDICATED") {
		return "dedicated"
	}

	return "server"
}

// Функция для получения bool из payload в legacy API
func getLegacyBoolFromPayload(payload map[string]interface{}, key string) bool {
	if val, exists := payload[key]; exists {
		switch v := val.(type) {
		case bool:
			return v
		case string:
			// В legacy API иногда bool может быть строкой
			return v == "true" || v == "1"
		}
	}
	return false
}

func ConvertLegacyServiceGroups(legacyGroups []legacy.ServiceGroup) []models.ServiceGroup {
	result := make([]models.ServiceGroup, len(legacyGroups))
	for i, legacyGroup := range legacyGroups {
		result[i] = ConvertLegacyServiceGroup(legacyGroup)
	}
	return result
}
