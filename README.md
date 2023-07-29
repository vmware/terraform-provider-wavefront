
# Wavefront Terraform Provider

Thanks for stopping by and considering contributing to the Wavefront Terraform Provider!!! The Terraform Provider is used to manage resources in Wavefront and currently supports creating, maintaining and destroying Alerts, Alert Targets and Dashboards Wavefront resources.

We do our best to prioritize, review, and merge all requests. Generally, adding missing features (as per the [Wavefront API](https://www.wavefront.com/api/) or bug fixes are welcomed. However, functional changes may require some discussion first.

We make use of [go-wavefront-management-api](https://github.com/WavefrontHQ/go-wavefront-management-api) to abstract the API from the provider. New features (and possibly bug fixes) will likely require updates to go-wavefront.

## Resources

Below you will find some resources to help you get acquainted with Terraform and the concept of Terraform providers.

* This is a good [blog post](https://www.terraform.io/guides/writing-custom-terraform-providers.html?) by Hashicorp to get started.
* Looking at how existing [Providers](https://github.com/terraform-providers) work can be useful.
* This is a good [blog post](https://opencredo.com/blogs/running-a-terraform-provider-with-a-debugger/) to get some details on how to debug custom terraform provider.

## Requirements
* Go version 1.13 or higher
* Terraform 0.10.0 or higher (Custom providers were released at 0.10.0)
* [govendor](https://github.com/kardianos/govendor) for dependency management


## Installing the Plugin
First ensure you have Go [installed and setup correctly](https://golang.org/doc/install).

Then locally fetch your forked repo - [repository](https://github.com/vmware/terraform-provider-wavefront)
`go get github.com/<your_account>/terraform-provider-wavefront`

*Note*: If you experience the following error message:
```
module declares its path as: github.com/vmware/terraform-provider-wavefront
but was required as: github.com/McCoyAle/terraform-provider-wavefront
```
This could be due to the [go.mod](https://github.com/vmware/terraform-provider-wavefront/blob/master/go.mod) file import reading the repository as VMware. You may need to update this to the name of your local repository. But do not submit this change upstream.

Next, you'll need to use the local `build.sh` script to build the current version binary. This will create two binaries in the form of terraform-provider-wavefront_version_os_arch in the root of the repository, one for Darwin amd64 and one for Linux amd64. We release darwin and linux amd64 packages on the [releases page](https://github.com/vmware/terraform-provider-wavefront/releases). If you require a different architecture you will need to build the plugin from source and then remove the `_os_arch` from the end of the file name and place it in `~/.terraform.d/plugins` which is where `terraform init` will look for plugins.

Valid provider filenames are `terraform-provider-NAME_X.X.X` or `terraform-provider-NAME_vX.X.X`

*Note*: If you're using a different operating system or architecture then you will need to update the build step of the makefile to also [build a binary for your OS and architecture](https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04).

It is also important to know you can also utilize `go build` to build the binaries as well. However, using the `build.sh` script will append the appropriate version identified in the [version](https://github.com/vmware/terraform-provider-wavefront/blob/master/version) directory.

Now that you have a binary you should attempt to run it and expect to see a message similar to the one below.

```
./terraform-provider-wavefront_v0.1.2_darwin_amd64
This binary is a plugin. The plugins are not meant to be executed directly.
Please execute the program that consumes these plugins, which will
load any plugins automatically.
```

If you experience any issues, please do not hesitate to submit an issue and we will prioritize accordingly!!!

### Running the Plugin

Use the main.tf file to create a test config, such as the following below:

```
 provider "wavefront" {
   address = "cluster.wavefront.com"
 }

 resource "wavefront_alert" "test_alert" {
   name = "Terraform Test Alert"
   target = "test@example.com,target:alert-target-id"
   condition = "100-ts(\"cpu.usage_idle\", environment=flamingo-int and cpu=cpu-total and service=game-service) > 80"
   display_expression = "100-ts(\"cpu.usage_idle\", environment=flamingo-int and cpu=cpu-total and service=game-service)"
   minutes = 5
   resolve_after_minutes = 5
   severity = "WARN"
   tags = [
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

Update main.tf to change a value, the run plan and apply again to check that updates work.

Run `terraform destroy` to test deleting resources.

### Acceptance Tests
Acceptance tests are run against the Wavefront API so you'll need an account to use. Run at your own risk.

You need to supply the `WAVEFRONT_TOKEN` and `WAVEFRONT_ADDRESS` environment variables

To run the tests run
`make testacc`

## Contributing

Please review the [CONTRIBUTOR.md](CONTRIBUTOR.md) document for more information on contributing.
