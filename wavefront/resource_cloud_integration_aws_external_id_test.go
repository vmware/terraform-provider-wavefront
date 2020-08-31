package wavefront

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccWavefrontCloudIntegrationAwsExternalId_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontCloudIntegrationAwsExternalIDDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWavefrontCloudIntegrationAwsExternalIDBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontCloudIntegrationAwsExternalIDExists(),
				),
			},
		},
	})
}

func testAccCheckWavefrontCloudIntegrationAwsExternalIDExists() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources["wavefront_cloud_integration_aws_external_id.external_id"]

		if !ok {
			return fmt.Errorf("not found: %s", "wavefront_cloud_integration_aws_external_id.external_id")
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Record ID is set")
		}

		ci := testAccProvider.Meta().(*wavefrontClient).client.CloudIntegrations()
		err := ci.VerifyAwsExternalID(rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("error finding Wavefront Cloud Integration %s", err)
		}

		return nil
	}
}

func testAccCheckWavefrontCloudIntegrationAwsExternalIDDestroy(s *terraform.State) error {
	integrations := testAccProvider.Meta().(*wavefrontClient).client.CloudIntegrations()
	for _, rs := range s.RootModule().Resources {
		if !strings.Contains(rs.Type, "wavefront_cloud_integration") {
			continue
		}

		err := integrations.VerifyAwsExternalID(rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("aws external id still exists")
		}
	}
	return nil
}

func testAccCheckWavefrontCloudIntegrationAwsExternalIDBasic() string {
	return `
resource "wavefront_cloud_integration_aws_external_id" "external_id" {
}
`
}
