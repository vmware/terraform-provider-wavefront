---
layout: "wavefront"
page_title: "Wavefront: User"
description: |-
  Provides a Wavefront User Resource. This allows user accounts to be created, updated, and deleted.
---

# Resource : wavefront_user

Provides a Wavefront User Resource. This allows user accounts to be created, updated, and deleted.

## Example usage

```hcl
resource "wavefront_user" "basic" {
  email  = "test+tftesting@example.com"
}
```

## Argument Reference

The following arguments are supported:

* `email` - (Required) The unique identifier of the user account to create. Must be a valid email address.
* `permissions` - (Optional) A list of permissions to assign to this role. Valid options are:
`agent_management`, `alerts_management`, `application_management`, `batch_query_priority`,
`dashboard_management`, `derived_metrics_management`, `embedded_charts`, `events_management`,
`external_links_management`, `host_tag_management`, `ingestion`, `metrics_management`,
`monitored_application_service_management`, `saml_sso_management`, `token_management`,
`user_management`.
* `user_groups` - (Optional) List of user groups to this user.

### Example

```hcl

resource "wavefront_user_group" "test_group" {
  name        = "Test Group"
  description = "Test Group"
}

resource "wavefront_user" "basic" {
  email       = "test+tftesting@example.com"
  permissions = [
    "agent_management",
    "events_management",
  ]
  user_groups = [
    wavefront_user_group.test_group.id
  ]
}
```

## Attribute Reference

* `customer`- The customer the user is associated with.

## Import

Users can be imported by using the `id`, e.g.:

```
$ terraform import wavefront_user.some_user test@example.com
```