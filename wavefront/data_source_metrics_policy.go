package wavefront

import (
	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"time"
)

const (
	updaterIdKey          = "updater_id"
	updatedEpochMillisKey = "updated_epoch_millis"
	accountsKey           = "accounts"
	tagsKey               = "tags"
	prefixesKey           = "prefixes"
	tagsAndedKey          = "tags_anded"
	accessTypeKey         = "access_type"
	userGroupsKey         = "user_groups"
	policyRulesKey        = "policy_rules"
)

func dataSourceMetricsPolicy() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceMetricsPolicyRead,
		Schema: dataSourceMetricsPolicySchema(),
	}
}

func dataSourceMetricsPolicyRead(d *schema.ResourceData, m interface{}) error {
	metrics := m.(*wavefrontClient).client.MetricsPolicyAPI()
	metricsPolicy, err := metrics.Get()
	if err != nil {
		return err
	}
	log.Printf("policy docs: %v", metricsPolicy)
	d.SetId(time.Now().UTC().String())
	if err := d.Set(policyRulesKey, flattenPolicyRules(metricsPolicy.PolicyRules)); err != nil {
		return err
	}
	if err := d.Set(customerKey, metricsPolicy.Customer); err != nil {
		return err
	}
	if err := d.Set(updaterIdKey, metricsPolicy.UpdaterId); err != nil {
		return err
	}
	if err := d.Set(updatedEpochMillisKey, metricsPolicy.UpdatedEpochMillis); err != nil {
		return err
	}
	return nil
}

func flattenPolicyRules(policy []wavefront.PolicyRule) []map[string]interface{} {
	tfMaps := make([]map[string]interface{}, len(policy))
	for i, v := range policy {
		tfMaps[i] = flattenPolicyRule(&v)
	}
	return tfMaps
}

func flattenPolicyRule(policy *wavefront.PolicyRule) map[string]interface{} {
	tfMap := make(map[string]interface{})
	tfMap[accountsKey] = policy.Accounts
	tfMap[userGroupsKey] = flattenUserGroups(policy.UserGroups)
	tfMap[rolesKey] = policy.Roles
	tfMap[nameKey] = policy.Name
	tfMap[tagsKey] = policy.Tags
	tfMap[descriptionKey] = policy.Description
	tfMap[prefixesKey] = policy.Prefixes
	tfMap[tagsAndedKey] = policy.TagsAnded
	tfMap[accessTypeKey] = policy.AccessType
	return tfMap
}

func flattenUserGroups(user []wavefront.UserGroup) []map[string]interface{} {
	tfMaps := make([]map[string]interface{}, len(user))
	for i, v := range user {
		tfMaps[i] = flattenUserGroup(&v)
	}
	return tfMaps
}

func flattenUserGroup(userG *wavefront.UserGroup) map[string]interface{} {
	tfMap := make(map[string]interface{})
	tfMap[nameKey] = userG.Name
	tfMap[idKey] = userG.ID
	tfMap[descriptionKey] = userG.Description
	return tfMap
}

func dataSourceMetricsPolicySchema() map[string]*schema.Schema {
	policyRulesSchema := policyRulesSchema()
	return map[string]*schema.Schema{
		policyRulesKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: policyRulesSchema,
			},
		},
		customerKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		updaterIdKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		updatedEpochMillisKey: {
			Type:     schema.TypeInt,
			Computed: true,
		},
	}
}

func policyRulesSchema() map[string]*schema.Schema {
	userGroupSchema := userGroupSchema()
	return map[string]*schema.Schema{
		accountsKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		userGroupsKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Resource{Schema: userGroupSchema},
		},
		rolesKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		nameKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		tagsKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		descriptionKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		prefixesKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		tagsAndedKey: {
			Type:     schema.TypeBool,
			Computed: true,
		},
		accessTypeKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func userGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		nameKey: {
			Type:     schema.TypeString,
			Required: true,
		},
		idKey: {
			Type:     schema.TypeString,
			Required: true,
		},
		descriptionKey: {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}
