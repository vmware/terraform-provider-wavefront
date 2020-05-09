---
layout: "wavefront"
page_title: "Wavefront: Cloud Integration NewRelic"
description: |-
  Provides a Wavefront Cloud Integration for NewRelic. This allows NewRelic cloud integrations to be created,
  updated, and deleted.
---

# Resource : wavefront_cloud_integration_newrelic

Provides a Wavefront Cloud Integration for NewRelic. This allows NewRelic cloud integrations to be created,
updated, and deleted.

## Example usage

```hcl
resource "wavefront_cloud_integration_newrelic" "newrelic" {
  name              = "Test Integration"
  api_key           = "example-api-key"
}
```

## Argument Reference

The following arguments are supported:

* `service` - (Required) A value denoting which cloud service this service integrates with
* `name` - (Required) The human-readable name of this integration
* `additional_tags` - (Optional) A list of point tag key-values to add to every point ingested using this integration
* `force_save` - (Optional) Forces this resource to save, even if errors are present.
* `api_key` - (Required) NewRelic REST api key
* `app_filter_regex` - (Optional) A regular expression that an application name must match (case-insensitively) iun order to collect metrics
* `host_filter_regex` - (Optional) A regular expression that a host name must match (case-insensitively) in order to collect metrics 
* `metric_filter` - (Optional) See [Metric Filter](#metric-filter)

### Metric Filter

The `metric_filter` mapping supports the following:

* `app_name` - (Required) The name of a NewRelic App
* `metric_filter_regex` - (Required) A regular expression that a metric name must match (case-insensitively) in order to be ingested

### Example

```hcl
resource "wavefront_cloud_integration_newrelic" "newrelic" {
  name              = "Test Integration"
  force_save        = true
  additional_tags = {
    "tag1" = "value1"
    "tag2" = "value2"
  }
  api_key           = "example-api-key"
  app_filter_regex  = "^someApp.*$"
  host_filter_regex = "^prod-env.*$"
  metric_filter {
    app_name            = "app1"
    metric_filter_regex = "^cpu.*?"
  }
  metric_filter {
    app_name            = "app2"
    metric_filter_regex = "^mem.*?"
  }
}
```

## Import

NewRelic Integrations can be imported using the `id`, e.g.

```
$ terraform import wavefront_cloud_integration_newrelic.newrelic a411c16b-3cf7-4f03-bf11-8ca05aab898d
```