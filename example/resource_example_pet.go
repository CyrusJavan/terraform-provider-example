package example

import (
	"context"
	"log"
	"strconv"

	"github.com/CyrusJavan/terraform-provider-example/client"
	"github.com/CyrusJavan/terraform-provider-example/client/pet"
	"github.com/CyrusJavan/terraform-provider-example/models"
	"github.com/go-openapi/runtime"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceExamplePet() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Pet name.",
			},
			"nicknames": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Description: "Pet nicknames.",
			},
			"favorite_foods": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Description: "Pet's favorite foods, in ranked order.",
			},
		},
		CreateContext: resourceExamplePetCreate,
		ReadContext:   resourceExamplePetRead,
		UpdateContext: resourceExamplePetUpdate,
		DeleteContext: resourceExamplePetDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Description: "Pet represents a single pet resource.",
	}
}

func resourceExamplePetCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Pet).Pet
	// Could we generate the marshal funcs?
	petData := marshalExamplePet(d)

	resp, err := c.AddPet(&pet.AddPetParams{
		Body:    petData,
		Context: ctx,
	})

	if err != nil {
		return diag.Errorf("could not add pet: %v", err)
	}
	d.SetId(strconv.FormatInt(resp.Payload.ID, 10))

	return nil
}

func resourceExamplePetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Pet).Pet
	petData := marshalExamplePet(d)

	resp, err := c.GetPetByID(&pet.GetPetByIDParams{
		PetID:   petData.ID,
		Context: ctx,
	})

	if apiErr, ok := err.(*runtime.APIError); ok && apiErr.Code == 404 {
		d.SetId("")
		return nil
	}

	if err != nil {
		return diag.Errorf("could not get pet: %v", err)
	}

	d.Set("name", *resp.Payload.Name)
	if resp.Payload.Nicknames == nil {
		resp.Payload.Nicknames = []string{}
	}
	err = d.Set("nicknames", resp.Payload.Nicknames)
	if err != nil {
		return diag.Errorf("could not set 'nicknames': %v", err)
	}
	if resp.Payload.FavoriteFoods == nil {
		resp.Payload.FavoriteFoods = []string{}
	}
	err = d.Set("favorite_foods", resp.Payload.FavoriteFoods)
	if err != nil {
		return diag.Errorf("could not set 'favorite_foods': %v", err)
	}
	d.SetId(strconv.FormatInt(resp.Payload.ID, 10))

	return nil
}

func resourceExamplePetUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Pet).Pet
	petData := marshalExamplePet(d)

	_, err := c.AddPet(&pet.AddPetParams{
		Body:    petData,
		Context: ctx,
	})

	if err != nil {
		return diag.Errorf("could not update pet: %v", err)
	}

	return nil
}

func resourceExamplePetDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Pet).Pet
	petData := marshalExamplePet(d)

	_, err := c.DeletePet(&pet.DeletePetParams{
		PetID:   petData.ID,
		Context: ctx,
	})

	if err != nil {
		return diag.Errorf("could not delete pet: %v", err)
	}

	return nil
}

func marshalExamplePet(d *schema.ResourceData) *models.Pet {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		log.Printf("received invalid ID: %v\n", id)
	}
	name := d.Get("name").(string)
	var nicknames []string
	for _, v := range d.Get("nicknames").(*schema.Set).List() {
		nicknames = append(nicknames, v.(string))
	}
	var favoriteFoods []string
	for _, v := range d.Get("favorite_foods").([]interface{}) {
		favoriteFoods = append(favoriteFoods, v.(string))
	}
	return &models.Pet{
		ID:            int64(id),
		Name:          &name,
		Nicknames:     nicknames,
		FavoriteFoods: favoriteFoods,
	}
}
