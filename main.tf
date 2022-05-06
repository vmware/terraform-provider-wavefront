terraform {
  required_providers {
    wavefront = {
      source = "vmware/wavefront"
      version = "3.1.0"
    }
  }
}

# sanity
data "wavefront_default_user_group" "everyone_group" {}



data "wavefront_roles" "roles" {}

data "wavefront_user" "user_1" {
  email = "sean.norris@woven-planet.global"
}
#// 27230d83-6ea1-49c3-892c-1e8686756047
#resource "wavefront_role" "tf-test-role" {
#  name = "tf-test-role"
#}
#
#// 8a382be4-7005-404d-afcd-307e09b227bc
#resource "wavefront_role" "test-role" {
#  name = "test-role"
#}

#resource "wavefront_user" "test" {
#  email = "test@woven-planet.global"
#  permissions = [wavefront_role.test.id]
#  user_groups = [data.wavefront_default_user_group.everyone_group.group_id]
#}
#
#resource "wavefront_role" "test" {
#  name = "test-role"
#  assignees = [data.wavefront_user.user_1.id]
#}

# import import wavefront_metrics_policy.main 1651683740449

#resource "wavefront_metrics_policy" "main" {
#  policy_rules {
#    name        = "Deny Sean gloo metrics - custom test"
#    description = "deny sean test"
#    prefixes    = ["api.gloo.*"]
#    tags {
#      key = "env"
#      value = "prod"
#    }
#    tags {
#      key = "region"
#      value = "us-east-1"
#    }
#    tags_anded  = false
#    access_type = "BLOCK"
#    account_ids    = [data.wavefront_user.user_1.id]
#    user_group_ids = []
#    role_ids       = []
#  }
#  policy_rules {
#    name        = "Allow All Metrics"
#    description = "Predefined policy rule. Allows access to all metrics (timeseries, histograms, and counters) for all accounts. If this rule is removed, all accounts can access all metrics if there are no matching blocking rules."
#    prefixes    = ["*"]
#    tags_anded  = false
#    access_type = "ALLOW"
#    account_ids    = []
#    user_group_ids = [data.wavefront_default_user_group.everyone_group.group_id]
#    role_ids       = []
#  }
#}

#data "wavefront_users" "users" {}

#data "wavefront_roles" "roles" {}

#data "wavefront_metrics_policy" "policies" {}

#output "groups" {
#  value = data.wavefront_default_user_group.everyone_group.group_id
#}
#
output "user_1" {
  value = data.wavefront_user.user_1
}

#output "all_users" {
#  value = data.wavefront_users.users.users
#}

#output "all_roles" {
#  value = data.wavefront_roles.roles.roles
#}

#output "all_rules" {
#  value = data.wavefront_metrics_policy.policies
#}