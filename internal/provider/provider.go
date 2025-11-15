package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vbobroff-app/terraform-provider-aeza/internal/data-sources"
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
				Default:     "https://my.aeza.net/api",
				Description: "Base URL for Aeza API",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"aeza_service": resources.ServiceResource(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			"aeza_services":      data_sources.ServicesDataSource(),
			"aeza_products":      data_sources.ProductsDataSource(),
			"aeza_service_types": data_sources.ServiceTypesDataSource(),
		},

		ConfigureContextFunc: configureProvider,
	}
}

func configureProvider(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	config := &Config{
		APIKey:  d.Get("api_key").(string),
		BaseURL: d.Get("base_url").(string),
	}

	return config, nil
}
