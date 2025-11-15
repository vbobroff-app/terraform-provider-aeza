package data_sources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vbobroff-app/terraform-provider-aeza/internal/provider"
)

func ServicesDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving list of Aeza services",

		ReadContext: servicesRead,

		Schema: map[string]*schema.Schema{
			"services": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of services",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Service ID",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Service name",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Service type (vps, hicpu, etc.)",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Service status",
						},
						"product_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Product name",
						},
						"location_code": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Location code",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Creation timestamp",
						},
					},
				},
			},
		},
	}
}

func servicesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	config := meta.(*provider.Config)
	apiClient, err := config.Client()
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to create client: %w", err))
	}

	services, err := apiClient.GetServices(ctx)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to get services: %w", err))
	}

	serviceList := make([]map[string]interface{}, len(services.Data.Items))
	for i, service := range services.Data.Items {
		serviceList[i] = map[string]interface{}{
			"id":            service.ID,
			"name":          service.Name,
			"type":          service.Type,
			"status":        service.Status,
			"product_name":  service.ProductName,
			"location_code": service.LocationCode,
			"created_at":    service.CreatedAt.Format("2006-01-02T15:04:05Z"),
		}
	}

	if err := d.Set("services", serviceList); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set services: %w", err))
	}

	d.SetId("services")

	return diags
}
