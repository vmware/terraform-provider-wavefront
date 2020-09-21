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

func TestAccWavefrontMaintenanceWindow_Basic(t *testing.T) {
	var record wavefront.MaintenanceWindow

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontMaintenanceWindowDestroy,
		Steps: []resource.TestStep{
			{
				Config: `
resource "wavefront_maintenance_window" "basic" {
    reason = "A good reason"
    title = "A good title"
    start_time_in_seconds = 1600123456
    end_time_in_seconds = "1601123456"
    relevant_customer_tags = ["customer_tag_1", "customer_tag_2"]
    relevant_host_tags = ["host_tag_1"]
    relevant_host_names = ["host_names_1", "host_names_2", "host_names_3"]
    relevant_host_tags_anded = true
    host_tag_group_host_names_group_anded = true
}`,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontMaintenanceWindowExists(
						"wavefront_maintenance_window.basic", &record),
					testAccWavefrontMaintenanceWindowOptions(
						&wavefront.MaintenanceWindowOptions{
							Reason:             "A good reason",
							Title:              "A good title",
							StartTimeInSeconds: 1600123456,
							EndTimeInSeconds:   1601123456,
							RelevantCustomerTags: []string{
								"customer_tag_1",
								"customer_tag_2",
							},
							RelevantHostTags: []string{"host_tag_1"},
							RelevantHostNames: []string{
								"host_names_1",
								"host_names_2",
								"host_names_3",
							},
							RelevantHostTagsAnded:           true,
							HostTagGroupHostNamesGroupAnded: true,
						},
						&record,
					),

					// Check against state that the attributes are as we expect
					resource.TestCheckResourceAttr(
						"wavefront_maintenance_window.basic", "reason", "A good reason"),
					resource.TestCheckResourceAttr(
						"wavefront_maintenance_window.basic", "title", "A good title"),
					resource.TestCheckResourceAttr(
						"wavefront_maintenance_window.basic", "start_time_in_seconds", "1600123456"),
					resource.TestCheckResourceAttr(
						"wavefront_maintenance_window.basic", "end_time_in_seconds", "1601123456"),
					resource.TestCheckResourceAttr(
						"wavefront_maintenance_window.basic", "relevant_customer_tags.#", "2"),
					resource.TestCheckResourceAttr(
						"wavefront_maintenance_window.basic", "relevant_host_tags.#", "1"),
					resource.TestCheckResourceAttr(
						"wavefront_maintenance_window.basic", "relevant_host_names.#", "3"),
					resource.TestCheckResourceAttr(
						"wavefront_maintenance_window.basic", "relevant_host_tags_anded", "true"),
					resource.TestCheckResourceAttr(
						"wavefront_maintenance_window.basic", "host_tag_group_host_names_group_anded", "true"),
				),
			},
			{
				Config: `
resource "wavefront_maintenance_window" "basic" {
    reason = "A better reason"
    title = "A better title"
    start_time_in_seconds = 1603123456
    end_time_in_seconds = "1604123456"
    relevant_customer_tags = ["customer_tag_1", "customer_tag_2"]
}`,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontMaintenanceWindowExists(
						"wavefront_maintenance_window.basic", &record),
					testAccWavefrontMaintenanceWindowOptions(
						&wavefront.MaintenanceWindowOptions{
							Reason:             "A better reason",
							Title:              "A better title",
							StartTimeInSeconds: 1603123456,
							EndTimeInSeconds:   1604123456,
							RelevantCustomerTags: []string{
								"customer_tag_1",
								"customer_tag_2",
							},
						},
						&record,
					),

					// Check against state that the attributes are as we expect
					resource.TestCheckResourceAttr(
						"wavefront_maintenance_window.basic", "reason", "A better reason"),
					resource.TestCheckResourceAttr(
						"wavefront_maintenance_window.basic", "title", "A better title"),
					resource.TestCheckResourceAttr(
						"wavefront_maintenance_window.basic", "start_time_in_seconds", "1603123456"),
					resource.TestCheckResourceAttr(
						"wavefront_maintenance_window.basic", "end_time_in_seconds", "1604123456"),
					resource.TestCheckResourceAttr(
						"wavefront_maintenance_window.basic", "relevant_customer_tags.#", "2"),
					resource.TestCheckResourceAttr(
						"wavefront_maintenance_window.basic", "relevant_host_tags.#", "0"),
					resource.TestCheckResourceAttr(
						"wavefront_maintenance_window.basic", "relevant_host_names.#", "0"),
					resource.TestCheckResourceAttr(
						"wavefront_maintenance_window.basic", "relevant_host_tags_anded", "false"),
					resource.TestCheckResourceAttr(
						"wavefront_maintenance_window.basic", "host_tag_group_host_names_group_anded", "false"),
				),
			},
		},
	})
}

func testAccCheckWavefrontMaintenanceWindowDestroy(s *terraform.State) error {

	maintenanceWindows := testAccProvider.Meta().(*wavefrontClient).client.MaintenanceWindows()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "wavefront_maintenance_window" {
			continue
		}
		_, err := maintenanceWindows.GetByID(rs.Primary.ID)
		if wavefront.NotFound(err) {
			continue
		}
		if err != nil {
			return fmt.Errorf("error finding Wavefront Maintenance Window, %s", err)
		}
		return fmt.Errorf("maintenance window still exists, %s", rs.Primary.ID)
	}
	return nil
}

func testAccCheckWavefrontMaintenanceWindowExists(
	n string, mw *wavefront.MaintenanceWindow) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Record ID is set")
		}

		maintenanceWindows := testAccProvider.Meta().(*wavefrontClient).client.MaintenanceWindows()

		result, err := maintenanceWindows.GetByID(rs.Primary.ID)
		if wavefront.NotFound(err) {
			return fmt.Errorf("maintenance window not found %s", rs.Primary.ID)
		}
		if err != nil {
			return fmt.Errorf("error finding Wavefront maintenance window %s", err)
		}
		*mw = *result
		return nil
	}
}

func testAccWavefrontMaintenanceWindowOptions(
	expected *wavefront.MaintenanceWindowOptions,
	mw *wavefront.MaintenanceWindow) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		sortStringSlices := cmpopts.SortSlices(func(lhs, rhs string) bool {
			return lhs < rhs
		})
		result := cmp.Diff(
			expected, mw.Options(), sortStringSlices, cmpopts.EquateEmpty())
		if result == "" {
			return nil
		}
		return fmt.Errorf("options differ: %s", result)
	}
}
