---
layout: "wavefront"
page_title: "Provider: Wavefront"
description: |-
    The Wavefront provider is used to interact with many resources supported by Wavefront.  The provider needs to be configured with the proper credentials before it can be used.
---

# Wavefront Provider

The Wavefront provider is used to interact with the Wavefront monitoring service. The
provider needs to be configured with the roper credentials before it can be used.

Use the navigation to the left to read about the available resources. 

## Example Usage

```hcl
# Configure the Wavefront provider
provider "wavefront" {
  version = "~> 2.0"
}

resource "wavefront_alert" "test_alert" {
  name                   = "High CPU Alert"
  condition              = "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total ) > 80"
  additional_information = "This is an Alert"
  display_expression     = "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total )"
  minutes                = 5
  severity               = "WARN"
  tags = [
    "env.preprod",
    "cpu.total"
  ]
}
```

## Authentication

The Wavefront provider offers two means of providing credentials for authentication.

* Static credentials
* Environment variables

### Static credentials
!> **Warning:** Hard-coding credentials into any Terraform configuration is not
recommended, and risks secret leakage should this file ever be committed to a 
public version control system.
 
Static credentials can be provided by adding an `address` and `token` in-line in 
the Wavefront provider block. 

```hcl
provider "wavefront" {
  address = "cluster.wavefront.com"
  token   = "your-wf-token-secret"
}
```

### Environment Variables

You can provide your credentials via the `WAVEFRONT_ADDRESS` and `WAVEFRONT_TOKEN`
, environment variables.  

```hcl
provider "wavefront" {}
```

Usage:

```sh
$ export WAVEFRONT_ADDRESS="cluster.wavefront.com"
$ export WAVEFRONT_TOKEN="your-wf-token-secret"
$ terraform plan
```

## Argument Reference
In addition to [generic `provider` arguments](https://www.terraform.io/docs/configuration/providers.html)
(e.g. `alias` and `version`), the following arguments are supported in the Wavefront 
`provider` block:

* `address` - (Optional) this is the URL of your cluster that you access Wavefront from without the 
leading `https://` or trailing `/` (e.g. `https://longboard.wavefront.com/` becomes `longboard.wavefront.com`)

* `token` - (Optional) this is a either a Users token or Service Account token with permissions necessary 
to manage your Wavefront account. 

* `http_proxy` - (Optional) The proxy type is determined by the URL scheme. `http`, `https`, and `socks5` are supported.  
If the scheme is empty `http` is assumed.
