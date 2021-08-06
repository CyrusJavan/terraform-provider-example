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
	nicknames := []string{"bubba", "buddy", "kiddo"}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccExamplePetConfig(rName, nicknames[0], nicknames[1], nicknames[2]),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "nicknames.0", nicknames[0]),
					resource.TestCheckResourceAttr(resourceName, "nicknames.1", nicknames[1]),
					resource.TestCheckResourceAttr(resourceName, "nicknames.2", nicknames[2]),
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

func testAccExamplePetConfig(rName, nName0, nName1, nName2 string) string {
	return fmt.Sprintf(`
resource example_pet test {
  name      = "%s"
  nicknames = ["%s", "%s", "%s"]
}
	`, rName, nName0, nName1, nName2)
}
