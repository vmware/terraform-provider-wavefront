package wavefront_plugin

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceCloudIntegrationAppDynamics() *schema.Resource {
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
			"service": serviceSchemaDefinition(wfAppDynamics),
			"service_refresh_rate_in_minutes": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"user_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"controller_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"encrypted_password": {
				Type:      schema.TypeString,
				Sensitive: true,
				Required:  true,
			},
			"enable_rollup": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"enable_error_metrics": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"enable_business_trx_metrics": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"enable_backend_metrics": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"enable_overall_perf_metrics": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"enable_individual_node_metrics": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"enable_app_infra_metrics": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"enable_service_endpoint_metrics": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"app_filter_regex": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}
