---
layout: "wavefront"
page_title: "Wavefront: Roles"
description: |-
    Get all Roles from Wavefront
---

# Data Source: wavefront_roles

Use this data source to get all Roles in Wavefront. 

## Argument Reference
* `limit` - (Optional) Limit is the maximum number of results to be returned. Defaults to 100.
* `offset` - (Optional) Offset is the offset from the first result to be returned. Defaults to 0.

## Example Usage

```hcl
# Get all Roles
data "wavefront_roles" "roles" {
  limit = 10
  offset = 0
}
```

## Attribute Reference
* `roles` - List of Wavefront Roles.
  * `id` - The Role ID.
  * `name` - The Role Name.
  * `description` - The Role's description.
  * `permissions` - List of Permissions (Strings) associated with Role.