package wavefront_plugin

import (
	"fmt"
	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"testing"
)

func TestAccWavefrontCloudIntegrationGcpBilling_Basic(t *testing.T) {
	var record wavefront.CloudIntegration
	resourcePrefix := "wavefront_cloud_integration_gcp_billing.gcp_billing"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontCloudIntegrationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWavefrontCloudIntegrationGcpBilling_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontCloudIntegrationExists(resourcePrefix, &record),
					testAccCheckWavefrontCloudIntegrationAttributes(&record, wfGcpBilling),
					// Check against state that the attributes are as we expect
					testAccCheckWavefrontCloudIntegrationResourceAttributes(resourcePrefix, wfGcpBilling),
					resource.TestCheckResourceAttr(resourcePrefix, "project_id", "example-gcp-project"),
					resource.TestCheckResourceAttr(resourcePrefix, "api_key", "example-api-key"),
					resource.TestCheckResourceAttr(resourcePrefix, "json_key", testGcpJsonKey("example-gcp-project")),
				),
			},
		},
	})
}

func TestAccWavefrontCloudIntegrationGcpBilling_BasicChanged(t *testing.T) {
	var record wavefront.CloudIntegration
	resourcePrefix := "wavefront_cloud_integration_gcp_billing.gcp_billing"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontCloudIntegrationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWavefrontCloudIntegrationGcpBilling_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontCloudIntegrationExists(resourcePrefix, &record),
					testAccCheckWavefrontCloudIntegrationAttributes(&record, wfGcpBilling),
					// Check against state that the attributes are as we expect
					testAccCheckWavefrontCloudIntegrationResourceAttributes(resourcePrefix, wfGcpBilling),
					resource.TestCheckResourceAttr(resourcePrefix, "project_id", "example-gcp-project"),
					resource.TestCheckResourceAttr(resourcePrefix, "api_key", "example-api-key"),
					resource.TestCheckResourceAttr(resourcePrefix, "json_key", testGcpJsonKey("example-gcp-project")),
				),
			},
			{
				Config: testAccCheckWavefrontCloudIntegrationGcpBilling_basicChanged(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontCloudIntegrationExists(resourcePrefix, &record),
					testAccCheckWavefrontCloudIntegrationAttributes(&record, wfGcpBilling),
					// Check against state that the attributes are as we expect
					testAccCheckWavefrontCloudIntegrationResourceAttributes(resourcePrefix, wfGcpBilling),
					resource.TestCheckResourceAttr(resourcePrefix, "project_id", "example-gcp-project"),
					resource.TestCheckResourceAttr(resourcePrefix, "api_key", "example-api-key"),
					resource.TestCheckResourceAttr(resourcePrefix, "json_key", testGcpJsonKey("example-gcp-project2")),
				),
			},
		},
	})
}

func testAccCheckWavefrontCloudIntegrationGcpBilling_basic() string {
	return fmt.Sprintf(`
resource "wavefront_cloud_integration_gcp_billing" "gcp_billing" {
  name                = "Test Integration"
  force_save          = true
  additional_tags = {
    "tag1" = "value1"
    "tag2" = "value2"
  }
  project_id          = "example-gcp-project"
  api_key             = "example-api-key"
  json_key            = <<EOF
%s
EOF
}`, testGcpJsonKey("example-gcp-project"))
}

func testAccCheckWavefrontCloudIntegrationGcpBilling_basicChanged() string {
	return fmt.Sprintf(`
resource "wavefront_cloud_integration_gcp_billing" "gcp_billing" {
  name                = "Test Integration"
  force_save          = true
  additional_tags = {
    "tag1" = "value1"
    "tag2" = "value2"
  }
  project_id          = "example-gcp-project"
  api_key             = "example-api-key"
  json_key            = <<EOF
%s
EOF
}`, testGcpJsonKey("example-gcp-project2"))
}
