package wavefront

import (
	"fmt"
	"testing"

	"github.com/WavefrontHQ/go-wavefront-management-api/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	testDashboardID1   = "id-1"
	testDashboardID2   = "id-2"
	testDashboardDesc1 = "desc-1"
	testDashboardDesc2 = "desc-2"
	testKey1           = "key-1"
	testKey2           = "key-2"
	testKey3           = "key-3"
	testVal1           = "val-1"
	testVal2           = "val-2"
	testVal3           = "val-3"
)

func TestDecodeAlertTriageDashboards(t *testing.T) {
	testData := map[string]interface{}{
		"name": "Test Alert",
		"alert_triage_dashboards": []interface{}{
			map[string]interface{}{
				"dashboard_id": testDashboardID1,
				"description":  testDashboardDesc1,
				"parameters": []interface{}{
					map[string]interface{}{
						"constants": map[string]interface{}{
							testKey1: testVal1,
							testKey2: testVal2,
						},
					},
				},
			},
			map[string]interface{}{
				"dashboard_id": testDashboardID2,
				"description":  testDashboardDesc2,
				"parameters": []interface{}{
					map[string]interface{}{
						"constants": map[string]interface{}{
							testKey3: testVal3,
						},
					},
				},
			},
		},
	}

	resourceData := schema.TestResourceDataRaw(t, resourceAlert().Schema, testData)

	alertTriageDashboards := decodeAlertTriageDashboards(resourceData)

	// Check the length of the result slice
	if len(alertTriageDashboards) != 2 {
		t.Errorf("Expected 2 AlertTriageDashboard items, but got %d", len(alertTriageDashboards))
	}

	// Check the content of the first item
	if alertTriageDashboards[0].DashboardId != testDashboardID1 {
		t.Errorf("Expected DashboardId to be '%s', but got %s", testDashboardID1, alertTriageDashboards[0].DashboardId)
	}
	if alertTriageDashboards[0].Description != testDashboardDesc1 {
		t.Errorf("Expected Description to be '%s', but got %s", testDashboardDesc1, alertTriageDashboards[0].Description)
	}
	if alertTriageDashboards[0].Parameters["constants"][testKey1] != testVal1 {
		t.Errorf("Expected constants.%s to be '%s', but got %s", testKey1, testVal1, alertTriageDashboards[0].Parameters["constants"][testKey1])
	}
	if alertTriageDashboards[0].Parameters["constants"][testKey2] != testVal2 {
		t.Errorf("Expected constants.%s to be '%s', but got %s", testKey2, testVal2, alertTriageDashboards[0].Parameters["constants"][testKey2])
	}

	// Check the content of the second item
	if alertTriageDashboards[1].DashboardId != testDashboardID2 {
		t.Errorf("Expected DashboardId to be '%s', but got %s", testDashboardID2, alertTriageDashboards[1].DashboardId)
	}
	if alertTriageDashboards[1].Description != testDashboardDesc2 {
		t.Errorf("Expected Description to be '%s', but got %s", testDashboardDesc2, alertTriageDashboards[1].Description)
	}
	if alertTriageDashboards[1].Parameters["constants"][testKey3] != testVal3 {
		t.Errorf("Expected constants.%s to be '%s', but got %s", testKey3, testVal3, alertTriageDashboards[1].Parameters["constants"][testKey3])
	}
}

