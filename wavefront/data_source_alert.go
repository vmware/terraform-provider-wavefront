package wavefront

import (
	"fmt"
	"time"

	"github.com/WavefrontHQ/go-wavefront-management-api/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAlert() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceAlertRead,
		Schema: dataSourceAlertSchema(),
	}
}

func dataSourceAlertSchema() map[string]*schema.Schema {

	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"alert_type": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"additional_information": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"target": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"targets": {
			Type:     schema.TypeMap,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"condition": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"conditions": {
			Type:     schema.TypeMap,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},

		"display_expression": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"minutes": {
			Type:     schema.TypeInt,
			Computed: true,
		},

		"resolve_after_minutes": {
			Type:     schema.TypeInt,
			Computed: true,
		},

		"notification_resend_frequency_minutes": {
			Type:     schema.TypeInt,
			Computed: true,
		},

		"severity": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"severity_list": {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},

		"status": {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},

		"tags": {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},

		"can_view": {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},

		"can_modify": {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},

		"process_rate_minutes": {
			Type:     schema.TypeInt,
			Computed: true,
		},

		"evaluate_realtime_data": {
			Type:     schema.TypeBool,
			Computed: true,
		},

		"include_obsolete_metrics": {
			Type:     schema.TypeBool,
			Computed: true,
		},

		"failing_host_label_pairs": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: sourceLabelSchema(),
			},
		},

		"in_maintenance_host_label_pairs": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: sourceLabelSchema(),
			},
		},

		"runbook_links": {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},

		"alert_triage_dashboards": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: alertTriageDashboardSchema(),
			},
		},
	}

}

func dataSourceAlertRead(d *schema.ResourceData, m interface{}) error {
	alertClient := m.(*wavefrontClient).client.Alerts()
	id, ok := d.GetOk("id")
	if !ok {
		return fmt.Errorf("required parameter '%s' not set", idKey)
	}

	idStr := fmt.Sprintf("%s", id)
	alert := wavefront.Alert{ID: &idStr}
	if err := alertClient.Get(&alert); err != nil {
		return err
	}

	// Data Source ID is set to current time to always refresh
	d.SetId(time.Now().UTC().String())
	return setAlertAttributes(d, alert)
}

func setAlertAttributes(d *schema.ResourceData, alert wavefront.Alert) error {
	if err := d.Set("id", *alert.ID); err != nil {
		return err
	}
	if err := d.Set("name", alert.Name); err != nil {
		return err
	}
	if err := d.Set("alert_type", alert.AlertType); err != nil {
		return err
	}
	if err := d.Set("additional_information", alert.AdditionalInfo); err != nil {
		return err
	}
	if err := d.Set("target", alert.Target); err != nil {
		return err
	}
	if err := d.Set("targets", alert.Targets); err != nil {
		return err
	}
	if err := d.Set("condition", alert.Condition); err != nil {
		return err
	}
	if err := d.Set("conditions", alert.Conditions); err != nil {
		return err
	}
	if err := d.Set("display_expression", alert.DisplayExpression); err != nil {
		return err
	}
	if err := d.Set("minutes", alert.Minutes); err != nil {
		return err
	}
	if err := d.Set("resolve_after_minutes", alert.ResolveAfterMinutes); err != nil {
		return err
	}
	if err := d.Set("notification_resend_frequency_minutes", alert.NotificationResendFrequencyMinutes); err != nil {
		return err
	}
	if err := d.Set("severity", alert.Severity); err != nil {
		return err
	}
	if err := d.Set("severity_list", alert.SeverityList); err != nil {
		return err
	}
	if err := d.Set("status", alert.Status); err != nil {
		return err
	}
	if err := d.Set("tags", alert.Tags); err != nil {
		return err
	}
	if err := d.Set("runbook_links", alert.RunbookLinks); err != nil {
		return err
	}
	if err := d.Set("can_view", alert.ACL.CanView); err != nil {
		return err
	}
	if err := d.Set("can_modify", alert.ACL.CanModify); err != nil {
		return err
	}
	if err := d.Set("process_rate_minutes", alert.CheckingFrequencyInMinutes); err != nil {
		return err
	}
	if err := d.Set("evaluate_realtime_data", alert.EvaluateRealtimeData); err != nil {
		return err
	}
	if err := d.Set("include_obsolete_metrics", alert.IncludeObsoleteMetrics); err != nil {
		return err
	}
	if err := d.Set("failing_host_label_pairs", flattenHostLabelPairs(alert.FailingHostLabelPairs)); err != nil {
		return err
	}
	if err := d.Set("in_maintenance_host_label_pairs", flattenHostLabelPairs(alert.InMaintenanceHostLabelPairs)); err != nil {
		return err
	}
	if err := d.Set("alert_triage_dashboards", parseAlertTriageDashboards(alert.AlertTriageDashboards)); err != nil {
		return err
	}
	return d.Set("process_rate_minutes", alert.CheckingFrequencyInMinutes)
}

func flattenHostLabelPairs(pairs []wavefront.SourceLabelPair) interface{} {
	tfMaps := make([]map[string]interface{}, len(pairs))
	for i, v := range pairs {
		tfMaps[i] = flattenHostLabelPair(v)
	}
	return tfMaps
}

func flattenHostLabelPair(pair wavefront.SourceLabelPair) map[string]interface{} {
	tfMap := make(map[string]interface{})
	tfMap["firing"] = pair.Firing
	tfMap["host"] = pair.Host
	return tfMap
}

func parseAlertTriageDashboards(alertTriageDashboards []wavefront.AlertTriageDashboard) []map[string]interface{} {
	triageDashboards := make([]map[string]interface{}, len(alertTriageDashboards))

	for i, alertTriageDashboard := range alertTriageDashboards {
		dashboardData := map[string]interface{}{
			"dashboard_id": alertTriageDashboard.DashboardId,
			"description":  alertTriageDashboard.Description,
		}

		parameters := make([]map[string]interface{}, 1)

		for key, value := range alertTriageDashboard.Parameters {
			parameters[0] = map[string]interface{}{key: value}
		}

		dashboardData["parameters"] = parameters
		triageDashboards[i] = dashboardData
	}

	return triageDashboards
}
