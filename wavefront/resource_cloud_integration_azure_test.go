package wavefront_plugin

import (
	"fmt"
	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"testing"
)

func TestAccWavefrontCloudIntegrationAzure_Basic(t *testing.T) {
	var record wavefront.CloudIntegration

	resourcePrefix := "wavefront_cloud_integration_azure.azure"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontCloudIntegrationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWavefrontCloudIntegrationAzure_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontCloudIntegrationExists(resourcePrefix, &record),
					testAccCheckWavefrontCloudIntegrationAttributes(&record, wfAzure),

					// Check against state that the attributes are as we expect
					testAccCheckWavefrontCloudIntegrationResourceAttributes(resourcePrefix, wfAzure),
					resource.TestCheckResourceAttr(resourcePrefix, "client_id", "client-id"),
					resource.TestCheckResourceAttr(resourcePrefix, "client_secret", "client-secret"),
					resource.TestCheckResourceAttr(resourcePrefix, "tenant", "my-tenant"),
					resource.TestCheckResourceAttr(resourcePrefix, "additional_tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourcePrefix, "additional_tags.tag2", "value2"),
					resource.TestCheckResourceAttr(resourcePrefix, "metric_filter_regex", "^.*?$"),
				),
			},
		},
	})
}

func TestAccWavefrontCloudIntegrationAzure_BasicChange(t *testing.T) {
	var record wavefront.CloudIntegration

	resourcePrefix := "wavefront_cloud_integration_azure.azure"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontCloudIntegrationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWavefrontCloudIntegrationAzure_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontCloudIntegrationExists(resourcePrefix, &record),
					testAccCheckWavefrontCloudIntegrationAttributes(&record, wfAzure),

					// Check against state that the attributes are as we expect
					testAccCheckWavefrontCloudIntegrationResourceAttributes(resourcePrefix, wfAzure),
					resource.TestCheckResourceAttr(resourcePrefix, "client_id", "client-id"),
					resource.TestCheckResourceAttr(resourcePrefix, "client_secret", "client-secret"),
					resource.TestCheckResourceAttr(resourcePrefix, "tenant", "my-tenant"),
					resource.TestCheckResourceAttr(resourcePrefix, "metric_filter_regex", "^.*?$"),
					resource.TestCheckResourceAttr(resourcePrefix, "additional_tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourcePrefix, "additional_tags.tag2", "value2"),
					resource.TestCheckResourceAttr(resourcePrefix, "resource_group_filter.#", "1"),
				),
			},
			{
				Config: testAccCheckWavefrontCloudIntegrationAzure_basicChangeAdd(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontCloudIntegrationExists(resourcePrefix, &record),
					testAccCheckWavefrontCloudIntegrationAttributes(&record, wfAzure),

					// Check against state that the attributes are as we expect
					testAccCheckWavefrontCloudIntegrationResourceAttributes(resourcePrefix, wfAzure),
					resource.TestCheckResourceAttr(resourcePrefix, "client_id", "client-id2"),
					resource.TestCheckResourceAttr(resourcePrefix, "client_secret", "client-secret2"),
					resource.TestCheckResourceAttr(resourcePrefix, "tenant", "my-tenant2"),
					resource.TestCheckResourceAttr(resourcePrefix, "additional_tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourcePrefix, "additional_tags.tag2", "value2"),
					resource.TestCheckResourceAttr(resourcePrefix, "additional_tags.tag3", "value3"),
					resource.TestCheckResourceAttr(resourcePrefix, "metric_filter_regex", "^[A-Z]+"),
					resource.TestCheckResourceAttr(resourcePrefix, "resource_group_filter.#", "2"),
				),
			},
		},
	})
}

func testAccCheckWavefrontCloudIntegrationAzure_basic() string {
	return fmt.Sprintf(`
resource "wavefront_cloud_integration_azure" "azure" {
	name 				= "Test Integration"
	force_save 			= true
	additional_tags     = {
		"tag1" = "value1"
		"tag2" = "value2"
    }
	resource_group_filter  = ["a"]
	metric_filter_regex = "^.*?$"
	client_id 			= "client-id"
	client_secret		= "client-secret"
	tenant				= "my-tenant"
}
`)
}

func testAccCheckWavefrontCloudIntegrationAzure_basicChangeAdd() string {
	return fmt.Sprintf(`
resource "wavefront_cloud_integration_azure" "azure" {
	name 			= "Test Integration"
	force_save 		= true
	additional_tags = {
		"tag1" = "value1"
		"tag2" = "value2"
		"tag3" = "value3"
    }
	metric_filter_regex    = "^[A-Z]+"
	resource_group_filter  = ["a", "b"]
	client_id 			   = "client-id2"
	client_secret		   = "client-secret2"
	tenant				   = "my-tenant2"
}
`)
}
