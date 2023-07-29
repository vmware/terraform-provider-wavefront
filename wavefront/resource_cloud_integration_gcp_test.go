package wavefront

import (
	"fmt"
	"strings"
	"testing"

	"github.com/WavefrontHQ/go-wavefront-management-api/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccWavefrontCloudIntegrationGcp_Basic(t *testing.T) {
	var record wavefront.CloudIntegration
	resourcePrefix := "wavefront_cloud_integration_gcp.gcp"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontCloudIntegrationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWavefrontCloudIntegrationGcpBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontCloudIntegrationExists(resourcePrefix, &record),
					testAccCheckWavefrontCloudIntegrationAttributes(&record, wfGcp),
					// Check against state that the attributes are as we expect
					testAccCheckWavefrontCloudIntegrationResourceAttributes(resourcePrefix, wfGcp),
					resource.TestCheckResourceAttr(resourcePrefix, "metric_filter_regex", "^(exampleMetricRegex).*?"),
					resource.TestCheckResourceAttr(resourcePrefix, "project_id", "example-gcp-project"),
					resource.TestCheckResourceAttr(resourcePrefix, "json_key", testGcpJSONKey("example-gcp-project")),
					resource.TestCheckResourceAttr(resourcePrefix, "categories.#", "1"),
				),
			},
		},
	})
}

func TestAccWavefrontCloudIntegrationGcp_BasicChanged(t *testing.T) {
	var record wavefront.CloudIntegration
	resourcePrefix := "wavefront_cloud_integration_gcp.gcp"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontCloudIntegrationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWavefrontCloudIntegrationGcpBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontCloudIntegrationExists(resourcePrefix, &record),
					testAccCheckWavefrontCloudIntegrationAttributes(&record, wfGcp),
					// Check against state that the attributes are as we expect
					testAccCheckWavefrontCloudIntegrationResourceAttributes(resourcePrefix, wfGcp),
					resource.TestCheckResourceAttr(resourcePrefix, "metric_filter_regex", "^(exampleMetricRegex).*?"),
					resource.TestCheckResourceAttr(resourcePrefix, "project_id", "example-gcp-project"),
					resource.TestCheckResourceAttr(resourcePrefix, "json_key", testGcpJSONKey("example-gcp-project")),
					resource.TestCheckResourceAttr(resourcePrefix, "categories.#", "1"),
				),
			},
			{
				Config: testAccCheckWavefrontCloudIntegrationGcpBasicChanged(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontCloudIntegrationExists(resourcePrefix, &record),
					testAccCheckWavefrontCloudIntegrationAttributes(&record, wfGcp),
					// Check against state that the attributes are as we expect
					testAccCheckWavefrontCloudIntegrationResourceAttributes(resourcePrefix, wfGcp),
					resource.TestCheckResourceAttr(resourcePrefix, "metric_filter_regex", "^(exampleMetricRegex).*?"),
					resource.TestCheckResourceAttr(resourcePrefix, "project_id", "example-gcp-project"),
					resource.TestCheckResourceAttr(resourcePrefix, "json_key", testGcpJSONKey("example-gcp-project")),
					resource.TestCheckResourceAttr(resourcePrefix, "categories.#", "2"),
				),
			},
		},
	})
}

func testAccCheckWavefrontCloudIntegrationGcpBasic() string {
	return fmt.Sprintf(`
resource "wavefront_cloud_integration_gcp" "gcp" {
  name                = "Test Integration"
  force_save          = true
  additional_tags = {
    "tag1" = "value1"
    "tag2" = "value2"
  }
  metric_filter_regex = "^(exampleMetricRegex).*?"
  project_id          = "example-gcp-project"
  json_key            = <<EOF
%s
EOF
  categories          = ["APPENGINE"]
}`, testGcpJSONKey("example-gcp-project"))
}

func testAccCheckWavefrontCloudIntegrationGcpBasicChanged() string {
	return fmt.Sprintf(`
resource "wavefront_cloud_integration_gcp" "gcp" {
  name                = "Test Integration"
  force_save          = true
  additional_tags = {
    "tag1" = "value1"
    "tag2" = "value2"
  }
  metric_filter_regex = "^(exampleMetricRegex).*?"
  project_id          = "example-gcp-project"
  json_key            = <<EOF
%s
EOF
  categories          = ["APPENGINE", "BIGQUERY"]
}`, testGcpJSONKey("example-gcp-project"))
}

func testGcpJSONKey(pid string) string {
	return strings.TrimSpace(fmt.Sprintf(`{
  "project-id": "%s"
}`, pid))
}
