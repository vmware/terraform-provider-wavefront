package wavefront_plugin

import (
	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"testing"
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
				Config: testAccCheckWavefrontRole_basic(),
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
