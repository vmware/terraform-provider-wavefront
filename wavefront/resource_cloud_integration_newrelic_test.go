package wavefront_plugin

import (
	"fmt"
	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"testing"
)

func TestAccWavefrontCloudIntegrationNewRelic_Basic(t *testing.T) {
	var record wavefront.CloudIntegration
	resourcePrefix := "wavefront_cloud_integration_newrelic.newrelic"
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontCloudIntegrationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWavefrontCloudIntegrationNewRelic_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontCloudIntegrationExists(resourcePrefix, &record),
					testAccCheckWavefrontCloudIntegrationAttributes(&record, wfNewRelic),
					// Check the attributes...
					testAccCheckWavefrontCloudIntegrationResourceAttributes(resourcePrefix, wfNewRelic),
					resource.TestCheckResourceAttr(resourcePrefix, "api_key", "example-api-key"),
					resource.TestCheckResourceAttr(resourcePrefix, "app_filter_regex", "^someApp.*$"),
					resource.TestCheckResourceAttr(resourcePrefix, "host_filter_regex", "^prod-env.*$"),
					resource.TestCheckResourceAttr(resourcePrefix, "metric_filter.#", "2"),
					resource.TestCheckResourceAttr(resourcePrefix, "metric_filter.0.app_name", "app1"),
					resource.TestCheckResourceAttr(resourcePrefix, "metric_filter.0.metric_filter_regex", "^cpu.*?"),
					resource.TestCheckResourceAttr(resourcePrefix, "metric_filter.1.app_name", "app2"),
					resource.TestCheckResourceAttr(resourcePrefix, "metric_filter.1.metric_filter_regex", "^mem.*?"),
				),
			},
		},
	})
}

func TestAccWavefrontCloudIntegrationNewRelic_BasicChanged(t *testing.T) {
	var record wavefront.CloudIntegration
	resourcePrefix := "wavefront_cloud_integration_newrelic.newrelic"
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontCloudIntegrationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWavefrontCloudIntegrationNewRelic_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontCloudIntegrationExists(resourcePrefix, &record),
					testAccCheckWavefrontCloudIntegrationAttributes(&record, wfNewRelic),
					// Check the attributes...
					testAccCheckWavefrontCloudIntegrationResourceAttributes(resourcePrefix, wfNewRelic),
					resource.TestCheckResourceAttr(resourcePrefix, "api_key", "example-api-key"),
					resource.TestCheckResourceAttr(resourcePrefix, "app_filter_regex", "^someApp.*$"),
					resource.TestCheckResourceAttr(resourcePrefix, "host_filter_regex", "^prod-env.*$"),
					resource.TestCheckResourceAttr(resourcePrefix, "metric_filter.#", "2"),
					resource.TestCheckResourceAttr(resourcePrefix, "metric_filter.0.app_name", "app1"),
					resource.TestCheckResourceAttr(resourcePrefix, "metric_filter.0.metric_filter_regex", "^cpu.*?"),
					resource.TestCheckResourceAttr(resourcePrefix, "metric_filter.1.app_name", "app2"),
					resource.TestCheckResourceAttr(resourcePrefix, "metric_filter.1.metric_filter_regex", "^mem.*?"),
				),
			},
			{
				Config: testAccCheckWavefrontCloudIntegrationNewRelic_basicChanged(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontCloudIntegrationExists(resourcePrefix, &record),
					testAccCheckWavefrontCloudIntegrationAttributes(&record, wfNewRelic),
					// Check the attributes...
					testAccCheckWavefrontCloudIntegrationResourceAttributes(resourcePrefix, wfNewRelic),
					resource.TestCheckResourceAttr(resourcePrefix, "api_key", "example-api-key"),
					resource.TestCheckResourceAttr(resourcePrefix, "app_filter_regex", "^dev.*$"),
					resource.TestCheckResourceAttr(resourcePrefix, "host_filter_regex", "^dev-env.*$"),
					resource.TestCheckResourceAttr(resourcePrefix, "metric_filter.#", "1"),
					resource.TestCheckResourceAttr(resourcePrefix, "metric_filter.0.app_name", "dev1"),
					resource.TestCheckResourceAttr(resourcePrefix, "metric_filter.0.metric_filter_regex", "^mem.*?"),
				),
			},
		},
	})
}

func testAccCheckWavefrontCloudIntegrationNewRelic_basic() string {
	return fmt.Sprintf(`
resource "wavefront_cloud_integration_newrelic" "newrelic" {
  name              = "Test Integration"
  force_save        = true
  additional_tags = {
    "tag1" = "value1"
    "tag2" = "value2"
  }
  api_key           = "example-api-key"
  app_filter_regex  = "^someApp.*$"
  host_filter_regex = "^prod-env.*$"
  metric_filter {
    app_name            = "app1"
    metric_filter_regex = "^cpu.*?"
  }
  metric_filter {
    app_name            = "app2"
    metric_filter_regex = "^mem.*?"
  }
}
`)
}

func testAccCheckWavefrontCloudIntegrationNewRelic_basicChanged() string {
	return fmt.Sprintf(`
resource "wavefront_cloud_integration_newrelic" "newrelic" {
  name              = "Test Integration"
  force_save        = true
  additional_tags = {
    "tag1" = "value1"
    "tag2" = "value2"
  }
  api_key           = "example-api-key"
  app_filter_regex  = "^dev.*$"
  host_filter_regex = "^dev-env.*$"
  metric_filter {
    app_name            = "dev1"
    metric_filter_regex = "^mem.*?"
  }
}
`)
}
