package codeclimate

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/tema-ai/terraform-provider-codeclimate/codeclimateclient"
)

func resourceRepository() *schema.Resource {
	return &schema.Resource{
		Read:   resourceRepositoryRead,
		Create: resourceRepositoryCreateForOrganization,
		Delete: resourceRepositoryDelete,

		Schema: map[string]*schema.Schema{
			"codeclimate_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"test_reporter_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"repository_url": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"organization_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"branch": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"human_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"services": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"self": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"web_coverage": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"web_issues": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"maintainability_badge": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"test_coverage_badge": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceRepositoryRead(d *schema.ResourceData, client interface{}) error {
	repositoryId := d.Id()

	c := client.(*codeclimateclient.Client)
	repository, err := c.GetRepositoryById(repositoryId)
	if err != nil {
		return err
	}

	d.SetId(repositoryId)
	err = d.Set("test_reporter_id", repository.TestReporterId)
	if err != nil {
		return err
	}
	err = d.Set("codeclimate_id", repository.Id)
	if err != nil {
		return err
	}
	err = d.Set("repository_url", repository.RepositoryURL)
	if err != nil {
		return err
	}
	err = d.Set("organization_id", repository.Organization)
	if err != nil {
		return err
	}
	err = d.Set("branch", repository.Branch)
	if err != nil {
		return err
	}
	err = d.Set("human_name", repository.HumanName)
	if err != nil {
		return err
	}
	err = d.Set("services", repository.LinkServices)
	if err != nil {
		return err
	}
	err = d.Set("self", repository.LinkSelf)
	if err != nil {
		return err
	}
	err = d.Set("web_coverage", repository.LinkWebCoverage)
	if err != nil {
		return err
	}
	err = d.Set("web_issues", repository.LinkWebIssues)
	if err != nil {
		return err
	}
	err = d.Set("maintainability_badge", repository.LinkMaintainabilityBadge)
	if err != nil {
		return err
	}
	err = d.Set("test_coverage_badge", repository.LinkTestCoverageBadge)
	return err
}

func resourceRepositoryCreateForOrganization(d *schema.ResourceData, client interface{}) error {
	repositoryUrl := d.Get("repository_url").(string)
	organizationId := d.Get("organization_id").(string)

	c := client.(*codeclimateclient.Client)

	repository, err := c.CreateOrganizationRepository(organizationId, repositoryUrl)
	if err != nil {
		return err
	}

	d.SetId(repository.Id)
	err = d.Set("test_reporter_id", repository.TestReporterId)
	if err != nil {
		return err
	}
	err = d.Set("codeclimate_id", repository.Id)

	return err
}

func resourceRepositoryDelete(d *schema.ResourceData, client interface{}) error {
	repositoryID := d.Get("codeclimate_id").(string)

	c := client.(*codeclimateclient.Client)
	err := c.DeleteOrganizationRepository(repositoryID)
	if err != nil {
		return err
	}

	d.SetId("")

	return err
}
