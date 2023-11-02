# Contributing

We welcome contributors to this Terraform Provider, and we'll do our best to review and merge all requests.
Adding missing features as per the [Wavefront API](https://www.wavefront.com/api/) or bug fixes will be welcomed.
Any functional changes will require discussion.

We make use of [go-wavefront-management-api](https://github.com/WavefrontHQ/go-wavefront-management-api) to abstract the API from the provider.
New features and bug fixes will likely require updates to go-wavefront client.

## Opening Issues
If you encounter a bug or you are making a feature request, please open an issue in this repo.

## Making Pull Requests
1. Fork the repository
1. Create a new branch for your change
1. Make your changes and submit a [Pull Request](https://help.github.com/articles/creating-a-pull-request-from-a-fork/)

Before submitting a pull request, please ensure that unit tests pass. Refer to the [README.md](./README.md) for instructions on running unit tests.

We will review your pull request and provide feedback.

## Versioning

We use [Semantic Versioning](http://semver.org/) on this project. The version is located inside the `version` file, in the root of the repository, in the format `vMajor.Minor.Patch`. Update this version as required.

## Creating a new Release

1. Update the CHANGELOG.md
1. Update the `version` file to vX.Y.Z
1. Commit changes
1. Make a new tag (`git tag vX.Y.Z`)
1. Push changes / tag vX.Y.Z (`git push --tags`)
    1. A GitHub Action should generate the necessary binaries and a Release on GitHub.
        1. Binaries need to be complete before a TF Registry Resync will work.
1. Update the GitHub Release text to match past releases. (aka version as title and summary as body)
1. Ask the HashiCorp team to Resync the provider by sending an email to support@hashicorp.com, or by using https://support.hashicorp.com/hc/en-us


## Helpful Resources for Provider Development

* This is a good [blog post](https://www.terraform.io/guides/writing-custom-terraform-providers.html?) by Hashicorp to get started.
* Looking at how existing [Providers](https://github.com/terraform-providers) work can be useful.
* This is a good [blog post](https://opencredo.com/blogs/running-a-terraform-provider-with-a-debugger/) to get some details on how to debug custom terraform provider.

