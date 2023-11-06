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
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"alert_type": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          wavefront.AlertTypeClassic,
				DiffSuppressFunc: suppressCase,
			},
			"target": {
				Type:             schema.TypeString,
				Optional:         true,
				StateFunc:        trimSpaces,
				DiffSuppressFunc: suppressSpaces,
				ValidateFunc:     validateAlertTarget,
			},
			"condition": {
				Type:             schema.TypeString,
				Optional:         true,
				StateFunc:        trimSpaces,
				DiffSuppressFunc: suppressAlertConditionOnType,
			},
			"conditions": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"threshold_targets": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"additional_information": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: suppressSpaces,
			},
			"display_expression": {
				Type:             schema.TypeString,
				Optional:         true,
				StateFunc:        trimSpaces,
				DiffSuppressFunc: suppressSpaces,
			},
			"minutes": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"resolve_after_minutes": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"notification_resend_frequency_minutes": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"severity": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tags": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"can_view": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"can_modify": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			"process_rate_minutes": {
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
	alertType := strings.ToUpper(d.Get("alert_type").(string))

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
	runbookLinks := decodeRunbookLinks(d)
	alertTriageDashboards := decodeAlertTriageDashboards(d)

	a := &wavefront.Alert{
		Name:                               d.Get("name").(string),
		AdditionalInfo:                     trimSpaces(d.Get("additional_information")),
		DisplayExpression:                  trimSpaces(d.Get("display_expression")),
		Minutes:                            d.Get("minutes").(int),
		ResolveAfterMinutes:                d.Get("resolve_after_minutes").(int),
		NotificationResendFrequencyMinutes: d.Get("notification_resend_frequency_minutes").(int),
		Tags:                               tags,
		CheckingFrequencyInMinutes:         d.Get("process_rate_minutes").(int),
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
		return fmt.Errorf("error creating Alert %s. %s", d.Get("name"), err)
	}

	d.SetId(*a.ID)

	canView, canModify := decodeAccessControlList(d)
	if d.HasChanges("can_view", "can_modify") {
		err = alerts.SetACL(*a.ID, canView, canModify)
		if err != nil {
			return fmt.Errorf("error setting ACL on Alert %s. %s", d.Get("name"), err)
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
	d.Set("name", tmpAlert.Name)
	if tmpAlert.Target != "" && tmpAlert.AlertType == wavefront.AlertTypeClassic {
		d.Set("target", tmpAlert.Target)
	}
	if tmpAlert.Severity != "" && tmpAlert.AlertType == wavefront.AlertTypeClassic {
		d.Set("severity", tmpAlert.Severity)
	}
	d.Set("condition", trimSpaces(tmpAlert.Condition))
	d.Set("additional_information", trimSpaces(tmpAlert.AdditionalInfo))
	d.Set("display_expression", trimSpaces(tmpAlert.DisplayExpression))
	d.Set("minutes", tmpAlert.Minutes)
	d.Set("resolve_after_minutes", tmpAlert.ResolveAfterMinutes)
	d.Set("notification_resend_frequency_minutes", tmpAlert.NotificationResendFrequencyMinutes)
	d.Set("tags", tmpAlert.Tags)
	d.Set("alert_type", tmpAlert.AlertType)
	d.Set("conditions", tmpAlert.Conditions)
	d.Set("threshold_targets", tmpAlert.Targets)
	d.Set("can_view", tmpAlert.ACL.CanView)
	d.Set("can_modify", tmpAlert.ACL.CanModify)
	d.Set("process_rate_minutes", tmpAlert.CheckingFrequencyInMinutes)
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
	runbookLinks := decodeRunbookLinks(d)
	alertTriageDashboards := decodeAlertTriageDashboards(d)
	canView, canModify := decodeAccessControlList(d)

	a := tmpAlert
	a.Name = d.Get("name").(string)
	a.AdditionalInfo = trimSpaces(d.Get("additional_information").(string))
	a.DisplayExpression = trimSpaces(d.Get("display_expression").(string))
	a.Minutes = d.Get("minutes").(int)
	a.ResolveAfterMinutes = d.Get("resolve_after_minutes").(int)
	a.NotificationResendFrequencyMinutes = d.Get("notification_resend_frequency_minutes").(int)
	a.Tags = tags
	a.RunbookLinks = runbookLinks
	a.AlertTriageDashboards = alertTriageDashboards
	a.CheckingFrequencyInMinutes = d.Get("process_rate_minutes").(int)

	err = validateAlertConditions(&a, d)
	if err != nil {
		return err
	}

	// Update the alert on Wavefront
	err = alerts.Update(&a)
	if err != nil {
		return fmt.Errorf("error Updating Alert %s. %s", d.Get("name"), err)
	}

	// Update the ACLs on the alert in Wavefront
	if d.HasChanges("can_view", "can_modify") {
		err = alerts.SetACL(*a.ID, canView, canModify)
		if err != nil {
			return fmt.Errorf("error updating ACLs on Alert %s. %s", d.Get("name"), err)
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
	alertType := strings.ToUpper(d.Get("alert_type").(string))
	if alertType == wavefront.AlertTypeThreshold {
		a.AlertType = wavefront.AlertTypeThreshold

		// v2 alerts now force sync `condition` the same as `display_expression`
		// for multi-threshold alerts
		a.Condition = d.Get("display_expression").(string)

		if conditions, ok := d.GetOk("conditions"); ok {
			a.Conditions = trimSpacesMap(conditions.(map[string]interface{}))
			err := validateThresholdLevels(a.Conditions)
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("conditions must be supplied for threshold alerts")
		}

		if targets, ok := d.GetOk("threshold_targets"); ok {
			a.Targets = trimSpacesMap(targets.(map[string]interface{}))
			return validateThresholdLevels(a.Targets)
		}

	} else if alertType == wavefront.AlertTypeClassic {
		a.AlertType = wavefront.AlertTypeClassic

		if d.Get("condition") == "" {
			return fmt.Errorf("condition must be supplied for classic alerts")
		}
		a.Condition = trimSpaces(d.Get("condition").(string))

		if d.Get("severity") == "" {
			return fmt.Errorf("severity must be supplied for classic alerts")
		}
		a.Severity = d.Get("severity").(string)
		a.Target = d.Get("target").(string)
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

// Decodes the runbook links from the state and returns a []string of links
func decodeRunbookLinks(d *schema.ResourceData) (links []string) {
	for _, link := range d.Get(runbookLinksKey).([]interface{}) {
		links = append(links, link.(string))
	}
	return links
}

func decodeAlertTriageDashboards(d *schema.ResourceData) []wavefront.AlertTriageDashboard {
	alertTriageDashboards := []wavefront.AlertTriageDashboard{}

	if dashboards, ok := d.Get(alertTriageDashboardsKey).([]interface{}); ok {
		for _, dashboard := range dashboards {
			dashboardData := dashboard.(map[string]interface{})
			alertTriageDashboard := wavefront.AlertTriageDashboard{
				DashboardId: dashboardData[dashboardIDKey].(string),
				Description: dashboardData[descriptionKey].(string),
				Parameters:  make(map[string]map[string]string),
			}

			if parameters, ok := dashboardData[parametersKey].([]interface{}); ok && len(parameters) > 0 {
				// Assuming there should be only one parameters block
				parametersData := parameters[0].(map[string]interface{})

				if constants, ok := parametersData[constantsKey].(map[string]interface{}); ok {
					alertTriageDashboard.Parameters[constantsKey] = make(map[string]string)
					for key, value := range constants {
						alertTriageDashboard.Parameters[constantsKey][key] = value.(string)
					}
				}
			}

			alertTriageDashboards = append(alertTriageDashboards, alertTriageDashboard)
		}
	}

	return alertTriageDashboards
}
