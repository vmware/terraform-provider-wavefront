package wavefront

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
	d.SetId(strconv.Itoa(metricsPolicy.UpdatedEpochMillis))
	if err := d.Set(policyRulesKey, flattenPolicyRules(metricsPolicy.PolicyRules)); err != nil {
		return err
	}
	if err := d.Set(customerKey, metricsPolicy.Customer); err != nil {
		return err
	}
	if err := d.Set(updaterIDKey, metricsPolicy.UpdaterId); err != nil {
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
	tfMap[accountsKey] = flattenAccounts(policy.Accounts)
	tfMap[userGroupsKey] = flattenPolicyUserGroups(policy.UserGroups)
	tfMap[roleIdsTagKey] = flattenPolicyRole(policy.Roles)
	tfMap[nameKey] = policy.Name
	tfMap[tagsKey] = flattenPolicyTags(policy.Tags)
	tfMap[descriptionKey] = policy.Description
	tfMap[prefixesKey] = policy.Prefixes
	tfMap[tagsAndedKey] = policy.TagsAnded
	tfMap[accessTypeKey] = policy.AccessType
	return tfMap
}

func flattenAccounts(accounts []wavefront.PolicyUser) []string {
	var accountIds []string
	for _, v := range accounts {
		accountIds = append(accountIds, v.ID)
	}
	return accountIds
}

func flattenPolicyUserGroups(users []wavefront.PolicyUserGroup) []string {
	var groupIds []string
	for _, v := range users {
		groupIds = append(groupIds, v.ID)
	}
	return groupIds
}

func flattenPolicyRole(roles []wavefront.Role) []string {
	var roleIds []string
	for _, v := range roles {
		roleIds = append(roleIds, v.ID)
	}
	return roleIds
}

func flattenPolicyTags(tags []wavefront.PolicyTag) []map[string]interface{} {
	tfMaps := make([]map[string]interface{}, 0)
	for _, tag := range tags {
		tfMap := make(map[string]interface{})
		tfMap[policyTagKey] = tag.Key
		tfMap[policyTagValue] = tag.Value
		tfMaps = append(tfMaps, tfMap)
	}
	return tfMaps
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
		updaterIDKey: {
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
			Optional: true,
			ForceNew: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		roleIdsTagKey: {
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
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					policyTagKey: {
						Type:     schema.TypeString,
						Required: true,
						ForceNew: true,
					},
					policyTagValue: {
						Type:     schema.TypeString,
						Required: true,
						ForceNew: true,
					},
				},
			},
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
			Type:             schema.TypeString,
			Required:         true,
			ForceNew:         true,
			ValidateDiagFunc: validateAccessTypeVal,
		},
	}
}

func validateAccessTypeVal(v interface{}, p cty.Path) diag.Diagnostics {
	var diags diag.Diagnostics

	accessType := fmt.Sprintf("%v", v)
	switch accessType {
	case "BLOCK", "ALLOW":
	default:
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("%s is an invalid access_type", accessType),
			Detail:   "access_type must be either 'BLOCK' or 'ALLOW'",
		})
	}
	return diags
}

func resourceMetricsPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	metrics := meta.(*wavefrontClient).client.MetricsPolicyAPI()
	rawPolicy := d.Get(policyRulesKey)
	policy, err := parsePolicyRules(rawPolicy)
	if err != nil {
		return err
	}
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
	d.SetId(strconv.Itoa(updatedPolicy.UpdatedEpochMillis))

	return resourceMetricsPolicyRead(d, meta)
}

func parsePolicyRules(raw interface{}) ([]wavefront.PolicyRuleRequest, error) {
	var rules []wavefront.PolicyRuleRequest

	rawArr := raw.([]interface{})
	for _, r := range rawArr {
		rule := r.(map[string]interface{})

		accountIds := parseStrArr(rule[accountsKey])
		userGroupIds := parseStrArr(rule[userGroupsKey])
		roleIds := parseStrArr(rule[roleIdsTagKey])

		if len(accountIds)+len(userGroupIds)+len(roleIds) < 1 {
			return nil, errors.New("policy_rule must have at least one associated account, user group, or role")
		}

		newRule := wavefront.PolicyRuleRequest{
			AccountIds:   accountIds,
			UserGroupIds: userGroupIds,
			RoleIds:      roleIds,
			Name:         rule[nameKey].(string),
			Tags:         parsePolicyTagsArr(rule[tagsKey]),
			Description:  rule[descriptionKey].(string),
			Prefixes:     parseStrArr(rule[prefixesKey]),
			TagsAnded:    rule[tagsAndedKey].(bool),
			AccessType:   rule[accessTypeKey].(string),
		}

		rules = append(rules, newRule)
	}
	return rules, nil
}

func parsePolicyTagsArr(raw interface{}) []wavefront.PolicyTag {
	var arr []wavefront.PolicyTag
	if raw != nil && len(raw.([]interface{})) > 0 {
		for _, v := range raw.([]interface{}) {
			kv := v.(map[string]interface{})
			key := kv[policyTagKey].(string)
			value := kv[policyTagValue].(string)
			arr = append(arr, wavefront.PolicyTag{
				Key:   key,
				Value: value,
			})
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
