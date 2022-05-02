package wavefront

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMetricsPolicy() *schema.Resource {
	return &schema.Resource{
		Read:   resourceMetricsPolicyRead,
		Schema: resourceMetricsPolicySchema(),
	}
}
