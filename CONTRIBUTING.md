# Contributing

Thanks for stopping by and considering contributing to the Wavefront Terraform Provider!!! We do our best to prioritze, review, and merge all requests. Generally, adding missing features (as per the [Wavefront API](https://www.wavefront.com/api/) or bug fixes are welcomed. However, functional changes may require some discussion first.

We make use of [go-wavefront-management-api](https://github.com/WavefrontHQ/go-wavefront-management-api) to abstract the API from the provider. New features (and possibly bug fixes) will likely require updates to go-wavefront.

### Steps

1. Open an Issue - Be descriptive as possible!
2. Fork the VMware Terraform Repository
3. Make your changes to your forked repository
4. Submit a [Pull Request](https://help.github.com/articles/creating-a-pull-request-from-a-fork/)

## Resources

* This is a good [blog post](https://www.terraform.io/guides/writing-custom-terraform-providers.html?) by Hashicorp to get started.
* Looking at how existing [Providers](https://github.com/terraform-providers) work can be useful.
* This is a good [blog post](https://opencredo.com/blogs/running-a-terraform-provider-with-a-debugger/) to get some details on how to debug custom terraform provider.

## Setup

First ensure you have Go [installed and setup correctly](https://golang.org/doc/install).

Then locally fetch your forked repo - [repository](https://github.com/vmware/terraform-provider-wavefront)
`go get github.com/<your_account>/terraform-provider-wavefront`

*Note*: If you experience the following error message:
```
module declares its path as: github.com/vmware/terraform-provider-wavefront
but was required as: github.com/McCoyAle/terraform-provider-wavefront
```
This could be due to the [go.mod](https://github.com/vmware/terraform-provider-wavefront/blob/master/go.mod) file import reading the repository as VMware. You may need to update this to the name of your local repository. But do not submit this change upstream. 

Next, you'll need to use the local `build.sh` script to build the current version binary. This will create two binaries in the form of terraform-provider-wavefront_version_os_arch in the root of the repository, one for Darwin amd64 and one for Linux amd64. 

*Note*: If you're using a different operating system or architecture then you will need to update the build step of the makefile to also [build a binary for your OS and architecture](https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04).

It is also important to know you can also utilize `go build` to build the binaries as well. However, using the `build.sh` script will append the appropriate version identified in the [version](https://github.com/vmware/terraform-provider-wavefront/blob/master/version) directory. 

Now that you have a binary you should attempt to run it and expect to see a message similar to the one below.

```
./terraform-provider-wavefront_v0.1.2_darwin_amd64
This binary is a plugin. These are not meant to be executed directly.
Please execute the program that consumes these plugins, which will
load any plugins automatically
```

## Versioning

We use [Semantic Versioning](http://semver.org/) on this project. The version is located inside the `version` file, in the root of the repository, in the format `vMajor.Minor.Patch`, update this version as required.

## Dependencies

We use [Go Vendor](https://github.com/kardianos/govendor) to manage dependencies. Any dependencies that you add or update should be reflected within vendor and pushed along with your changes.

## Unit Testing

Unit Tests should be written where required and can be run from `make test`. The core functionality of the provider (Read, Create, Update, Delete and Import of resources is best tested via integration tests) but any supporting function should be unit tested.

`make test` does not run acceptance tests.

## Acceptance Testing

Acceptance Tests are required for the Read, Create, Update, Delete and Import of resources. To run the acceptance tests you should have access to a Wavefront account.

The `WAVEFRONT_ADDRESS` and `WAVEFRONT_TOKEN` environment variables are required in order for the tests to run.

```
export WAVEFRONT_ADDRESS=<your-account>.wavefront.com
export WAVEFRONT_TOKEN=<your-wavefront-token>

make testacc
```

## Formatting

`make fmt` will ensure that your code is correctly formatted. The build, test and acceptance stages of make will also run a fmt check and let you know if you need to run `make fmt`

## Running Locally

You can test the changes locally by building the plugin `make build`. Once you have the plugin you should remove the `_os_arch` from the end of the file name and place it in `~/.terraform.d/plugins` which is where `terraform init` will look for plugins.

Valid provider file names are `terraform-provider-NAME_X.X.X` or `terraform-provider-NAME_vX.X.X`
