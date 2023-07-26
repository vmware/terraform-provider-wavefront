package wavefront

import (
	"fmt"
	"testing"

	"github.com/WavefrontHQ/go-wavefront-management-api/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
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
				Config: testAccCheckWavefrontCloudIntegrationGcpBillingBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontCloudIntegrationExists(resourcePrefix, &record),
					testAccCheckWavefrontCloudIntegrationAttributes(&record, wfGcpBilling),
					// Check against state that the attributes are as we expect
					testAccCheckWavefrontCloudIntegrationResourceAttributes(resourcePrefix, wfGcpBilling),
					resource.TestCheckResourceAttr(resourcePrefix, "project_id", "example-gcp-project"),
					resource.TestCheckResourceAttr(resourcePrefix, "api_key", "example-api-key"),
					resource.TestCheckResourceAttr(resourcePrefix, "json_key", testGcpJSONKey("example-gcp-project")),
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
				Config: testAccCheckWavefrontCloudIntegrationGcpBillingBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontCloudIntegrationExists(resourcePrefix, &record),
					testAccCheckWavefrontCloudIntegrationAttributes(&record, wfGcpBilling),
					// Check against state that the attributes are as we expect
					testAccCheckWavefrontCloudIntegrationResourceAttributes(resourcePrefix, wfGcpBilling),
					resource.TestCheckResourceAttr(resourcePrefix, "project_id", "example-gcp-project"),
					resource.TestCheckResourceAttr(resourcePrefix, "api_key", "example-api-key"),
					resource.TestCheckResourceAttr(resourcePrefix, "json_key", testGcpJSONKey("example-gcp-project")),
				),
			},
			{
				Config: testAccCheckWavefrontCloudIntegrationGcpBillingBasicChanged(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontCloudIntegrationExists(resourcePrefix, &record),
					testAccCheckWavefrontCloudIntegrationAttributes(&record, wfGcpBilling),
					// Check against state that the attributes are as we expect
					testAccCheckWavefrontCloudIntegrationResourceAttributes(resourcePrefix, wfGcpBilling),
					resource.TestCheckResourceAttr(resourcePrefix, "project_id", "example-gcp-project"),
					resource.TestCheckResourceAttr(resourcePrefix, "api_key", "example-api-key"),
					resource.TestCheckResourceAttr(resourcePrefix, "json_key", testGcpJSONKey("example-gcp-project2")),
				),
			},
		},
	})
}

func testAccCheckWavefrontCloudIntegrationGcpBillingBasic() string {
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
}`, testGcpJSONKey("example-gcp-project"))
}

func testAccCheckWavefrontCloudIntegrationGcpBillingBasicChanged() string {
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
}`, testGcpJSONKey("example-gcp-project2"))
}
