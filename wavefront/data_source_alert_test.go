package wavefront

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccAlertIDRequired(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccAlertIDRequiredFailConfig,
				ExpectError: regexp.MustCompile("The argument \"id\" is required, but no definition was found."),
			},
		},
	})
}

const testAccAlertIDRequiredFailConfig = `
data "wavefront_alert" "test_alert" {
}
`
