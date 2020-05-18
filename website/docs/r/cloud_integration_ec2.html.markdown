---
layout: "wavefront"
page_title: "Wavefront: Cloud Integration EC2"
description: |-
  Provides a Wavefront Cloud Integration for EC2. This allows EC2 cloud integrations to be created,
  updated, and delete
---

# Resource : wavefront_cloud_integration_ec2

Provides a Wavefront Cloud Integration for EC2. This allows EC2 cloud integrations to be created,
updated, and delete

## Example usage

```hcl
resource "wavefront_cloud_integration_aws_external_id" "ext_id" { 
}

resource "wavefront_cloud_integration_ec2" "ec2" {
  name        = "Test Integration"
  role_arn    = "arn:aws::1234567:role/example-arn"
  external_id = wavefront_cloud_integration_aws_external_id.ext_id.id
}
```

## Argument Reference

The following arguments are supported:

* `service` - (Required) A value denoting which cloud service this service integrates with
* `name` - (Required) The human-readable name of this integration
* `additional_tags` - (Optional) A list of point tag key-values to add to every point ingested using this integration
* `force_save` - (Optional) Forces this resource to save, even if errors are present
* `service_refresh_rate_in_minutes` - (Optional) How often, in minutes, to refresh the service
* `role_arn` - (Required) The external id corresponding to the Role ARN
* `external_id` - (Required) The Role ARN that the customer has created in AWS IAM to allow access to Wavefront 
* `hostname_tags` - (Optional) 

### Example

```hcl
resource "wavefront_cloud_integration_aws_external_id" "ext_id" {
}

resource "wavefront_cloud_integration_ec2" "ec2" {
  name       = "Test Integration"
  force_save = true
  additional_tags = {
    "tag1" = "value1"
    "tag2" = "value2"
  }
  role_arn    = "arn:aws::1234567:role/example-arn"
  external_id = wavefront_cloud_integration_aws_external_id.ext_id.id
  hostname_tags  = ["host", "source", "name"]
  service_refresh_rate_in_minutes = 10
}
```

## Import

EC2 Cloud Integrations can be imported using the `id`, e.g.

```
$ terraform import wavefront_cloud_integration_ec2.ec2 a411c16b-3cf7-4f03-bf11-8ca05aab898d
```