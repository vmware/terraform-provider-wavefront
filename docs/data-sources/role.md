---
layout: "wavefront"
page_title: "Wavefront: Role"
description: |-
    Get the information about a specific Wavefront role.
---

# Data Source: wavefront_role

Use this data source to get information about a Wavefront role by its ID.

## Argument Reference
* `id` - (Required) The ID associated with the role data to be fetched.

## Example Usage

```hcl
# Get the information about the role.
data "wavefront_role" "example" {
  id = "role-id"
}
```

## Attribute Reference

* `id` - The ID of the role in Wavefront.
* `name` - The name of the role in Wavefront.
* `description` - Human-readable description of the role.
* `permissions` - The list of permissions associated with role.