---
layout: "wavefront"
page_title: "Wavefront: User Group"
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
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the user group
* `description` - (Required) A short description of the user group

### Example

```hcl
resource "wavefront_user_group" "basic" {
  name        = "Basic User Group"
  description = "Basic User Group for Unit Tests"
}
```

## Import

User Groups can be imported using the `id`, e.g.

```
$ terraform import wavefront_user_group.some_group a411c16b-3cf7-4f03-bf11-8ca05aab898d
```