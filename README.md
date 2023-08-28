# Aria Operations for Applications Terraform Provider

Thanks for stopping by and considering contributing to the Terraform Provider for interacting with Aria Operations for Applications!!! The Terraform Provider is used to manage resources in Aria Operations for Applications (AOA) and currently supports creating, maintaining and destroying Alerts, Alert Targets, and Dashboards for AOA resources.

We do our best to prioritize, review, and merge all requests. Generally, adding missing features (as per the [Wavefront API](https://www.wavefront.com/api/)) or bug fixes are welcomed. However, functional changes may require some discussion first.

We make use of [go-wavefront-management-api](https://github.com/WavefrontHQ/go-wavefront-management-api) to abstract the API from the provider. New features (and possibly bug fixes) will likely require updates to go-wavefront.

## Resources

Below, you will find some resources to help you get acquainted with Terraform and the concept of Terraform providers.

* This is a good [blog post](https://www.terraform.io/guides/writing-custom-terraform-providers.html?) by Hashicorp to get started.
* Looking at how existing [Providers](https://github.com/terraform-providers) work can be useful.
* This is a good [blog post](https://opencredo.com/blogs/running-a-terraform-provider-with-a-debugger/) to get some details on how to debug custom terraform provider.

## Requirements

* Go version 1.13 or higher [installed and setup correctly](https://golang.org/doc/install).
* Terraform 0.10.0 or higher (Custom providers were released at 0.10.0)
* [govendor](https://github.com/kardianos/govendor) for dependency management

## Developing the Provider
1. To use your local copy of the provider, you first need to build it.
    ```shell
    make build
    ```
   This will install the provider to `/Users/<USERNAME>/go/bin`.
2. Then add a `.terraformrc` file to your home directory tha references the locally built provider:
    ```shell
    provider_installation {
      dev_overrides {
        "vmware/wavefront" = "/Users/<USERNAME>/go/bin"
      }
    }
    ```
    You may need to run `terragform init -upgrade` to switch between local and remote versions of the plugin.

    * For more information on how the dev_overrides works, see [Development Overrides for Provider Developers](https://developer.hashicorp.com/terraform/cli/config/config-file#development-overrides-for-provider-developers).

### Linting

1. Install `golangci-lint`: https://golangci-lint.run/usage/install
1. Run
    ```shell
    make lint
    ```

### Running Tests

1. Run
    ```shell
    make test
    ```

### Acceptance Tests

Acceptance tests are run against the Wavefront API, so you'll need an account to use them. Run at your own risk.

You need to supply the `WAVEFRONT_TOKEN` and `WAVEFRONT_ADDRESS` environment variables

To run the tests run
`make testacc`

### Using a local Go cli

1. To test your local copy of the go-wavefront-management-api client library, add this to the provider's `go.mod` file:
    ```text
    replace github.com/WavefrontHQ/go-wavefront-management-api/vX vX.Y.Z => /Users/<USERNAME>/workspace/go-wavefront-management-api
    ```

### Building the provider for different architectures

You can use the local `build.sh` script to build specific versions of the binary. By default, this will create two binaries in the form of `terraform-provider-wavefront_<version_os_arch>` in the root of the repository, one for `Darwin amd64` and one for `Linux amd64`. We release darwin and linux amd64 packages on the [releases page](https://github.com/vmware/terraform-provider-wavefront/releases). If you require a different architecture, you can add it to the `build.sh` script to be built.

Now that you have a binary, you test that it was built correctly by attempting to run it and confirming that you see a message similar to the one below.

```shell
./terraform-provider-wavefront_v0.1.2_darwin_amd64
This binary is a plugin. The plugins are not meant to be executed directly.
Please execute the program that consumes these plugins, which will
load any plugins automatically.
```

If you experience any issues, please do not hesitate to submit an issue, and we will prioritize accordingly!!!

## Using the Plugin

Use the `main.tf` file to create a test config, such as the following below:

```terraform
 provider "wavefront" {
  address = "cluster.wavefront.com"
}

resource "wavefront_alert" "test_alert" {
  name                  = "Terraform Test Alert"
  target                = "test@example.com,target:alert-target-id"
  condition             = "100-ts(\"cpu.usage_idle\", environment=flamingo-int and cpu=cpu-total and service=game-service) > 80"
  display_expression    = "100-ts(\"cpu.usage_idle\", environment=flamingo-int and cpu=cpu-total and service=game-service)"
  minutes               = 5
  resolve_after_minutes = 5
  severity              = "WARN"
  tags                  = [
    "terraform",
    "flamingo"
  ]
}
```

Export your wavefront token `export WAVEFRONT_TOKEN=<token>` You could also configure the `token` in the provider section of main.tf, but best not to.

*Note*: If you are not familiar with the process for creating a token please review the following [page](https://docs.wavefront.com/wavefront_api.html)

Run `terraform init` to load your provider.

Run `terraform plan` to show the plan.

Run `terraform apply` to apply the test configuration and then check the results in Wavefront.

Update main.tf to change a value, then run plan and apply again to check that the updates work.

Run `terraform destroy` to test deleting resources.

## Contributing

Please review the [CONTRIBUTOR.md](CONTRIBUTOR.md) document for more information on contributing.
