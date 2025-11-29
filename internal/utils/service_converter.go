// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/utils/service_converter.go
package utils

import (
	"fmt"
	"time"

	"github.com/vbobroff-app/terraform-provider-aeza/internal/models"
	"github.com/vbobroff-app/terraform-provider-aeza/internal/models/legacy"
	"github.com/vbobroff-app/terraform-provider-aeza/internal/models/next"
)

func ConvertNextService(nextService next.Service) models.TerraformService {
	return models.TerraformService{
		ID:           nextService.ID,
		Name:         nextService.Name,
		IP:           nextService.IP,
		Price:        nextService.Price,
		PaymentTerm:  nextService.PaymentTerm,
		AutoProlong:  nextService.AutoProlong,
		CreatedAt:    nextService.CreatedAt,
		ExpiresAt:    nextService.ExpiresAt,
		Status:       nextService.Status,
		TypeSlug:     nextService.TypeSlug,
		ProductName:  nextService.ProductName,
		LocationCode: nextService.LocationCode,
	}
}

func ConvertLegacyService(legacyService legacy.ServiceVPS) models.TerraformService {
	tfService := models.TerraformService{
		ID:           legacyService.ID,
		Name:         legacyService.Name,
		IP:           legacyService.IP,
		PaymentTerm:  legacyService.PaymentTerm,
		AutoProlong:  legacyService.AutoProlong,
		Status:       legacyService.Status,
		LocationCode: legacyService.LocationCode,
	}

	// Получение ProductName из вложенного Product
	if legacyService.Product.Name != "" {
		tfService.ProductName = legacyService.Product.Name
	} else {
		tfService.ProductName = "Unknown"
	}

	// Получение TypeSlug из продукта
	if legacyService.Product.Type != "" {
		tfService.TypeSlug = legacyService.Product.Type
	} else {
		tfService.TypeSlug = "unknown"
	}

	// Преобразование времени
	tfService.CreatedAt = convertUnixToISO(legacyService.Timestamps.CreatedAt)
	tfService.ExpiresAt = convertUnixToISO(legacyService.Timestamps.ExpiresAt)

	// Расчет цены из доступных источников
	tfService.Price = calculatePriceFromLegacy(legacyService)

	return tfService
}

// calculatePriceFromLegacy вычисляет актуальную цену из Legacy структуры
func calculatePriceFromLegacy(legacyService legacy.ServiceVPS) float64 {
	// Приоритет 1: individualPrices
	if legacyService.IndividualPrices != nil {
		if price, exists := legacyService.IndividualPrices[legacyService.PaymentTerm]; exists {
			return float64(price)
		}
	}

	// Приоритет 2: rawPrices услуги
	if legacyService.RawPrices != nil {
		if price, exists := legacyService.RawPrices[legacyService.PaymentTerm]; exists {
			return float64(price)
		}
	}

	// Приоритет 3: prices из продукта
	if legacyService.Product.Prices.Hour > 0 || legacyService.Product.Prices.Month > 0 {
		switch legacyService.PaymentTerm {
		case "hour":
			return legacyService.Product.Prices.Hour
		case "month":
			return legacyService.Product.Prices.Month
		case "year":
			return legacyService.Product.Prices.Year
		case "half_year":
			return legacyService.Product.Prices.HalfYear
		case "quarter_year":
			return legacyService.Product.Prices.QuarterYear
		}
	}

	// Приоритет 4: rawPrices продукта
	if legacyService.Product.RawPrices != (legacy.ProductPrices{}) {
		switch legacyService.PaymentTerm {
		case "hour":
			if legacyService.Product.RawPrices.Hour > 0 {
				return legacyService.Product.RawPrices.Hour
			}
		case "month":
			if legacyService.Product.RawPrices.Month > 0 {
				return legacyService.Product.RawPrices.Month
			}
		case "year":
			if legacyService.Product.RawPrices.Year > 0 {
				return legacyService.Product.RawPrices.Year
			}
		case "half_year":
			if legacyService.Product.RawPrices.HalfYear > 0 {
				return legacyService.Product.RawPrices.HalfYear
			}
		case "quarter_year":
			if legacyService.Product.RawPrices.QuarterYear > 0 {
				return legacyService.Product.RawPrices.QuarterYear
			}
		}
	}

	// Приоритет 5: individualPrices продукта
	if legacyService.Product.IndividualPrices != (legacy.ProductPrices{}) {
		switch legacyService.PaymentTerm {
		case "hour":
			if legacyService.Product.IndividualPrices.Hour > 0 {
				return legacyService.Product.IndividualPrices.Hour
			}
		case "month":
			if legacyService.Product.IndividualPrices.Month > 0 {
				return legacyService.Product.IndividualPrices.Month
			}
		case "year":
			if legacyService.Product.IndividualPrices.Year > 0 {
				return legacyService.Product.IndividualPrices.Year
			}
		case "half_year":
			if legacyService.Product.IndividualPrices.HalfYear > 0 {
				return legacyService.Product.IndividualPrices.HalfYear
			}
		case "quarter_year":
			if legacyService.Product.IndividualPrices.QuarterYear > 0 {
				return legacyService.Product.IndividualPrices.QuarterYear
			}
		}
	}

	// Если цены не найдены, возвращаем 0
	return 0
}

