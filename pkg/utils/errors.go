package utils

import (
	"fmt"
	"strings"
)

// WrapError оборачивает ошибку с контекстом
func WrapError(err error, context string) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", context, err)
}

// WrapErrorf оборачивает ошибку с форматированным контекстом
func WrapErrorf(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	context := fmt.Sprintf(format, args...)
	return fmt.Errorf("%s: %w", context, err)
}

// IsNotFoundError проверяет, является ли ошибка ошибкой "не найдено"
func IsNotFoundError(err error) bool {
	if err == nil {
		return false
	}

	// Проверяем сообщение об ошибке на наличие ключевых слов
	errorStr := err.Error()
	return strings.Contains(errorStr, "not found") ||
		strings.Contains(errorStr, "not exist") ||
		strings.Contains(errorStr, "404")
}
