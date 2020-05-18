package wavefront_plugin

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceCloudIntegrationCloudTrail() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudIntegrationCreate,
		Read:   resourceCloudIntegrationRead,
		Update: resourceCloudIntegrationUpdate,
		Delete: resourceCloudIntegrationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"additional_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"force_save": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"service": serviceSchemaDefinition(wfCloudTrail),
			"service_refresh_rate_in_minutes": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"region": {
				Type:     schema.TypeString,
				Required: true,
			},
			"prefix": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"bucket_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"filter_rule": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"role_arn": {
				Type:     schema.TypeString,
				Required: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
