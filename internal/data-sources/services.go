// internal/data-sources/services.go
package data_sources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vbobroff-app/terraform-provider-aeza/internal/interfaces"
)

// Убираем зависимость от provider, используем интерфейс
type servicesDataSource struct {
	client interfaces.DataClient
}

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
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"current_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"product_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"owner_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"location_code": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"payment_term": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"auto_prolong": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"backups": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"created_at": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"updated_at": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func servicesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(interfaces.DataClient)

	services, err := client.ListServices(ctx)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error fetching services: %w", err))
	}

	serviceList := make([]map[string]interface{}, len(services))
	for i, service := range services {
		serviceList[i] = map[string]interface{}{
			"id":             service.ID,
			"name":           service.Name,
			"status":         service.Status,
			"current_status": service.CurrentStatus,
			"last_status":    service.LastStatus,
			"ip":             service.IP,
			"product_id":     service.ProductID,
			"owner_id":       service.OwnerID,
			"location_code":  service.LocationCode,
			"payment_term":   service.PaymentTerm,
			"auto_prolong":   service.AutoProlong,
			"backups":        service.Backups,
			"created_at":     service.CreatedAt,
			"updated_at":     service.UpdatedAt,
		}
	}

	if err := d.Set("services", serviceList); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("services")
	return nil
}
