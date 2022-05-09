---
layout: "wavefront"
page_title: "Wavefront: Metrics Policy"
description: |-
Provides a Wavefront Metrics Policy Resource. This allows management of Metrics Policy to control access to time series, histograms, and delta counters
---

# Resource : wavefront_metrics_policy

Provides a Wavefront Metrics Policy Resource. This allows management of Metrics Policy to control access to time series, histograms, and delta counters

## Example usage

```hcl
data "wavefront_default_user_group" "everyone" {}

resource "wavefront_metrics_policy" "main" {
  policy_rules {
    name        = "Allow All Metrics"
    description = "Predefined policy rule. Allows access to all metrics (timeseries, histograms, and counters) for all accounts. If this rule is removed, all accounts can access all metrics if there are no matching blocking rules."
    prefixes    = ["*"]
    tags_anded  = false
    access_type = "ALLOW"
    user_group_ids = [data.wavefront_default_user_group.everyone.group_id]
  }
}
```

## Argument Reference

The following arguments are supported:
* `policy_rules` - (Required) List of Metrics Policies, must have at least one entry.
  * `name` - (Required) The unique name identifier for a Metrics Policy. The name is visible on the Metrics Security Policy page.
  * `description` - (Required) A detailed description of the Metrics Policy. The description is visible only when you edit the rule.
  * `account_ids` - (Optional) List of account ids to apply Metrics Policy to. Must have at least one associated account_id, user_group_id, or role_id.
  * `user_group_ids` - (Optional) List of user group ids to apply Metrics Policy to. Must have at least one associated account_id, user_group_id, or role_id.
  * `role_ids` -(Optional) List of role ids to apply Metrics Policy to. Must have at least one associated account_id, user_group_id, or role_id.
  * `prefixes` - (Required) List of prefixes to match metrics on. You can specify the full metric name or use a wildcard character in metric names, sources, or point tags. The wildcard character alone (*) means all metrics.
  * `tags` - (Optional) List of Key/Value tags to select target metrics for policy.
    * `key` - (Required) The tag's key.
    * `value` - (Required) The tag's value.
  * `tags_anded` - (Required) Bool where `true` require all tags are met by selected metrics, else `false` select metrics that match any give tag.
  * `access_type` (Required) Valid options are `ALLOW` and `BLOCK`.

### Example

```hcl
data "wavefront_default_user_group" "everyone" {}

resource "wavefront_role" "test" {
  name = "test-role"
  assignees = [data.wavefront_user.example.id]
}

resource "wavefront_user" "example" {
  email = "example@example.com"
  user_groups = [data.wavefront_default_user_group.everyone.group_id]
}

resource "wavefront_metrics_policy" "main" {
  policy_rules {
    name        = "Deny example role metrics"
    description = "deny example role test"
    prefixes    = ["example.api.*"]
    tags_anded  = false
    access_type = "BLOCK"
    role_ids       = [wavefront_role.test.id]
  }
  policy_rules {
    name        = "Deny example user metrics"
    description = "deny example user test"
    prefixes    = ["example.system.*"]
    tags {
      key = "env"
      value = "prod"
    }
    tags {
      key = "region"
      value = "us-east-1"
    }

    tags_anded  = true
    access_type = "BLOCK"
    account_ids    = [wavefront_user.example.id]
  }
  policy_rules {
    name        = "Allow All Metrics"
    description = "Predefined policy rule. Allows access to all metrics (timeseries, histograms, and counters) for all accounts. If this rule is removed, all accounts can access all metrics if there are no matching blocking rules."
    prefixes    = ["*"]
    tags_anded  = false
    access_type = "ALLOW"
    user_group_ids = [data.wavefront_default_user_group.everyone.group_id]
  }
}
```

## Attribute Reference

* `customer`- The customer the user is associated with.
* `updater_id` - The account_id who applied the current policy.
* `updated_epoch_millis` - When the policy was applied in epoch_millis.

## Import

Users can be imported by using the `updated_epoch_millis`, e.g.:

```
$ terraform import wavefront_metrics_policy.some_metrics_policy 1651846476678
```