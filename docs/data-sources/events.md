---
layout: "wavefront"
page_title: "Wavefront: Events"
description: |-
    Get the information about all Wavefront events.
---

# Data Source: wavefront_events

Use this data source to get information about all Wavefront events.


## Argument Reference
* `latest_start_time_epoch_millis` - (Required) Latest start time in epoch milliseconds.
* `earliest_start_time_epoch_millis` - Earliest start time in epoch milliseconds.
* `limit` - (Optional) Limit is the maximum number of results to be returned. Defaults to 100.
* `offset` - (Optional) Offset is the offset from the first result to be returned. Defaults to 0.


## Example Usage

```hcl
# Get the information about all events
data "wavefront_events" "example" {
}
```

## Attribute Reference

* `events` - List of all events in Wavefront. For each event you will see a list of attributes.
  * `name` - The name of the event in Wavefront.
  * `id` - The ID of the event in Wavefront.
  * `start_time`- The start time of the event in epoch milliseconds.
  * `end_time` - The end time of the event in epoch milliseconds.
  * `severity` - The severity category of the event.
  * `type` - The type of the event.
  * `details` - The description of the event.
  * `is_ephemeral` - A Boolean flag. If set to `true`, creates a point-in-time event (i.e. with no duration).
  * `annotations` - Annotations associated with the event.
  * `tags` - A set of tags assigned to the event.
* `latest_start_time_epoch_millis` - Latest start time in epoch milliseconds.
* `earliest_start_time_epoch_millis` - Earliest start time in epoch milliseconds.