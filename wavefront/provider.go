package wavefront

import (
	"fmt"
	"sync"

	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
			"wavefront_cloud_integration_tesla":              resourceCloudIntegrationTesla(),
			"wavefront_dashboard":                            resourceDashboard(),
			"wavefront_dashboard_json":                       resourceDashboardJSON(),
			"wavefront_derived_metric":                       resourceDerivedMetric(),
			"wavefront_external_link":                        resourceExternalLink(),
			"wavefront_ingestion_policy":                     resourceIngestionPolicy(),
			"wavefront_maintenance_window":                   resourceMaintenanceWindow(),
			"wavefront_service_account":                      resourceServiceAccount(),
			"wavefront_role":                                 resourceRole(),
			"wavefront_user":                                 resourceUser(),
			"wavefront_user_group":                           resourceUserGroup(),
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

// MutexKV is a simple key/value store for arbitrary mutexes. It can be used to
// serialize changes across arbitrary collaborators that share knowledge of the
// keys they must serialize on.
//
// The initial use case is to let aws_security_group_rule resources serialize
// their access to individual security groups based on SG ID.
type MutexKV struct {
	lock  sync.Mutex
	store map[string]*sync.Mutex
}

// Locks the mutex for the given key. Caller is responsible for calling Unlock
// for the same key
func (m *MutexKV) Lock(key string) {
	m.get(key).Lock()
}

// Unlock the mutex for the given key. Caller must have called Lock for the same key first
func (m *MutexKV) Unlock(key string) {
	m.get(key).Unlock()
}

// Returns a mutex for the given key, no guarantee of its lock status
func (m *MutexKV) get(key string) *sync.Mutex {
	m.lock.Lock()
	defer m.lock.Unlock()
	mutex, ok := m.store[key]
	if !ok {
		mutex = &sync.Mutex{}
		m.store[key] = mutex
	}
	return mutex
}

// Returns a properly initalized MutexKV
func NewMutexKV() *MutexKV {
	return &MutexKV{
		store: make(map[string]*sync.Mutex),
	}
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
	wfTesla            string = "TESLA"
)
