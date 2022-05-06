package wavefront

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccWavefrontDataSourceUser_Ok(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontMetricsPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWavefrontMetricsPolicyBasic(),
				Check:  testAccCheckWavefrontCustomMetricsPolicy(testAccGetBasicPolicy),
			},
		},
	})
}

func TestAccWavefrontPolicy_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontMetricsPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWavefrontMetricsPolicyBasic(),
				Check:  testAccCheckWavefrontCustomMetricsPolicy(testAccGetBasicPolicy),
			},
		},
	})
}

func testAccCheckWavefrontDataSourceUserOkPolicy() string {
	return `
`
}
