package wavefront

import (
	"os"
	"strings"
	"testing"

	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccUser_Basic(t *testing.T) {
	resourceName := "wavefront_user.basic"
	var record wavefront.User

	config := testAccCheckWavefrontUserImporterBasic()
	if os.Getenv("TF_ACC") == "1" {
		replace := "tftesting"
		newCustomer := getCustomerName()
		config = strings.Replace(config, replace, newCustomer, 1)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontUserExists("wavefront_user.basic", &record),
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

func testAccCheckWavefrontUserImporterBasic() string {
	return `resource "wavefront_user" "basic" {
	       email  = "test+tftesting@example.com"
	       permissions = [
	               "agent_management",
	               "alerts_management",
	       ]
	}`
}
