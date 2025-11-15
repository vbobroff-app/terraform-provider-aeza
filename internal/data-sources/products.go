package data_sources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vbobroff-app/terraform-provider-aeza/internal/client"
	"github.com/vbobroff-app/terraform-provider-aeza/internal/provider"
)

func ProductsDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving list of Aeza products",

		ReadContext: productsRead,

		Schema: map[string]*schema.Schema{
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter products by type (vps, hicpu, etc.)",
			},
			"location": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter products by location",
			},
			"products": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of products",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Product ID",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Product name",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Product type",
						},
						"cpu": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "CPU cores",
						},
						"ram": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "RAM in MB",
						},
						"disk": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Disk space in GB",
						},
						"price_hourly": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Hourly price in kopecks",
						},
						"price_monthly": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Monthly price in kopecks",
						},
						"location": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Product location",
						},
						"mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Product mode (shared, dedicated, etc.)",
						},
					},
				},
			},
		},
	}
}

func productsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	config := meta.(*provider.Config)
	apiClient, err := config.Client()
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to create client: %w", err))
	}

	products, err := apiClient.GetProducts(ctx)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to get products: %w", err))
	}

	filterType := d.Get("type").(string)
	filterLocation := d.Get("location").(string)

	var filteredProducts []client.Product
	for _, product := range products.Data.Items {
		// Apply filters
		if filterType != "" && product.Type != filterType {
			continue
		}
		if filterLocation != "" {
			if location, ok := product.Group.Payload["label"].(string); ok {
				if location != filterLocation {
					continue
				}
			}
		}
		filteredProducts = append(filteredProducts, product)
	}

	productList := make([]map[string]interface{}, len(filteredProducts))
	for i, product := range filteredProducts {
		// Extract configuration
		cpu, ram, disk := extractProductConfiguration(product.Configuration)

		// Extract location from group payload
		location := ""
		if label, ok := product.Group.Payload["label"].(string); ok {
			location = label
		}

		// Extract mode from group payload
		mode := ""
		if modeVal, ok := product.Group.Payload["mode"].(string); ok {
			mode = modeVal
		}

		productList[i] = map[string]interface{}{
			"id":            product.ID,
			"name":          product.Name,
			"type":          product.Type,
			"cpu":           cpu,
			"ram":           ram,
			"disk":          disk,
			"price_hourly":  product.Prices.Hour,
			"price_monthly": product.Prices.Month,
			"location":      location,
			"mode":          mode,
		}
	}

	if err := d.Set("products", productList); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set products: %w", err))
	}

	d.SetId("products")

	return diags
}

func extractProductConfiguration(config []client.ProductConfig) (cpu, ram, disk int) {
	for _, item := range config {
		switch item.Slug {
		case "cpu":
			cpu = item.Base
		case "ram":
			ram = item.Base
		case "rom":
			disk = item.Base
		}
	}
	return
}
