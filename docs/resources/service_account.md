---
layout: "wavefront"
page_title: "Wavefront: Service Account"
description: |-
  Provides a Wavefront Service Account Resource. This allows service accounts to be created, updated, and deleted.
---

# Resource : wavefront_service_account

Provides a Wavefront Service Account Resource. This allows service accounts to be created, updated, and deleted.

## Example usage

```hcl
resource "wavefront_service_account" "basic" {
  identifier  = "sa::tftesting"
  active = true
}
```

## Argument Reference

The following arguments are supported:

* `identifier` - (Required) The (unique) identifier of the service account to create. Must start with sa::
* `active` - (Required) Whether or not the service account is active
* `description` - (Optional) The description of the service account
* `permissions` - (Optional) List of permission to grant to this service account.  Valid options are
`agent_management`, `alerts_management`, `dashboard_management`, `embedded_charts`, `events_management`, `external_links_management`,
`host_tag_management`, `metrics_management`, `user_management`
* `user_groups` - (Optional) List of user groups for this service account

### Example

```hcl

resource "wavefront_user_group" "test_group" {
  name        = "Test Group"
  description = "Test Group"
}

resource "wavefront_service_account" "basic" {
  identifier  = "sa::tftesting"
  active      = true
  description = "A service account description"
  permissions = [
    "agent_management",
    "events_management",
  ]
  user_groups = [
    wavefront_user_group.test_group.id
  ]
}
```

## Import

Service accounts can be imported using `identifier`, e.g.

```
$ terraform import wavefront_service_account.basic sa::tftesting
```
