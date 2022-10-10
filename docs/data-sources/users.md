---
layout: "wavefront"
page_title: "Wavefront: Users"
description: |-
    Get all users from Wavefront
---

# Data Source: wavefront_users

Use this data source to get all users in Wavefront. 

## Argument Reference
* `limit` - (Optional) Limit is the maximum number of results to be returned. Defaults to 100.
* `offset` - (Optional) Offset is the offset from the first result to be returned. Defaults to 0.


## Example Usage

```hcl
# Get all users
data "wavefront_users" "users" {
  limit = 10
  offset = 0
}
```

## Attribute Reference

* `users` - List of all users in Wavefront.
  * `permissions` - List of permissions granted to a user.
  * `user_group_ids` - List of User Group Ids the user is a member of.
  * `customer`- The customer the user is associated with.
  * `last_successful_login` - When the user last logged in to Wavefront.