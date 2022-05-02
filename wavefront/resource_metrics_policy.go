package wavefront

import (
	"fmt"
	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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

func resourceMetricsPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceMetricsPolicyUpdate,
		Read:   resourceMetricsPolicyRead,
		Update: resourceMetricsPolicyUpdate,
		Delete: resourceMetricsPolicyDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourceMetricsPolicySchema(),
	}
}

func resourceMetricsPolicyRead(d *schema.ResourceData, meta interface{}) error {
	metrics := meta.(*wavefrontClient).client.MetricsPolicyAPI()
	metricsPolicy, err := metrics.Get()
	if err != nil {
		return fmt.Errorf("error retrieving metrics policy: %d", err)
	}
	d.SetId(string(rune(metricsPolicy.UpdatedEpochMillis)))
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

func resourceMetricsPolicySchema() map[string]*schema.Schema {
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

func resourceMetricsPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	metrics := meta.(*wavefrontClient).client.MetricsPolicyAPI()
	// todo do we need to parse / validate metrics policy to give better error? Maybe add validation to tf
	customRules := d.Get(policyRulesKey).([]wavefront.PolicyRule)
	if len(customRules) < 1 {
		return fmt.Errorf("error updating Metrics Policy, no valid Policy Rules set")
	}
	newPolicyRules := &wavefront.UpdateMetricsPolicyRequest{
		PolicyRules: customRules,
	}
	updatedPolicy, err := metrics.Update(newPolicyRules)
	if err != nil {
		return fmt.Errorf("error updating metrics policy: %d", err)
	}
	d.SetId(string(rune(updatedPolicy.UpdatedEpochMillis)))

	return resourceMetricsPolicyRead(d, meta)
}

// resourceMetricsPolicyDelete reverts metrics policy to default predefined policy rule allowing access to all metrics for everyone
func resourceMetricsPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	metrics := meta.(*wavefrontClient).client.MetricsPolicyAPI()

	defaultPolicyRules := &wavefront.UpdateMetricsPolicyRequest{
		PolicyRules: []wavefront.PolicyRule{{
			//Accounts: []string{}, TODO validate unneeded
			UserGroups: []wavefront.UserGroup{{
				Name:        "Everyone",
				Description: "System group which contains all users",
			}},
			//Roles:       []string{},TODO validate unneeded
			Name: "Allow All Metrics",
			//Tags:        []string{},TODO validate unneeded
			Description: "Predefined policy rule. Allows access to all metrics (timeseries, histograms, and counters) for all accounts. If this rule is removed, all accounts can access all metrics if there are no matching blocking rules.",
			Prefixes:    []string{"*"},
			TagsAnded:   false,
			AccessType:  "ALLOW",
		}},
	}
	_, err := metrics.Update(defaultPolicyRules)
	if err != nil {
		return fmt.Errorf("error deleting custom metrics policy: %d", err)
	}
	d.SetId("")
	return nil
}
