package wavefront

import (
	"fmt"
	"time"

	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceUserGroup() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceUserGroupRead,
		Schema: userGroupNewSchema(),
	}
}

func userGroupNewSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		// Computed Values
		idKey: {
			Type:     schema.TypeString,
			Required: true,
		},
		nameKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		descriptionKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		rolesKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		usersKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
	}
}

func dataSourceUserGroupRead(d *schema.ResourceData, m interface{}) error {
	userGroupClient := m.(*wavefrontClient).client.UserGroups()
	id, ok := d.GetOk("id")
	if !ok {
		return fmt.Errorf("required parameter '%s' not set", idKey)
	}

	idStr := fmt.Sprintf("%s", id)
	userGroup := wavefront.UserGroup{ID: &idStr}
	if err := userGroupClient.Get(&userGroup); err != nil {
		return err
	}

	// Data Source ID is set to current time to always refresh
	d.SetId(time.Now().UTC().String())
	return setUserGroupAttributes(d, userGroup)
}

func setUserGroupAttributes(d *schema.ResourceData, userGroup wavefront.UserGroup) error {
	if err := d.Set(idKey, userGroup.ID); err != nil {
		return err
	}
	if err := d.Set(nameKey, userGroup.Name); err != nil {
		return err
	}
	if err := d.Set(descriptionKey, userGroup.Description); err != nil {
		return err
	}
	if err := d.Set(rolesKey, flattenUserGroupRoles(userGroup.Roles)); err != nil {
		return err
	}
	return d.Set(usersKey, userGroup.Users)
}
