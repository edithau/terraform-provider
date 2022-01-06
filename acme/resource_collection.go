package acme

import (
	"context"
	"strconv"

	//hc "github.com/hashicorp-demoapp/hashicups-client-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCollection() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCollectionCreate,
		//ReadContext:   resourceCollectionRead,
		//UpdateContext: resourceCollectionUpdate,
		//DeleteContext: resourceCollectionDelete,
		Schema: map[string]*schema.Schema{
			"collection": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceCollectionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)


	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	collectionId := c.CreateCollection()
	d.SetId(strconv.Itoa(collectionId))

	return diags
}

