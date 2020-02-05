package wavefront_plugin

import (
	"fmt"
	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"time"
)

func dataSourceDefaultUserGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDefaultUserGroupRead,
		Schema: map[string]*schema.Schema{
			"group_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceDefaultUserGroupRead(d *schema.ResourceData, m interface{}) error {
	userGroups := m.(*wavefrontClient).client.UserGroups()

	results, err := userGroups.Find(
		[]*wavefront.SearchCondition{
			{
				Key:            "name",
				Value:          "Everyone",
				MatchingMethod: "EXACT",
			},
		},
	)

	if err != nil {
		return fmt.Errorf("error reading Default UserGroup 'Everyone' in Wavefront, %s", err)
	}

	if len(results) != 1 {
		return fmt.Errorf("error finding default UserGroup 'Everyone' in Wavefront")
	}

	userGroup := results[0]
	d.SetId(time.Now().UTC().String())
	d.Set("group_id", userGroup.ID)

	return nil
}
