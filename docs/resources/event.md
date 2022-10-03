---
layout: "wavefront"
page_title: "Wavefront: Event"
description: |-
    Provides a Wavefront Event resource. This allows events to be created, updated, and deleted.
---

# Resource : wavefront_event

Provides a Wavefront event resource. This allows events to be created, updated, and deleted.

## Example Usage

```hcl
resource "wavefront_event" "event" {
  name = "terraform-test"
  annotations = {
    severity = "info"
    type = "event type"
    details = "description"
  }
  tag = [
    "eventTag1"
  ]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the event as it is displayed in Wavefront.
* `tags` - (Optional) A set of tags to assign to this resource.
* `annotations` - (Required) The annotations associated with the event.
* `start_time`- (Optional) The start time of the event in epoch milliseconds.
* `end_time` - (Optional) The end time of the event in epoch milliseconds.

### Example
```hcl
resource "wavefront_event" "event" {
  name = "terraform-test"
  annotations = {
    severity = "info"
    type = "event type"
    details = "description"
  }
  tag = [
    "eventTag1"
  ]
  start_time = 1490000000000
  end_time = 1490000000001
}
```

## Import

You can import events by using the id, for example:

```
$ terraform import wavefront_event.event 1479868728473
```