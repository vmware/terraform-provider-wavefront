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

* `identifier` - (Required) The unique identifier of the service account to create. Must have the prefix `sa::`.
* `active` - (Required) Whether or not the service account is active.
* `description` - (Optional) The description of the service account.
* `permissions` - (Optional) A list of permissions to assign to this role. Valid options are:
`agent_management`, `alerts_management`, `application_management`, `batch_query_priority`,
`dashboard_management`, `derived_metrics_management`, `embedded_charts`, `events_management`,
`external_links_management`, `host_tag_management`, `ingestion`, `metrics_management`,
`monitored_application_service_management`, `saml_sso_management`, `token_management`,
`user_management`.
* `user_groups` - (Optional) List of user groups for this service account.
* `ingestion_policy` - (Optional) ID of ingestion policy.

### Example

```hcl

resource "wavefront_user_group" "test_group" {
  name        = "Test Group"
  description = "Test Group"
}

resource "wavefront_ingestion_policy" "test_ingestion" {
  name = "test_ingestion"
  description = "An ingestion policy for testing"
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
  ingestion_policy = wavefront_ingestion_policy.test_ingestion.id
}
```

## Import

Service accounts can be imported by using `identifier`, e.g.:

```
$ terraform import wavefront_service_account.basic sa::tftesting
```
