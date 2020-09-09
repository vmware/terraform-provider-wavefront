---
layout: "wavefront"
page_title: "Wavefront: Role"
description: |-
  Provides a Wavefront Role Resource. This allows roles to be created, updated, and deleted.
---

# Resource : wavefront_user_group

Provides a Wavefront Role Resource. This allows user groups to be created, updated, and deleted.

## Example usage

```hcl
resource "wavefront_role" "role" {
  name        = "Test Role"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the user group
* `description` - (Optional) A short description of the user group
* `permissions` - (Optional) A list of permissions to assign to this role. Valid options are 
`agent_management`, `alerts_management`, `dashboard_management`, `embedded_charts`, `events_management`, `external_links_management`,
`host_tag_management`, `metrics_management`, `user_management`
* `assignees` - (Optional) A list of user groups or accounts to assign to this role. 


### Example

```hcl
resource "wavefront_user_group" "agents_group" {
  name        = "Test Group"
  description = "Test Group"
}

resource "wavefront_user" "basic" {
  email       = "test+tftesting@example.com"
  user_groups = [
    wavefront_user_group.test_group.id
  ]
}

resource "wavefront_role" "agent_management" {
  name        = "Agent Management Role"
  description = "Agent Management Role for Testing"
  permissions = [
    "agent_management"
  ]
  assignees   = [
    wavefront_user_group.test_group.id
  ]
}
```

## Import

User Groups can be imported using the `id`, e.g.

```
$ terraform import wavefront_user_group.some_group a411c16b-3cf7-4f03-bf11-8ca05aab898d
```