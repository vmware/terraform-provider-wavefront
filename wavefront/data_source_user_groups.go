package wavefront

import (
	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"time"
)

func dataSourceUserGroups() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceUserGroupsRead,
		Schema: dataSourceUserGroupsSchema(),
	}
}

func dataSourceUserGroupsSchema() map[string]*schema.Schema {
	userGroupsSchema := userGroupSchema()
	return map[string]*schema.Schema{
		userGroupsListKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: userGroupsSchema,
			},
		},
	}
}

func userGroupSchema() map[string]*schema.Schema {
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

func dataSourceUserGroupsRead(d *schema.ResourceData, m interface{}) error {
	d.SetId(time.Now().UTC().String())
	var allGroups []*wavefront.UserGroup
	groupsClient := m.(*wavefrontClient).client.UserGroups()

	cont := true
	offset := 0

	for cont {
		filter := []*wavefront.SearchCondition{
			{Key: "limit", Value: string(rune(pageSize)), MatchingMethod: exactMatching},
			{Key: "offset", Value: string(rune(offset)), MatchingMethod: exactMatching},
		}

		groups, err := groupsClient.Find(filter)
		if err != nil {
			return err
		}
		allGroups = append(allGroups, groups...)

		if len(groups) < pageSize {
			cont = false
		} else {
			offset += pageSize
		}
	}

	if err := d.Set(userGroupsListKey, flattenUserGroups(allGroups)); err != nil {
		return err
	}
	return nil
}

func flattenUserGroups(users []*wavefront.UserGroup) []map[string]interface{} {
	tfMaps := make([]map[string]interface{}, len(users))
	for i, v := range users {
		tfMaps[i] = flattenUserGroup(v)
	}
	return tfMaps
}

func flattenUserGroup(group *wavefront.UserGroup) map[string]interface{} {
	tfMap := make(map[string]interface{})
	tfMap[idKey] = *group.ID
	tfMap[nameKey] = group.Name
	tfMap[descriptionKey] = group.Description
	tfMap[rolesKey] = flattenUserGroupRoles(group.Roles)
	tfMap[usersKey] = group.Users

	return tfMap
}

func flattenUserGroupRoles(roles []wavefront.Role) []string {
	var roleIds []string

	for _, r := range roles {
		roleIds = append(roleIds, r.ID)
	}
	return roleIds
}
