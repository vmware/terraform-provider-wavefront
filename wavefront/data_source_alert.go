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
		nameKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		idKey: {
			Type:     schema.TypeString,
			Required: true,
		},
		alertTypeKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		additionalInformationKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		targetKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		targetsKey: {
			Type:     schema.TypeMap,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		conditionKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		conditionsKey: {
			Type:     schema.TypeMap,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},

		displayExpressionKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		minutesKey: {
			Type:     schema.TypeInt,
			Computed: true,
		},

		resolveAfterMinutesKey: {
			Type:     schema.TypeInt,
			Computed: true,
		},

		notificationResendFrequencyMinutesKey: {
			Type:     schema.TypeInt,
			Computed: true,
		},

		severityKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		severityListKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},

		statusKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},

		tagsKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},

		canViewKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},

		canModifyKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},

		processRateMinutesKey: {
			Type:     schema.TypeInt,
			Computed: true,
		},

		evaluateRealtimeDataKey: {
			Type:     schema.TypeBool,
			Computed: true,
		},

		includeObsoleteMetricsKey: {
			Type:     schema.TypeBool,
			Computed: true,
		},

		failingHostLabelPairsKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: sourceLabelSchema(),
			},
		},

		inMaintenanceHostLabelPairsKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: sourceLabelSchema(),
			},
		},

		runbookLinksKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},

		alertTriageDashboardsKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: alertTriageDashboardSchema(),
			},
		},
	}
}

func alertTriageDashboardSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		dashboardIDKey: {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Dashboard ID",
		},
		descriptionKey: {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Dashboard Description",
		},
		parametersKey: {
			MaxItems: 1, // There should be only one "parameters" block
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					constantsKey: { // Currently, only "constants" are supported
						Type:     schema.TypeMap,
						Optional: true,
						Elem:     schema.TypeString,
					},
				},
			},
		},
	}
}

func sourceLabelSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		hostKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		firingKey: {
			Type:     schema.TypeInt,
			Computed: true,
		},
	}
}

func dataSourceAlertRead(d *schema.ResourceData, m interface{}) error {
	alertClient := m.(*wavefrontClient).client.Alerts()
	id, ok := d.GetOk(idKey)
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
	if err := d.Set(idKey, *alert.ID); err != nil {
		return err
	}
	if err := d.Set(nameKey, alert.Name); err != nil {
		return err
	}
	if err := d.Set(alertTypeKey, alert.AlertType); err != nil {
		return err
	}
	if err := d.Set(additionalInformationKey, alert.AdditionalInfo); err != nil {
		return err
	}
	if err := d.Set(targetKey, alert.Target); err != nil {
		return err
	}
	if err := d.Set(targetsKey, alert.Targets); err != nil {
		return err
	}
	if err := d.Set(conditionKey, alert.Condition); err != nil {
		return err
	}
	if err := d.Set(conditionsKey, alert.Conditions); err != nil {
		return err
	}
	if err := d.Set(displayExpressionKey, alert.DisplayExpression); err != nil {
		return err
	}
	if err := d.Set(minutesKey, alert.Minutes); err != nil {
		return err
	}
	if err := d.Set(resolveAfterMinutesKey, alert.ResolveAfterMinutes); err != nil {
		return err
	}
	if err := d.Set(notificationResendFrequencyMinutesKey, alert.NotificationResendFrequencyMinutes); err != nil {
		return err
	}
	if err := d.Set(severityKey, alert.Severity); err != nil {
		return err
	}
	if err := d.Set(severityListKey, alert.SeverityList); err != nil {
		return err
	}
	if err := d.Set(statusKey, alert.Status); err != nil {
		return err
	}
	if err := d.Set(tagsKey, alert.Tags); err != nil {
		return err
	}
	if err := d.Set(runbookLinksKey, alert.RunbookLinks); err != nil {
		return err
	}
	if err := d.Set(canViewKey, alert.ACL.CanView); err != nil {
		return err
	}
	if err := d.Set(canModifyKey, alert.ACL.CanModify); err != nil {
		return err
	}
	if err := d.Set(processRateMinutesKey, alert.CheckingFrequencyInMinutes); err != nil {
		return err
	}
	if err := d.Set(evaluateRealtimeDataKey, alert.EvaluateRealtimeData); err != nil {
		return err
	}
	if err := d.Set(includeObsoleteMetricsKey, alert.IncludeObsoleteMetrics); err != nil {
		return err
	}
	if err := d.Set(failingHostLabelPairsKey, flattenHostLabelPairs(alert.FailingHostLabelPairs)); err != nil {
		return err
	}
	if err := d.Set(inMaintenanceHostLabelPairsKey, flattenHostLabelPairs(alert.InMaintenanceHostLabelPairs)); err != nil {
		return err
	}
	if err := d.Set(alertTriageDashboardsKey, parseAlertTriageDashboards(alert.AlertTriageDashboards)); err != nil {
		return err
	}
	return nil
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
	tfMap[firingKey] = pair.Firing
	tfMap[hostKey] = pair.Host
	return tfMap
}

func parseAlertTriageDashboards(alertTriageDashboards []wavefront.AlertTriageDashboard) (dashboards []map[string]interface{}) {
	for _, alertTriageDashboard := range alertTriageDashboards {
		dashboards = append(dashboards, parseAlertTriageDashboard(alertTriageDashboard))
	}
	return dashboards
}

func parseAlertTriageDashboard(alertTriageDashboard wavefront.AlertTriageDashboard) (dashboard map[string]interface{}) {
	dashboard = map[string]interface{}{
		dashboardIDKey: alertTriageDashboard.DashboardId,
		descriptionKey: alertTriageDashboard.Description,
		parametersKey:  parseAlertTriageDashboardParameters(alertTriageDashboard.Parameters),
	}
	return dashboard
}

func parseAlertTriageDashboardParameters(alertTriageDashboardParameters map[string]map[string]string) (parameters []map[string]interface{}) {
	for key, value := range alertTriageDashboardParameters {
		parameters = append(parameters, map[string]interface{}{key: value})
	}
	return parameters
}
