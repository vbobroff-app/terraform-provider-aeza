package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// ParseTimestamp парсит timestamp из строки
func ParseTimestamp(timestamp string) (time.Time, error) {
	// Пробуем разные форматы
	formats := []string{
		time.RFC3339,
		"2006-01-02T15:04:05.999999999Z07:00",
		"2006-01-02 15:04:05",
		time.RFC1123,
	}

	for _, format := range formats {
		if t, err := time.Parse(format, timestamp); err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("unable to parse timestamp: %s", timestamp)
}

// PriceToRubles конвертирует цену в копейках в рубли
func PriceToRubles(kopecks int) float64 {
	return float64(kopecks) / 100.0
}

// StringToInt безопасно конвертирует строку в int
func StringToInt(s string) (int, error) {
	if s == "" {
		return 0, nil
	}
	return strconv.Atoi(s)
}

// IntToString конвертирует int в строку
func IntToString(i int) string {
	return strconv.Itoa(i)
}

// MapToString преобразует map в строку для отладки
func MapToString(m map[string]interface{}) string {
	if m == nil {
		return "nil"
	}

	var parts []string
	for k, v := range m {
		parts = append(parts, fmt.Sprintf("%s:%v", k, v))
	}
	return "{" + strings.Join(parts, ", ") + "}"
}
