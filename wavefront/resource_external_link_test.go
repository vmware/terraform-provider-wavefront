package wavefront

import (
	"fmt"
	"testing"

	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccWavefrontExternalLink_Basic(t *testing.T) {
	var record wavefront.ExternalLink

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontExternalLinkDestroy,
		Steps: []resource.TestStep{
			{
				Config: `
resource "wavefront_external_link" "basic" {
    name = "External Link"
    description = "A description of the link"
    template = "https://example.com/source={{{source}}}&startTime={{startEpochMillis}}"
    metric_filter_regex = "^.*$"
    source_filter_regex = "^prod.*$"
    point_tag_filter_regexes = {
        service = "^query$"
    }
    is_log_integration = true
}`,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontExternalLinkExists(
						"wavefront_external_link.basic", &record),
					testAccWavefrontExternalLinkEquals(
						&wavefront.ExternalLink{
							Name:                  "External Link",
							Description:           "A description of the link",
							Template:              "https://example.com/source={{{source}}}&startTime={{startEpochMillis}}",
							MetricFilterRegex:     "^.*$",
							SourceFilterRegex:     "^prod.*$",
							PointTagFilterRegexes: map[string]string{"service": "^query$"},
							IsLogIntegration:      true,
						},
						&record,
					),
					// Check against state that the attributes are as we expect
					resource.TestCheckResourceAttr(
						"wavefront_external_link.basic", "name", "External Link"),
					resource.TestCheckResourceAttr(
						"wavefront_external_link.basic", "description", "A description of the link"),
					resource.TestCheckResourceAttr(
						"wavefront_external_link.basic", "template", "https://example.com/source={{{source}}}&startTime={{startEpochMillis}}"),
					resource.TestCheckResourceAttr(
						"wavefront_external_link.basic", "metric_filter_regex", "^.*$"),
					resource.TestCheckResourceAttr(
						"wavefront_external_link.basic", "source_filter_regex", "^prod.*$"),
					resource.TestCheckResourceAttr(
						"wavefront_external_link.basic", "is_log_integration", "true"),
				),
			},
			{
				Config: `
resource "wavefront_external_link" "basic" {
    name = "Internal Link"
    description = "An internal link"
    template = "http://example.com/source={{{source}}}&startTime={{startEpochMillis}}"
    metric_filter_regex = "^metric.*$"
    source_filter_regex = "^source.*$"
    point_tag_filter_regexes = {
        service = "^ts$"
    }
    is_log_integration = false
}`,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontExternalLinkExists(
						"wavefront_external_link.basic", &record),
					testAccWavefrontExternalLinkEquals(
						&wavefront.ExternalLink{
							Name:                  "Internal Link",
							Description:           "An internal link",
							Template:              "http://example.com/source={{{source}}}&startTime={{startEpochMillis}}",
							MetricFilterRegex:     "^metric.*$",
							SourceFilterRegex:     "^source.*$",
							PointTagFilterRegexes: map[string]string{"service": "^ts$"},
							IsLogIntegration:      false,
						},
						&record,
					),
					// Check against state that the attributes are as we expect
					resource.TestCheckResourceAttr(
						"wavefront_external_link.basic", "name", "Internal Link"),
					resource.TestCheckResourceAttr(
						"wavefront_external_link.basic", "description", "An internal link"),
					resource.TestCheckResourceAttr(
						"wavefront_external_link.basic", "template", "http://example.com/source={{{source}}}&startTime={{startEpochMillis}}"),
					resource.TestCheckResourceAttr(
						"wavefront_external_link.basic", "metric_filter_regex", "^metric.*$"),
					resource.TestCheckResourceAttr(
						"wavefront_external_link.basic", "source_filter_regex", "^source.*$"),
					resource.TestCheckResourceAttr(
						"wavefront_external_link.basic", "is_log_integration", "false"),
				),
			},
		},
	})
}

func testAccCheckWavefrontExternalLinkDestroy(s *terraform.State) error {

	externalLinks := testAccProvider.Meta().(*wavefrontClient).client.ExternalLinks()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "wavefront_external_link" {
			continue
		}
		id := rs.Primary.ID
		el := wavefront.ExternalLink{ID: &id}
		err := externalLinks.Get(&el)
		if wavefront.NotFound(err) {
			continue
		}
		if err != nil {
			return fmt.Errorf("error finding Wavefront External Link, %s", err)
		}
		return fmt.Errorf("external link still exists, %s", rs.Primary.ID)
	}
	return nil
}

func testAccCheckWavefrontExternalLinkExists(
	n string, el *wavefront.ExternalLink) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Record ID is set")
		}

		externalLinks := testAccProvider.Meta().(*wavefrontClient).client.ExternalLinks()
		id := rs.Primary.ID
		*el = wavefront.ExternalLink{ID: &id}
		err := externalLinks.Get(el)
		if wavefront.NotFound(err) {
			return fmt.Errorf("external link not found %s", rs.Primary.ID)
		}
		if err != nil {
			return fmt.Errorf("error finding Wavefront external link %s", err)
		}
		return nil
	}
}

func externalLinkZeroExtraFields(el *wavefront.ExternalLink) *wavefront.ExternalLink {
	elCopy := *el
	elCopy.ID = nil
	elCopy.CreatorId = ""
	elCopy.UpdaterId = ""
	elCopy.UpdatedEpochMillis = 0
	elCopy.CreatedEpochMillis = 0
	return &elCopy
}

func testAccWavefrontExternalLinkEquals(
	expected *wavefront.ExternalLink,
	el *wavefront.ExternalLink) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		result := cmp.Diff(
			externalLinkZeroExtraFields(expected),
			externalLinkZeroExtraFields(el),
			cmpopts.EquateEmpty())
		if result == "" {
			return nil
		}
		return fmt.Errorf("options differ: %s", result)
	}
}
