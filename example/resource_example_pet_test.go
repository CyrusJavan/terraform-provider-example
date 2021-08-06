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
	favoriteFoods := []string{"chicken", "goat", "zebra mussel"}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccExamplePetConfig(rName, nicknames[0], nicknames[1], nicknames[2], favoriteFoods[0], favoriteFoods[1], favoriteFoods[2]),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "nicknames.0", nicknames[0]),
					resource.TestCheckResourceAttr(resourceName, "nicknames.1", nicknames[1]),
					resource.TestCheckResourceAttr(resourceName, "nicknames.2", nicknames[2]),
					resource.TestCheckResourceAttr(resourceName, "favorite_foods.0", favoriteFoods[0]),
					resource.TestCheckResourceAttr(resourceName, "favorite_foods.1", favoriteFoods[1]),
					resource.TestCheckResourceAttr(resourceName, "favorite_foods.2", favoriteFoods[2]),
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

func testAccExamplePetConfig(rName, nName0, nName1, nName2, fFood0, fFood1, fFood2 string) string {
	return fmt.Sprintf(`
resource example_pet test {
  name           = "%s"
  nicknames      = ["%s", "%s", "%s"]
  favorite_foods = ["%s", "%s", "%s"]
}
	`, rName, nName0, nName1, nName2, fFood0, fFood1, fFood2)
}
