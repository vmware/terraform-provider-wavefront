---
layout: "wavefront"
page_title: "Wavefront: External Link"
description: |-
    Get the information about a specific Wavefront external link.
---

# Data Source: wavefront_external_link

Use this data source to get information about a Wavefront external link by its ID.

## Argument Reference
* `id` - (Required) The ID of the external link.

## Example Usage

```hcl
# Get the information about a specific external links.
data "wavefront_external_link" "example" {
  id = "sample-external-link-id"
}
```

## Attribute Reference

* `name` - The name of the external link.
* `id` -  The ID of the external link.
* `description` - Human-readable description for this link.
* `template` - The mustache template for this link. The template must expand to a full URL, including scheme, origin, etc.
* `metric_filter_regex` - Controls whether a link is displayed in the context menu of a highlighted series. If present, the metric name of the highlighted series must match this regular expression in order for the link to be displayed.
* `source_filter_regex` - Controls whether a link is displayed in the context menu of a highlighted series. If present, the source name of the highlighted series must match this regular expression in order for the link to be displayed.
* `point_tag_filter_regexes` - (Optional) Controls whether a link is displayed in the context menu of a highlighted
      series. This is a map from string to regular expression. The highlighted series must contain point tags whose
      keys are present in the keys of this map and whose values match the regular expressions associated with those
      keys in order for the link to be displayed.
* `is_log_integration` - Whether this is a "Log Integration" subType of external link.
* `created_epoch_millis` - The timestamp in epoch milliseconds indicating when the external link is created.
* `updated_epoch_millis` - The timestamp in epoch milliseconds indicating when the external link is updated.
* `creator_id` - The ID of the user who created the external link.
* `updater_id` - The ID of the user who updated the external link.
