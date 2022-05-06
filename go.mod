module github.com/vmware/terraform-provider-wavefront

require (
	github.com/WavefrontHQ/go-wavefront-management-api v1.14.0
	github.com/google/go-cmp v0.5.5
	github.com/hashicorp/go-cty v1.4.1-0.20200414143053-d3edf31b6320
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.6.1
)

replace github.com/WavefrontHQ/go-wavefront-management-api => github.com/ssnorrizaurous/go-wavefront-management-api v1.14.3-0.20220506102536-86212642d25c

go 1.16
