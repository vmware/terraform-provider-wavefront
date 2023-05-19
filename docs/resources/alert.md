---
layout: "wavefront"
page_title: "Wavefront: Alert"
description: |-
  Provides a Wavefront Alert resource.  This allows alerts to be created, updated, and deleted.
---

# Resource : wavefront_alert

Provides a Wavefront Alert resource.  This allows alerts to be created, updated, and deleted.

## Example usage

```hcl
resource "wavefront_alert" "foobar" {
  name = "Test Alert"
  target = "test@example.com,target:alert-target-id"
  condition = "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total ) > 80"
  display_expression = "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total )"
  minutes = 5
  resolve_after_minutes = 5
  severity = "WARN"
  tags = [
    "terraform",
    "test"
  ]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the alert as it is displayed in Wavefront.
* `tags` - (Required) A set of tags to assign to this resource.
* `alert_type` - (Optional) The type of alert in Wavefront.  Either `CLASSIC` (default) 
or `THRESHOLD`.
* `minutes` - (Required) The number of consecutive minutes that a series matching the condition query must 
evaluate to "true" (non-zero value) before the alert fires.
* `target` - (Optional, `CLASSIC` alerts only) A comma-separated list of the email address or integration endpoint 
(such as PagerDuty or webhook) to notify when the alert status changes. Multiple target types can be in the list.
Alert target format: ({email}|pd:{pd_key}|target:{alert-target-id}).
* `condition` - (Optional) A Wavefront query that is evaluated at regular intervals (default is 1 minute).
The alert fires and notifications are triggered when a data series matching this query evaluates 
to a non-zero value for a set number of consecutive minutes. 
* `conditions` - (Optional, `THRESHOLD` alerts only) a string->string map of `severity` to `condition` 
for which this alert will trigger.
* `threshold_targets` - (Optional, `THRESHOLD` alerts only) A string to string map of Targets for severity.
* `additional_information` - (Optional) User-supplied additional explanatory information for this alert.
Useful for linking runbooks, migrations, etc.
* `display_expression` - (Optional) A second query whose results are displayed in the alert user
interface instead of the condition query.  This field is often used to display a version
of the condition query with Boolean operators removed so that numerical values are plotted.
* `resolve_after_minutes` - (Optional) The number of consecutive minutes that a firing series matching the condition
query must evaluate to "false" (zero value) before the alert resolves.  When unset, this defaults to
the same value as `minutes`.
* `notification_resend_frequency_minutes` - (Optional) How often to re-trigger a continually failing alert. 
If absent or <= 0, no re-triggering occurs.  
* `severity` - (Optional, `CLASSIC` alerts only) - Severity of the alert, valid values are `INFO`, `SMOKE`, `WARN`, `SEVERE`.
* `can_view` - (Optional) A list of valid users or groups that can view this resource on a tenant. Default is Empty list.
* `can_modify` - (Optional) A list of valid users or groups that can modify this resource on a tenant.
* `process_rate_minutes` - (Optional) The specified query is executed every `process_rate_minutes` minutes. Default value is 5 minutes.


### Example
```hcl
resource "wavefront_alert" "test_alert" {
  name = "Terraform Test Alert"
  target = "test@example.com,target:alert-target-id"
  condition = "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total ) > 80"
  additional_information = "This is a Terraform Test Alert"
  display_expression = "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total )"
  minutes = 5
  resolve_after_minutes = 5
  severity = "WARN"
  tags = [
    "terraform",
    "test"
  ]

  can_view = [
    "test@example.com",
  ]
  process_rate_minutes = 4
}

resource "wavefront_alert_target" "test_target" {
  name = "Terraform Test Target"
  description = "Test target"
  method = "EMAIL"
  recipient = "test@example.com"
  email_subject = "This is a test"
  is_html_content = true
  template = "{}"
  triggers = [
    "ALERT_OPENED",
    "ALERT_RESOLVED"
  ]
}


resource "wavefront_alert" "test_threshold_alert" {
  name = "Terraform Test Alert"
  alert_type = "THRESHOLD"
  additional_information = "This is a Terraform Test Alert"
  display_expression = "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total )"
  minutes = 5
  resolve_after_minutes = 5

  conditions = {
    "severe" = "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total ) > 80"
    "warn" = "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total ) > 60"
    "info" = "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total ) > 50"
  }

  threshold_targets = {
	"severe" = "target:${wavefront_alert_target.test_target.id}"
  }
  
  tags = [
    "terraform"
  ]
}
```

## Import

Alerts can be imported using the `id`, e.g.

```
$ terraform import wavefront_alert_target.alert_target 1479868728473
```