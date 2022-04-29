package wavefront

import (
	"fmt"
	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	userIdKey              = "id"
	identifierKey          = "identifier"
	customerKey            = "customer"
	lastSuccessfulLoginKey = "lastSuccessfulLogin"
	//groupsKey = "groups"
	//userGroupsKey = "userGroups"
	//ingestionPoliciesKey = "ingestionPolicies"
	//rolesKey = "roles"
)

func DataSourceUser() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceUserRead,
		Schema: map[string]*schema.Schema{
			// Query Parameters
			userIdKey: {
				Type:     schema.TypeString,
				Optional: false,
			},
			// Computed Values
			identifierKey: {
				Type:     schema.TypeString,
				Computed: true,
			},
			customerKey: {
				Type:     schema.TypeString,
				Computed: true,
			},
			lastSuccessfulLoginKey: {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceUserRead(d *schema.ResourceData, m interface{}) error {
	userClient := m.(*wavefrontClient).client.Users()
	id, ok := d.GetOk(userIdKey)
	if !ok {
		return fmt.Errorf("required parameter '%s' not set", userIdKey)
	}
	idStr := fmt.Sprintf("%s", id)
	user := wavefront.User{ID: &idStr}
	if err := userClient.Get(&user); err != nil {
		return err
	}

	return userAttributes(d, user)
}

func userAttributes(d *schema.ResourceData, user wavefront.User) error {
	d.SetId(*user.ID)
	if err := d.Set(identifierKey, *user.ID); err != nil {
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
