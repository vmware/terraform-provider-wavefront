package wavefront

import (
	"testing"

	"github.com/WavefrontHQ/go-wavefront-management-api/v2"
	"github.com/stretchr/testify/assert"
)

func TestFlattenAlert(t *testing.T) {
	var testAlertID = "test-alert-id"
	// Create a sample wavefront.Alert object
	alert := &wavefront.Alert{
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
	flattened := flattenAlerts([]*wavefront.Alert{alert})[0]

	// Check if the flattened map matches the expected values
	assert.Equal(t, alert.Name, flattened["name"])
	assert.Equal(t, *alert.ID, flattened["id"])
	assert.Equal(t, alert.AlertType, flattened["alert_type"])
	assert.Equal(t, alert.AdditionalInfo, flattened["additional_information"])
	assert.Equal(t, alert.Target, flattened["target"])
	assert.Equal(t, alert.Targets, flattened["targets"])
	assert.Equal(t, alert.Condition, flattened["condition"])
	assert.Equal(t, alert.Conditions, flattened["conditions"])
	assert.Equal(t, alert.DisplayExpression, flattened["display_expression"])
	assert.Equal(t, alert.Minutes, flattened["minutes"])
	assert.Equal(t, alert.ResolveAfterMinutes, flattened["resolve_after_minutes"])
	assert.Equal(t, alert.NotificationResendFrequencyMinutes, flattened["notification_resend_frequency_minutes"])
	assert.Equal(t, alert.Severity, flattened["severity"])
	assert.Equal(t, alert.SeverityList, flattened["severity_list"])
	assert.Equal(t, alert.Status, flattened["status"])
	assert.Equal(t, alert.Tags, flattened["tags"])
	assert.Equal(t, alert.ACL.CanView, flattened["can_view"])
	assert.Equal(t, alert.ACL.CanModify, flattened["can_modify"])
	assert.Equal(t, alert.CheckingFrequencyInMinutes, flattened["process_rate_minutes"])
	assert.Equal(t, alert.EvaluateRealtimeData, flattened["evaluate_realtime_data"])
	assert.Equal(t, alert.IncludeObsoleteMetrics, flattened["include_obsolete_metrics"])
	assert.Equal(t, len(alert.InMaintenanceHostLabelPairs), len(flattened["failing_host_label_pairs"].([]map[string]interface{})))
	assert.Equal(t, len(alert.InMaintenanceHostLabelPairs), len(flattened["in_maintenance_host_label_pairs"].([]map[string]interface{})))
	assert.Equal(t, alert.RunbookLinks, flattened[runbookLinksKey])
	assert.Equal(t, alert.AlertTriageDashboards, flattened[alertTriageDashboardsKey].([]wavefront.AlertTriageDashboard))
}
