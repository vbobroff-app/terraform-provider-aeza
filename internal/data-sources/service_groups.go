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
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Service type slug (e.g., 'vps', 'dedicated')",
						},
						"location": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Location label (e.g., 'NL-SHARED', 'US-DEDICATED')",
						},
						"country_code": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Country code (e.g., 'nl', 'de', 'fr')",
						},
						"server_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Server type (e.g., 'shared', 'dedicated')",
						},
						"is_disabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the group is disabled",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Group description",
						},
						"features": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Localized features description",
						},
						"service_handler": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Service handler type",
						},
						"cpu_model": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CPU model",
						},
						"cpu_frequency": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CPU frequency",
						},
						"network_speed": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Network speed",
						},
						"ipv4_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of IPv4 addresses",
						},
						"ipv6_subnet": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv6 subnet",
						},
					},
				},
			},
		},
	}
}

func serviceGroupsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(interfaces.DataClient)

	serviceType := ""
	if serviceTypeVal, ok := d.Get("service_type").(string); ok {
		serviceType = serviceTypeVal
	}

	groups, err := client.ListServiceGroups(ctx, serviceType)
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to read service groups: %w", err))
	}

	groupList := make([]map[string]interface{}, len(groups))
	for i, group := range groups {
		groupList[i] = map[string]interface{}{
			"id":              group.ID,
			"name":            group.Name,
			"type":            group.Type,
			"location":        group.Location,
			"country_code":    group.CountryCode,
			"server_type":     group.ServerType,
			"is_disabled":     group.IsDisabled,
			"description":     group.Description,
			"features":        group.Features,
			"service_handler": group.ServiceHandler,
			"cpu_model":       group.CPUModel,
			"cpu_frequency":   group.CPUFrequency,
			"network_speed":   group.NetworkSpeed,
			"ipv4_count":      group.IPv4Count,
			"ipv6_subnet":     group.IPv6Subnet,
		}
	}

	d.SetId("service_groups")
	if err := d.Set("groups", groupList); err != nil {
		return diag.FromErr(fmt.Errorf("unable to set groups in state: %w", err))
	}

	return nil
}
