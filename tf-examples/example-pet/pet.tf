provider "example" {}

terraform {
  required_providers {
    example = {
      source = "example.com/example/example"
    }
  }
}

resource "example_pet" "pet_1" {
  name = "sharik"
}
