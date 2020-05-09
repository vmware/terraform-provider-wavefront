---
layout: "wavefront"
page_title: "Wavefront: "
description: |-
  Provides a Wavefront User Group Resource. This allows user groups to be created, updated, and deleted.
---

# Resource : wavefront_user_group

Provides a Wavefront User Group Resource. This allows user groups to be created, updated, and deleted.

## Example usage

```hcl
resource "wavefront_user_group" "basic" {
  name        = "Basic User Group"
  description = "Basic User Group for Unit Tests"
  permissions = [
    "alerts_management",
	"events_management"
  ]
}
```

## Argument Reference

The following arguments are supported:

* `permissions` - (Required) 
* `name` - (Required) 
* `description` - (Required) 

### Example

```hcl
resource "wavefront_user_group" "basic" {
  name        = "Basic User Group"
  description = "Basic User Group for Unit Tests"
  permissions = [
    "alerts_management",
	"events_management"
  ]
}
```

## Import

User Groups can be imported using the `id`, e.g.

```
$ terraform import wavefront_user_group.some_group a411c16b-3cf7-4f03-bf11-8ca05aab898d
```