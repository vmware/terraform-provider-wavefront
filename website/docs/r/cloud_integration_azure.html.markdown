---
layout: "wavefront"
page_title: "Wavefront: Cloud Integration Azure"
description: |-
  Provides a Wavefront Cloud Integration for Azure. This allows azure cloud integrations to be created,
  updated, and deleted.
---

# Resource : wavefront_cloud_integration_azure

Provides a Wavefront Cloud Integration for Azure. This allows azure cloud integrations to be created,
updated, and deleted.

## Example usage

```hcl
resource "wavefront_cloud_integration_azure_activity_log" "azure_activity_log" {
  name          = "Test Integration"
  client_id     = "client-id2"
  client_secret = "client-secret2"
  tenant        = "my-tenant2"
}
```


## Argument Reference

The following arguments are supported:

* `service` - (Required) A value denoting which cloud service this service integrates with
* `name` - (Required) The human-readable name of this integration
* `additional_tags` - (Optional) A list of point tag key-values to add to every point ingested using this integration
* `force_save` - (Optional) Forces this resource to save, even if errors are present.
* `client_secret` - (Required) Client secret for an Azure service account within your project
* `client_id` - (Required) Client id for an azure service account within your project
* `tenant` - (Required)  Tenant Id for an Azure service account within your project
* `resource_group_filter` - (Optional) A list of Azure resource groups from which to pull metrics
* `metric_filter_regex` - (Optional) A regular expression that a metric name must match (case-insensitively) in order to be ingested
* `category_filter` - (Optional) A list of Azure Activity Log categories.

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
}
```

## Import

Azure Cloud Integrations can be imported using the `id`, e.g.

```
$ terraform import wavefront_cloud_integration_azure.azure a411c16b-3cf7-4f03-bf11-8ca05aab898d
```