package wavefront

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/WavefrontHQ/go-wavefront-management-api/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAlerts() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceAlertsRead,
		Schema: dataSourceAlertsSchema(),
	}
}

func dataSourceAlertsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		// Computed Values
		alertsKey: {
			Type:     schema.TypeSet,
			Computed: true,
			Elem: &schema.Resource{
				Schema: dataSourceAlertSchema(),
			},
		},
		limitKey: {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  100,
		},
		offsetKey: {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
	}

}

func dataSourceAlertsRead(d *schema.ResourceData, m interface{}) error {

	var allAlerts []*wavefront.Alert
	limit := d.Get(limitKey).(int)
	offset := d.Get(offsetKey).(int)

	if err := json.Unmarshal(searchAll(limit, offset, "alert", nil, nil, m), &allAlerts); err != nil {
		return fmt.Errorf("Response is invalid JSON")
	}

	// Data Source ID is set to current time to always refresh
	d.SetId(time.Now().UTC().String())

	return d.Set(alertsKey, flattenAlerts(allAlerts))
}

func flattenAlerts(alerts []*wavefront.Alert) []map[string]interface{} {
	tfMaps := make([]map[string]interface{}, len(alerts))
	for i, v := range alerts {
		tfMaps[i] = flattenAlert(v)
	}
	return tfMaps
}

func flattenAlert(alert *wavefront.Alert) map[string]interface{} {
	tfMap := make(map[string]interface{})
	tfMap[idKey] = *alert.ID
	tfMap[nameKey] = alert.Name
	tfMap[alertTypeKey] = alert.AlertType
	tfMap[additionalInformationKey] = alert.AdditionalInfo
	tfMap[targetKey] = alert.Target
	tfMap[targetsKey] = alert.Targets
	tfMap[conditionKey] = alert.Condition
	tfMap[conditionsKey] = alert.Conditions
	tfMap[displayExpressionKey] = alert.DisplayExpression
	tfMap[minutesKey] = alert.Minutes
	tfMap[resolveAfterMinutesKey] = alert.ResolveAfterMinutes
	tfMap[notificationResendFrequencyMinutesKey] = alert.NotificationResendFrequencyMinutes
	tfMap[severityKey] = alert.Severity
	tfMap[severityListKey] = alert.SeverityList
	tfMap[statusKey] = alert.Status
	tfMap[tagsKey] = alert.Tags
	tfMap[canViewKey] = alert.ACL.CanView
	tfMap[canModifyKey] = alert.ACL.CanModify
	tfMap[processRateMinutesKey] = alert.CheckingFrequencyInMinutes
	tfMap[evaluateRealtimeDataKey] = alert.EvaluateRealtimeData
	tfMap[includeObsoleteMetricsKey] = alert.IncludeObsoleteMetrics
	tfMap[failingHostLabelPairsKey] = flattenHostLabelPairs(alert.FailingHostLabelPairs)
	tfMap[inMaintenanceHostLabelPairsKey] = flattenHostLabelPairs(alert.InMaintenanceHostLabelPairs)
	tfMap[runbookLinksKey] = alert.RunbookLinks
	tfMap[alertTriageDashboardsKey] = alert.AlertTriageDashboards

	return tfMap
}
