package example

import (
	"context"

	"github.com/CyrusJavan/terraform-provider-example/client"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{},
		ResourcesMap: map[string]*schema.Resource{
			"example_pet": resourceExamplePet(),
		},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: configure,
	}
}

func configure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	httpClient := client.NewHTTPClient(strfmt.Default)
	return httpClient, nil
}
