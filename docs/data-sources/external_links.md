---
layout: "wavefront"
page_title: "Wavefront: External Links"
description: |-
    Get the information about all Wavefront external links.
---

# Data Source: wavefront_external_links

Use this data source to get information about all Wavefront external links.

## Argument Reference
* `limit` - (Optional) Limit is the maximum number of results to be returned. Defaults to 100.
* `offset` - (Optional) Offset is the offset from the first result to be returned. Defaults to 0.

## Example Usage

```hcl
# Get the information about all external links.
data "wavefront_external_links" "example" {
  limit = 10
  offset = 0
}
```

## Attribute Reference

* `external_links` - List of all external links in Wavefront. For each external link you will see a list of attributes.
  * `name` - The name of the external link.
  * `id` -  The ID of the external link.
  * `description` - Human-readable description of the link.
  * `template` - The mustache template for the link. The template must expand to a full URL, including scheme, origin, etc.
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
