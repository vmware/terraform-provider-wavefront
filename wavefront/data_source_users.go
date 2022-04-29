package wavefront

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceUsers() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDefaultUserGroupRead,
		Schema: map[string]*schema.Schema{
			"users": {
				Type:     schema.TypeList,
				Computed: true,
			},
		},
	}
}
