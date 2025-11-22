// internal/utils/product_converter.go
package utils

import (
	"github.com/vbobroff-app/terraform-provider-aeza/internal/models"
	"github.com/vbobroff-app/terraform-provider-aeza/internal/models/legacy"
)

func ConvertLegacyProduct(legacyProduct legacy.Product) models.Product {
	product := models.Product{
		ID:             int64(legacyProduct.ID),
		Type:           legacyProduct.Type,
		Name:           legacyProduct.Name,
		ServiceHandler: legacyProduct.ServiceHandler,
	}

	// GroupID (может быть null в legacy)
	if legacyProduct.GroupID != nil {
		product.GroupID = int64(*legacyProduct.GroupID)
	}

	// Prices - конвертируем ProductPrices в map[string]float64
	product.Prices = make(map[string]float64)
	if legacyProduct.Prices.Hour > 0 {
		product.Prices["hour"] = float64(legacyProduct.Prices.Hour)
	}
	if legacyProduct.Prices.Month > 0 {
		product.Prices["month"] = float64(legacyProduct.Prices.Month)
	}
	if legacyProduct.Prices.Year > 0 {
		product.Prices["year"] = float64(legacyProduct.Prices.Year)
	}
	if legacyProduct.Prices.HalfYear > 0 {
		product.Prices["half_year"] = float64(legacyProduct.Prices.HalfYear)
	}
	if legacyProduct.Prices.QuarterYear > 0 {
		product.Prices["quarter_year"] = float64(legacyProduct.Prices.QuarterYear)
	}

	// Configuration
	product.Configuration = make([]models.ConfigurationItem, len(legacyProduct.Configuration))
	for i, config := range legacyProduct.Configuration {
		product.Configuration[i] = models.ConfigurationItem{
			Max:    float64(config.Max),
			Base:   float64(config.Base),
			Slug:   config.Slug,
			Type:   config.Type,
			Prices: config.Prices,
		}
	}

	// SummaryConfiguration
	product.SummaryConfiguration = make(map[string]models.ConfigurationSummary)
	for key, summary := range legacyProduct.SummaryConfiguration {
		if summaryMap, ok := summary.(map[string]interface{}); ok {
			configSummary := models.ConfigurationSummary{
				Slug: getStringValue(summaryMap, "slug"),
				Type: getStringValue(summaryMap, "type"),
			}

			// Обрабатываем числа (могут быть int или float64)
			if max, ok := summaryMap["max"]; ok {
				if f, ok := max.(float64); ok {
					configSummary.Max = f
				} else if i, ok := max.(int); ok {
					configSummary.Max = float64(i)
				} else if i, ok := max.(int64); ok {
					configSummary.Max = float64(i)
				}
			}

			if base, ok := summaryMap["base"]; ok {
				if f, ok := base.(float64); ok {
					configSummary.Base = f
				} else if i, ok := base.(int); ok {
					configSummary.Base = float64(i)
				} else if i, ok := base.(int64); ok {
					configSummary.Base = float64(i)
				}
			}

			product.SummaryConfiguration[key] = configSummary
		}
	}

	return product
}

// Вспомогательная функция для безопасного получения строковых значений
func getStringValue(m map[string]interface{}, key string) string {
	if val, ok := m[key]; ok {
		if str, ok := val.(string); ok {
			return str
		}
	}
	return ""
}
