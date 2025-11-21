// internal/utils/service_converter.go
package utils

import (
	"strings"

	"github.com/vbobroff-app/terraform-provider-aeza/internal/models"
	"github.com/vbobroff-app/terraform-provider-aeza/internal/models/next"
)

// ConvertNextServiceGroup преобразует группу услуг из API v2 в Terraform ServiceGroup
func ConvertNextServiceGroup(apiGroup next.ServiceGroup) models.ServiceGroup {
	return models.ServiceGroup{
		ID:             apiGroup.ID,
		Name:           apiGroup.Name,
		Type:           apiGroup.Type.Slug,
		ServiceHandler: apiGroup.Type.ServiceHandler,
		Description:    apiGroup.Description,
		Location:       getStringFromPayload(apiGroup.Payload, "label"),
		CountryCode:    getStringFromPayload(apiGroup.Payload, "code"),
		ServerType:     getStringFromPayload(apiGroup.Payload, "mode"),
		IsDisabled:     getBoolFromPayload(apiGroup.Payload, "isDisabled"),
		Features:       getFeatures(apiGroup.LocaledPayload),
		CPUModel:       getCPUModel(apiGroup.LocaledPayload),
		CPUFrequency:   getCPUFrequency(apiGroup.LocaledPayload),
		NetworkSpeed:   getNetworkSpeed(apiGroup.LocaledPayload),
		IPv4Count:      getIPv4Count(apiGroup.LocaledPayload),
		IPv6Subnet:     getIPv6Subnet(apiGroup.LocaledPayload),
	}
}

// Универсальная функция для извлечения features
func getFeatures(localed map[string]interface{}) string {
	if localed == nil {
		return ""
	}

	// Пробуем разные форматы features

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

	return ""
}

// Остальные вспомогательные функции
func getStringFromPayload(payload map[string]interface{}, key string) string {
	if val, exists := payload[key]; exists {
		if str, ok := val.(string); ok {
			return str
		}
	}
	return ""
}

func getBoolFromPayload(payload map[string]interface{}, key string) bool {
	if val, exists := payload[key]; exists {
		if b, ok := val.(bool); ok {
			return b
		}
	}
	return false
}

func getCPUModel(localed map[string]interface{}) string {
	features := getFeatures(localed)
	return extractCPUModel(features)
}

func getCPUFrequency(localed map[string]interface{}) string {
	features := getFeatures(localed)
	return extractCPUFrequency(features)
}

func getNetworkSpeed(localed map[string]interface{}) string {
	features := getFeatures(localed)
	return extractNetworkSpeed(features)
}

func getIPv4Count(localed map[string]interface{}) int {
	features := getFeatures(localed)
	return extractIPv4Count(features)
}

func getIPv6Subnet(localed map[string]interface{}) string {
	features := getFeatures(localed)
	return extractIPv6Subnet(features)
}

// Функции парсинга (оставляем без изменений)
func extractCPUModel(features string) string {
	lines := strings.Split(features, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.Contains(line, "Процессор") {
			return strings.TrimPrefix(line, "Процессор ")
		}
		if strings.Contains(line, "Processor") {
			return strings.TrimPrefix(line, "Processor ")
		}
	}
	return ""
}

func extractCPUFrequency(features string) string {
	lines := strings.Split(features, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.Contains(line, "Частота") {
			return strings.TrimPrefix(line, "Частота ")
		}
		if strings.Contains(line, "Frequency") {
			return strings.TrimPrefix(line, "Frequency ")
		}
	}
	return ""
}

func extractNetworkSpeed(features string) string {
	lines := strings.Split(features, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.Contains(line, "Интернет до") {
			return strings.TrimPrefix(line, "Интернет до ")
		}
		if strings.Contains(line, "Internet up to") {
			return strings.TrimPrefix(line, "Internet up to ")
		}
	}
	return ""
}

func extractIPv4Count(features string) int {
	lines := strings.Split(features, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.Contains(line, "1 адрес IPv4") || strings.Contains(line, "1 IPv4 address") {
			return 1
		}
	}
	return 0
}

func extractIPv6Subnet(features string) string {
	lines := strings.Split(features, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.Contains(line, "/48") {
			return "/48"
		}
		if strings.Contains(line, "/64") {
			return "/64"
		}
	}
	return ""
}

// ConvertNextServiceGroups преобразует массив групп услуг из API v2 в Terraform ServiceGroup
func ConvertNextServiceGroups(nextGroups []next.ServiceGroup) []models.ServiceGroup {
	result := make([]models.ServiceGroup, len(nextGroups))
	for i, apiGroup := range nextGroups {
		result[i] = ConvertNextServiceGroup(apiGroup)
	}
	return result
}
