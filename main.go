package main

import (
	"github.com/CyrusJavan/terraform-provider-example/example"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)


//go:generate terraform fmt -recursive ./tf-examples/

//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: example.Provider,
	})
}
