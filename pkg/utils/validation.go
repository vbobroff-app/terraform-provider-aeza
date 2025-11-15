package utils

import (
	"fmt"
	"regexp"
	"strings"
)

// ValidateServiceName проверяет корректность имени сервиса
func ValidateServiceName(name string) error {
	if len(name) < 1 || len(name) > 64 {
		return fmt.Errorf("service name must be between 1 and 64 characters")
	}

	// Проверяем на допустимые символы
	validChars := regexp.MustCompile(`^[a-zA-Z0-9\-_\. ]+$`)
	if !validChars.MatchString(name) {
		return fmt.Errorf("service name can only contain alphanumeric characters, hyphens, underscores, dots and spaces")
	}

	return nil
}

// ValidateProductID проверяет корректность ID продукта
func ValidateProductID(productID int) error {
	if productID <= 0 {
		return fmt.Errorf("product ID must be positive")
	}
	return nil
}

// NormalizeServiceType нормализует тип сервиса
func NormalizeServiceType(serviceType string) string {
	return strings.ToLower(strings.TrimSpace(serviceType))
}

// IsValidServiceType проверяет валидность типа сервиса
func IsValidServiceType(serviceType string) bool {
	validTypes := map[string]bool{
		"vps":       true,
		"hicpu":     true,
		"dedicated": true,
		"domain":    true,
		"vpn":       true,
		"soft":      true,
	}
	return validTypes[NormalizeServiceType(serviceType)]
}
