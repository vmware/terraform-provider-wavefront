package wavefront

import (
	"testing"

	"github.com/WavefrontHQ/go-wavefront-management-api/v2"
	"github.com/stretchr/testify/assert"
)

func TestFlattenAlert(t *testing.T) {
	var testAlertID = "test-alert-id"
	// Create a sample wavefront.Alert object
	expected := &wavefront.Alert{
		Name:           testAlertName,
		ID:             &testAlertID,
		AlertType:      "CLASSIC",
		AdditionalInfo: "Additional information about the alert",
		Target:         "example.target",
		Targets: map[string]string{
			"severe": "severeTarget",
			"warn":   "warnTarget",
		},
		Condition: "metric > threshold",
		Conditions: map[string]string{
			"severe": "severeCondition",
			"warn":   "warnCondition",
		},
		DisplayExpression:                  "ts(query)",
		Minutes:                            5,
		ResolveAfterMinutes:                10,
		NotificationResendFrequencyMinutes: 15,
		Severity:                           "SEVERE",
		SeverityList:                       []string{"SEVERE", "WARN"},
		Status:                             []string{"Active"},
		Tags:                               []string{"tag1", "tag2"},
		RunbookLinks:                       []string{"link1", "link2"},
		ACL: wavefront.AccessControlList{
			CanView:   []string{"user1", "user2"},
			CanModify: []string{"user3", "user4"},
		},
		CheckingFrequencyInMinutes: 5,
		EvaluateRealtimeData:       true,
		IncludeObsoleteMetrics:     true,
		FailingHostLabelPairs: []wavefront.SourceLabelPair{
			{
				Host:   "host1",
				Firing: 1,
			},
			{
				Host:   "host2",
				Firing: 0,
			},
		},
		InMaintenanceHostLabelPairs: []wavefront.SourceLabelPair{
			{
				Host:   "host3",
				Firing: 1,
			},
			{
				Host:   "host4",
				Firing: 0,
			},
		},
		AlertTriageDashboards: []wavefront.AlertTriageDashboard{
			{
				DashboardId: "dashboard1",
				Parameters:  map[string]map[string]string{constantsKey: {testKey1: testVal1, testKey2: testVal2}},
				Description: "Dashboard 1",
			},
			{
				DashboardId: "dashboard2",
				Parameters:  map[string]map[string]string{constantsKey: {testKey1: testVal1, testKey2: testVal2}},
				Description: "Dashboard 2",
			},
		},
	}

	// Call the function to flatten the alert
	actual := flattenAlerts([]*wavefront.Alert{expected})[0]

	// Check if the flattened map matches the expected values
	assert.Equal(t, expected.Name, actual["name"])
	assert.Equal(t, *expected.ID, actual["id"])
	assert.Equal(t, expected.AlertType, actual["alert_type"])
	assert.Equal(t, expected.AdditionalInfo, actual["additional_information"])
	assert.Equal(t, expected.Target, actual["target"])
	assert.Equal(t, expected.Targets, actual["targets"])
	assert.Equal(t, expected.Condition, actual["condition"])
	assert.Equal(t, expected.Conditions, actual["conditions"])
	assert.Equal(t, expected.DisplayExpression, actual["display_expression"])
	assert.Equal(t, expected.Minutes, actual["minutes"])
	assert.Equal(t, expected.ResolveAfterMinutes, actual["resolve_after_minutes"])
	assert.Equal(t, expected.NotificationResendFrequencyMinutes, actual["notification_resend_frequency_minutes"])
	assert.Equal(t, expected.Severity, actual["severity"])
	assert.Equal(t, expected.SeverityList, actual["severity_list"])
	assert.Equal(t, expected.Status, actual["status"])
	assert.Equal(t, expected.Tags, actual["tags"])
	assert.Equal(t, expected.ACL.CanView, actual["can_view"])
	assert.Equal(t, expected.ACL.CanModify, actual["can_modify"])
	assert.Equal(t, expected.CheckingFrequencyInMinutes, actual["process_rate_minutes"])
	assert.Equal(t, expected.EvaluateRealtimeData, actual["evaluate_realtime_data"])
	assert.Equal(t, expected.IncludeObsoleteMetrics, actual["include_obsolete_metrics"])
	assert.Equal(t, len(expected.InMaintenanceHostLabelPairs), len(actual["failing_host_label_pairs"].([]map[string]interface{})))
	assert.Equal(t, len(expected.InMaintenanceHostLabelPairs), len(actual["in_maintenance_host_label_pairs"].([]map[string]interface{})))
	assert.Equal(t, expected.RunbookLinks, actual[runbookLinksKey])
	assert.Equal(t, expected.AlertTriageDashboards, actual[alertTriageDashboardsKey].([]wavefront.AlertTriageDashboard))
}
