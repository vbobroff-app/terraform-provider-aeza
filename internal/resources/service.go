package resources

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/vbobroff-app/terraform-provider-aeza/internal/client"
	"github.com/vbobroff-app/terraform-provider-aeza/internal/provider"
)

func ServiceResource() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for managing Aeza services",

		CreateContext: serviceCreate,
		ReadContext:   serviceRead,
		DeleteContext: serviceDelete,

		Schema: map[string]*schema.Schema{
			"product_id": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "Product ID for the service",
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  "Name of the service",
				ValidateFunc: validation.StringLenBetween(1, 64),
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Type of service (vps, hicpu, etc.)",
				Default:     "vps",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Current status of the service",
			},
			"product_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the product",
			},
			"location_code": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Location code of the service",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Creation timestamp",
			},
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},
	}
}

func serviceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*provider.Config)
	apiClient, err := config.Client()
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to create client: %w", err))
	}

	createRequest := &client.CreateServiceRequest{
		ProductID: d.Get("product_id").(int),
		Name:      d.Get("name").(string),
		Type:      d.Get("type").(string),
	}

	service, err := apiClient.CreateService(ctx, createRequest)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to create service: %w", err))
	}

	d.SetId(service.Data.ID)

	return serviceRead(ctx, d, meta)
}

func serviceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	// Find the service by ID
	var foundService *client.Service
	for _, service := range services.Data.Items {
		if service.ID == d.Id() {
			foundService = &service
			break
		}
	}

	if foundService == nil {
		d.SetId("")
		return diags
	}

	// Set attributes
	if err := d.Set("name", foundService.Name); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set name: %w", err))
	}
	if err := d.Set("type", foundService.Type); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set type: %w", err))
	}
	if err := d.Set("status", foundService.Status); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set status: %w", err))
	}
	if err := d.Set("product_name", foundService.ProductName); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set product_name: %w", err))
	}
	if err := d.Set("location_code", foundService.LocationCode); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set location_code: %w", err))
	}
	if err := d.Set("created_at", foundService.CreatedAt.Format(time.RFC3339)); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set created_at: %w", err))
	}

	return diags
}

func serviceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	config := meta.(*provider.Config)
	apiClient, err := config.Client()
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to create client: %w", err))
	}

	err = apiClient.DeleteService(ctx, d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to delete service: %w", err))
	}

	d.SetId("")
	return diags
}
