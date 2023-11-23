package wavefront

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

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
			"csp_address": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CSP_ADDRESS", ""),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"wavefront_alert":                                resourceAlert(),
			"wavefront_alert_target":                         resourceTarget(),
			"wavefront_cloud_integration_app_dynamics":       resourceCloudIntegrationAppDynamics(),
			"wavefront_cloud_integration_aws_external_id":    resourceCloudIntegrationAwsExternalID(),
			"wavefront_cloud_integration_azure":              resourceCloudIntegrationAzure(),
			"wavefront_cloud_integration_azure_activity_log": resourceCloudIntegrationAzureActivityLog(),
			"wavefront_cloud_integration_cloudwatch":         resourceCloudIntegrationCloudWatch(),
			"wavefront_cloud_integration_cloudtrail":         resourceCloudIntegrationCloudTrail(),
			"wavefront_cloud_integration_ec2":                resourceCloudIntegrationEc2(),
			"wavefront_cloud_integration_gcp":                resourceCloudIntegrationGcp(),
			"wavefront_cloud_integration_gcp_billing":        resourceCloudIntegrationGcpBilling(),
			"wavefront_cloud_integration_newrelic":           resourceCloudIntegrationNewRelic(),
			"wavefront_dashboard":                            resourceDashboard(),
			"wavefront_dashboard_json":                       resourceDashboardJSON(),
			"wavefront_derived_metric":                       resourceDerivedMetric(),
			"wavefront_external_link":                        resourceExternalLink(),
			"wavefront_ingestion_policy":                     resourceIngestionPolicy(),
			"wavefront_maintenance_window":                   resourceMaintenanceWindow(),
			"wavefront_metrics_policy":                       resourceMetricsPolicy(),
			"wavefront_service_account":                      resourceServiceAccount(),
			"wavefront_role":                                 resourceRole(),
			"wavefront_user":                                 resourceUser(),
			"wavefront_user_group":                           resourceUserGroup(),
			"wavefront_event":                                resourceEvent(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"wavefront_default_user_group":     dataSourceDefaultUserGroup(),
			"wavefront_metrics_policy":         dataSourceMetricsPolicy(),
			"wavefront_role":                   dataSourceRole(),
			"wavefront_roles":                  dataSourceRoles(),
			"wavefront_user":                   dataSourceUser(),
			"wavefront_user_group":             dataSourceUserGroup(),
			"wavefront_user_groups":            dataSourceUserGroups(),
			"wavefront_users":                  dataSourceUsers(),
			"wavefront_external_links":         dataSourceExternalLinks(),
			"wavefront_external_link":          dataSourceExternalLink(),
			"wavefront_maintenance_window_all": dataSourceMaintenanceWindows(),
			"wavefront_maintenance_window":     dataSourceMaintenanceWindow(),
			"wavefront_alerts":                 dataSourceAlerts(),
			"wavefront_alert":                  dataSourceAlert(),
			"wavefront_derived_metrics":        dataSourceDerivedMetrics(),
			"wavefront_derived_metric":         dataSourceDerivedMetric(),
			"wavefront_event":                  dataSourceEvent(),
			"wavefront_events":                 dataSourceEvents(),
			"wavefront_dashboard":              dataSourceDashboard(),
			"wavefront_dashboards":             dataSourceDashboards(),
		},
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	address := ""
	token := ""
	proxy := ""
	cspAddress := ""

	if v, ok := d.GetOk("address"); ok {
		address = v.(string)
	}

	if v, ok := d.GetOk("token"); ok {
		token = v.(string)
	}

	if v, ok := d.GetOk("proxy"); ok {
		proxy = v.(string)
	}

	if v, ok := d.GetOk("csp_address"); ok {
		cspAddress = v.(string)
	}

	return newWavefrontClient(address, token, proxy, cspAddress)
}

var wfMutexKV = NewMutexKV()

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
)
