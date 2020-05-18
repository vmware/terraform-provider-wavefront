package wavefront_plugin

import (
	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"testing"
)

func TestAccCloudIntegration_AppDynamics(t *testing.T) {
	var record wavefront.CloudIntegration
	resourceName := "wavefront_cloud_integration_app_dynamics.app_dynamics"
	testCloudIntegrationCommon(t, resourceName, record, testAccCheckWavefrontCloudIntegrationAppDynamics_basic())
}

func TestAccCloudIntegration_Azure(t *testing.T) {
	var record wavefront.CloudIntegration
	resourceName := "wavefront_cloud_integration_azure.azure"
	testCloudIntegrationCommon(t, resourceName, record, testAccCheckWavefrontCloudIntegrationAzure_basic())
}

func TestAccCloudIntegration_AzureActivityLog(t *testing.T) {
	var record wavefront.CloudIntegration
	resourceName := "wavefront_cloud_integration_azure_activity_log.azure_activity_log"
	testCloudIntegrationCommon(t, resourceName, record, testAccCheckWavefrontCloudIntegrationAzureActivityLog_basic())
}

func TestAccCloudIntegration_CloudTrail(t *testing.T) {
	var record wavefront.CloudIntegration
	resourceName := "wavefront_cloud_integration_cloudtrail.cloudtrail"
	testCloudIntegrationCommon(t, resourceName, record, testAccCheckWavefrontCloudIntegrationCloudTrail_basic())
}

func TestAccCloudIntegration_CloudWatch(t *testing.T) {
	var record wavefront.CloudIntegration
	resourceName := "wavefront_cloud_integration_cloudwatch.cloudwatch"
	testCloudIntegrationCommon(t, resourceName, record, testAccCheckWavefrontCloudIntegrationCloudWatch_basic())
}

func TestAccCloudIntegration_Ec2(t *testing.T) {
	var record wavefront.CloudIntegration
	resourceName := "wavefront_cloud_integration_ec2.ec2"
	testCloudIntegrationCommon(t, resourceName, record, testAccCheckWavefrontCloudIntegrationEc2_basic())
}

func TestAccCloudIntegration_Gcp(t *testing.T) {
	var record wavefront.CloudIntegration
	resourceName := "wavefront_cloud_integration_gcp.gcp"
	testCloudIntegrationCommon(t, resourceName, record, testAccCheckWavefrontCloudIntegrationGcp_basic())
}

func TestAccCloudIntegration_GcpBilling(t *testing.T) {
	var record wavefront.CloudIntegration
	resourceName := "wavefront_cloud_integration_gcp_billing.gcp_billing"
	testCloudIntegrationCommon(t, resourceName, record, testAccCheckWavefrontCloudIntegrationGcpBilling_basic())
}

func TestAccCloudIntegration_NewRelic(t *testing.T) {
	var record wavefront.CloudIntegration
	resourceName := "wavefront_cloud_integration_newrelic.newrelic"
	testCloudIntegrationCommon(t, resourceName, record, testAccCheckWavefrontCloudIntegrationNewRelic_basic())
}

func TestAccCloudIntegration_Tesla(t *testing.T) {
	var record wavefront.CloudIntegration
	resourceName := "wavefront_cloud_integration_tesla.tesla"
	testCloudIntegrationCommon(t, resourceName, record, testAccCheckWavefrontCloudIntegrationTesla_basic())
}

func testCloudIntegrationCommon(t *testing.T, resourceName string, record wavefront.CloudIntegration, config string) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontCloudIntegrationDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontCloudIntegrationExists(resourceName, &record),
				),
			},
			{
				ResourceName: resourceName,
				ImportState:  true,
				ImportStateVerifyIgnore: []string{
					"force_save", "encrypted_password", "client_secret", "json_key", "api_key", "password",
				},
				ImportStateVerify: true,
			},
		},
	})
}
