---
layout: "wavefront"
page_title: "Wavefront: Default User Group"
description: |-
    Get the default user group `Everyone` from Wavefront
---

# Data Source: wavefront_default_user_group

Use this data source to get the Group ID of the `Everyone` group in Wavefront.

## Example Usage

```hcl
# Get the default user group "Everyone"
data "wavefront_default_user_group" "everyone_group" {
}
```

## Attribute Reference

* `group_id` - Set to the Group ID of the `Everyone` group, suitable for referencing
  in other resources that support group memberships.