package codeclimate

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/tema-ai/terraform-provider-codeclimate/codeclimateclient"
)

func dataSourceOrganization() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceOrganizationRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceOrganizationRead(d *schema.ResourceData, client interface{}) error {
	organizationName := d.Get("name").(string)

	c := client.(*codeclimateclient.Client)
	organization, err := c.GetOrganization(organizationName)
	if err != nil {
		return err
	}

	d.SetId(organization.Id)
	err = d.Set("name", organization.Name)

	return err
}
