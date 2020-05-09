---
layout: "wavefront"
page_title: "Wavefront: Cloud Integration CloudWatch"
description: |-
  Provides a Wavefront Cloud Integration for CloudWatch. This allows CloudWatch cloud integrations to be created,
  updated, and delete
---

# Resource : wavefront_cloud_integration_cloudwatch

Provides a Wavefront Cloud Integration for CloudTrail. This allows CloudTrail cloud integrations to be created,
updated, and delete
  
## Example usage

```hcl
resource "wavefront_cloud_integration_aws_external_id" "ext_id" { 
}

resource "wavefront_cloud_integration_cloudwatch" "cloudwatch" {
  name        = "Test Integration"
  force_save  = true
  role_arn    = "arn:aws::1234567:role/example-arn"
  external_id = wavefront_cloud_integration_aws_external_id.ext_id.id
}
```

## Argument Reference

The following arguments are supported:

* `service` - (Required) A value denoting which cloud service this service integrates with
* `name` - (Required) The human-readable name of this integration
* `additional_tags` - (Optional) A list of point tag key-values to add to every point ingested using this integration
* `force_save` - (Optional) Forces this resource to save, even if errors are present.
* `role_arn` - (Required) The external id corresponding to the Role ARN
* `external_id` - (Required) The Role ARN that the customer has created in AWS IAM to allow access to Wavefront
* `point_tag_filter_regex` - (Optional) A regular expression that AWS tag key name must match (case-insensitively)
in order to be ingested 
* `volume_selection_tags` - (Optional) A string->string map of whitelist of volume tag-value pairs (in AWS).
If the volume's AWS tags match this whitelist, CloudWatch data about this volume is ingested. 
Multiple entries are OR'ed
* `instance_selection_tags` - (Optional) A string->string map whitelist of instance tag-value pairs (in AWS).
If the instance's AWS tags match this whitelist, CloudWatch data about this instance is ingested. 
Multiple entries are OR'ed 
* `metric_filter_regex` - (Optional) A regular expression that a CloudWatch metric name must match (case-insensitively) in order to be ingested
* `namespaces` - (Optional) A list of namespaces that limit what we query from CloudWatch

### Example

```hcl
resource "wavefront_cloud_integration_aws_external_id" "ext_id" {
}

resource "wavefront_cloud_integration_cloudwatch" "cloudwatch" {
  name       = "Test Integration"
  force_save = true
  additional_tags = {
    "tag1" = "value1"
    "tag2" = "value2"
  }
  role_arn    = "arn:aws::1234567:role/example-arn"
  external_id = wavefront_cloud_integration_aws_external_id.ext_id.id
  namespaces  = ["ec2", "elb", "route53"]
  instance_selection_tags = {
    "env"    = "dev"
    "mirror" = "b"
  }
  volume_selection_tags = {
    "env" = "dev"
  }
  point_tag_filter_regex = "^dev"
  metric_filter_regex    = "^.*?\\.cpu.*$"
}
```

## Import

CloudWatch Cloud Integrations can be imported using the `id`, e.g.

```
$ terraform import wavefront_cloud_integration_cloudwatch.cloudwatch a411c16b-3cf7-4f03-bf11-8ca05aab898d
```