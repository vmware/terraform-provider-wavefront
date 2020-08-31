package wavefront

import (
	"testing"

	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccCloudIntegration_AppDynamics(t *testing.T) {
	var record wavefront.CloudIntegration
	resourceName := "wavefront_cloud_integration_app_dynamics.app_dynamics"
	testCloudIntegrationCommon(t, resourceName, record, testAccCheckWavefrontCloudIntegrationAppDynamicsBasic())
}

func TestAccCloudIntegration_Azure(t *testing.T) {
	var record wavefront.CloudIntegration
	resourceName := "wavefront_cloud_integration_azure.azure"
	testCloudIntegrationCommon(t, resourceName, record, testAccCheckWavefrontCloudIntegrationAzureBasic())
}

func TestAccCloudIntegration_AzureActivityLog(t *testing.T) {
	var record wavefront.CloudIntegration
	resourceName := "wavefront_cloud_integration_azure_activity_log.azure_activity_log"
	testCloudIntegrationCommon(t, resourceName, record, testAccCheckWavefrontCloudIntegrationAzureActivityLogBasic())
}

func TestAccCloudIntegration_CloudTrail(t *testing.T) {
	var record wavefront.CloudIntegration
	resourceName := "wavefront_cloud_integration_cloudtrail.cloudtrail"
	testCloudIntegrationCommon(t, resourceName, record, testAccCheckWavefrontCloudIntegrationCloudTrailBasic())
}

func TestAccCloudIntegration_CloudWatch(t *testing.T) {
	var record wavefront.CloudIntegration
	resourceName := "wavefront_cloud_integration_cloudwatch.cloudwatch"
	testCloudIntegrationCommon(t, resourceName, record, testAccCheckWavefrontCloudIntegrationCloudWatchBasic())
}

func TestAccCloudIntegration_Ec2(t *testing.T) {
	var record wavefront.CloudIntegration
	resourceName := "wavefront_cloud_integration_ec2.ec2"
	testCloudIntegrationCommon(t, resourceName, record, testAccCheckWavefrontCloudIntegrationEc2Basic())
}

func TestAccCloudIntegration_Gcp(t *testing.T) {
	var record wavefront.CloudIntegration
	resourceName := "wavefront_cloud_integration_gcp.gcp"
	testCloudIntegrationCommon(t, resourceName, record, testAccCheckWavefrontCloudIntegrationGcpBasic())
}

func TestAccCloudIntegration_GcpBilling(t *testing.T) {
	var record wavefront.CloudIntegration
	resourceName := "wavefront_cloud_integration_gcp_billing.gcp_billing"
	testCloudIntegrationCommon(t, resourceName, record, testAccCheckWavefrontCloudIntegrationGcpBillingBasic())
}

func TestAccCloudIntegration_NewRelic(t *testing.T) {
	var record wavefront.CloudIntegration
	resourceName := "wavefront_cloud_integration_newrelic.newrelic"
	testCloudIntegrationCommon(t, resourceName, record, testAccCheckWavefrontCloudIntegrationNewRelicBasic())
}

func TestAccCloudIntegration_Tesla(t *testing.T) {
	var record wavefront.CloudIntegration
	resourceName := "wavefront_cloud_integration_tesla.tesla"
	testCloudIntegrationCommon(t, resourceName, record, testAccCheckWavefrontCloudIntegrationTeslaBasic())
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
