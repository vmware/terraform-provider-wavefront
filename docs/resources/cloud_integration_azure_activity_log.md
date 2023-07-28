---
layout: "wavefront"
page_title: "Wavefront: Cloud Integration Azure Activity Logs"
description: |-
  Provides a Wavefront Cloud Integration for Azure Activity Logs. This allows Azure activity log cloud integrations to be created,
  updated, and deleted.
---

# Resource : wavefront_cloud_integration_azure_activity_log

Provides a Wavefront Cloud Integration for Azure Activity Logs. This allows Azure activity log cloud integrations to be created,
updated, and deleted.

## Example usage

```hcl
resource "wavefront_cloud_integration_azure_activity_log" "azure_activity_log" {
  name            = "Test Integration"
  category_filter = ["ADMINISTRATIVE"]
  client_id       = "client-id2"
  client_secret   = "client-secret2"
  tenant          = "my-tenant2"
}
```

## Argument Reference

The following arguments are supported:

* `service` - (Required) A value denoting which cloud service this service integrates with.
* `name` - (Required) The human-readable name of this integration.
* `additional_tags` - (Optional) A list of point tag key-values to add to every point ingested using this integration.
* `force_save` - (Optional) Forces this resource to save, even if errors are present.
* `service_refresh_rate_in_minutes` - (Optional) How often, in minutes, to refresh the service.
* `client_secret` - (Required) Client secret for an Azure service account within your project.
* `client_id` - (Required) Client ID for an Azure service account within your project.
* `tenant` - (Required)  Tenant ID for an Azure service account within your project.
* `category_filter` - (Optional) A list of Azure services (such as Microsoft.Compute/virtualMachines) from which to pull metrics.

### Example

```hcl
resource "wavefront_cloud_integration_azure_activity_log" "azure_activity_log" {
  name       = "Test Integration"
  force_save = true
  additional_tags = {
    "tag1" = "value1"
    "tag2" = "value2"
    "tag3" = "value3"
  }
  category_filter = ["ADMINISTRATIVE", "SERVICEHEALTH"]
  client_id       = "client-id2"
  client_secret   = "client-secret2"
  tenant          = "my-tenant2"
  service_refresh_rate_in_minutes = 10
}
```

## Import

Azure Activity Log Cloud Integrations can be imported by using the `id`, e.g.:

```
$ terraform import wavefront_cloud_integration_azure_activity_log.azure_al a411c16b-3cf7-4f03-bf11-8ca05aab898d
```