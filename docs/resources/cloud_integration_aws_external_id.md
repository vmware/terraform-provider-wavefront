---
layout: "wavefront"
page_title: "Wavefront: AWS External ID"
description: |-
  Provides an External ID for use in AWS IAM Roles.  This allows External IDs to be created and deleted. 
---

# Resource : wavefront_cloud_integration_aws_external_id

Provides an External ID for use in AWS IAM Roles.  This allows External IDs to be created and deleted.

## Example usage

```hcl
resource "wavefront_cloud_integration_aws_external_id" "external_id" {
}
```

## Attributes Reference

* `id` - The External ID created in Wavefront

## Import

External IDs can be imported using the `id`, e.g.

```
$ terraform import wavefront_cloud_integration_aws_external_id.external_id uGJdkH3k
```