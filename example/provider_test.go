package example

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

var providerFactories = map[string]func() (*schema.Provider, error){
	"example": func() (*schema.Provider, error) {
		return Provider(), nil
	},
}
