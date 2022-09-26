---
layout: "wavefront"
page_title: "Wavefront: Event"
description: |-
Get the information about a certain Wavefront event.
---

# Data Source: wavefront_event

Use this data source to get information about a certain Wavefront event.

## Argument Reference
* `id` - (Required) The ID associated with the event data to be fetched.

## Example Usage

```hcl
# Get the information about a Wavefront event by its ID.
data "wavefront_event" "example" {
   id = "sample-event-id"
}
```

## Attribute Reference

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
