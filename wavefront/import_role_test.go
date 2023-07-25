package wavefront

import (
	"testing"

	"github.com/WavefrontHQ/go-wavefront-management-api/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccRole_basic(t *testing.T) {
	var record wavefront.Role
	resourceName := "wavefront_role.role"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontCloudIntegrationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWavefrontRoleBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontRoleExists(resourceName, &record),
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
