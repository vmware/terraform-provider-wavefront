package wavefront_plugin

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func resourceCloudIntegrationNewRelic() *schema.Resource {
	newrelicMetricFilters := &schema.Schema{
		Type:       schema.TypeList,
		ConfigMode: schema.SchemaConfigModeAttr,
		Optional:   true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"app_name": {
					Type:     schema.TypeString,
					Required: true,
				},
				"metric_filter_regex": {
					Type:     schema.TypeString,
					Required: true,
				},
			},
		},
	}
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
			"service": serviceSchemaDefinition(wfNewRelic),
			"service_refresh_rate_in_minutes": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"api_key": {
				Sensitive: true,
				Type:      schema.TypeString,
				Required:  true,
			},
			"app_filter_regex": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"host_filter_regex": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"metric_filter": newrelicMetricFilters,
		},
	}
}
