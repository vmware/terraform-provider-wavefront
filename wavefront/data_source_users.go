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
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		userGroupsKey: {
			Type:     schema.TypeList,
			Computed: true,
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
	tfMap[emailKey] = *user.ID
	tfMap[permissionsKey] = user.Permissions
	tfMap[userGroupsKey] = flattenUserGroupsToIds(user.Groups)
	tfMap[customerKey] = user.Customer
	tfMap[lastSuccessfulLoginKey] = int(user.LastSuccessfulLogin)

	return tfMap
}
