package wavefront

import (
	"testing"

	"github.com/WavefrontHQ/go-wavefront-management-api/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccWavefrontCloudIntegrationAzureActivityLog_Basic(t *testing.T) {
	var record wavefront.CloudIntegration

	resourcePrefix := "wavefront_cloud_integration_azure_activity_log.azure_activity_log"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontCloudIntegrationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWavefrontCloudIntegrationAzureActivityLogBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontCloudIntegrationExists(resourcePrefix, &record),
					testAccCheckWavefrontCloudIntegrationAttributes(&record, wfAzureActivityLog),

					// Check against state that the attributes are as we expect
					testAccCheckWavefrontCloudIntegrationResourceAttributes(resourcePrefix, wfAzureActivityLog),
					resource.TestCheckResourceAttr(resourcePrefix, "client_id", "client-id"),
					resource.TestCheckResourceAttr(resourcePrefix, "client_secret", "client-secret"),
					resource.TestCheckResourceAttr(resourcePrefix, "tenant", "my-tenant"),
					resource.TestCheckResourceAttr(resourcePrefix, "additional_tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourcePrefix, "additional_tags.tag2", "value2"),
					resource.TestCheckResourceAttr(resourcePrefix, "category_filter.#", "1"),
				),
			},
		},
	})
}

func TestAccWavefrontCloudIntegrationAzureActivityLog_BasicChange(t *testing.T) {
	var record wavefront.CloudIntegration

	resourcePrefix := "wavefront_cloud_integration_azure_activity_log.azure_activity_log"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontCloudIntegrationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWavefrontCloudIntegrationAzureActivityLogBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontCloudIntegrationExists(resourcePrefix, &record),
					testAccCheckWavefrontCloudIntegrationAttributes(&record, wfAzureActivityLog),

					// Check against state that the attributes are as we expect
					testAccCheckWavefrontCloudIntegrationResourceAttributes(resourcePrefix, wfAzureActivityLog),
					resource.TestCheckResourceAttr(resourcePrefix, "client_id", "client-id"),
					resource.TestCheckResourceAttr(resourcePrefix, "client_secret", "client-secret"),
					resource.TestCheckResourceAttr(resourcePrefix, "tenant", "my-tenant"),
					resource.TestCheckResourceAttr(resourcePrefix, "additional_tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourcePrefix, "additional_tags.tag2", "value2"),
					resource.TestCheckResourceAttr(resourcePrefix, "category_filter.#", "1"),
				),
			},
			{
				Config: testAccCheckWavefrontCloudIntegrationAzureActivityLogBasicChangeAdd(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontCloudIntegrationExists(resourcePrefix, &record),
					testAccCheckWavefrontCloudIntegrationAttributes(&record, wfAzureActivityLog),

					// Check against state that the attributes are as we expect
					testAccCheckWavefrontCloudIntegrationResourceAttributes(resourcePrefix, wfAzureActivityLog),
					resource.TestCheckResourceAttr(resourcePrefix, "client_id", "client-id2"),
					resource.TestCheckResourceAttr(resourcePrefix, "client_secret", "client-secret2"),
					resource.TestCheckResourceAttr(resourcePrefix, "tenant", "my-tenant2"),
					resource.TestCheckResourceAttr(resourcePrefix, "additional_tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourcePrefix, "additional_tags.tag2", "value2"),
					resource.TestCheckResourceAttr(resourcePrefix, "additional_tags.tag3", "value3"),
					resource.TestCheckResourceAttr(resourcePrefix, "category_filter.#", "2"),
				),
			},
		},
	})
}

func testAccCheckWavefrontCloudIntegrationAzureActivityLogBasic() string {
	return `
resource "wavefront_cloud_integration_azure_activity_log" "azure_activity_log" {
	name 				= "Test Integration"
	force_save 			= true
	additional_tags     = {
		"tag1" = "value1"
		"tag2" = "value2"
    }
	category_filter  = ["ADMINISTRATIVE"]
	client_id 			= "client-id"
	client_secret		= "client-secret"
	tenant				= "my-tenant"
}
`
}

func testAccCheckWavefrontCloudIntegrationAzureActivityLogBasicChangeAdd() string {
	return `
resource "wavefront_cloud_integration_azure_activity_log" "azure_activity_log" {
	name 			= "Test Integration"
	force_save 		= true
	additional_tags = {
		"tag1" = "value1"
		"tag2" = "value2"
		"tag3" = "value3"
    }
	category_filter  = ["ADMINISTRATIVE", "SERVICEHEALTH"]
	client_id 		 = "client-id2"
	client_secret	 = "client-secret2"
	tenant			 = "my-tenant2"
}
`
}
