package wavefront

import (
	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"time"
)

const (
	pageSize                = 100
	rolesKey                = "roles"
	createdEpochMillisKey   = "created_epoch_millis"
	lastUpdatedMsKey        = "last_updated_ms"
	sampleLinkedGroupsKey   = "sample_linked_groups"
	sampleLinkedAccountsKey = "sample_linked_accounts"
	linkedGroupsCountKey    = "linked_groups_count"
	linkedAccountsCountKey  = "linked_accounts_count"
	lastUpdatedAccountIdKey = "last_updated_account_id"
	nameKey                 = "name"
	permissionsKey          = "permissions"
	descriptionKey          = "description"
	exactMatching           = "EXACT"
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

	contPaging := true
	offset := 0
	for contPaging {
		filter := []*wavefront.SearchCondition{
			{Key: "limit", Value: string(rune(pageSize)), MatchingMethod: exactMatching},
			{Key: "offset", Value: string(rune(offset)), MatchingMethod: exactMatching},
		}

		roles, err := userClient.Find(filter)
		if err != nil {
			return err
		}
		for _, v := range roles {
			allRoles = append(allRoles, v)
		}

		if len(roles) < pageSize {
			contPaging = false
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
	tfMap[createdEpochMillisKey] = role.CreatedEpochMillis
	tfMap[lastUpdatedMsKey] = role.LastUpdatedMs
	//tfMap[sampleLinkedGroupsKey] = role.SampleLinkedGroups TODO impl
	if role.SampleLinkedAccounts != nil {
		tfMap[sampleLinkedAccountsKey] = *role.SampleLinkedAccounts //todo fix?
	}
	tfMap[linkedGroupsCountKey] = role.LinkedGroupsCount
	tfMap[linkedAccountsCountKey] = role.LinkedAccountsCount
	tfMap[customerKey] = role.Customer
	tfMap[lastUpdatedAccountIdKey] = role.LastUpdatedAccountId
	tfMap[nameKey] = role.Name
	tfMap[permissionsKey] = role.Permissions
	tfMap[descriptionKey] = role.Description
	return tfMap
}

//func mapStrList(strs []*string) []string {
//	var mapped []string
//	for _, v := range strs {
//		mapped = append(mapped, *v)
//	}
//	return mapped
//}

func rolesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		// Computed Values
		idKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		createdEpochMillisKey: {
			Type:     schema.TypeInt,
			Computed: true,
		},
		lastUpdatedMsKey: {
			Type:     schema.TypeInt,
			Computed: true,
		},
		//SampleLinkedGroupsKey: { //TODO look into existing impl
		//	Type:     schema.TypeList,
		//	Computed: true,
		//	Elem:
		//},
		sampleLinkedAccountsKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		linkedGroupsCountKey: {
			Type:     schema.TypeInt,
			Computed: true,
		},
		linkedAccountsCountKey: {
			Type:     schema.TypeInt,
			Computed: true,
		},
		customerKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		lastUpdatedAccountIdKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		nameKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		permissionsKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		descriptionKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}
