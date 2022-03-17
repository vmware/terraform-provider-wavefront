---
layout: "wavefront"
page_title: "Wavefront: Cloud Integration Tesla"
description: |-
  Provides a Wavefront Cloud Integration for Tesla. This allows Tesla cloud integrations to be created,
  updated, and deleted.
---

# Resource : wavefront_tesla

Provides a Wavefront Cloud Integration for Tesla. This allows Tesla cloud integrations to be created,
updated, and deleted.

## Example usage

```hcl
resource "wavefront_cloud_integration_tesla" "tesla" {
  name              = "Test Integration"
  email    = "email@example.com"
  password = "password"
}
```

## Argument Reference

The following arguments are supported:

* `service` - (Required) A value denoting which cloud service this service integrates with.
* `name` - (Required) The human-readable name of this integration.
* `additional_tags` - (Optional) A list of point tag key-values to add to every point ingested using this integration.
* `force_save` - (Optional) Forces this resource to save, even if errors are present.
* `service_refresh_rate_in_minutes` - (Optional) How often, in minutes, to refresh the service.
* `email` - (Required) A Tesla account login email address.
* `password` - (Required) The Tesla account login password.

### Example

```hcl
resource "wavefront_cloud_integration_tesla" "tesla" {
  name              = "Test Integration"
  force_save        = true
  additional_tags = {
    "tag1" = "value1"
    "tag2" = "value2"
  }
  email    = "email@example.com"
  password = "password"
  service_refresh_rate_in_minutes = 10
}
```

## Import

Tesla Integrations can be imported by using the `id`, e.g.:

```
$ terraform import wavefront_cloud_integration_tesla.tesla a411c16b-3cf7-4f03-bf11-8ca05aab898d
```