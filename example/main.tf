terraform {
  required_providers {
    wavefront = {
      source  = "vmware/wavefront",
      version = ">= 3.0.0"
    }
  }
}

provider "wavefront" {
  address = "<YOUR_WAVEFRONT_CLUSTER_ADDRESS>"
  token   = "<YOUR_WAVEFRONT_CLUSTER_TOKEN>"
}

resource "wavefront_alert" "mac_cpu_usage_over_ninety_percent" {
  name                   = "Mac CPU Usage Over 90%"
  alert_type             = "THRESHOLD"
  display_expression     = "100 - avg(ts(\"mac.cpu.usage.idle\", source=\"mac.host\"))"
  conditions             = {
    "severe"             = "100 - avg(ts(\"mac.cpu.usage.idle\", source=\"mac.host\")) > 90"
  }
  additional_information = "This is a sample alert created by terraform. It monitors the CPU Usage and fires when it's over 90%."
  minutes                = 5
  resolve_after_minutes  = 5
  threshold_targets      = {
    "severe"             = "alert_target@example.com"
  }
  tags                   = [
    "example",
    "cpu.usage"
  ]
}
