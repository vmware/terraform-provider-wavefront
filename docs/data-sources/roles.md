---
layout: "wavefront"
page_title: "Wavefront: Roles"
description: |-
    Get all Roles from Wavefront
---

# Data Source: wavefront_roles

Use this data source to get all Roles in Wavefront. 

## Example Usage

```hcl
# Get all Roles
data "wavefront_roles" "roles" {
}
```

## Attribute Reference
* `roles` - List of Wavefront Roles.
  * `id` - The Role ID.
  * `name` - The Role Name.
  * `description` - The Role's description.
  * `permissions` - List of Permissions (Strings) associated with Role.
