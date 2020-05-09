---
layout: "wavefront"
page_title: "Wavefront: User"
description: |-
  Provides a Wavefront User Resource. This allows users to be created, updated, and deleted.
---

# Resource : wavefront_user

Provides a Wavefront User Resource. This allows users to be created, updated, and deleted.

## Example usage

```hcl
resource "wavefront_user" "basic" {
  email  = "test+tftesting@example.com"
}
```

## Argument Reference

The following arguments are supported:

* `email` - (Required) The (unique) identifier of the user to create. Must be a valid email address
* `groups` - (Optional) List of permission groups to grant to this user.  Valid options are
* `user_groups` - (Optional) List of user groups to this user
`agent_management`, `alerts_management`, `dashboard_management`, `embedded_charts`, `events_management`, `external_links_management`,
`host_tag_management`, `metrics_management`, `user_management`

### Example

```hcl

resource "wavefront_user_group" "test_group" {
  name = "Test Group"
}

resource "wavefront_user" "basic" {
  email  = "test+tftesting@example.com"
  groups = [
    "agent_management",
    "events_management",
  ]
  user_groups = [
    wavefront_user_group.test_group.id
  ]
}
```

## Attribute Reference

* `customer`- The customer the user is associated with 

## Import

Users can be imported using the `id`, e.g.

```
$ terraform import wavefront_user.some_user test@example.com
```