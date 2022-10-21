package wavefront

import (
	"fmt"
	"time"

	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRole() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceRoleRead,
		Schema: roleSchema(),
	}
}

func roleSchema() map[string]*schema.Schema {
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
		permissionsKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
	}
}

func dataSourceRoleRead(d *schema.ResourceData, m interface{}) error {
	roleClient := m.(*wavefrontClient).client.Roles()
	id, ok := d.GetOk("id")
	if !ok {
		return fmt.Errorf("required parameter '%s' not set", idKey)
	}

	idStr := fmt.Sprintf("%s", id)
	role := wavefront.Role{ID: idStr}
	if err := roleClient.Get(&role); err != nil {
		return err
	}

	// Data Source ID is set to current time to always refresh
	d.SetId(time.Now().UTC().String())
	return setRoleAttributes(d, role)
}

func setRoleAttributes(d *schema.ResourceData, role wavefront.Role) error {
	if err := d.Set(idKey, role.ID); err != nil {
		return err
	}
	if err := d.Set(nameKey, role.Name); err != nil {
		return err
	}
	if err := d.Set(descriptionKey, role.Description); err != nil {
		return err
	}
	if err := d.Set(permissionsKey, role.Permissions); err != nil {
		return err
	}

	return nil
}
