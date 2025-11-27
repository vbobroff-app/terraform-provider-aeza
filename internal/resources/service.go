package resources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vbobroff-app/terraform-provider-aeza/internal/interfaces"
	"github.com/vbobroff-app/terraform-provider-aeza/internal/models"
)

func ServiceResource() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for managing Aeza services",

		CreateContext: serviceCreate,
		ReadContext:   serviceRead,
		UpdateContext: serviceUpdate,
		DeleteContext: serviceDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"product_id": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"location_code": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"payment_term": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "month",
			},
			"auto_prolong": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"os": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"recipe": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"iso_url": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			// ✅ Вычисляемые поля из ServiceCreateResponse
			"order_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"product_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"group_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"product_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"location_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"price": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"transaction_amount": {
				Type:     schema.TypeString,
				Computed: true,
			},
			// ✅ Поля из ServiceGetResponse
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip": {
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

	req := models.ServiceCreateRequest{
		Name:         d.Get("name").(string),
		ProductID:    int64(d.Get("product_id").(int)),
		LocationCode: d.Get("location_code").(string),
		PaymentTerm:  d.Get("payment_term").(string),
		AutoProlong:  d.Get("auto_prolong").(bool),
		OS:           d.Get("os").(string),
		Recipe:       d.Get("recipe").(string),
		IsoURL:       d.Get("iso_url").(string),
	}

	resp, err := client.CreateService(ctx, req)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating service: %w", err))
	}

	if resp.ID == 0 {
		return diag.FromErr(fmt.Errorf("API returned ID=0, creation failed"))
	}

	// ✅ Устанавливаем ID созданной услуги
	d.SetId(fmt.Sprintf("%d", resp.ID))

	// ✅ Сохраняем все вычисляемые поля из ответа создания
	if err := setServiceCreateResponse(d, resp); err != nil {
		return diag.FromErr(fmt.Errorf("error setting service response data: %w", err))
	}

	return nil
}

func serviceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(interfaces.ResourceClient)

	id, err := parseIntID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	// ✅ Получаем детальную информацию об услуге
	resp, err := client.GetService(ctx, id)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error reading service: %w", err))
	}

	service := resp.Service

	// ✅ Устанавливаем поля из ServiceGetResponse
	if err := setServiceGetResponse(d, service); err != nil {
		return diag.FromErr(fmt.Errorf("error setting service data: %w", err))
	}

	return nil
}

func serviceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(interfaces.ResourceClient)

	id, err := parseIntID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	if d.HasChanges("name", "payment_term", "auto_prolong") {
		req := models.ServiceCreateRequest{
			Name:        d.Get("name").(string),
			PaymentTerm: d.Get("payment_term").(string),
			AutoProlong: d.Get("auto_prolong").(bool),
			// Не передаем ForceNew поля
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

// ✅ Вспомогательные функции для установки данных

func setServiceCreateResponse(d *schema.ResourceData, resp *models.ServiceCreateResponse) error {
	if err := d.Set("order_id", resp.OrderID); err != nil {
		return fmt.Errorf("error setting order_id: %w", err)
	}
	if err := d.Set("date", resp.Date); err != nil {
		return fmt.Errorf("error setting date: %w", err)
	}
	if err := d.Set("product_type", resp.ProductType); err != nil {
		return fmt.Errorf("error setting product_type: %w", err)
	}
	if resp.GroupId != nil {
		if err := d.Set("group_id", *resp.GroupId); err != nil {
			return fmt.Errorf("error setting group_id: %w", err)
		}
	}
	if err := d.Set("product_name", resp.ProductName); err != nil {
		return fmt.Errorf("error setting product_name: %w", err)
	}
	if err := d.Set("location_name", resp.LocationName); err != nil {
		return fmt.Errorf("error setting location_name: %w", err)
	}
	if err := d.Set("price", resp.Price); err != nil {
		return fmt.Errorf("error setting price: %w", err)
	}
	if err := d.Set("transaction_amount", resp.TransactionAmount); err != nil {
		return fmt.Errorf("error setting transaction_amount: %w", err)
	}
	if err := d.Set("status", resp.Status); err != nil {
		return fmt.Errorf("error setting status: %w", err)
	}

	return nil
}

func setServiceGetResponse(d *schema.ResourceData, service models.Service) error {
	if err := d.Set("name", service.Name); err != nil {
		return fmt.Errorf("error setting name: %w", err)
	}
	if err := d.Set("product_id", service.ProductID); err != nil {
		return fmt.Errorf("error setting product_id: %w", err)
	}
	if err := d.Set("payment_term", service.PaymentTerm); err != nil {
		return fmt.Errorf("error setting payment_term: %w", err)
	}
	if err := d.Set("auto_prolong", service.AutoProlong); err != nil {
		return fmt.Errorf("error setting auto_prolong: %w", err)
	}
	if err := d.Set("status", service.Status); err != nil {
		return fmt.Errorf("error setting status: %w", err)
	}
	if err := d.Set("ip", service.IP); err != nil {
		return fmt.Errorf("error setting ip: %w", err)
	}
	if err := d.Set("created_at", service.CreatedAt); err != nil {
		return fmt.Errorf("error setting created_at: %w", err)
	}
	if err := d.Set("updated_at", service.UpdatedAt); err != nil {
		return fmt.Errorf("error setting updated_at: %w", err)
	}

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
