package resources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vbobroff-app/terraform-provider-aeza/internal/interfaces"
	"github.com/vbobroff-app/terraform-provider-aeza/internal/source-models"
)

func ServiceResource() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for managing Aeza services",

		CreateContext: serviceCreate,
		ReadContext:   serviceRead,
		UpdateContext: serviceUpdate,
		DeleteContext: serviceDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"location_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"product_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"status": {
				Type:     schema.TypeString,
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
	}
}

func serviceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(interfaces.ResourceClient)

	req := source_models.ServiceCreateRequest{
		Name:       d.Get("name").(string),
		Type:       d.Get("type").(string),
		LocationID: int64(d.Get("location_id").(int)),
		ProductID:  int64(d.Get("product_id").(int)),
	}

	resp, err := client.CreateService(ctx, req)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating service: %w", err))
	}

	d.SetId(fmt.Sprintf("%d", resp.ID))
	return serviceRead(ctx, d, meta)
}

func serviceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(interfaces.ResourceClient)

	id, err := parseIntID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	resp, err := client.GetService(ctx, id)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error reading service: %w", err))
	}

	service := resp.Service
	d.Set("name", service.Name)
	d.Set("type", service.Type)
	d.Set("location_id", service.LocationID)
	d.Set("product_id", service.ProductID)
	d.Set("status", service.Status)
	d.Set("created_at", service.CreatedAt)
	d.Set("updated_at", service.UpdatedAt)

	return nil
}

func serviceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(interfaces.ResourceClient)

	id, err := parseIntID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	if d.HasChanges("name", "type", "location_id", "product_id") {
		req := source_models.ServiceCreateRequest{
			Name:       d.Get("name").(string),
			Type:       d.Get("type").(string),
			LocationID: int64(d.Get("location_id").(int)),
			ProductID:  int64(d.Get("product_id").(int)),
		}

		err := client.UpdateService(ctx, id, req)
		if err != nil {
			return diag.FromErr(fmt.Errorf("error updating service: %w", err))
		}
	}

	return serviceRead(ctx, d, meta)
}

func serviceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(interfaces.ResourceClient)

	id, err := parseIntID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	err = client.DeleteService(ctx, id)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting service: %w", err))
	}

	d.SetId("")
	return nil
}

func parseIntID(id string) (int64, error) {
	var intID int64
	_, err := fmt.Sscanf(id, "%d", &intID)
	if err != nil {
		return 0, fmt.Errorf("invalid ID format: %s", id)
	}
	return intID, nil
}
