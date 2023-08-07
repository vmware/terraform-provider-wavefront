---
layout: "wavefront"
page_title: "Wavefront: Maintenance Window"
description: |-
  Provides a Wavefront Maintenance Window Resource. This allows maintenance windows to be created, updated, and deleted.
---

# Resource : wavefront_maintenance_window

Provides a Wavefront Maintenance Window Resource. This allows maintenance windows to be created, updated, and deleted.

## Example usage

```hcl
resource "wavefront_maintenance_window" "basic" {
  reason  = "Routine maintenance for 2020"
  title = "Routine maintenance"
  start_time_in_seconds = 1600123456
  end_time_in_seconds = 1601123456
  relevant_host_names = ["my_hostname", "my_other_hostname"]
}
```

## Argument Reference

The following arguments are supported:

* `reason` - (Required) The reason for the maintenance window.
* `title` - (Required) The title of the maintenance window.
* `start_time_in_seconds` - (Required) start time in seconds after 1 Jan 1970 GMT.
* `end_time_in_seconds` - (Required) end time in seconds after 1 Jan 1970 GMT.
* `relevant_customer_tags` - List of alert tags whose matching alerts will be put into maintenance because
  of this maintenance window. At least one of `relevant_customer_tags`, `relevant_host_tags`, or `relevant_host_names`
  is required.
* `relevant_host_tags` - List of source/host tags whose matching sources/hosts will be put into maintenance
  because of this maintenance window. At least one of `relevant_customer_tags`, `relevant_host_tags`, or
  `relevant_host_names` is required.
* `relevant_host_names` - List of source/host names that will be put into maintenance because of this
  maintenance window. At least one of `relevant_customer_tags`, `relevant_host_tags`, or `relevant_host_names`
  is required.
* `relevant_host_tags_anded` - (Optional) Whether to AND source/host tags listed in `relevant_host_tags`.
  If `true`, a source/host must contain all tags in order for the maintenance window to apply. If `false`,
  the tags are OR'ed, and a source/host must contain one of the tags. Default: `false`.
* `host_tag_group_host_names_group_anded` - (Optional) If `true`, a source/host must be in `relevant_host_names`
  and have tags matching the specification formed by `relevant_host_tags` and `relevant_host_tags_anded` in
  order for this maintenance window to apply. If `false`, a source/host must either be in `relevant_host_names`
  or match `relevant_host_tags` and `relevant_host_tags_anded`. Default: `false`.

### Example

```hcl

resource "wavefront_maintenance_window" "basic" {
  reason  = "Routine maintenance for 2020"
  title = "Routine maintenance"
  start_time_in_seconds = 1600123456
  end_time_in_seconds = 1601123456
  relevant_host_tags = ["my_host_tag1", "my_host_tag2"]
  relevant_host_tags_anded = true
}
```

## Import

Maintenance windows can be imported using the `id`, e.g.

```
$ terraform import wavefront_maintenance_window.basic 1600383357095 
```
