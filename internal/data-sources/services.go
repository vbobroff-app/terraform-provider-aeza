// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/data-sources/services.go
package data_sources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vbobroff-app/terraform-provider-aeza/internal/interfaces"
	"github.com/vbobroff-app/terraform-provider-aeza/internal/models"
	"github.com/vbobroff-app/terraform-provider-aeza/internal/utils"
)

func ServicesDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving list of Aeza services",

		ReadContext: servicesRead,

		Schema: map[string]*schema.Schema{
			"services": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Порядок здесь определяет порядок вывода в Terraform!
						"id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Unique identifier of the service",
						},
						"type_slug": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of service (vps, dns, vpn, etc.)",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the service",
						},
						"ip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Primary IP address of the service",
						},
						"price_raw": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Price in cents for calculations",
						},
						"price_display": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Formatted price with currency",
						},
						"payment_term": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Payment term (hour, month, year, etc.)",
						},
						"auto_prolong": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether auto-prolong is enabled",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Service creation timestamp in formatted format",
						},
						"expires_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Service expiration timestamp in formatted format",
						},
						"product_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the product",
						},
						"location_code": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Location code where service is deployed",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Current status of the service",
						},
					},
				},
			},
		},
	}
}

func servicesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(interfaces.DataClient)

	// Получаем услуги в формате TerraformService
	services, err := client.ListServices(ctx)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error fetching services: %w", err))
	}

	// Преобразуем в формат для Terraform state
	if err := flattenTerraformServices(d, services); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("services")
	return nil
}

// flattenTerraformServices преобразует TerraformService в Terraform state с форматированием
func flattenTerraformServices(d *schema.ResourceData, services []models.TerraformService) error {
	serviceList := make([]map[string]interface{}, len(services))

	for i, service := range services {
		serviceList[i] = map[string]interface{}{
			// Важно: порядок должен соответствовать порядку в схеме выше!
			"id":            service.ID,
			"type_slug":     service.TypeSlug,
			"name":          service.Name,
			"ip":            service.IP,
			"price_raw":     service.Price,                    // Числовая цена для вычислений
			"price_display": utils.FormatPrice(service.Price), // Форматированная цена для отображения
			"payment_term":  service.PaymentTerm,
			"auto_prolong":  service.AutoProlong,
			"created_at":    utils.FormatDate(service.CreatedAt),
			"expires_at":    utils.FormatDate(service.ExpiresAt),
			"product_name":  service.ProductName,
			"location_code": service.LocationCode,
			"status":        service.Status,
		}
	}

	if err := d.Set("services", serviceList); err != nil {
		return fmt.Errorf("error setting services: %v", err)
	}

	return nil
}
