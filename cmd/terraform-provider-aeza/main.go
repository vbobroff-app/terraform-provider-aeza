package main

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/vbobroff-app/terraform-provider-aeza/internal/provider"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			log.Printf("❌❌❌ PANIC in provider: %v", r)
		}
	}()

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.New,
	})
}
