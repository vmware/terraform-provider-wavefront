---
layout: "wavefront"
page_title: "Wavefront: Cloud Integration App Dynamics"
description: |-
  Provides a Wavefront Cloud Integration for App Dynamics. This allows app dynamics cloud integrations to be created,
  updated, and deleted.
---

# Resource : wavefront_cloud_integration_app_dynamics

Provides a Wavefront Cloud Integration for App Dynamics. This allows app dynamics cloud integrations to be created, 
updated, and deleted.

## Example usage

```hcl
resource "wavefront_cloud_integration_app_dynamics" "app_dynamics" {
  name 				= "Test Integration"
  user_name 			= "example"
  controller_name 	= "exampleController"
  encrypted_password 	= "encryptedPassword"	
}
```

## Argument Reference

The following arguments are supported:

* `service` - (Required) A value denoting which cloud service this service integrates with
* `name` - (Required) The human-readable name of this integration
* `additional_tags` - (Optional) A list of point tag key-values to add to every point ingested using this integration
* `force_save` - (Optional) Forces this resource to save, even if errors are present
* `service_refresh_rate_in_minutes` - (Optional) How often, in minutes, to refresh the service
* `user_name` - (Required) Username is a combination of userName and the account name
* `controller_name` - (Required) Name of the SaaS controller
* `encrypted_password` - (Required) Password for AppDynamics user
* `app_filter_regex` - (Optional)  List of regular expressions that a application name must match (case-insensitively) 
in order to be ingested
* `enable_rollup` - (Optional) Set this to `false` to get separate results for all values within the time range, 
by default it is `true` 
* `enable_error_metrics` - (Optional) Boolean flag to control Error metric injection
* `enable_business_trx_metrics` - (Optional) Boolean flag to control Business Transaction metric injection 
* `enable_backend_metrics` - (Optional) Boolean flag to control Backend metric injection
* `enable_overall_perf_metrics` - (Optional) Boolean flag to control Overall Performance metric injection
* `enable_individual_node_metrics` - (Optional) Boolean flag to control Individual Node metric injection
* `enable_app_infra_metrics` - (Optional) Boolean flag to control Application Infrastructure metric injection
* `enable_service_endpoint_metrics` - (Optional) Boolean flag to control Service End point metric injection


### Example

```hcl
resource "wavefront_cloud_integration_app_dynamics" "app_dynamics" {
  name       = "Test Integration"
  force_save = true
  additional_tags = {
    "tag1" = "value1"
    "tag2" = "value2"
  }
  user_name                       = "example2"
  controller_name                 = "exampleController2"
  encrypted_password              = "encryptedPassword"
  enable_rollup                   = false
  enable_error_metrics            = true
  enable_business_trx_metrics     = true
  enable_backend_metrics          = true
  enable_overall_perf_metrics     = true
  enable_individual_node_metrics  = true
  enable_app_infra_metrics        = true
  enable_service_endpoint_metrics = true
}
```

## Import

App Dynamic Cloud Integrations can be imported using the `id`, e.g.

```
$ terraform import wavefront_cloud_integration_app_dynamics.app_dynamics a411c16b-3cf7-4f03-bf11-8ca05aab898d
```
