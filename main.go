package main

import (
	"github.com/CyrusJavan/terraform-provider-example/example"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: example.Provider,
	})
}
