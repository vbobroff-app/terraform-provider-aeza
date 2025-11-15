package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/vbobroff-app/terraform-provider-aeza/internal/client"
)

type Config struct {
	APIKey  string
	BaseURL string
}

func (c *Config) Client() (*client.Client, diag.Diagnostics) {
	if c.APIKey == "" {
		return nil, diag.FromErr(fmt.Errorf("API Key cannot be empty"))
	}

	return client.NewClient(c.BaseURL, c.APIKey), nil
}
