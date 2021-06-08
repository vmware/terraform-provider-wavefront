package wavefront

import (
	"testing"

	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDerivedMetric_Basic(t *testing.T) {
	resourceName := "wavefront_derived_metric.derived"
	var record wavefront.DerivedMetric

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontDerivedMetricDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWavefrontDerivedMetricImporterBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontDerivedMetricExists("wavefront_derived_metric.derived", &record),
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

func testAccCheckWavefrontDerivedMetricImporterBasic() string {
	return `
resource "wavefront_derived_metric" "derived" {
  name                   = "dummy derived metric"
  minutes                = 5
  query                  = "aliasMetric(5, \"some.metric\")"
  additional_information = "this is a dummy derived metric"
  tags = [
    "somemetric",
    "thatistagged",
    "withmytags"
  ]
}
`
}
