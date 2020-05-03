package wavefront_plugin

import (
	"fmt"
	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"testing"
)

func TestAccWavefrontCloudIntegrationAppDynamics_Basic(t *testing.T) {
	var record wavefront.CloudIntegration

	resourcePrefix := "wavefront_cloud_integration_app_dynamics.app_dynamics"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontCloudIntegrationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWavefrontCloudIntegrationAppDynamics_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontCloudIntegrationExists(resourcePrefix, &record),
					testAccCheckWavefrontCloudIntegrationAttributes(&record, wfAppDynamics),

					// Check against state that the attributes are as we expect
					testAccCheckWavefrontCloudIntegrationResourceAttributes(resourcePrefix, wfAppDynamics),
					resource.TestCheckResourceAttr(resourcePrefix, "controller_name", "exampleController"),
					resource.TestCheckResourceAttr(resourcePrefix, "user_name", "example"),
				),
			},
		},
	})
}

func TestAccWavefrontCloudIntegrationAppDynamics_BasicChange(t *testing.T) {
	var record wavefront.CloudIntegration

	resourcePrefix := "wavefront_cloud_integration_app_dynamics.app_dynamics"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontCloudIntegrationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWavefrontCloudIntegrationAppDynamics_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontCloudIntegrationExists(resourcePrefix, &record),
					testAccCheckWavefrontCloudIntegrationAttributes(&record, wfAppDynamics),

					// Check against state that the attributes are as we expect
					testAccCheckWavefrontCloudIntegrationResourceAttributes(resourcePrefix, wfAppDynamics),
					resource.TestCheckResourceAttr(resourcePrefix, "controller_name", "exampleController"),
					resource.TestCheckResourceAttr(resourcePrefix, "user_name", "example"),
					resource.TestCheckResourceAttr(resourcePrefix, "enable_rollup", "true"),
					resource.TestCheckResourceAttr(resourcePrefix, "enable_error_metrics", "false"),
					resource.TestCheckResourceAttr(resourcePrefix, "enable_business_trx_metrics", "false"),
					resource.TestCheckResourceAttr(resourcePrefix, "enable_backend_metrics", "false"),
					resource.TestCheckResourceAttr(resourcePrefix, "enable_overall_perf_metrics", "false"),
					resource.TestCheckResourceAttr(resourcePrefix, "enable_individual_node_metrics", "false"),
					resource.TestCheckResourceAttr(resourcePrefix, "enable_app_infra_metrics", "false"),
					resource.TestCheckResourceAttr(resourcePrefix, "enable_service_endpoint_metrics", "false"),
				),
			},
			{
				Config: testAccCheckWavefrontCloudIntegrationAppDynamics_basicChanges(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontCloudIntegrationExists(resourcePrefix, &record),
					testAccCheckWavefrontCloudIntegrationAttributes(&record, wfAppDynamics),

					// Check against state that the attributes are as we expect
					testAccCheckWavefrontCloudIntegrationResourceAttributes(resourcePrefix, wfAppDynamics),
					resource.TestCheckResourceAttr(resourcePrefix, "controller_name", "exampleController2"),
					resource.TestCheckResourceAttr(resourcePrefix, "user_name", "example2"),
					resource.TestCheckResourceAttr(resourcePrefix, "enable_rollup", "false"),
					resource.TestCheckResourceAttr(resourcePrefix, "enable_error_metrics", "true"),
					resource.TestCheckResourceAttr(resourcePrefix, "enable_business_trx_metrics", "true"),
					resource.TestCheckResourceAttr(resourcePrefix, "enable_backend_metrics", "true"),
					resource.TestCheckResourceAttr(resourcePrefix, "enable_overall_perf_metrics", "true"),
					resource.TestCheckResourceAttr(resourcePrefix, "enable_individual_node_metrics", "true"),
					resource.TestCheckResourceAttr(resourcePrefix, "enable_app_infra_metrics", "true"),
					resource.TestCheckResourceAttr(resourcePrefix, "enable_service_endpoint_metrics", "true"),
				),
			},
		},
	})
}

func testAccCheckWavefrontCloudIntegrationAppDynamics_basic() string {
	return fmt.Sprintf(`
resource "wavefront_cloud_integration_app_dynamics" "app_dynamics" {
	name 				= "Test Integration"
	force_save 			= true
	additional_tags     = {
		"tag1" = "value1"
		"tag2" = "value2"
    }
	user_name 			= "example"
	controller_name 	= "exampleController"
	encrypted_password 	= "encryptedPassword"	
}
`)
}

func testAccCheckWavefrontCloudIntegrationAppDynamics_basicChanges() string {
	return fmt.Sprintf(`
resource "wavefront_cloud_integration_app_dynamics" "app_dynamics" {
  name       = "Test Integration"
  force_save = true
  additional_tags = {
    "tag1" = "value1"
    "tag2" = "value2"
  }
  user_name                       = "example2"
  controller_name                 = "exampleController2"
  encrypted_password              = "encryptedPassword"
  enable_rollup                   = false
  enable_error_metrics            = true
  enable_business_trx_metrics     = true
  enable_backend_metrics          = true
  enable_overall_perf_metrics     = true
  enable_individual_node_metrics  = true
  enable_app_infra_metrics        = true
  enable_service_endpoint_metrics = true
}
`)
}
