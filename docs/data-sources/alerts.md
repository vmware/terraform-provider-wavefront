---
layout: "wavefront"
page_title: "Wavefront: Alerts"
description: |-
    Get the information about all Wavefront alerts.
---

# Data Source: wavefront_alerts

Use this data source to get information about all Wavefront alerts.

## Example Usage

```hcl
# Get the information about all alerts.
data "wavefront_alerts" "example" {
}
```

## Attribute Reference

* `alerts` - List of all alerts in Wavefront. For each alert you will see a list of attributes.
  * `name` - The name of the alert as it is displayed in Wavefront.
  * `id` - The ID of the alert in Wavefront.
  * `alert_type`- The type of alert in Wavefront.
  * `additional_information` - User-supplied additional explanatory information about this alert.
  * `target` - An email address or integration endpoint (such as PagerDuty or webhook) to notify when the alert status changes.
  * `targets` - A comma-separated list of the email addresses or integration endpoints (such as PagerDuty or webhook) to notify when the alert status changes. Multiple target types can be in the list.
  * `condition` - A Wavefront query that is evaluated at regular intervals (default is 1 minute). The alert fires and notifications are triggered when a data series matching this query evaluates to a non-zero value for a set number of consecutive minutes.
  * `conditions` - A map of severity to condition for which this alert will trigger.
  * `display_expression` - A second query the results of which are displayed in the alert user interface instead of the condition query.
  * `minutes` - The number of consecutive minutes that a series matching the condition query must evaluate to "true" (non-zero value) before the alert fires.
  * `resolve_after_minutes` - The number of consecutive minutes that a firing series matching the condition query must evaluate to "false" (zero value) before the alert resolves.
  * `notification_resend_frequency_minutes` - How often to re-trigger a continually failing alert.
  * `severity` - The severity of the alert.
  * `status` - The status of the alert.
  * `tags` - A set of tags assigned to the alert.
  * `can_view` - A list of users or groups that can view the alert.
  * `can_modify` - A list of users or groups that can modify the alert.
  * `process_rate_minutes` - The specified query is executed every `process_rate_minutes` minutes.
  * `evaluate_realtime_data` - A Boolean flag to enable real-time evaluation.
  * `include_obsolete_metrics` - A Boolean flag indicating whether to include obsolete metrics or not.
  * `failing_host_label_pairs` - A list of failing host label pairs.
  * `in_maintenance_host_label_pairs` - A list of in maintenance host label pairs.
