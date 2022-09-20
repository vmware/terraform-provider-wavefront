---
layout: "wavefront"
page_title: "Wavefront: Event"
description: |-
Get the information for a given event from Wavefront
---

# Data Source: wavefront_event

Use this data source to get information for a given event from Wavefront.

## Argument Reference
* `id` - (Required) The id associated with the event data to be fetched.

## Example Usage

```hcl
# Get the info for all derived metrics
data "wavefront_event" "example" {
   id = "sample-event-id"
}
```

## Attribute Reference

* `name` - The name of the Event in Wavefront.
* `id` - The id of the Event in Wavefront.
* `start_time`- The start time of the Event in epoch milliseconds.
* `end_time` - The end time of the Event in epoch milliseconds.
* `severit` - The severity category of the Event.
* `type` - The type of Event.
* `details` - The description about the Event.
* `is_ephemeral` - A boolean flag, if true, creates a point-in-time Event( i.e. with no duration)
* `annotations` - Annotations associated with the Event.
* `tags` - A set of tags assigned to the Event.
