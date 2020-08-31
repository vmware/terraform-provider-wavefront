package wavefront

import (
	"testing"

	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccTarget_Basic(t *testing.T) {
	resourceName := "wavefront_alert_target.foobar"
	var record wavefront.Target

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontTargetDestroy,
		Steps: []resource.TestStep{
			{
				ExpectNonEmptyPlan: true,
				Config:             testAccCheckWavefrontTargetImporterBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontTargetExists("wavefront_alert_target.foobar", &record),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"custom_headers.Testing"},
			},
		},
	})
}

func testAccCheckWavefrontTargetImporterBasic() string {
	return `
	resource "wavefront_alert_target" "foobar" {
	  name = "Terraform Test Target"
		description = "Test target"
		method = "WEBHOOK"
		recipient = "https://hooks.slack.com/services/test/me"
		content_type = "application/json"
		custom_headers = {
			"Testing" = "true"
		}
		template = "{}"
		triggers = [
			"ALERT_OPENED",
			"ALERT_RESOLVED"
		]
	}
	`
}
