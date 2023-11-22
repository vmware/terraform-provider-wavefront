package wavefront

import (
	"fmt"
	"strings"

	"github.com/WavefrontHQ/go-wavefront-management-api/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAlert() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlertCreate,
		Read:   resourceAlertRead,
		Update: resourceAlertUpdate,
		Delete: resourceAlertDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			nameKey: {
				Type:     schema.TypeString,
				Required: true,
			},
			alertTypeKey: {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          wavefront.AlertTypeClassic,
				DiffSuppressFunc: suppressCase,
			},
			targetKey: {
				Type:             schema.TypeString,
				Optional:         true,
				StateFunc:        trimSpaces,
				DiffSuppressFunc: suppressSpaces,
				ValidateFunc:     validateAlertTarget,
			},
			conditionKey: {
				Type:             schema.TypeString,
				Optional:         true,
				StateFunc:        trimSpaces,
				DiffSuppressFunc: suppressAlertConditionOnType,
			},
			conditionsKey: {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			thresholdTargetsKey: {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			additionalInformationKey: {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: suppressSpaces,
			},
			displayExpressionKey: {
				Type:             schema.TypeString,
				Optional:         true,
				StateFunc:        trimSpaces,
				DiffSuppressFunc: suppressSpaces,
			},
			minutesKey: {
				Type:     schema.TypeInt,
				Required: true,
			},
			resolveAfterMinutesKey: {
				Type:     schema.TypeInt,
				Optional: true,
			},
			notificationResendFrequencyMinutesKey: {
				Type:     schema.TypeInt,
				Optional: true,
			},
			severityKey: {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			tagsKey: {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			canViewKey: {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			canModifyKey: {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			processRateMinutesKey: {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			runbookLinksKey: {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			alertTriageDashboardsKey: {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: alertTriageDashboardSchema(),
				},
			},
		},
	}
}

func validateAlertTarget(val interface{}, _ string) (warnings []string, errors []error) {
	target := val.(string)
	if target == "" {
		return nil, nil
	}

	targets := strings.Split(target, ",")
	for _, t := range targets {
		if strings.HasPrefix(t, "pd:") || strings.HasPrefix(t, "target:") ||
			strings.Contains(t, "@") {
			continue
		}
		errors = append(errors,
			fmt.Errorf("valid alert targets must be prefixed with pd:, target:, or be a valid email address"))

		break
	}

	return warnings, errors
}

func suppressAlertConditionOnType(k, old, new string, d *schema.ResourceData) bool {
	alertType := strings.ToUpper(d.Get(alertTypeKey).(string))

	// after v2 alerts, `condition` has been force sync with `display_expression`
	// in multi-threshold alert
	// terraform does not support to infer attribute value from other attributes
	// thus we suppress the diff check for `condition` in multi-threshold alert.
	if alertType == wavefront.AlertTypeThreshold {
		return true
	}

	return suppressSpaces(k, old, new, d)
}

func resourceAlertCreate(d *schema.ResourceData, meta interface{}) error {
	alerts := meta.(*wavefrontClient).client.Alerts()

	tags := decodeTags(d)
	runbookLinks := decodeRunbookLinks(d.Get(runbookLinksKey).([]interface{}))
	alertTriageDashboards := decodeAlertTriageDashboards(d.Get(alertTriageDashboardsKey).([]interface{}))

	a := &wavefront.Alert{
		Name:                               d.Get(nameKey).(string),
		AdditionalInfo:                     trimSpaces(d.Get(additionalInformationKey)),
		DisplayExpression:                  trimSpaces(d.Get(displayExpressionKey)),
		Minutes:                            d.Get(minutesKey).(int),
		ResolveAfterMinutes:                d.Get(resolveAfterMinutesKey).(int),
		NotificationResendFrequencyMinutes: d.Get(notificationResendFrequencyMinutesKey).(int),
		Tags:                               tags,
		CheckingFrequencyInMinutes:         d.Get(processRateMinutesKey).(int),
		RunbookLinks:                       runbookLinks,
		AlertTriageDashboards:              alertTriageDashboards,
	}

	err := validateAlertConditions(a, d)
	if err != nil {
		return err
	}

	// Create the alert on Wavefront
	err = alerts.Create(a)
	if err != nil {
		return fmt.Errorf("error creating Alert %s. %s", d.Get(nameKey), err)
	}

	d.SetId(*a.ID)

	canView, canModify := decodeAccessControlList(d)
	if d.HasChanges(canViewKey, canModifyKey) {
		err = alerts.SetACL(*a.ID, canView, canModify)
		if err != nil {
			return fmt.Errorf("error setting ACL on Alert %s. %s", d.Get(nameKey), err)
		}
	}
	return nil
}

func resourceAlertRead(d *schema.ResourceData, meta interface{}) error {
	alerts := meta.(*wavefrontClient).client.Alerts()

	alertID := d.Id()
	tmpAlert := wavefront.Alert{ID: &alertID}
	err := alerts.Get(&tmpAlert)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("error finding Wavefront Alert %s. %s", d.Id(), err)
	}

	// Use the Wavefront ID as the Terraform ID
	d.SetId(*tmpAlert.ID)
	d.Set(nameKey, tmpAlert.Name)
	if tmpAlert.Target != "" && tmpAlert.AlertType == wavefront.AlertTypeClassic {
		d.Set(targetKey, tmpAlert.Target)
	}
	if tmpAlert.Severity != "" && tmpAlert.AlertType == wavefront.AlertTypeClassic {
		d.Set(severityKey, tmpAlert.Severity)
	}
	d.Set(conditionKey, trimSpaces(tmpAlert.Condition))
	d.Set(additionalInformationKey, trimSpaces(tmpAlert.AdditionalInfo))
	d.Set(displayExpressionKey, trimSpaces(tmpAlert.DisplayExpression))
	d.Set(minutesKey, tmpAlert.Minutes)
	d.Set(resolveAfterMinutesKey, tmpAlert.ResolveAfterMinutes)
	d.Set(notificationResendFrequencyMinutesKey, tmpAlert.NotificationResendFrequencyMinutes)
	d.Set(tagsKey, tmpAlert.Tags)
	d.Set(alertTypeKey, tmpAlert.AlertType)
	d.Set(conditionsKey, tmpAlert.Conditions)
	d.Set(thresholdTargetsKey, tmpAlert.Targets)
	d.Set(canViewKey, tmpAlert.ACL.CanView)
	d.Set(canModifyKey, tmpAlert.ACL.CanModify)
	d.Set(processRateMinutesKey, tmpAlert.CheckingFrequencyInMinutes)
	d.Set(runbookLinksKey, tmpAlert.RunbookLinks)
	d.Set(alertTriageDashboardsKey, parseAlertTriageDashboards(tmpAlert.AlertTriageDashboards))

	return nil
}

func resourceAlertUpdate(d *schema.ResourceData, meta interface{}) error {
	alerts := meta.(*wavefrontClient).client.Alerts()

	alertID := d.Id()
	tmpAlert := wavefront.Alert{ID: &alertID}
	err := alerts.Get(&tmpAlert)

	if err != nil {
		d.SetId("")
		return nil
	}

	tags := decodeTags(d)
	runbookLinks := decodeRunbookLinks(d.Get(runbookLinksKey).([]interface{}))
	alertTriageDashboards := decodeAlertTriageDashboards(d.Get(alertTriageDashboardsKey).([]interface{}))
	canView, canModify := decodeAccessControlList(d)

	a := tmpAlert
	a.Name = d.Get(nameKey).(string)
	a.AdditionalInfo = trimSpaces(d.Get(additionalInformationKey).(string))
	a.DisplayExpression = trimSpaces(d.Get(displayExpressionKey).(string))
	a.Minutes = d.Get(minutesKey).(int)
	a.ResolveAfterMinutes = d.Get(resolveAfterMinutesKey).(int)
	a.NotificationResendFrequencyMinutes = d.Get(notificationResendFrequencyMinutesKey).(int)
	a.Tags = tags
	a.RunbookLinks = runbookLinks
	a.AlertTriageDashboards = alertTriageDashboards
	a.CheckingFrequencyInMinutes = d.Get(processRateMinutesKey).(int)

	err = validateAlertConditions(&a, d)
	if err != nil {
		return err
	}

	// Update the alert on Wavefront
	err = alerts.Update(&a)
	if err != nil {
		return fmt.Errorf("error Updating Alert %s. %s", d.Get(nameKey), err)
	}

	// Update the ACLs on the alert in Wavefront
	if d.HasChanges(canViewKey, canModifyKey) {
		err = alerts.SetACL(*a.ID, canView, canModify)
		if err != nil {
			return fmt.Errorf("error updating ACLs on Alert %s. %s", d.Get(nameKey), err)
		}
	}

	return nil
}

func resourceAlertDelete(d *schema.ResourceData, meta interface{}) error {
	alerts := meta.(*wavefrontClient).client.Alerts()

	alertID := d.Id()
	tmpAlert := wavefront.Alert{ID: &alertID}
	err := alerts.Get(&tmpAlert)
	if err != nil {
		return fmt.Errorf("error finding Wavefront Alert %s. %s", d.Id(), err)
	}
	a := tmpAlert

	// Delete the Alert
	err = alerts.Delete(&a, true)
	if err != nil {
		return fmt.Errorf("failed to delete Alert %s. %s", d.Id(), err)
	}
	d.SetId("")
	return nil
}

func validateAlertConditions(a *wavefront.Alert, d *schema.ResourceData) error {
	alertType := strings.ToUpper(d.Get(alertTypeKey).(string))
	if alertType == wavefront.AlertTypeThreshold {
		a.AlertType = wavefront.AlertTypeThreshold

		// v2 alerts now force sync `condition` the same as `display_expression`
		// for multi-threshold alerts
		a.Condition = d.Get(displayExpressionKey).(string)

		if conditions, ok := d.GetOk(conditionsKey); ok {
			a.Conditions = trimSpacesMap(conditions.(map[string]interface{}))
			err := validateThresholdLevels(a.Conditions)
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("conditions must be supplied for threshold alerts")
		}

		if targets, ok := d.GetOk(thresholdTargetsKey); ok {
			a.Targets = trimSpacesMap(targets.(map[string]interface{}))
			return validateThresholdLevels(a.Targets)
		}

	} else if alertType == wavefront.AlertTypeClassic {
		a.AlertType = wavefront.AlertTypeClassic

		if d.Get(conditionKey) == "" {
			return fmt.Errorf("condition must be supplied for classic alerts")
		}
		a.Condition = trimSpaces(d.Get(conditionKey).(string))

		if d.Get(severityKey) == "" {
			return fmt.Errorf("severity must be supplied for classic alerts")
		}
		a.Severity = d.Get(severityKey).(string)
		a.Target = d.Get(targetKey).(string)
	} else {
		return fmt.Errorf("alert_type must be CLASSIC or THRESHOLD")
	}

	return nil
}

func validateThresholdLevels(m map[string]string) error {
	for key := range m {
		ok := false
		for _, level := range []string{"severe", "warn", "info", "smoke"} {
			if key == level {
				ok = true
				break
			}
		}
		if !ok {
			return fmt.Errorf("invalid severity: %s", key)
		}
	}
	return nil
}

func decodeRunbookLinks(rawRunbookLinks []interface{}) (links []string) {
	for _, link := range rawRunbookLinks {
		links = append(links, link.(string))
	}
	return links
}

func decodeAlertTriageDashboards(rawDashboards []interface{}) (alertTriageDashboards []wavefront.AlertTriageDashboard) {
	for _, rawDashboard := range rawDashboards {
		dashboardData := rawDashboard.(map[string]interface{})

		alertTriageDashboard := wavefront.AlertTriageDashboard{
			DashboardId: dashboardData[dashboardIDKey].(string),
			Description: dashboardData[descriptionKey].(string),
			Parameters:  decodeAlertTriageDashboardParameters(dashboardData[parametersKey].([]interface{})),
		}

		alertTriageDashboards = append(alertTriageDashboards, alertTriageDashboard)
	}

	return alertTriageDashboards
}

func decodeAlertTriageDashboardParameters(parameterData []interface{}) map[string]map[string]string {
	parameters := make(map[string]map[string]string)
	if len(parameterData) > 0 {
		for _, parameterBlocks := range parameterData {
			for parameterBlockType, parameterBlockValue := range parameterBlocks.(map[string]interface{}) {
				parameters[parameterBlockType] = make(map[string]string)
				for parameterKey, parameterValue := range parameterBlockValue.(map[string]interface{}) {
					parameters[parameterBlockType][parameterKey] = parameterValue.(string)
				}
			}
		}
	}
	return parameters
}
