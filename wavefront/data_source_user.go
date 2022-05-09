package wavefront

import (
	"fmt"
	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"time"
)

const (
	emailKey               = "email"
	idKey                  = "id"
	customerKey            = "customer"
	lastSuccessfulLoginKey = "last_successful_login"
)

func dataSourceUser() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceUserRead,
		Schema: dataSourceUserSchema(),
	}
}

func dataSourceUserSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		// Query Values
		emailKey: {
			Type:     schema.TypeString,
			Required: true,
		},
		// Computed Values
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

func dataSourceUserRead(d *schema.ResourceData, m interface{}) error {
	userClient := m.(*wavefrontClient).client.Users()
	id, ok := d.GetOk(emailKey)
	if !ok {
		return fmt.Errorf("required parameter '%s' not set", emailKey)
	}
	idStr := fmt.Sprintf("%s", id)
	user := wavefront.User{ID: &idStr}
	if err := userClient.Get(&user); err != nil {
		return err
	}

	d.SetId(time.Now().UTC().String())
	return setUserAttributes(d, user)
}

func setUserAttributes(d *schema.ResourceData, user wavefront.User) error {
	if err := d.Set(emailKey, *user.ID); err != nil {
		return err
	}
	if err := d.Set(permissionsKey, user.Permissions); err != nil {
		return err
	}
	if err := d.Set(userGroupsKey, flattenUserGroupsToIds(user.Groups)); err != nil {
		return err
	}
	if err := d.Set(customerKey, user.Customer); err != nil {
		return err
	}
	if err := d.Set(lastSuccessfulLoginKey, int(user.LastSuccessfulLogin)); err != nil {
		return err
	}
	return nil
}

// flattenUserGroups extracts user_group Ids from list of user_group objects
func flattenUserGroupsToIds(groups wavefront.UserGroupsWrapper) []string {
	var ids []string
	for _, group := range groups.UserGroups {
		ids = append(ids, *group.ID)
	}
	return ids
}
