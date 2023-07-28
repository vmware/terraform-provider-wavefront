---
layout: "wavefront"
page_title: "Wavefront: Derived Metric"
description: |-
  Provides a Wavefront Derived Metric Resource. This allows derived metrics to be created,
  updated, and deleted.
---

# Resource : wavefront_derived_metric

Provides a Wavefront Derived Metric Resource. This allows derived metrics to be created,
updated, and deleted.
  

## Example usage

```hcl
resource "wavefront_derived_metric" "derived" {
  name    = "dummy derived metric"
  minutes = 5
  query   = "aliasMetric(5, \"some.metric\")"
}
```
## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the Derived Metric in Wavefront.
* `query` - (Required) A Wavefront query that is evaluated at regular intervals (default is 1 minute).
* `minutes` - (Required) How frequently the query generating the derived metric is run.
* `additional_information` - (Optional) User-supplied additional explanatory information for the derived metric.
* `tags` - (Optional) A set of tags to assign to this resource.

### Example

```hcl
resource "wavefront_derived_metric" "derived" {
  name                   = "dummy derived metric"
  minutes                = 10
  query                  = "aliasMetric(mavg(5m, ts(cpu.idle)), \"cpu.5m-avg\")"
  additional_information = "this is a dummy derived metric"
  tags = [
    "cpu.5m.avg",
    "derived"
  ]
}
```

## Import

Derived Metrics can be imported by using the `id`, e.g.:

```
$ terraform import wavefront_derived_metric.derived_metric 1577102900578
```