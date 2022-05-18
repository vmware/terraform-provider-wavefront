package wavefront

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMetricsPolicy() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceMetricsPolicyRead,
		Schema: dataSourceMetricsPolicySchema(),
	}
}

func dataSourceMetricsPolicyRead(d *schema.ResourceData, meta interface{}) error {
	metrics := meta.(*wavefrontClient).client.MetricsPolicyAPI()
	metricsPolicy, err := metrics.Get()
	if err != nil {
		return fmt.Errorf("error retrieving metrics policy: %d", err)
	}
	// Data Source ID is set to current time to always refresh
	d.SetId(time.Now().UTC().String())
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

func dataSourceMetricsPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		// Computed Values
		policyRulesKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: dataSourcePolicyRulesSchema(),
			},
		},
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

func dataSourcePolicyRulesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		accountsKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		userGroupsKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		roleIdsTagKey: {
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
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					policyTagKey: {
						Type:     schema.TypeString,
						Computed: true,
					},
					policyTagValue: {
						Type:     schema.TypeString,
						Computed: true,
					},
				},
			},
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
