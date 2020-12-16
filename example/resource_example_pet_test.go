package example

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccExamplePet(t *testing.T) {
	rName := fmt.Sprintf("%s", acctest.RandString(5))
	resourceName := "example_pet.test"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccExamplePetConfig(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccExamplePetConfig(rName string) string {
	return fmt.Sprintf(`
resource example_pet test {
  name = "%s"
}
	`, rName)
}
