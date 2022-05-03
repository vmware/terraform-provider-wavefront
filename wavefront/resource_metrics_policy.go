package wavefront

import (
	"fmt"
	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

const (
	updaterIdKey          = "updater_id"
	updatedEpochMillisKey = "updated_epoch_millis"
	accountsKey           = "accounts"
	tagsKey               = "tags"
	prefixesKey           = "prefixes"
	tagsAndedKey          = "tags_anded"
	accessTypeKey         = "access_type"
	userGroupsKey         = "user_group_ids"
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

func flattenUserGroups(user []wavefront.UserGroup) []string {
	var groupIds []string
	for _, v := range user {
		groupIds = append(groupIds, *v.ID)
	}
	return groupIds
}

func resourceMetricsPolicySchema() map[string]*schema.Schema {
	policyRulesSchema := policyRulesSchema()
	return map[string]*schema.Schema{
		// User specified value
		policyRulesKey: {
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Resource{
				Schema: policyRulesSchema,
			},
		},
		// Computed Values
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
	return map[string]*schema.Schema{
		accountsKey: {
			Type:     schema.TypeList,
			Optional: true,
			ForceNew: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		userGroupsKey: {
			Type:     schema.TypeList,
			Required: true,
			ForceNew: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		rolesKey: {
			Type:     schema.TypeList,
			Optional: true,
			ForceNew: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		nameKey: {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		tagsKey: {
			Type:     schema.TypeList,
			Optional: true,
			ForceNew: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		descriptionKey: {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		prefixesKey: {
			Type:     schema.TypeList,
			Required: true,
			ForceNew: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		tagsAndedKey: {
			Type:     schema.TypeBool,
			Required: true,
			ForceNew: true,
		},
		accessTypeKey: {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
	}
}

func resourceMetricsPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	metrics := meta.(*wavefrontClient).client.MetricsPolicyAPI()
	rawPolicy := d.Get(policyRulesKey)
	log.Printf("recieved: %v", rawPolicy)
	policy := parsePolicyRules(rawPolicy)
	if len(policy) < 1 {
		return fmt.Errorf("error updating Metrics Policy, no valid Policy Rules set")
	}
	newPolicyRules := &wavefront.UpdateMetricsPolicyRequest{
		PolicyRules: policy,
	}
	updatedPolicy, err := metrics.Update(newPolicyRules)
	if err != nil {
		return fmt.Errorf("error updating metrics policy: %v", err)
	}
	d.SetId(string(rune(updatedPolicy.UpdatedEpochMillis)))

	return resourceMetricsPolicyRead(d, meta)
}

func parsePolicyRules(raw interface{}) []wavefront.PolicyRuleRequest {
	var rules []wavefront.PolicyRuleRequest

	rawArr := raw.([]interface{})
	for _, r := range rawArr {
		rule := r.(map[string]interface{})
		log.Printf("rule: %v", rule)

		newRule := wavefront.PolicyRuleRequest{
			Accounts:     parseStrArr(rule[accountsKey]),
			UserGroupIds: parseStrArr(rule[userGroupsKey]),
			Roles:        parseStrArr(rule[rolesKey]),
			Name:         rule[nameKey].(string),
			Tags:         parseStrArr(rule[tagsKey]),
			Description:  rule[descriptionKey].(string),
			Prefixes:     parseStrArr(rule[prefixesKey]),
			TagsAnded:    rule[tagsAndedKey].(bool),
			AccessType:   rule[accessTypeKey].(string),
		}

		rules = append(rules, newRule)
	}
	return rules
}

func parseStrArr(raw interface{}) []string {
	var arr []string
	if len(raw.([]interface{})) > 0 {
		for _, v := range raw.([]interface{}) {
			arr = append(arr, v.(string))
		}

	}
	return arr
}

// resourceMetricsPolicyDelete reverts metrics policy to default predefined policy rule allowing access to all metrics for everyone
func resourceMetricsPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	// needed to lookup default 'everyone' group assignment
	groups := meta.(*wavefrontClient).client.UserGroups()
	groupResults, err := groups.Find(
		[]*wavefront.SearchCondition{
			{
				Key:            "name",
				Value:          "Everyone",
				MatchingMethod: "EXACT",
			},
		},
	)
	if err != nil {
		return fmt.Errorf("error reading Default UserGroup 'Everyone' in Wavefront, %s", err)
	}

	if len(groupResults) != 1 {
		return fmt.Errorf("error finding default UserGroup 'Everyone' in Wavefront")
	}

	defaultGroup := groupResults[0]

	metrics := meta.(*wavefrontClient).client.MetricsPolicyAPI()

	defaultPolicyRules := &wavefront.UpdateMetricsPolicyRequest{
		PolicyRules: []wavefront.PolicyRuleRequest{{
			UserGroupIds: []string{*defaultGroup.ID},
			Name:         "Allow All Metrics",
			Description:  "Predefined policy rule. Allows access to all metrics (timeseries, histograms, and counters) for all accounts. If this rule is removed, all accounts can access all metrics if there are no matching blocking rules.",
			Prefixes:     []string{"*"},
			TagsAnded:    false,
			AccessType:   "ALLOW",
		}},
	}
	_, err = metrics.Update(defaultPolicyRules)
	if err != nil {
		return fmt.Errorf("error deleting custom metrics policy: %d", err)
	}
	d.SetId("")
	return nil
}
