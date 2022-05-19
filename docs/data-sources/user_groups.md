---
layout: "wavefront"
page_title: "Wavefront: User Groups"
description: |-
    Get all User Groups from Wavefront
---

# Data Source: wavefront_user_groups

Use this data source to get all User Groups in Wavefront. 

## Example Usage

```hcl
# Get all user groups
data "wavefront_user_groups" "groups" {
}
```

## Attribute Reference

* `user_groups` - List of user groups.
  * `id` - The group ID.
  * `name` - The group name.
  * `description` - The group description.
  * `roles` - List of roles associated with the group.
  * `users` - List of users assigned to the group.
