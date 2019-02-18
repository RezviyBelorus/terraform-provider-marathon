package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/rolyv/terraform-provider-marathon/marathon"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: marathon.Provider,
	})
}
