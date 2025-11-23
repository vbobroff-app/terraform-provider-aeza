// internal/provider/provider.go
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vbobroff-app/terraform-provider-aeza/internal/client"
	data_sources "github.com/vbobroff-app/terraform-provider-aeza/internal/data-sources"
	"github.com/vbobroff-app/terraform-provider-aeza/internal/resources"
)

func New() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("AEZA_API_KEY", nil),
				Description: "API Key for Aeza provider",
				Sensitive:   true,
			},
			"base_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("AEZA_BASE_URL", "https://my.aeza.net/api"),
				Description: "Base URL for Aeza API",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"aeza_service": resources.ServiceResource(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			"aeza_services":       data_sources.ServicesDataSource(),
			"aeza_products":       data_sources.ProductsDataSource(),
			"aeza_service_types":  data_sources.ServiceTypesDataSource(),
			"aeza_service_groups": data_sources.ServiceGroupsData(),
			"aeza_os_list":        data_sources.OSDataSource(),
		},

		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	apiKey := d.Get("api_key").(string)
	baseURL := d.Get("base_url").(string)

	var diags diag.Diagnostics

	// Добавляем валидацию
	if apiKey == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "API Key is required",
		})
		return nil, diags
	}

	client, err := client.NewClient(baseURL, apiKey)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create Aeza client",
			Detail:   err.Error(),
		})
		return nil, diags
	}

	return client, diags
}
