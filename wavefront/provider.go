package wavefront_plugin

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/mutexkv"

	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

type wavefrontClient struct {
	client wavefront.Client
}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"address": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("WAVEFRONT_ADDRESS", ""),
			},
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("WAVEFRONT_TOKEN", ""),
			},
			"http_proxy": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"wavefront_alert":                                resourceAlert(),
			"wavefront_dashboard":                            resourceDashboard(),
			"wavefront_dashboard_json":                       resourceDashboardJson(),
			"wavefront_derived_metric":                       resourceDerivedMetric(),
			"wavefront_alert_target":                         resourceTarget(),
			"wavefront_user":                                 resourceUser(),
			"wavefront_role":                                 resourceRole(),
			"wavefront_user_group":                           resourceUserGroup(),
			"wavefront_cloud_integration_cloudwatch":         resourceCloudIntegrationCloudWatch(),
			"wavefront_cloud_integration_cloudtrail":         resourceCloudIntegrationCloudTrail(),
			"wavefront_cloud_integration_ec2":                resourceCloudIntegrationEc2(),
			"wavefront_cloud_integration_gcp":                resourceCloudIntegrationGcp(),
			"wavefront_cloud_integration_gcp_billing":        resourceCloudIntegrationGcpBilling(),
			"wavefront_cloud_integration_newrelic":           resourceCloudIntegrationNewRelic(),
			"wavefront_cloud_integration_app_dynamics":       resourceCloudIntegrationAppDynamics(),
			"wavefront_cloud_integration_tesla":              resourceCloudIntegrationTesla(),
			"wavefront_cloud_integration_azure":              resourceCloudIntegrationAzure(),
			"wavefront_cloud_integration_azure_activity_log": resourceCloudIntegrationAzureActivityLog(),
			"wavefront_cloud_integration_aws_external_id":    resourceCloudIntegrationAwsExternalId(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"wavefront_default_user_group": dataSourceDefaultUserGroup(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := &wavefront.Config{
		Address:   d.Get("address").(string),
		Token:     d.Get("token").(string),
		HttpProxy: d.Get("http_proxy").(string),
	}
	wFClient, err := wavefront.NewClient(config)
	if err != nil {
		return nil, fmt.Errorf("failed to configure Wavefront Client %s", err)
	}
	return &wavefrontClient{
		client: *wFClient,
	}, nil

}

var wfMutexKV = mutexkv.NewMutexKV()

const (
	wfAppDynamics      string = "APPDYNAMICS"
	wfAzure            string = "AZURE"
	wfAzureActivityLog string = "AZUREACTIVITYLOG"
	wfCloudTrail       string = "CLOUDTRAIL"
	wfCloudWatch       string = "CLOUDWATCH"
	wfEc2              string = "EC2"
	wfGcp              string = "GCP"
	wfGcpBilling       string = "GCPBILLING"
	wfNewRelic         string = "NEWRELIC"
	wfTesla            string = "TESLA"
)
