# Aria Operations for Applications Terraform Provider

Thanks for stopping by and considering contributing to the Terraform Provider for interacting with Aria Operations for Applications! This Provider is used to manage resources in Aria Operations for Applications (AOA). 

We welcome contributors to this provider! Please see our [CONTRIBUTING.md](./CONTRIBUTING.md) for more details on contributing.

## Requirements

* Go version 1.13 or higher [installed and setup correctly](https://golang.org/doc/install).
* Terraform 0.10.0 or higher (Custom providers were released at 0.10.0).
* Install [`golangci-lint`](https://golangci-lint.run/usage/install).

## Provider Development and Installation

In this section, you'll find information about how to develop, build and install your custom terraform provider locally.
The examples will assume `MacOS`, but will provide instructions for how to build for other platforms.

### Building the Provider

1. To use your local copy of the provider, you first need to build it.
    ```shell
    make build
    ```
   This will install the provider to `/Users/<USERNAME>/go/bin`.
1. Then add a `.terraformrc` file to your home directory that references the locally built provider:
    ```shell
    provider_installation {
      dev_overrides {
        "vmware/wavefront" = "/Users/<USERNAME>/go/bin"
      }
    
    direct {}
    }
    ```
    You may need to run `terraform init -upgrade` to switch between local and remote versions of the plugin.

    * For more information on how the dev_overrides works, see [Development Overrides for Provider Developers](https://developer.hashicorp.com/terraform/cli/config/config-file#development-overrides-for-provider-developers).
1. Now, when running `terraform` commands - such as `plan` or `apply` - your local copy of the provider will be used. If everything has worked as expected, you should see output similar to the following when running `terraform` commands:
```text
╷
│ Warning: Provider development overrides are in effect
│
│ The following provider development overrides are set in the CLI configuration:
│  - vmware/wavefront in /Users/<USERNAME>/go/bin
│
│ The behavior may therefore not match any released version of the provider and applying changes may cause the state to become incompatible with published releases.
╵
```

### Unit Tests

Unit Tests should be written where required and can be run from `make test`. The core functionality of the provider (Read, Create, Update, Delete and Import of resources is best tested via integration tests), but any supporting function should be unit tested.

`make test` does not run acceptance tests.

### Acceptance Tests

Acceptance Tests are required for the Read, Create, Update, Delete and Import of resources. Acceptance tests are run against the Wavefront API, so you'll need an account to use them. Run at your own risk.

The `WAVEFRONT_ADDRESS` and `WAVEFRONT_TOKEN` environment variables are required in order for the tests to run.

```shell
export WAVEFRONT_ADDRESS=<your-account>.wavefront.com
export WAVEFRONT_TOKEN=<your-wavefront-token>

make testacc
```

### Linting and Formatting

1. Run
    ```shell
    make lint
    make fmt
    ```

### Using a local copy of the Wavefront Golang Client Library

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

Once you have the plugin you should remove the `_os_arch` from the end of the file name.

## Using the Plugin

To see instructions for using this provider to manage resources in Aria Operations for Applications, please see [the documentation in the terraform registry](https://registry.terraform.io/providers/vmware/wavefront/latest/docs).
