package wavefront_plugin

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func resourceCloudIntegrationGcpBilling() *schema.Resource {
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
			"service": serviceSchemaDefinition(wfGcpBilling),
			"service_refresh_rate_in_minutes": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"api_key": {
				Sensitive: true,
				Type:      schema.TypeString,
				Required:  true,
			},
			"json_key": {
				Sensitive: true,
				Type:      schema.TypeString,
				Required:  true,
				StateFunc: trimSpaces,
			},
		},
	}
}
