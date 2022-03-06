package aquasec

import (
	"context"
	"log"

	"github.com/aquasecurity/terraform-provider-aquasec/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceUsers() *schema.Resource {
	return &schema.Resource{
		Description: "The data source `aquasec_users` provides a method to query all users within the Aqua " +
			"users database. The fields returned from this query are detailed in the Schema section below.",
		ReadContext: resourceRead,
		Schema: map[string]*schema.Schema{
			"users": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"user_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"email": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"first_time": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_super": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"ui_access": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"role": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"roles": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"plan": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Println("[DEBUG]  inside dataUser")
	c := m.(*client.Client)
	result, err := c.GetUsers()
	if err == nil {
		users, id := flattenUsersData(&result)
		d.SetId(id)
		if err := d.Set("users", users); err != nil {
			return diag.FromErr(err)
		}
	} else {
		return diag.FromErr(err)
	}

	return nil
}