func TestValidateAlertTarget(t *testing.T) {
	email := "example@wavefront.com"
	pdKey := "pd:not-a-real-pagerduty-key"
	targetID := "target:efKj3H9aF"
	w, e := validateAlertTarget(email, "target")
	if len(w) != 0 || len(e) != 0 {
		t.Fatal("expected no errors on email address validation")
	}
	w, e = validateAlertTarget(pdKey, "target")
	if len(w) != 0 && len(e) != 0 {
		t.Fatal("expected no errors on pager-duty key target validation")
	}
	w, e = validateAlertTarget(targetID, "target")
	if len(w) != 0 && len(e) != 0 {
		t.Fatal("expected no errors on alert target validation")
	}
	_, e = validateAlertTarget("totally,invalid,target", "target")
	if len(e) == 0 {
		t.Fatal("expected error on invalid alert targets")
	}
	w, e = validateAlertTarget(fmt.Sprintf("%s,%s,%s", email, pdKey, targetID), "target")
	if len(w) != 0 && len(e) != 0 {
		t.Fatal("expected no errors on multiple alert target validation")
	}
}
func TestAccWavefrontAlert_Basic(t *testing.T) {
	var record wavefront.Alert

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontAlertDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWavefrontAlertBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontAlertExists("wavefront_alert.test_alert", &record),
					testAccCheckWavefrontAlertAttributes(&record),
					// Check against state that the attributes are as we expect
					resource.TestCheckResourceAttr(
						"wavefront_alert.test_alert", "name", "Terraform Test Alert"),
					resource.TestCheckResourceAttr(
						"wavefront_alert.test_alert", "target", "test@example.com"),
					resource.TestCheckResourceAttr(
						"wavefront_alert.test_alert", "condition", "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total ) > 80"),
					resource.TestCheckResourceAttr(
						"wavefront_alert.test_alert", "additional_information", "This is a Terraform Test Alert"),
					resource.TestCheckResourceAttr(
						"wavefront_alert.test_alert", "display_expression", "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total )"),
					resource.TestCheckResourceAttr(
						"wavefront_alert.test_alert", "minutes", "5"),
					resource.TestCheckResourceAttr(
						"wavefront_alert.test_alert", "resolve_after_minutes", "5"),
					resource.TestCheckResourceAttr(
						"wavefront_alert.test_alert", "severity", "WARN"),
					resource.TestCheckResourceAttr(
						"wavefront_alert.test_alert", "tags.#", "5"),
					resource.TestCheckResourceAttr(
						"wavefront_alert.test_alert", "process_rate_minutes", "2"),
				),
			},
		},
	})
}
func TestAccWavefrontAlert_RequiredAttributes(t *testing.T) {
	var record wavefront.Alert
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontAlertDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWavefrontAlertRequiredAttributes(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontAlertExists("wavefront_alert.test_alert_required", &record),
					testAccCheckWavefrontAlertAttributes(&record),
					// Check against state that the attributes are as we expect
					resource.TestCheckResourceAttr(
						"wavefront_alert.test_alert_required", "name", "Terraform Test Alert Required Attributes Only"),
					resource.TestCheckResourceAttr(
						"wavefront_alert.test_alert_required", "target", "test@example.com"),
					resource.TestCheckResourceAttr(
						"wavefront_alert.test_alert_required", "condition", "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total ) > 80"),
					resource.TestCheckResourceAttr(
						"wavefront_alert.test_alert_required", "additional_information", "This is a Terraform Test Alert Required"),
					resource.TestCheckResourceAttr(
						"wavefront_alert.test_alert_required", "minutes", "5"),
					resource.TestCheckResourceAttr(
						"wavefront_alert.test_alert_required", "severity", "WARN"),
					resource.TestCheckResourceAttr(
						"wavefront_alert.test_alert_required", "tags.#", "2"),
				),
			},
		},
	})
}
func TestAccWavefrontAlert_Updated(t *testing.T) {
	var record wavefront.Alert

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontAlertDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWavefrontAlertBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontAlertExists("wavefront_alert.test_alert", &record),
					testAccCheckWavefrontAlertAttributes(&record),
					resource.TestCheckResourceAttr(
						"wavefront_alert.test_alert", "target", "test@example.com"),
				),
			},
			{
				Config: testAccCheckWavefrontAlertNewValue(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontAlertExists("wavefront_alert.test_alert", &record),
					testAccCheckWavefrontAlertAttributesUpdated(&record),
					resource.TestCheckResourceAttr(
						"wavefront_alert.test_alert", "target", "terraform@example.com"),
				),
			},
		},
	})
}
func TestAccWavefrontAlert_RemoveOptionalAttribute(t *testing.T) {
	var record wavefront.Alert
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontAlertDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWavefrontAlertRemoveAttributes(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontAlertExists("wavefront_alert.test_alert_required", &record),
					testAccCheckWavefrontAlertAttributesRemoved(&record),
					resource.TestCheckResourceAttr(
						"wavefront_alert.test_alert_required", "target", "test@example.com"),
				),
			},
			{
				Config: testAccCheckWavefrontAlertUpdatedRemoveAttributes(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontAlertExists("wavefront_alert.test_alert_required", &record),
					testAccCheckWavefrontAlertAttributesRemovedUpdated(&record),
					resource.TestCheckResourceAttr(
						"wavefront_alert.test_alert_required", "target", "terraform@example.com"),
				),
			},
		},
	})
}
func TestAccWavefrontAlert_Multiple(t *testing.T) {
	var record wavefront.Alert
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontAlertDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWavefrontAlertMultiple(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontAlertExists("wavefront_alert.test_alert1", &record),
					testAccCheckWavefrontAlertAttributes(&record),
					resource.TestCheckResourceAttr(
						"wavefront_alert.test_alert1", "name", "Terraform Test Alert 1"),
					resource.TestCheckResourceAttr(
						"wavefront_alert.test_alert2", "name", "Terraform Test Alert 2"),
					resource.TestCheckResourceAttr(
						"wavefront_alert.test_alert3", "name", "Terraform Test Alert 3"),
				),
			},
		},
	})
}
func TestAccWavefrontAlert_Threshold(t *testing.T) {
	var record wavefront.Alert
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontAlertDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWavefrontAlertThreshold(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontAlertExists("wavefront_alert.test_threshold_alert", &record),
					testAccCheckWavefrontThresholdAlertAttributes(&record),
					//Check against state that the attributes are as we expect
					resource.TestCheckResourceAttr(
						"wavefront_alert.test_threshold_alert", "conditions.%", "3"),
					resource.TestCheckResourceAttr(
						"wavefront_alert.test_threshold_alert", "threshold_targets.%", "1"),
				),
			},
		},
	})
}
func TestAccWavefrontAlert_ThresholdWithCondition(t *testing.T) {
	var record wavefront.Alert
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontAlertDestroy,
		Steps: []resource.TestStep{
			// change the condition and verify the condition change is ignored.
			{
				Config: testAccCheckWavefrontAlertThresholdChangeCondition(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontAlertExists(
						"wavefront_alert.test_threshold_alert_change_condition", &record),
					testAccCheckWavefrontThresholdAlertAttributes(&record),
					// Check against state that the attributes are as we expect
					// TODO: figure out why verification on `condition` using `resource.TestCheckResourceAttr` does not
					//  work.
					resource.TestCheckResourceAttr(
						"wavefront_alert.test_threshold_alert_change_condition", "conditions.%", "3"),
					resource.TestCheckResourceAttr(
						"wavefront_alert.test_threshold_alert_change_condition", "threshold_targets.%", "1"),
				),
			},
		},
	})
}
func TestResourceAlert_validateAlertConditions(t *testing.T) {
	cases := []struct {
		name         string
		conf         *schema.ResourceData
		errorMessage string
	}{
		{
			"invalid alert type",
			func() *schema.ResourceData {
				d := resourceAlert().TestResourceData()
				d.Set("alert_type", "WRONG")
				return d
			}(),
			"alert_type must be CLASSIC or THRESHOLD",
		},
		{
			"classic alert missing condition",
			func() *schema.ResourceData {
				d := resourceAlert().TestResourceData()
				d.Set("alert_type", "CLASSIC")
				d.Set("severity", "severe")
				return d
			}(),
			"condition must be supplied for classic alerts",
		},
		{
			"classic alert missing severity",
			func() *schema.ResourceData {
				d := resourceAlert().TestResourceData()
				d.Set("alert_type", "CLASSIC")
				d.Set("condition", "ts()")
				return d
			}(),
			"severity must be supplied for classic alerts",
		},
		{
			"classic alert",
			func() *schema.ResourceData {
				d := resourceAlert().TestResourceData()
				d.Set("alert_type", "CLASSIC")
				d.Set("condition", "ts()")
				d.Set("severity", "severe")
				return d
			}(),
			"",
		},
		{
			"threshold alert missing conditions",
			func() *schema.ResourceData {
				d := resourceAlert().TestResourceData()
				d.Set("alert_type", "THRESHOLD")
				return d
			}(),
			"conditions must be supplied for threshold alerts",
		},
		{
			"threshold alert",
			func() *schema.ResourceData {
				d := resourceAlert().TestResourceData()
				d.Set("alert_type", "THRESHOLD")
				d.Set("conditions", map[string]interface{}{"severe": "ts()"})
				return d
			}(),
			"",
		},
		{
			"threshold alert invalid condition",
			func() *schema.ResourceData {
				d := resourceAlert().TestResourceData()
				d.Set("alert_type", "THRESHOLD")
				d.Set("conditions", map[string]interface{}{"banana": "ts()"})
				return d
			}(),
			"invalid severity: banana",
		},
		{
			"threshold alert invalid target",
			func() *schema.ResourceData {
				d := resourceAlert().TestResourceData()
				d.Set("alert_type", "THRESHOLD")
				d.Set("conditions", map[string]interface{}{"severe": "ts()"})
				d.Set("threshold_targets", map[string]interface{}{"banana": "ts()"})
				return d
			}(),
			"invalid severity: banana",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := validateAlertConditions(&wavefront.Alert{}, c.conf)
			m := ""
			if err == nil {
				m = ""
			} else {
				m = err.Error()
			}
			if m != c.errorMessage {
				t.Errorf("expected error '%s', got '%s'", c.errorMessage, err.Error())
			}
		})
	}
}
func testAccCheckWavefrontAlertDestroy(s *terraform.State) error {
	alerts := testAccProvider.Meta().(*wavefrontClient).client.Alerts()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "wavefront_alert" {
			continue
		}
		tmpAlert := wavefront.Alert{ID: &rs.Primary.ID}
		err := alerts.Get(&tmpAlert)
		if err == nil {
			return fmt.Errorf("alert still exists")
		}
	}
	return nil
}
func testAccCheckWavefrontAlertAttributes(alert *wavefront.Alert) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if alert.Target != "test@example.com" && alert.Target != "test@example.com,foo@example.com" {
			return fmt.Errorf("bad value: %s", alert.Target)
		}
		return nil
	}
}
func testAccCheckWavefrontThresholdAlertAttributes(alert *wavefront.Alert) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if alert.Condition != "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total )" {
			return fmt.Errorf("bad value: %s", alert.Condition)
		}
		if val, ok := alert.Conditions["severe"]; ok {
			if val != "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total ) > 80" {
				return fmt.Errorf("bad value: %s", alert.Conditions["severe"])
			}
		} else {
			return fmt.Errorf("multi-threshold alert's conditions are not set")
		}
		return nil
	}
}
func testAccCheckWavefrontAlertAttributesUpdated(alert *wavefront.Alert) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if alert.Target != "terraform@example.com" {
			return fmt.Errorf("bad value: %s", alert.Target)
		}
		return nil
	}
}
func testAccCheckWavefrontAlertAttributesRemoved(alert *wavefront.Alert) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if alert.ResolveAfterMinutes != 5 {
			return fmt.Errorf("unexpected value for ResolveAfterMinutes %v, expected 5", alert.ResolveAfterMinutes)
		}
		return nil
	}
}
func testAccCheckWavefrontAlertAttributesRemovedUpdated(alert *wavefront.Alert) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if alert.ResolveAfterMinutes != 0 {
			return fmt.Errorf("unexpected value for ResolveAfterMinutes %v, expected 0", alert.ResolveAfterMinutes)
		}
		return nil
	}
}
func testAccCheckWavefrontAlertExists(n string, alert *wavefront.Alert) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no Record ID is set")
		}
		alerts := testAccProvider.Meta().(*wavefrontClient).client.Alerts()
		tmpAlert := wavefront.Alert{ID: &rs.Primary.ID}
		err := alerts.Get(&tmpAlert)
		if err != nil {
			return fmt.Errorf("error finding Wavefront Alert %s", err)
		}
		*alert = tmpAlert
		return nil
	}
}
func testAccCheckWavefrontAlertBasic() string {
	return `
resource "wavefront_user" "basic" {
	email  = "test+tftesting@example.com"
	permissions = [
		"agent_management",
		"alerts_management",
	]
}
resource "wavefront_alert" "test_alert" {
  name = "Terraform Test Alert"
  target = "test@example.com"
  condition = "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total ) > 80"
  additional_information = "This is a Terraform Test Alert"
  display_expression = "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total )"
  minutes = 5
  resolve_after_minutes = 5
  severity = "WARN"
  tags = [
	"b",
    "terraform",
    "c",
    "test",
    "a"
  ]
  can_view = [
    wavefront_user.basic.id,
  ]
  process_rate_minutes = 2
}
`
}
func testAccCheckWavefrontAlertRemoveAttributes() string {
	return `
resource "wavefront_alert" "test_alert_required" {
  name = "Terraform Test Alert Required Attributes Only"
  target = "test@example.com"
  condition = "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total ) > 80"
  additional_information = "This is a Terraform Test Alert Required"
  minutes = 5
  resolve_after_minutes = 5
  severity = "WARN"
  tags = [
    "terraform",
    "test"
  ]
}
`
}
func testAccCheckWavefrontAlertUpdatedRemoveAttributes() string {
	return `
resource "wavefront_alert" "test_alert_required" {
  name = "Terraform Test Alert Required Attributes Only"
  target = "terraform@example.com"
  condition = "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total ) > 80"
  additional_information = "This is a Terraform Test Alert Required"
  minutes = 5
  severity = "WARN"
  tags = [
    "terraform",
    "test"
  ]
}
`
}
func testAccCheckWavefrontAlertRequiredAttributes() string {
	return `
resource "wavefront_alert" "test_alert_required" {
  name = "Terraform Test Alert Required Attributes Only"
  target = "test@example.com"
  condition = "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total ) > 80"
  additional_information = "This is a Terraform Test Alert Required"
  minutes = 5
  severity = "WARN"
  tags = [
    "terraform",
    "test"
  ]
}
`
}
func testAccCheckWavefrontAlertNewValue() string {
	return `
resource "wavefront_alert" "test_alert" {
  name = "Terraform Test Alert"
  target = "terraform@example.com"
  condition = "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total ) > 80"
  additional_information = "This is a Terraform Test Alert"
  display_expression = "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total )"
  minutes = 5
  resolve_after_minutes = 5
  severity = "WARN"
  tags = [
    "terraform",
    "test"
  ]
}
`
}
func testAccCheckWavefrontAlertMultiple() string {
	return `
resource "wavefront_alert" "test_alert1" {
  name = "Terraform Test Alert 1"
  target = "test@example.com,foo@example.com"
  condition = "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total ) > 80"
  additional_information = "This is a Terraform Test Alert"
  display_expression = "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total )"
  minutes = 5
  resolve_after_minutes = 5
  severity = "WARN"
  tags = [
    "terraform1",
  ]
}
resource "wavefront_alert" "test_alert2" {
  name = "Terraform Test Alert 2"
  target = "test@example.com"
  condition = "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total ) > 80"
  additional_information = "This is a Terraform Test Alert"
  display_expression = "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total )"
  minutes = 5
  resolve_after_minutes = 5
  severity = "WARN"
  tags = [
    "terraform2",
    "test"
  ]
}
resource "wavefront_alert" "test_alert3" {
  name = "Terraform Test Alert 3"
  target = "test@example.com"
  condition = "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total ) > 80"
  additional_information = "This is a Terraform Test Alert"
  display_expression = "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total )"
  minutes = 5
  resolve_after_minutes = 5
  severity = "WARN"
  tags = [
    "terraform",
  ]
}
`
}
func testAccCheckWavefrontAlertThreshold() string {
	return `
resource "wavefront_alert_target" "test_target" {
  name = "Terraform Test Target"
  description = "Test target"
  method = "EMAIL"
  recipient = "test@example.com"
  email_subject = "This is a test"
  is_html_content = true
  template = "{}"
  triggers = [
    "ALERT_OPENED",
    "ALERT_RESOLVED"
  ]
}
resource "wavefront_alert" "test_threshold_alert" {
  name = "Terraform Test Alert"
  alert_type = "THRESHOLD"
  additional_information = "This is a Terraform Test Alert"
  display_expression = "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total )"
  minutes = 5
  resolve_after_minutes = 5
  conditions = {
    "severe" = "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total ) > 80"
    "warn" = "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total ) > 60"
    "info" = "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total ) > 50"
  }
  threshold_targets = {
	"severe" = "target:${wavefront_alert_target.test_target.id}"
  }
  
  tags = [
    "terraform"
  ]
}
`
}
func testAccCheckWavefrontAlertThresholdChangeCondition() string {
	return `
resource "wavefront_alert_target" "test_target" {
  name = "Terraform Test Target"
  description = "Test target"
  method = "EMAIL"
  recipient = "test@example.com"
  email_subject = "This is a test"
  is_html_content = true
  template = "{}"
  triggers = [
    "ALERT_OPENED",
    "ALERT_RESOLVED"
  ]
}
resource "wavefront_alert" "test_threshold_alert_change_condition" {
  name = "Terraform Test Alert"
  alert_type = "THRESHOLD"
  additional_information = "This is a Terraform Test Alert"
  # change in condition for multi-threshold alert takes no effect, wavefront backend will force sync the condition
  # with display_expression
  condition = "change_condition"
  display_expression = "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total )"
  minutes = 5
  resolve_after_minutes = 5
  conditions = {
    "severe" = "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total ) > 80"
    "warn" = "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total ) > 60"
    "info" = "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total ) > 50"
  }
  threshold_targets = {
	"severe" = "target:${wavefront_alert_target.test_target.id}"
  }
  tags = [
    "terraform"
  ]
}
`
}
