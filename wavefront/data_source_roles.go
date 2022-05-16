package wavefront

import (
	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"time"
)

func dataSourceRoles() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceRolesRead,
		Schema: dataSourceRolesSchema(),
	}
}

func dataSourceRolesSchema() map[string]*schema.Schema {
	rolesSchema := rolesSchema()
	return map[string]*schema.Schema{
		rolesKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Resource{Schema: rolesSchema},
		},
	}
}

func dataSourceRolesRead(d *schema.ResourceData, m interface{}) error {
	var allRoles []*wavefront.Role
	userClient := m.(*wavefrontClient).client.Roles()

	cont := true
	offset := 0
	for cont {
		filter := []*wavefront.SearchCondition{
			{Key: "limit", Value: string(rune(pageSize)), MatchingMethod: exactMatching},
			{Key: "offset", Value: string(rune(offset)), MatchingMethod: exactMatching},
		}

		roles, err := userClient.Find(filter)
		if err != nil {
			return err
		}
		allRoles = append(allRoles, roles...)

		if len(roles) < pageSize {
			cont = false
		} else {
			offset += pageSize
		}
	}

	d.SetId(time.Now().UTC().String())
	log.Printf("found_roles: %v", allRoles)
	if err := d.Set(rolesKey, flattenRoles(allRoles)); err != nil {
		return err
	}
	return nil
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
