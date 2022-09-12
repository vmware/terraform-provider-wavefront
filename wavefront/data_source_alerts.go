package wavefront

import (
	"time"

	"github.com/WavefrontHQ/go-wavefront-management-api"
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
	alertClient := m.(*wavefrontClient).client.Alerts()

	// Data Source ID is set to current time to always refresh
	d.SetId(time.Now().UTC().String())

	cont := true
	offset := 0

	for cont {
		filter := []*wavefront.SearchCondition{
			{Key: "limit", Value: string(rune(pageSize)), MatchingMethod: exactMatching},
			{Key: "offset", Value: string(rune(offset)), MatchingMethod: exactMatching},
		}

		alerts, err := alertClient.Find(filter)
		if err != nil {
			return err
		}

		allAlerts = append(allAlerts, alerts...)

		if len(allAlerts) < pageSize {
			cont = false
		} else {
			offset += pageSize
		}
	}

	if err := d.Set("alerts", flattenAlerts(allAlerts)); err != nil {
		return err
	}
	return nil
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
	tfMap["can_view"] = alert.ACL.CanView
	tfMap["can_modify"] = alert.ACL.CanModify
	tfMap["process_rate_minutes"] = alert.CheckingFrequencyInMinutes
	tfMap["evaluate_realtime_data"] = alert.EvaluateRealtimeData
	tfMap["include_obsolete_metrics"] = alert.IncludeObsoleteMetrics
	tfMap["failing_host_label_pairs"] = alert.FailingHostLabelPairs
	tfMap["in_maintenance_host_label_pairs"] = alert.InMaintenanceHostLabelPairs

	return tfMap
}
