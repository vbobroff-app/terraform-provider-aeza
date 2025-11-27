// internal/utils/service_create_utils.go
package utils

import (
	"github.com/vbobroff-app/terraform-provider-aeza/internal/models"
	"github.com/vbobroff-app/terraform-provider-aeza/internal/models/legacy"
)

// ConvertToLegacyServiceCreateRequest конвертирует Terraform модель в Legacy API запрос
func ConvertToLegacy_ServiceCreateRequest(req models.ServiceCreateRequest) legacy.ServiceCreateRequest {
	return legacy.ServiceCreateRequest{
		Method:      "balance",
		Count:       1,
		Term:        req.PaymentTerm,
		Name:        req.Name,
		ProductID:   req.ProductID,
		AutoProlong: req.AutoProlong,
		Backups:     false,
		Parameters: legacy.ServiceCreateParameters{
			OS:     req.OS,
			Recipe: stringToPtr(req.Recipe),
			IsoURL: req.IsoURL,
		},
	}
}

// ConvertFromLegacy_ServiceCreateResponse конвертирует Legacy API ответ в Terraform модель
func ConvertFromLegacy_ServiceCreateResponse(resp legacy.ServiceCreateResponse) models.ServiceCreateResponse {
	// Берем первый item из ответа (обычно там один)
	if len(resp.Data.Items) == 0 {
		return models.ServiceCreateResponse{
			ID:     0,
			Status: "error",
			Date:   FormatDateFromUnix(0),
		}
	}

	item := resp.Data.Items[0]
	transaction := resp.Data.Transaction

	// ID созданной услуги находится в createdServiceIds
	var serviceID int64
	if len(item.CreatedServiceIds) > 0 {
		serviceID = item.CreatedServiceIds[0]
	}

	// Извлекаем LocationName из product.group.payload.label
	locationName := extractLocationName(item.Product)

	// Конвертируем цены из копеек в евро
	price := FormatPrice(item.IndividualPrice)
	transactionAmount := FormatPrice(transaction.Amount)

	return models.ServiceCreateResponse{
		ID:                serviceID,
		OrderID:           item.ID,
		Status:            item.Status,
		Date:              FormatDateFromUnix(item.CreatedAt),
		ProductId:         item.ProductID,
		ProductType:       item.Product.Type,
		GroupId:           item.Product.GroupID,
		ProductName:       item.Product.Name,
		LocationName:      locationName,
		Term:              item.Term,
		Price:             price,
		TransactionAmount: transactionAmount,
	}
}

// extractLocationName извлекает название локации из структуры продукта
func extractLocationName(product legacy.Product) string {
	if product.Group.Payload != nil {
		if label, exists := product.Group.Payload["label"]; exists {
			if labelStr, ok := label.(string); ok {
				return labelStr
			}
		}
	}
	return ""
}
