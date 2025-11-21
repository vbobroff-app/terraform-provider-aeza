// internal/data-sources/service_groups.go
package data_sources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vbobroff-app/terraform-provider-aeza/internal/interfaces"
)

func ServiceGroupsData() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving Aeza service groups",

		ReadContext: serviceGroupsRead,

		Schema: map[string]*schema.Schema{
			"service_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter groups by service type (e.g., 'vps', 'dedicated', 'hicpu')",
			},
			"groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Group ID",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Group name",
						},
						"type_slug": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Service type slug (e.g., 'vps', 'dedicated')",
						},
						"type_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Service type display name",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Group description",
						},
						"service_handler": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Service handler type",
						},
						"payload": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Group payload data",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func serviceGroupsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(interfaces.DataClient)

	serviceType := d.Get("service_type").(string)

	groups, err := client.GetServiceGroups(ctx, serviceType)
	if err != nil {
		return diag.FromErr(err)
	}

	groupList := make([]map[string]interface{}, len(groups.Items))
	for i, group := range groups.Items {
		description := ""
		if group.Description != nil {
			description = *group.Description
		}

		// Конвертируем payload в map[string]string для Terraform
		payloadMap := make(map[string]string)
		for key, value := range group.Payload {
			switch v := value.(type) {
			case string:
				payloadMap[key] = v
			case bool:
				payloadMap[key] = fmt.Sprintf("%t", v)
			case float64:
				payloadMap[key] = fmt.Sprintf("%.0f", v)
			case int:
				payloadMap[key] = fmt.Sprintf("%d", v)
			default:
				// Пропускаем сложные типы
				continue
			}
		}

		groupList[i] = map[string]interface{}{
			"id":              group.ID,
			"name":            group.Name,
			"type_slug":       group.Type.Slug,
			"type_name":       group.Type.Name,
			"description":     description,
			"service_handler": group.Type.ServiceHandler,
			"payload":         payloadMap,
		}
	}

	d.SetId("service_groups")
	if err := d.Set("groups", groupList); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
