package wavefront

import (
	"testing"

	"github.com/WavefrontHQ/go-wavefront-management-api/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccWavefrontCloudIntegrationEc2_Basic(t *testing.T) {
	var record wavefront.CloudIntegration
	resourcePrefix := "wavefront_cloud_integration_ec2.ec2"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			err := testAccCheckWavefrontCloudIntegrationDestroy(state)
			if err != nil {
				return err
			}
			err = testAccCheckWavefrontCloudIntegrationAwsExternalIDDestroy(state)
			if err != nil {
				return err
			}
			return nil
		},
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWavefrontCloudIntegrationEc2Basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontCloudIntegrationExists(resourcePrefix, &record),
					testAccCheckWavefrontCloudIntegrationAttributes(&record, wfEc2),
					// Check against state that the attributes are as we expect
					testAccCheckWavefrontCloudIntegrationResourceAttributes(resourcePrefix, wfEc2),
					resource.TestCheckResourceAttr(resourcePrefix, "hostname_tags.#", "2"),
					testAccCheckWavefrontCloudIntegrationVerifyExtID(
						resourcePrefix, "wavefront_cloud_integration_aws_external_id.ext_id"),
					resource.TestCheckResourceAttr(resourcePrefix, "role_arn", "arn:aws::1234567:role/example-arn"),
				),
			},
		},
	})
}

func TestAccWavefrontCloudIntegrationEc2_BasicChanged(t *testing.T) {
	var record wavefront.CloudIntegration
	resourcePrefix := "wavefront_cloud_integration_ec2.ec2"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			err := testAccCheckWavefrontCloudIntegrationDestroy(state)
			if err != nil {
				return err
			}
			err = testAccCheckWavefrontCloudIntegrationAwsExternalIDDestroy(state)
			if err != nil {
				return err
			}
			return nil
		},
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWavefrontCloudIntegrationEc2Basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontCloudIntegrationExists(resourcePrefix, &record),
					// Check against state that the attributes are as we expect
					testAccCheckWavefrontCloudIntegrationResourceAttributes(resourcePrefix, wfEc2),
					testAccCheckWavefrontCloudIntegrationAttributes(&record, wfEc2),
					resource.TestCheckResourceAttr(resourcePrefix, "hostname_tags.#", "2"),
					testAccCheckWavefrontCloudIntegrationVerifyExtID(
						resourcePrefix, "wavefront_cloud_integration_aws_external_id.ext_id"),
					resource.TestCheckResourceAttr(resourcePrefix, "role_arn", "arn:aws::1234567:role/example-arn"),
				),
			},
			{
				Config: testAccCheckWavefrontCloudIntegrationEc2BasicChanged(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontCloudIntegrationExists(resourcePrefix, &record),
					testAccCheckWavefrontCloudIntegrationAttributes(&record, wfEc2),
					// Check against state that the attributes are as we expect
					testAccCheckWavefrontCloudIntegrationResourceAttributes(resourcePrefix, wfEc2),
					resource.TestCheckResourceAttr(resourcePrefix, "hostname_tags.#", "3"),
					testAccCheckWavefrontCloudIntegrationVerifyExtID(
						resourcePrefix, "wavefront_cloud_integration_aws_external_id.ext_id"),
					resource.TestCheckResourceAttr(
						resourcePrefix, "role_arn", "arn:aws::1234567:role/example-arn"),
				),
			},
		},
	})
}

func testAccCheckWavefrontCloudIntegrationEc2Basic() string {
	return `
resource "wavefront_cloud_integration_aws_external_id" "ext_id" { 
}

resource "wavefront_cloud_integration_ec2" "ec2" {
  name       = "Test Integration"
  force_save = true
  additional_tags = {
    "tag1" = "value1"
    "tag2" = "value2"
  }
  role_arn    = "arn:aws::1234567:role/example-arn"
  external_id = wavefront_cloud_integration_aws_external_id.ext_id.id
  hostname_tags  = ["host", "source"]
}`
}

func testAccCheckWavefrontCloudIntegrationEc2BasicChanged() string {
	return `
resource "wavefront_cloud_integration_aws_external_id" "ext_id" {
}

resource "wavefront_cloud_integration_ec2" "ec2" {
  name       = "Test Integration"
  force_save = true
  additional_tags = {
    "tag1" = "value1"
    "tag2" = "value2"
  }
  role_arn    = "arn:aws::1234567:role/example-arn"
  external_id = wavefront_cloud_integration_aws_external_id.ext_id.id
  hostname_tags  = ["host", "source", "name"]
}`
}
