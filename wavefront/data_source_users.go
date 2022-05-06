package wavefront

import (
	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"time"
)

const usersKey = "users"

func dataSourceUsers() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceUsersRead,
		Schema: dataSourceUsersSchema(),
	}
}

func dataSourceUsersSchema() map[string]*schema.Schema {
	userSchema := userSchema()
	return map[string]*schema.Schema{
		usersKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: userSchema,
			},
		},
	}
}

func userSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		// Computed Values
		emailKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		permissionsKey: {
			Type:     schema.TypeList,
			Optional: true,
			ForceNew: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		userGroupsKey: {
			Type:     schema.TypeList,
			Optional: true,
			ForceNew: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		customerKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		lastSuccessfulLoginKey: {
			Type:     schema.TypeInt,
			Computed: true,
		},
	}
}

//func dataSourceUsersSchema() map[string]*schema.Schema {
//	return map[string]*schema.Schema{
//		usersKey: {
//			Type:     schema.TypeList,
//			Computed: true,
//			Elem:     &schema.Schema{Type: schema.TypeString},
//		},
//	}
//}
//
//func dataSourceUsersRead(d *schema.ResourceData, m interface{}) error {
//	userClient := m.(*wavefrontClient).client.Users()
//
//	users, err := userClient.Find(nil)
//	if err != nil {
//		return err
//	}
//	d.SetId(time.Now().UTC().String())
//	if err := d.Set(usersKey, flattenUsers(users)); err != nil {
//		return err
//	}
//	return nil
//}
//
//func flattenUsers(users []*wavefront.User) []string {
//	var ids []string
//	for _, user := range users {
//		ids = append(ids, *user.ID)
//	}
//	return ids
//}

func dataSourceUsersRead(d *schema.ResourceData, m interface{}) error {
	userClient := m.(*wavefrontClient).client.Users()

	users, err := userClient.Find(nil)
	if err != nil {
		return err
	}
	d.SetId(time.Now().UTC().String())
	if err := d.Set(usersKey, flattenUsers(users)); err != nil {
		return err
	}
	return nil
}

func flattenUsers(users []*wavefront.User) []map[string]interface{} {
	tfMaps := make([]map[string]interface{}, len(users))
	for i, v := range users {
		tfMaps[i] = flattenUser(v)
	}
	return tfMaps
}

func flattenUser(user *wavefront.User) map[string]interface{} {
	tfMap := make(map[string]interface{})
	tfMap[identifierKey] = user.ID
	tfMap[customerKey] = user.Customer
	tfMap[lastSuccessfulLoginKey] = user.LastSuccessfulLogin
	// TODO do we need this? see https://github.com/hashicorp/terraform-provider-aws/blob/611b4737168f4f0051bb63ef221f0e76f156f392/internal/service/lakeformation/data_lake_settings.go#L271
	//if user == nil {
	//	return tfMap
	//}
	return tfMap
}
