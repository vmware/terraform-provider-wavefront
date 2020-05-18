---
layout: "wavefront"
page_title: "Wavefront: Alert Target"
description: |-
  Provides a wavefront Alert Target resource. This allows alert targets to created, updated, and deleted.
---

# Resource : wavefront_alert_target

Provides a wavefront Alert Target resource. This allows alert targets to created, updated, and deleted.

## Example usage

```hcl
resource "wavefront_alert_target" "test_target" {
  name = "Terraform Test Target"
  description = "Test target"
  method = "WEBHOOK"
  recipient = "https://hooks.slack.com/services/test/me"
  content_type = "application/json"
  custom_headers = {
    "Testing" = "true"
  }
  template = "{}"
  triggers = [
    "ALERT_OPENED",
    "ALERT_RESOLVED"
  ]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the alert target as it is displayed in wavefront
* `description` - (Required) Description describing this alert target.
* `triggers` - (Required) A list of occurrences on which this webhook will be fired. Valid values are `ALERT_OPENED`,
`ALERT_UPDATED`, `ALERT_RESOLVED`, `ALERT_MAINTENANCE`, `ALERT_SNOOZED`, `ALERT_INVALID`, `ALERT_NO_LONGER_INVALID`, 
`ALERT_RETRIGGERED`, `ALERT_NO_DATA`, `ALERT_NO_DATA_RESOLVED`, `ALERT_NO_DATA_MAINTENANCE`, `ALERT_SEVERITY_UPDATE`.
* `template` - (Required) A mustache template that will form the body of the POST request, email and summary of the PagerDuty.
* `recipient` - (Required) The end point for the notification Target.  `EMAIL`: email address. `PAGERDUTY`: PagerDuty 
routing key. `WEBHOOK`: URL endpoint. 
* `method` - (Optional) The notification method used for notification target. One of `WEBHOOK`, `EMAIL`, `PAGERDUTY`.
* `route` - (Optional) List of routing targets that this alert target will notify. See [Route](#route)
* `email_subject` - (Optional) The subject title of an email notification target.
* `is_html_content` - (Optional) Determine whether the email alert content is sent as HTML or text.
* `content_type` - (Optional) The value of the `Content-Type` header of the webhook.
* `custom_headers` - (Optional) A `string->string` map specifying the custome HTTP header key/value pairs that will be 
sent in the requests with a method of `WEBHOOK`.

### Route

The `route` mapping supports the following:

* `method` - (Required)  The notification method used for notification target. One of `WEBHOOK`, `EMAIL`, `PAGERDUTY`.
* `target` - (Required) The endpoint for the alert route. `EMAIL`: email address. `PAGERDUTY`: PagerDuty routing 
key. `WEBHOOK`: URL endpoint. 
* `filter` - (Required) String that filters the route. Space delimited.  Currently only allows a single key value pair.
(e.g. `env prod`)

### Example

```hcl
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

resource "wavefront_alert_target" "test_target" {
  name         = "Terraform Test Target"
  description  = "Test target"
  method       = "WEBHOOK"
  recipient    = "https://hooks.slack.com/services/test/me"
  content_type = "application/json"
  custom_headers = {
    "Testing" = "true"
  }
  template = "{}"
  triggers = [
    "ALERT_OPENED",
    "ALERT_RESOLVED",
  ]
  route {
    method = "WEBHOOK"
    target = "https://hooks.slack.com/services/test/me/prod"
    filter = {
      key   = "env"
      value = "prod"
    }
  }
  route {
    method = "WEBHOOK"
    target = "https://hooks.slack.com/services/test/me/dev"
    filter = {
      key   = "env"
      value = "dev"
    }
  }
}
```

## Import

Alert Targets can be imported using the `id`, e.g.

```
$ terraform import wavefront_alert_target.alert_target abcdEFGhijKLMNO
```