---
layout: "wavefront"
page_title: "Wavefront: Events"
description: |-
Get the information for all events from Wavefront
---

# Data Source: wavefront_events

Use this data source to get information for all events from Wavefront.

## Example Usage

```hcl
# Get the info for all derived metrics
data "wavefront_events" "example" {
}
```

## Attribute Reference

* `events` - List of all events in Wavefront.
    * `name` - The name of the Event in Wavefront.
    * `id` - The id of the Derived Metric in Wavefront.
    * `start_time`- The start time of the Event in epoch milliseconds.
    * `end_time` - The end time of the Event in epoch milliseconds.
    * `severity` - The severity category of the Event.
    * `type` - The type of Event.
    * `details` - The description about the Event.
    * `is_ephemeral` - A boolean flag, if true, creates a point-in-time Event( i.e. with no duration)
    * `annotations` - Annotations associated with the Event.
    * `tags` - A set of tags assigned to the Event.
