---
layout: "wavefront"
page_title: "Wavefront: Cloud Integration Cloudtrail"
description: |-
  Provides a Wavefront Cloud Integration for CloudTrail. This allows CloudTrail cloud integrations to be created,
  updated, and deleted.
---

# Resource : wavefront_cloud_integration_cloudtrail

Provides a Wavefront Cloud Integration for CloudTrail. This allows CloudTrail cloud integrations to be created,
updated, and deleted.

## Example usage

```hcl
resource "wavefront_cloud_integration_aws_external_id" "ext_id" { 
}

resource "wavefront_cloud_integration_cloudtrail" "cloudtrail" {
  name        = "Test Integration"
  role_arn    = "arn:aws::1234567:role/example-arn"
  external_id = wavefront_cloud_integration_aws_external_id.ext_id.id
  region      = "us-west-2"
  bucket_name = "example-s3-bucket"
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
* `region` - (Required) The AWS region of the S3 bucket where CloudTrail logs are stored
* `bucket_name` - (Required) Name of the S3 bucket where CloudTrail logs are stored
* `filter_rule` - (Optional) Rule to filter CloudTrail log event
* `prefix` - (Optional) The common prefix, if any, appended to all CloudTrail log files.

### Example
```hcl
resource "wavefront_cloud_integration_aws_external_id" "ext_id" {
}

resource "wavefront_cloud_integration_cloudtrail" "cloudtrail" {
  name       = "Test Integration"
  force_save = true
  additional_tags = {
    "tag1" = "value1"
    "tag2" = "value2"
  }
  region      = "us-west-1"
  bucket_name = "example-s3-bucket2"
  filter_rule = "someFilterRule"
  role_arn    = "arn:aws::1234567:role/example-arn"
  external_id = wavefront_cloud_integration_aws_external_id.ext_id.id
}
```

## Import

CloudTrail Cloud Integrations can be imported using the `id`, e.g.

```
$ terraform import wavefront_cloud_integration_cloudtrail.cloudtrail a411c16b-3cf7-4f03-bf11-8ca05aab898d
```