package data_sources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vbobroff-app/terraform-provider-aeza/internal/provider"
)

func ServiceTypesDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving list of Aeza service types",

		ReadContext: serviceTypesRead,

		Schema: map[string]*schema.Schema{
			"types": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of service types",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"slug": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Service type slug",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Service type name",
						},
						"order": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Display order",
						},
					},
				},
			},
		},
	}
}

func serviceTypesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	config := meta.(*provider.Config)
	apiClient, err := config.Client()
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to create client: %w", err))
	}

	serviceTypes, err := apiClient.GetServiceTypes(ctx)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to get service types: %w", err))
	}

	typeList := make([]map[string]interface{}, len(serviceTypes.Data.Items))
	for i, serviceType := range serviceTypes.Data.Items {
		typeList[i] = map[string]interface{}{
			"slug":  serviceType.Slug,
			"name":  serviceType.Name,
			"order": serviceType.Order,
		}
	}

	if err := d.Set("types", typeList); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set service types: %w", err))
	}

	d.SetId("service_types")

	return diags
}
