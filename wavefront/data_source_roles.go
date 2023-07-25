package wavefront

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/WavefrontHQ/go-wavefront-management-api/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRoles() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceRolesRead,
		Schema: dataSourceRolesSchema(),
	}
}

func dataSourceRolesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		rolesKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Resource{Schema: rolesSchema()},
		},
		limitKey: {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  100,
		},
		offsetKey: {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
	}
}

func dataSourceRolesRead(d *schema.ResourceData, m interface{}) error {
	var allRoles []*wavefront.Role

	limit := d.Get(limitKey).(int)
	offset := d.Get(offsetKey).(int)

	if err := json.Unmarshal(searchAll(limit, offset, "role", nil, nil, m), &allRoles); err != nil {
		return fmt.Errorf("Response is invalid JSON")
	}

	// Data Source ID is set to current time to always refresh
	d.SetId(time.Now().UTC().String())
	return d.Set(rolesKey, flattenRoles(allRoles))
}

func flattenRoles(roles []*wavefront.Role) []map[string]interface{} {
	tfMaps := make([]map[string]interface{}, len(roles))
	for i, v := range roles {
		tfMaps[i] = flattenRole(v)
	}
	return tfMaps
}

func flattenRole(role *wavefront.Role) map[string]interface{} {
	tfMap := make(map[string]interface{})
	tfMap[idKey] = role.ID
	tfMap[nameKey] = role.Name
	tfMap[descriptionKey] = role.Description
	tfMap[permissionsKey] = role.Permissions
	return tfMap
}

func rolesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		// Computed Values
		idKey: {
			Type:     schema.TypeString,
			Computed: true,
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
