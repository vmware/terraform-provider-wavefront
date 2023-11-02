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
		"alerts": {
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

func sourceLabelSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"host": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"firing": {
			Type:     schema.TypeInt,
			Computed: true,
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

	return d.Set("alerts", flattenAlerts(allAlerts))
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
	tfMap["id"] = *alert.ID
	tfMap["name"] = alert.Name
	tfMap["alert_type"] = alert.AlertType
	tfMap["additional_information"] = alert.AdditionalInfo
	tfMap["target"] = alert.Target
	tfMap["targets"] = alert.Targets
	tfMap["condition"] = alert.Condition
	tfMap["conditions"] = alert.Conditions
	tfMap["display_expression"] = alert.DisplayExpression
	tfMap["minutes"] = alert.Minutes
	tfMap["resolve_after_minutes"] = alert.ResolveAfterMinutes
	tfMap["notification_resend_frequency_minutes"] = alert.NotificationResendFrequencyMinutes
	tfMap["severity"] = alert.Severity
	tfMap["severity_list"] = alert.SeverityList
	tfMap["status"] = alert.Status
	tfMap["tags"] = alert.Tags
	tfMap["runbook_links"] = alert.RunbookLinks
	tfMap["can_view"] = alert.ACL.CanView
	tfMap["can_modify"] = alert.ACL.CanModify
	tfMap["process_rate_minutes"] = alert.CheckingFrequencyInMinutes
	tfMap["evaluate_realtime_data"] = alert.EvaluateRealtimeData
	tfMap["include_obsolete_metrics"] = alert.IncludeObsoleteMetrics
	tfMap["failing_host_label_pairs"] = flattenHostLabelPairs(alert.FailingHostLabelPairs)
	tfMap["in_maintenance_host_label_pairs"] = flattenHostLabelPairs(alert.InMaintenanceHostLabelPairs)
	tfMap["process_rate_minutes"] = alert.CheckingFrequencyInMinutes

	return tfMap
}
