package wavefront

import (
	"regexp"
	"testing"

	"github.com/WavefrontHQ/go-wavefront-management-api/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/assert"
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

func TestParseAlertTriageDashboards(t *testing.T) {
	var dashboards []wavefront.AlertTriageDashboard

	dashboard1 := wavefront.AlertTriageDashboard{}
	dashboard1.DashboardId = "foo"
	dashboard1.Description = "bar"
	dashboard1.Parameters = map[string]map[string]string{"constants": {"key1": "val1", "key2": "val2"}}
	dashboards = append(dashboards, dashboard1)

	dashboard2 := wavefront.AlertTriageDashboard{}
	dashboard2.DashboardId = "bar"
	dashboard2.Description = "foo"
	dashboard2.Parameters = map[string]map[string]string{"constants": {"key3": "val3", "key4": "val4"}}
	dashboards = append(dashboards, dashboard2)

	result := parseAlertTriageDashboards(dashboards)

	expected := []map[string]interface{}{
		{
			"dashboard_id": "foo",
			"description":  "bar",
			"parameters": []map[string]interface{}{
				{
					"constants": map[string]string{
						"key1": "val1",
						"key2": "val2",
					},
				},
			},
		},
		{
			"dashboard_id": "bar",
			"description":  "foo",
			"parameters": []map[string]interface{}{
				{
					"constants": map[string]string{
						"key3": "val3",
						"key4": "val4",
					},
				},
			},
		},
	}

	assert.Equal(t, expected, result)

}

const testAccAlertIDRequiredFailConfig = `
data "wavefront_alert" "test_alert" {
}
`