// ConvertNextServices преобразует список служб из API v2
func ConvertNextServices(nextServices []next.Service) []models.TerraformService {
	var result []models.TerraformService
	for _, service := range nextServices {
		result = append(result, ConvertNextService(service))
	}
	return result
}

// ConvertLegacyServices преобразует список служб из Legacy API
func ConvertLegacyServices(legacyServices []legacy.ServiceVPS) []models.TerraformService {
	var result []models.TerraformService
	for _, service := range legacyServices {
		result = append(result, ConvertLegacyService(service))
	}
	return result
}

// ConvertLegacyVPSToTerraform преобразует legacy.ServiceVPS в TerraformService
func ConvertLegacyVPSToTerraform(vpsService legacy.ServiceVPS) models.TerraformServiceDetailed {
	tfService := models.TerraformServiceDetailed{
		// Основная идентификация
		ID:           vpsService.ID,
		Name:         vpsService.Name,
		Status:       vpsService.Status,
		LocationCode: vpsService.LocationCode,
		AutoProlong:  vpsService.AutoProlong,
		Backups:      vpsService.Backups,
		PaymentTerm:  vpsService.PaymentTerm,
		Capabilities: vpsService.Capabilities,
		Payload:      vpsService.Payload,

		// Сетевые настройки
		IP:   vpsService.IP,
		IPs:  extractIPs(vpsService.IPs),
		IPv6: extractIPv6(vpsService.IPv6),

		// Продукт
		ProductName: vpsService.Product.Name,
		TypeSlug:    getTypeSlug(vpsService.Product),

		// Программное обеспечение - для ServiceVPS Parameters это структура legacy.Parameters
		OS:       getOSDisplayName(vpsService.Parameters.OS),
		Username: vpsService.Parameters.Username,
	}

	// Стоимость (евро)
	tfService.Price = getPriceInEuros(vpsService.RawPrices, vpsService.PaymentTerm)
	tfService.PriceDisplay = formatPriceDisplay(tfService.Price, vpsService.PaymentTerm)

	// Конфигурация оборудования
	if vpsService.SummaryConfiguration != nil {
		tfService.CPU = getCPUCount(vpsService.SummaryConfiguration["cpu"])
		tfService.RAM = getRAMCount(vpsService.SummaryConfiguration["ram"])
		tfService.Storage = getStorageCount(vpsService.SummaryConfiguration["rom"])
		tfService.IPCount = getIPCount(vpsService.SummaryConfiguration["ip"])
	}

	// Временные метки
	tfService.CreatedAt = convertUnixToISO(vpsService.Timestamps.CreatedAt)
	tfService.ExpiresAt = convertUnixToISO(vpsService.Timestamps.ExpiresAt)
	tfService.CreatedDate = formatDisplayDate(vpsService.Timestamps.CreatedAt)
	tfService.ExpiresDate = formatDisplayDate(vpsService.Timestamps.ExpiresAt)

	return tfService
}

// ConvertNextToTerraform преобразует next.Service в TerraformService
func ConvertNextToTerraform(nextService next.Service) models.TerraformServiceDetailed {
	return models.TerraformServiceDetailed{
		// Основная идентификация
		ID:           nextService.ID,
		Name:         nextService.Name,
		Status:       nextService.Status,
		LocationCode: nextService.LocationCode,
		AutoProlong:  nextService.AutoProlong,
		PaymentTerm:  nextService.PaymentTerm,
		Capabilities: nextService.Capabilities,
		Payload:      nextService.Payload,

		// Сетевые настройки
		IP: nextService.IP,

		// Продукт
		ProductName: nextService.ProductName,
		TypeSlug:    nextService.TypeSlug,

		// Стоимость (евро)
		Price:        float64(nextService.Price) / 100, // Конвертируем из центов
		PriceDisplay: formatPriceDisplay(float64(nextService.Price)/100, nextService.PaymentTerm),

		// Временные метки
		CreatedAt:   nextService.CreatedAt,
		ExpiresAt:   nextService.ExpiresAt,
		CreatedDate: formatISOToDisplayDate(nextService.CreatedAt),
		ExpiresDate: formatISOToDisplayDate(nextService.ExpiresAt),
	}
}

