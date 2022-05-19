---
layout: "wavefront"
page_title: "Wavefront: User"
description: |-
    Get the information for a given user from Wavefront
---

# Data Source: wavefront_user

Use this data source to get information for a given user by email from Wavefront. 

## Argument Reference
* `email` - The email associated with the user data to be fetched.

## Example Usage

```hcl
# Get the info for user "example.user@example.com"
data "wavefront_user" "example" {
  email = "example.user@example.com"
}
```

## Attribute Reference

* `permissions` - List of permissions granted to a user.
* `user_group_ids` - List of User Group Ids the user is a member of.
* `customer`- The customer the user is associated with.
* `last_successful_login` - When the user last logged in to Wavefront.
