package wavefront

import (
	"testing"

	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccUserGroup_Basic(t *testing.T) {
	resourceName := "wavefront_user_group.basic"
	var record wavefront.UserGroup

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontUserGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWavefrontUserGroupImporterBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontUserGroupExists("wavefront_user_group.basic", &record),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckWavefrontUserGroupImporterBasic() string {
	return `
resource "wavefront_user_group" "basic" {
  name        = "Basic User Group"
  description = "Basic User Group for Unit Tests"
}
`
}
