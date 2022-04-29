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

data "wavefront_user" "user_1" {
  id = "sean.norris@woven-planet.global"
}

#data "wavefront_users" "users" {}

#data "wavefront_roles" "roles" {}

data "wavefront_metrics_policy" "policies" {}

output "groups" {
  value = data.wavefront_default_user_group.everyone_group.group_id
}

output "user_1" {
  value = data.wavefront_user.user_1.identifier
}

#output "all_users" {
#  value = data.wavefront_users.users.users
#}

#output "all_roles" {
#  value = data.wavefront_roles.roles.roles
#}

output "all_rules" {
  value = data.wavefront_metrics_policy.policies
}