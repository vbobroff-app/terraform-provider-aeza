package resources

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vbobroff-app/terraform-provider-aeza/internal/interfaces"
	"github.com/vbobroff-app/terraform-provider-aeza/internal/models"
)

func ServiceActionsResource() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for managing service actions (prolong, reinstall, reboot)",

		CreateContext: actionsCreate,
		ReadContext:   actionsRead,
		UpdateContext: actionsUpdate,
		DeleteContext: actionsDelete,

		Schema: map[string]*schema.Schema{
			"service_id": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "ID of the service to manage",
			},
			"prolong": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Prolong service configuration",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"method": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "balance",
							Description: "Payment method (balance, sms, etc.)",
						},
						"term": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Payment term (hour, day, month)",
						},
						"count": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     1,
							Description: "Number of terms to prolong",
						},
						// Вычисляемые поля транзакции prolong
						"transaction_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "ID of the transaction created by the prolong operation",
						},
						"transaction_amount": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Amount of the transaction",
						},
						"transaction_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Status of the transaction (created, pending, completed)",
						},
						"transaction_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of the transaction (prolong, etc.)",
						},
						"transaction_created_at": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Timestamp when the transaction was created",
						},
						"force": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Force prolongation on every apply",
						},
					},
				},
			},
		},
	}
}

func actionsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(interfaces.ResourceClient)
	serviceID := int64(d.Get("service_id").(int))

	// Handle prolong operation
	if v, ok := d.GetOk("prolong"); ok {
		prolongConfig := v.([]interface{})[0].(map[string]interface{})

		req := models.ServiceProlongRequest{
			Method: prolongConfig["method"].(string),
			Term:   prolongConfig["term"].(string),
			Count:  int64(prolongConfig["count"].(int)),
		}

		resp, err := client.ProlongService(ctx, serviceID, req)
		if err != nil {
			return diag.FromErr(fmt.Errorf("error prolonging service: %w", err))
		}

		if resp != nil && resp.Transaction != nil {
			updatedProlongConfig := map[string]interface{}{
				"method":                 req.Method,
				"term":                   req.Term,
				"count":                  req.Count,
				"transaction_id":         resp.Transaction.ID,
				"transaction_amount":     resp.Transaction.Amount,
				"transaction_status":     resp.Transaction.Status,
				"transaction_type":       resp.Transaction.Type,
				"transaction_created_at": resp.Transaction.CreatedAt,
			}

			if err := d.Set("prolong", []interface{}{updatedProlongConfig}); err != nil {
				return diag.FromErr(fmt.Errorf("error updating prolong configuration: %w", err))
			}

			log.Printf("[INFO] Successfully prolonged service %d for %d %s", serviceID, req.Count, req.Term)
		}
	}

	d.SetId(fmt.Sprintf("%d", serviceID))
	return actionsRead(ctx, d, meta)
}

func actionsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// For service actions, we don't have much to read back from API
	// Just verify the service exists
	client := meta.(interfaces.ResourceClient)

	serviceID, err := parseIntID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	// Verify service exists
	_, err = client.GetService(ctx, serviceID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error reading service: %w", err))
	}

	return nil
}

func actionsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// For update, we can handle prolong operation changes
	if d.HasChange("prolong") {
		return actionsCreate(ctx, d, meta)
	}
	return nil
}

func actionsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// For service actions, delete doesn't actually delete anything
	// It just removes the resource from state
	d.SetId("")
	return nil
}