// Вспомогательные функции

// getPriceInEuros возвращает цену в евро
func getPriceInEuros(rawPrices map[string]int, paymentTerm string) float64 {
	if rawPrices == nil {
		return 0
	}

	priceCents := rawPrices[paymentTerm]
	if priceCents == 0 && paymentTerm != "month" {
		priceCents = rawPrices["month"] // fallback на месяц
	}

	return float64(priceCents) / 100
}

// formatPriceDisplay форматирует отображение цены
func formatPriceDisplay(price float64, paymentTerm string) string {
	periods := map[string]string{
		"hour":         "Час",
		"month":        "Месяц",
		"year":         "Год",
		"half_year":    "Полгода",
		"quarter_year": "Квартал",
	}

	period := periods[paymentTerm]
	if period == "" {
		period = paymentTerm
	}

	return fmt.Sprintf("%.2f € / %s", price, period)
}

// extractIPs извлекает все IP адреса
func extractIPs(ips []legacy.IPAddress) []string {
	var result []string
	for _, ip := range ips {
		result = append(result, ip.Value)
	}
	return result
}

// extractIPv6 извлекает IPv6 подсеть
func extractIPv6(ipv6 []legacy.IPv6Address) string {
	if len(ipv6) > 0 {
		return ipv6[0].Value
	}
	return ""
}

// getTypeSlug определяет тип услуги
func getTypeSlug(product legacy.Product) string {
	if product.Type != "" {
		return product.Type
	}
	if product.TypeObject.Slug != "" {
		return product.TypeObject.Slug
	}
	return "unknown"
}

// Конфигурация оборудования
func getCPUCount(config legacy.ConfigurationItem) int {
	return config.Count
}

func getRAMCount(config legacy.ConfigurationItem) int {
	return config.Count
}

func getStorageCount(config legacy.ConfigurationItem) int {
	return config.Count
}

func getIPCount(config legacy.ConfigurationItem) int {
	return config.Count
}

// OS mapping
func getOSDisplayName(osSlug string) string {
	mapping := map[string]string{
		"ubuntu_2404":  "Ubuntu 24.04",
		"ubuntu_2204":  "Ubuntu 22.04",
		"ubuntu_2004":  "Ubuntu 20.04",
		"centos_7":     "CentOS 7",
		"debian_11":    "Debian 11",
		"windows_2019": "Windows Server 2019",
	}

	if name, exists := mapping[osSlug]; exists {
		return name
	}
	return osSlug
}

// Форматирование дат
func convertUnixToISO(unixTimestamp int64) string {
	if unixTimestamp == 0 {
		return ""
	}
	return time.Unix(unixTimestamp, 0).UTC().Format(time.RFC3339)
}

func formatDisplayDate(unixTimestamp int64) string {
	if unixTimestamp == 0 {
		return ""
	}

	t := time.Unix(unixTimestamp, 0)
	monthNames := map[time.Month]string{
		1: "января", 2: "февраля", 3: "марта", 4: "апреля",
		5: "мая", 6: "июня", 7: "июля", 8: "августа",
		9: "сентября", 10: "октября", 11: "ноября", 12: "декабря",
	}

	return fmt.Sprintf("%d %s %d г.", t.Day(), monthNames[t.Month()], t.Year())
}

func formatISOToDisplayDate(isoDate string) string {
	if isoDate == "" {
		return ""
	}

	t, err := time.Parse(time.RFC3339, isoDate)
	if err != nil {
		return isoDate
	}

	monthNames := map[time.Month]string{
		1: "января", 2: "февраля", 3: "марта", 4: "апреля",
		5: "мая", 6: "июня", 7: "июля", 8: "августа",
		9: "сентября", 10: "октября", 11: "ноября", 12: "декабря",
	}

	return fmt.Sprintf("%d %s %d г.", t.Day(), monthNames[t.Month()], t.Year())
}
