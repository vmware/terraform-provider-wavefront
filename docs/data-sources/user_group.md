---
layout: "wavefront"
page_title: "Wavefront: User Group"
description: |-
Get the information about a specific Wavefront user group.
---

# Data Source: wavefront_user_group

Use this data source to get information about a Wavefront user group by its ID.

## Argument Reference
* `id` - (Required) The ID associated with the user group data to be fetched.

## Example Usage

```hcl
# Get the information about the role.
data "wavefront_user_group" "example" {
  id = "user-group-id"
}
```

## Attribute Reference

* `id` - The ID of the group in Wavefront.
* `name` - The name of the group in Wavefront.
* `description` - Human-readable description of the group.
* `roles` - The list of roles associated with the group.
* `users` - The list of users assigned to the group.
