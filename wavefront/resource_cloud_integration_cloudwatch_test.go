package wavefront

import (
	"testing"

	"github.com/WavefrontHQ/go-wavefront-management-api/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccWavefrontCloudIntegrationCloudWatch_Basic(t *testing.T) {
	var record wavefront.CloudIntegration
	resourcePrefix := "wavefront_cloud_integration_cloudwatch.cloudwatch"
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
				Config: testAccCheckWavefrontCloudIntegrationCloudWatchBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontCloudIntegrationExists(resourcePrefix, &record),
					testAccCheckWavefrontCloudIntegrationAttributes(&record, wfCloudWatch),
					// Check against state that the attributes are as we expect
					testAccCheckWavefrontCloudIntegrationResourceAttributes(resourcePrefix, wfCloudWatch),
					resource.TestCheckResourceAttr(resourcePrefix, "namespaces.#", "2"),
					resource.TestCheckResourceAttr(resourcePrefix, "instance_selection_tags.env", "prod"),
					resource.TestCheckResourceAttr(resourcePrefix, "volume_selection_tags.env", "prod"),
					resource.TestCheckResourceAttr(resourcePrefix, "point_tag_filter_regex", "^prod$"),
					resource.TestCheckResourceAttr(resourcePrefix, "metric_filter_regex", "^.*?\\.cpu.*$"),
					testAccCheckWavefrontCloudIntegrationVerifyExtID(
						resourcePrefix, "wavefront_cloud_integration_aws_external_id.ext_id"),
					resource.TestCheckResourceAttr(resourcePrefix, "role_arn", "arn:aws::1234567:role/example-arn"),
				),
			},
		},
	})
}

func TestAccWavefrontCloudIntegrationCloudWatch_BasicChanged(t *testing.T) {
	var record wavefront.CloudIntegration
	resourcePrefix := "wavefront_cloud_integration_cloudwatch.cloudwatch"
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
				Config: testAccCheckWavefrontCloudIntegrationCloudWatchBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontCloudIntegrationExists(resourcePrefix, &record),
					// Check against state that the attributes are as we expect
					testAccCheckWavefrontCloudIntegrationResourceAttributes(resourcePrefix, wfCloudWatch),
					testAccCheckWavefrontCloudIntegrationAttributes(&record, wfCloudWatch),
					resource.TestCheckResourceAttr(resourcePrefix, "namespaces.#", "2"),
					resource.TestCheckResourceAttr(resourcePrefix, "instance_selection_tags.env", "prod"),
					resource.TestCheckResourceAttr(resourcePrefix, "volume_selection_tags.env", "prod"),
					resource.TestCheckResourceAttr(resourcePrefix, "point_tag_filter_regex", "^prod$"),
					resource.TestCheckResourceAttr(resourcePrefix, "metric_filter_regex", "^.*?\\.cpu.*$"),
					testAccCheckWavefrontCloudIntegrationVerifyExtID(
						resourcePrefix, "wavefront_cloud_integration_aws_external_id.ext_id"),
					resource.TestCheckResourceAttr(resourcePrefix, "role_arn", "arn:aws::1234567:role/example-arn"),
				),
			},
			{
				Config: testAccCheckWavefrontCloudIntegrationCloudWatchBasicChanged(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontCloudIntegrationExists(resourcePrefix, &record),
					testAccCheckWavefrontCloudIntegrationAttributes(&record, wfCloudWatch),
					// Check against state that the attributes are as we expect
					testAccCheckWavefrontCloudIntegrationResourceAttributes(resourcePrefix, wfCloudWatch),
					resource.TestCheckResourceAttr(resourcePrefix, "namespaces.#", "3"),
					resource.TestCheckResourceAttr(resourcePrefix, "instance_selection_tags.env", "dev"),
					resource.TestCheckResourceAttr(resourcePrefix, "instance_selection_tags.mirror", "b"),
					resource.TestCheckResourceAttr(resourcePrefix, "volume_selection_tags.env", "dev"),
					resource.TestCheckResourceAttr(resourcePrefix, "point_tag_filter_regex", "^dev"),
					resource.TestCheckResourceAttr(resourcePrefix, "metric_filter_regex", "^.*?\\.cpu.*$"),
					testAccCheckWavefrontCloudIntegrationVerifyExtID(
						resourcePrefix, "wavefront_cloud_integration_aws_external_id.ext_id"),
					resource.TestCheckResourceAttr(
						resourcePrefix, "role_arn", "arn:aws::1234567:role/example-arn"),
				),
			},
		},
	})
}

func testAccCheckWavefrontCloudIntegrationCloudWatchBasic() string {
	return `
resource "wavefront_cloud_integration_aws_external_id" "ext_id" { 
}

resource "wavefront_cloud_integration_cloudwatch" "cloudwatch" {
  name       = "Test Integration"
  force_save = true
  additional_tags = {
    "tag1" = "value1"
    "tag2" = "value2"
  }
  role_arn    = "arn:aws::1234567:role/example-arn"
  external_id = wavefront_cloud_integration_aws_external_id.ext_id.id
  namespaces  = ["ec2", "elb"]
  instance_selection_tags = {
    "env"    = "prod"
    "mirror" = "a"
  }
  volume_selection_tags = {
    "env" = "prod"
  }
  point_tag_filter_regex = "^prod$"
  metric_filter_regex    = "^.*?\\.cpu.*$"
}`
}

func testAccCheckWavefrontCloudIntegrationCloudWatchBasicChanged() string {
	return `
resource "wavefront_cloud_integration_aws_external_id" "ext_id" {
}

resource "wavefront_cloud_integration_cloudwatch" "cloudwatch" {
  name       = "Test Integration"
  force_save = true
  additional_tags = {
    "tag1" = "value1"
    "tag2" = "value2"
  }
  role_arn    = "arn:aws::1234567:role/example-arn"
  external_id = wavefront_cloud_integration_aws_external_id.ext_id.id
  namespaces  = ["ec2", "elb", "route53"]
  instance_selection_tags = {
    "env"    = "dev"
    "mirror" = "b"
  }
  volume_selection_tags = {
    "env" = "dev"
  }
  point_tag_filter_regex = "^dev"
  metric_filter_regex    = "^.*?\\.cpu.*$"
}`
}
