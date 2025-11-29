// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/utils/format_utils.go
package utils

import (
	"fmt"
	"strings"
	"time"
)

// FormatPrice форматирует цену в "X,XX €"
func FormatPrice(price float64) string {
	// Предполагаем, что цена в копейках/центах (494 = 4.94 €)
	euros := float64(price) / 100.0
	// Заменяем точку на запятую для европейского формата
	formatted := fmt.Sprintf("%.2f", euros)
	formatted = strings.Replace(formatted, ".", ",", 1)
	return fmt.Sprintf("%s €", formatted)
}

// FormatDate преобразует ISO дату в формат "DD.MM.YYYY THH:MM:SS.MSZ"
func FormatDate(isoDate string) string {
	// Парсим ISO дату
	t, err := time.Parse(time.RFC3339, isoDate)
	if err != nil {
		return isoDate
	}

	// Форматируем в "DD.MM.YYYY THH:MM:SS.MSZ"
	return t.Format("02.01.2006 T15:04:05.000Z")
}

// FormatDateFromUnix преобразует Unix timestamp в формат "DD.MM.YYYY THH:MM:SS.MSZ"
func FormatDateFromUnix(unixTimestamp int64) string {
	if unixTimestamp == 0 {
		return ""
	}

	t := time.Unix(unixTimestamp, 0)
	return t.Format("02.01.2006 T15:04:05.000Z")
}

// stringToPtr преобразует строку в указатель
func stringToPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
