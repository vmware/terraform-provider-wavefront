---
layout: "wavefront"
page_title: "Wavefront: External Link"
description: |-
  Provides a Wavefront External Link Resource. This allows external links to be created, updated, and deleted.
---

# Resource : wavefront_external_link

Provides a Wavefront External Link Resource. This allows external links to be created, updated, and deleted.

## Example usage

```hcl
resource "wavefront_external_link" "basic" {
  name = "External Link"
  description = "An external link description"
  template = "https://example.com/source={{{source}}}&startTime={{startEpochMillis}}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the external link.
* `description` - (Required) Human-readable description for this link.
* `template` - (Required) The mustache template for this link. The template must expand to a full URL, including scheme, origin, etc.
* `metric_filter_regex` - (Optional) Controls whether a link is displayed in the context menu of a highlighted series. If present, the metric name of the highlighted series must match this regular expression in order for the link to be displayed.
* `source_filter_regex` - (Optional) Controls whether a link is displayed in the context menu of a highlighted series. If present, the source name of the highlighted series must match this regular expression in order for the link to be displayed.
* `point_tag_filter_regexes` - (Optional) Controls whether a link is displayed in the context menu of a highlighted
  series. This is a map from string to regular expression. The highlighted series must contain point tags whose
  keys are present in the keys of this map and whose values match the regular expressions associated with those
  keys in order for the link to be displayed.
* `is_log_integration` - Whether this is a "Log Integration" subType of external link.

### Example

```hcl

resource "wavefront_external_link" "basic" {
  name = "External Link"
  description = "An external link description"
  template = "https://example.com/source={{{source}}}&startTime={{startEpochMillis}}"
  metric_filter_regex = "^metric.*$"
  source_filter_regex = "^prod.*$"
  point_tag_filter_regexes = {
    service = "^query$"
  }
  is_log_integration = true
}
```

## Import

Maintenance windows can be imported by using the `id`, e.g.:

```
$ terraform import wavefront_external_link.basic fVj6fz6zYC4aBkID
```
