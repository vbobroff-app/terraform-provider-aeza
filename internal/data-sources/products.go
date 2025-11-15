package data_sources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vbobroff-app/terraform-provider-aeza/internal/interfaces"
)

func ProductsDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving list of Aeza products",

		ReadContext: productsRead,

		Schema: map[string]*schema.Schema{
			"products": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"price": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"currency": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"category": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func productsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(interfaces.DataClient)

	products, err := client.ListProducts(ctx)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error fetching products: %w", err))
	}

	productList := make([]map[string]interface{}, len(products))
	for i, product := range products {
		productList[i] = map[string]interface{}{
			"id":          product.ID,
			"name":        product.Name,
			"description": product.Description,
			"price":       product.Price,
			"currency":    product.Currency,
			"category":    product.Category,
		}
	}

	if err := d.Set("products", productList); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("products")
	return nil
}
