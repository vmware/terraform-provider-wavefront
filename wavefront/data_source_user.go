package wavefront

import (
	"fmt"
	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	idKey                  = "id"
	identifierKey          = "identifier"
	customerKey            = "customer"
	lastSuccessfulLoginKey = "last_successful_login"
	// only in getUsers? ssoIdKey  = "ssoId"
	//groupsKey = "groups"
	//userGroupsKey = "userGroups"
	//ingestionPoliciesKey = "ingestionPolicies"
	//rolesKey = "roles"
)

func dataSourceUser() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceUserRead,
		Schema: dataSourceUserSchema(),
	}
}

func dataSourceUserRead(d *schema.ResourceData, m interface{}) error {
	userClient := m.(*wavefrontClient).client.Users()
	id, ok := d.GetOk(idKey)
	if !ok {
		return fmt.Errorf("required parameter '%s' not set", idKey)
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

func dataSourceUserSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		// Query Values
		idKey: {
			Type:     schema.TypeString,
			Required: true,
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
	}
}

func userSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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
	}
}
