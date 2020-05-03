package wavefront_plugin

import (
	"fmt"
	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

func TestAccWavefrontCloudIntegrationCloudTrail_Basic(t *testing.T) {
	var record wavefront.CloudIntegration
	resourcePrefix := "wavefront_cloud_integration_cloudtrail.cloudtrail"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			err := testAccCheckWavefrontCloudIntegrationDestroy(state)
			if err != nil {
				return err
			}
			err = testAccCheckWavefrontCloudIntegrationAwsExternalIdDestroy(state)
			if err != nil {
				return err
			}
			return nil
		},
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWavefrontCloudIntegrationCloudTrail_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontCloudIntegrationExists(resourcePrefix, &record),
					testAccCheckWavefrontCloudIntegrationAttributes(&record, wfCloudTrail),
					testAccCheckWavefrontCloudIntegrationVerifyExtId(
						resourcePrefix, "wavefront_cloud_integration_aws_external_id.ext_id"),
					// Check against state that the attributes are as we expect
					testAccCheckWavefrontCloudIntegrationResourceAttributes(resourcePrefix, wfCloudTrail),
					resource.TestCheckResourceAttr(resourcePrefix, "role_arn", "arn:aws::1234567:role/example-arn"),
					resource.TestCheckResourceAttr(resourcePrefix, "region", "us-west-2"),
					resource.TestCheckResourceAttr(resourcePrefix, "bucket_name", "example-s3-bucket"),
				),
			},
		},
	})
}

func TestAccWavefrontCloudIntegrationCloudTrail_BasicChanged(t *testing.T) {
	var record wavefront.CloudIntegration
	resourcePrefix := "wavefront_cloud_integration_cloudtrail.cloudtrail"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			err := testAccCheckWavefrontCloudIntegrationDestroy(state)
			if err != nil {
				return err
			}
			err = testAccCheckWavefrontCloudIntegrationAwsExternalIdDestroy(state)
			if err != nil {
				return err
			}
			return nil
		},
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWavefrontCloudIntegrationCloudTrail_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontCloudIntegrationExists(resourcePrefix, &record),
					// Check against state that the attributes are as we expect
					testAccCheckWavefrontCloudIntegrationResourceAttributes(resourcePrefix, wfCloudTrail),
					testAccCheckWavefrontCloudIntegrationAttributes(&record, wfCloudTrail),
					resource.TestCheckResourceAttr(resourcePrefix, "bucket_name", "example-s3-bucket"),
					resource.TestCheckResourceAttr(resourcePrefix, "region", "us-west-2"),
					testAccCheckWavefrontCloudIntegrationVerifyExtId(
						resourcePrefix, "wavefront_cloud_integration_aws_external_id.ext_id"),
					resource.TestCheckResourceAttr(resourcePrefix, "role_arn", "arn:aws::1234567:role/example-arn"),
				),
			},
			{
				Config: testAccCheckWavefrontCloudIntegrationCloudTrail_basicChanged(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontCloudIntegrationExists(resourcePrefix, &record),
					testAccCheckWavefrontCloudIntegrationAttributes(&record, wfCloudTrail),
					// Check against state that the attributes are as we expect
					testAccCheckWavefrontCloudIntegrationResourceAttributes(resourcePrefix, wfCloudTrail),
					resource.TestCheckResourceAttr(resourcePrefix, "bucket_name", "example-s3-bucket2"),
					resource.TestCheckResourceAttr(resourcePrefix, "region", "us-west-1"),
					resource.TestCheckResourceAttr(resourcePrefix, "filter_rule", "someFilterRule"),
					testAccCheckWavefrontCloudIntegrationVerifyExtId(
						resourcePrefix, "wavefront_cloud_integration_aws_external_id.ext_id"),
					resource.TestCheckResourceAttr(resourcePrefix, "role_arn", "arn:aws::1234567:role/example-arn"),
				),
			},
		},
	})
}

func testAccCheckWavefrontCloudIntegrationVerifyExtId(resourcePrefix, n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Record ID is set")
		}

		if err := resource.TestCheckResourceAttr(resourcePrefix, "external_id", rs.Primary.ID)(s); err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckWavefrontCloudIntegrationCloudTrail_basic() string {
	return fmt.Sprintf(`
resource "wavefront_cloud_integration_aws_external_id" "ext_id" { 
}

resource "wavefront_cloud_integration_cloudtrail" "cloudtrail" {
  name       = "Test Integration"
  force_save = true
  additional_tags = {
    "tag1" = "value1"
    "tag2" = "value2"
  }
  role_arn    = "arn:aws::1234567:role/example-arn"
  external_id = wavefront_cloud_integration_aws_external_id.ext_id.id
  region      = "us-west-2"
  bucket_name = "example-s3-bucket"
}`)
}

func testAccCheckWavefrontCloudIntegrationCloudTrail_basicChanged() string {
	return fmt.Sprintf(`
resource "wavefront_cloud_integration_aws_external_id" "ext_id" {
}

resource "wavefront_cloud_integration_cloudtrail" "cloudtrail" {
  name       = "Test Integration"
  force_save = true
  additional_tags = {
    "tag1" = "value1"
    "tag2" = "value2"
  }
  region      = "us-west-1"
  bucket_name = "example-s3-bucket2"
  filter_rule = "someFilterRule"
  role_arn    = "arn:aws::1234567:role/example-arn"
  external_id = wavefront_cloud_integration_aws_external_id.ext_id.id
}
`)
}
