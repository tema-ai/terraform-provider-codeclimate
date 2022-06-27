package codeclimate

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/tema-ai/terraform-provider-codeclimate/codeclimateclient"
)

const codeClimateApiHost string = "https://api.codeclimate.com/v1"
const codeClimateDefaultToken string = "CODECLIMATE_TOKEN"

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc(codeClimateDefaultToken, nil),
				Description: "Token for the CodeClimate API.",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"codeclimate_repository":   dataSourceRepository(),
			"codeclimate_organization": dataSourceOrganization(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"codeclimate_repository": resourceRepository(),
		},

		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	client := codeclimateclient.Client{
		ApiKey:  d.Get("api_key").(string),
		BaseUrl: codeClimateApiHost,
	}

	return &client, nil
}
