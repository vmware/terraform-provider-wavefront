---
layout: "wavefront"
page_title: "Wavefront: Maintenance Windows"
description: |-
    Get the information about all Wavefront maintenance windows.
---

# Data Source: wavefront_maintenance_window_all

Use this data source to get information about all Wavefront maintenance windows.

## Argument Reference
* `limit` - (Optional) Limit is the maximum number of results to be returned. Defaults to 100.
* `offset` - (Optional) Offset is the offset from the first result to be returned. Defaults to 0.

## Example Usage

```hcl
# Get the information about all external links.
data "wavefront_maintenance_window_all" "example" {
  limit = 10
  offset = 0
}
```

## Attribute Reference

* `maintenance_windows` - List of all external links in Wavefront. For each external link you will see a list of attributes.
    * `id` -  The ID of the maintenance window.
    * `event_name` - The event name of the maintenance window.
    * `reason` - The reason for the maintenance window.
    * `title` - The title of the maintenance window.
    * `customer_id` - The ID of the customer in Wavefront.
    * `running_state` - The running state of the maintenance window.
    * `start_time_in_seconds` - The start time in seconds after 1 Jan 1970 GMT.
    * `end_time_in_seconds` - The end time in seconds after 1 Jan 1970 GMT.
    * `relevant_customer_tags` - The list of alert tags whose matching alerts will be put into maintenance because
      of this maintenance window. At least one of `relevant_customer_tags`, `relevant_host_tags`, or `relevant_host_names`
      is required.
    * `relevant_host_tags` - The list of source/host tags whose matching sources/hosts will be put into maintenance
      because of this maintenance window. At least one of `relevant_customer_tags`, `relevant_host_tags`, or
      `relevant_host_names` is required.
    * `relevant_host_names` - The list of source/host names that will be put into maintenance because of this
      maintenance window. At least one of `relevant_customer_tags`, `relevant_host_tags`, or `relevant_host_names`
      is required.
    * `relevant_host_tags_anded` - Whether to AND source/host tags listed in `relevant_host_tags`.
      If `true`, a source/host must contain all tags in order for the maintenance window to apply. If `false`,
      the tags are OR'ed, and a source/host must contain one of the tags. Default: `false`.
    * `host_tag_group_host_names_group_anded` - If `true`, a source/host must be in `relevant_host_names`
      and have tags matching the specification formed by `relevant_host_tags` and `relevant_host_tags_anded` in
      order for this maintenance window to apply. If `false`, a source/host must either be in `relevant_host_names`
      or match `relevant_host_tags` and `relevant_host_tags_anded`. Default: `false`.
    * `created_epoch_millis` - The timestamp in epoch milliseconds indicating when the external link is created.
    * `updated_epoch_millis` - The timestamp in epoch milliseconds indicating when the external link is updated.
    * `creator_id` - The ID of the user who created the external link.
    * `updater_id` - The ID of the user who updated the external link.
 