package wavefront

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccWavefrontPolicy_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontMetricsPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWavefrontMetricsPolicyBasic(),
				Check:  testAccCheckWavefrontCustomMetricsPolicy(testAccGetBasicPolicy),
			},
		},
	})
}

func TestAccWavefrontPolicy_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontMetricsPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWavefrontMetricsPolicyUpdate(),
				Check:  testAccCheckWavefrontCustomMetricsPolicy(testAccGetUpdatedPolicy),
			},
		},
	})
}

func TestAccWavefrontPolicy_BadAccessType(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontMetricsPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckWavefrontMetricsPolicyBadAccessPolicy(),
				Check:       testAccCheckWavefrontCustomMetricsPolicy(testAccGetUpdatedPolicy),
				ExpectError: regexp.MustCompile("access_type must be either 'BLOCK' or 'ALLOW'"),
			},
		},
	})
}

func TestAccWavefrontPolicy_NoSelector(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontMetricsPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckWavefrontMetricsPolicyNoSelectorSpecifiedPolicy(),
				Check:       testAccCheckWavefrontCustomMetricsPolicy(testAccGetUpdatedPolicy),
				ExpectError: regexp.MustCompile("policy_rule must have at least one associated account, user group, or role"),
			},
		},
	})
}

func testAccGetBasicPolicy() ([]wavefront.PolicyRule, error) {
	userGroups := testAccProvider.Meta().(*wavefrontClient).client.UserGroups()
	results, err := userGroups.Find(
		[]*wavefront.SearchCondition{
			{
				Key:            "name",
				Value:          "Everyone",
				MatchingMethod: "EXACT",
			},
		},
	)
	if err != nil {
		return nil, err
	}
	if len(results) != 1 {
		return nil, fmt.Errorf("unable to get default group, lookup return a match of %d groups", len(results))
	}
	defaultUser := results[0]
	return []wavefront.PolicyRule{{
		Accounts: []wavefront.PolicyUser{},
		UserGroups: []wavefront.PolicyUserGroup{{
			ID:          *defaultUser.ID,
			Name:        defaultUser.Name,
			Description: defaultUser.Description,
		}},
		Roles:       []wavefront.Role{},
		Name:        "Allow All Metrics",
		Tags:        []wavefront.PolicyTag{},
		Description: "Predefined policy rule. Allows access to all metrics (timeseries, histograms, and counters) for all accounts. If this rule is removed, all accounts can access all metrics if there are no matching blocking rules.",
		Prefixes:    []string{"*"},
		TagsAnded:   false,
		AccessType:  "ALLOW",
	}}, nil
}

func testAccGetUpdatedPolicy() ([]wavefront.PolicyRule, error) {
	users := testAccProvider.Meta().(*wavefrontClient).client.Users()
	userResp, err := users.Find(
		[]*wavefront.SearchCondition{
			{
				Key:            "id",
				Value:          "example@example.com",
				MatchingMethod: "EXACT",
			},
		},
	)
	if err != nil {
		return nil, err
	}
	if len(userResp) != 1 {
		return nil, fmt.Errorf("unable to get example user, lookup return a match of %d users", len(userResp))
	}
	testUser := userResp[0]

	roles := testAccProvider.Meta().(*wavefrontClient).client.Roles()
	rolesResp, err := roles.Find(
		[]*wavefront.SearchCondition{
			{
				Key:            "name",
				Value:          "test-role",
				MatchingMethod: "EXACT",
			},
		},
	)
	if err != nil {
		return nil, err
	}
	if len(rolesResp) != 1 {
		return nil, fmt.Errorf("unable to get test role, lookup return a match of %d roles", len(rolesResp))
	}
	testRole := rolesResp[0]

	basicArr, err := testAccGetBasicPolicy()
	if err != nil {
		return nil, err
	}

	return []wavefront.PolicyRule{
		{
			Accounts:   []wavefront.PolicyUser{},
			UserGroups: []wavefront.PolicyUserGroup{},
			Roles: []wavefront.Role{{
				ID:          testRole.ID,
				Name:        testRole.Name,
				Description: testRole.Description,
			}},
			Name:        "Deny example role metrics",
			Tags:        []wavefront.PolicyTag{},
			Description: "deny example role test",
			Prefixes:    []string{"example.api.*"},
			TagsAnded:   false,
			AccessType:  "BLOCK",
		},
		{
			Accounts: []wavefront.PolicyUser{{
				ID:   *testUser.ID,
				Name: *testUser.ID,
			}},
			UserGroups: []wavefront.PolicyUserGroup{},
			Roles:      []wavefront.Role{},
			Name:       "Deny example user metrics",
			Tags: []wavefront.PolicyTag{
				{Key: "env", Value: "prod"},
				{Key: "region", Value: "us-east-1"},
			},
			Description: "deny example user test",
			Prefixes:    []string{"example.system.*"},
			TagsAnded:   true,
			AccessType:  "BLOCK",
		},
		basicArr[0],
	}, nil
}

func testAccCheckWavefrontMetricsPolicyDestroy(s *terraform.State) error {
	policy := testAccProvider.Meta().(*wavefrontClient).client.MetricsPolicyAPI()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "wavefront_metrics_policy" {
			continue
		}

		results, err := policy.Get()

		if err != nil {
			return fmt.Errorf("error finding Wavefront Metrics Policy. %s", err)
		}
		if len(results.PolicyRules) != 1 || results.PolicyRules[0].Name != "Allow All Metrics" {
			return fmt.Errorf("default policy has not been restored")
		}
	}

	return nil
}

func testAccCheckWavefrontCustomMetricsPolicy(rulesProvider func() ([]wavefront.PolicyRule, error)) resource.TestCheckFunc {
	n := "wavefront_metrics_policy.main"
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("not found: %s", n)
		}
		desiredPolicies, err := rulesProvider()
		if err != nil {
			return fmt.Errorf("unable to get basic policy: %v", err)
		}
		metricsPolicy := testAccProvider.Meta().(*wavefrontClient).client.MetricsPolicyAPI()
		policy, err := metricsPolicy.Get()
		if err != nil {
			return fmt.Errorf("error finding Wavefront Metrics Policy %s", err)
		}
		if rs.Primary.ID != string(rune(policy.UpdatedEpochMillis)) {
			return fmt.Errorf("expected metric policy id of %s does not match %s", rs.Primary.ID, string(rune(policy.UpdatedEpochMillis)))
		}
		if len(desiredPolicies) != len(policy.PolicyRules) {
			return fmt.Errorf("expected policy count of %d does not match %d", len(desiredPolicies), len(policy.PolicyRules))
		}

		for k, v := range policy.PolicyRules {
			if fmt.Sprintf("%v", desiredPolicies[k]) != fmt.Sprintf("%v", v) {
				return fmt.Errorf("metrics policy %v does not match %v", desiredPolicies[k], v)
			}
		}
		return nil
	}
}

// Simulates replacement of default metrics policy with system owned matching policy
func testAccCheckWavefrontMetricsPolicyBasic() string {
	return `
data "wavefront_default_user_group" "everyone" {}

resource "wavefront_metrics_policy" "main" {
  policy_rules {
    name        = "Allow All Metrics"
    description = "Predefined policy rule. Allows access to all metrics (timeseries, histograms, and counters) for all accounts. If this rule is removed, all accounts can access all metrics if there are no matching blocking rules."
    prefixes    = ["*"]
    tags_anded  = false
    access_type = "ALLOW"
    user_group_ids = [data.wavefront_default_user_group.everyone.group_id]
  }
}`
}

// Exercises wavefront_metrics_policy resource policies scoped by accounts_ids and role_ids
func testAccCheckWavefrontMetricsPolicyUpdate() string {
	return `
data "wavefront_default_user_group" "everyone" {}


resource "wavefront_user" "example" {
 email = "example@example.com"
 user_groups = [data.wavefront_default_user_group.everyone.group_id]
}

resource "wavefront_role" "test" {
  name = "test-role"
  assignees = [wavefront_user.example.id]
}

resource "wavefront_metrics_policy" "main" {
  policy_rules {
    name        = "Deny example role metrics"
    description = "deny example role test"
    prefixes    = ["example.api.*"]
    tags_anded  = false
    access_type = "BLOCK"
    role_ids       = [wavefront_role.test.id]
  }
  policy_rules {
    name        = "Deny example user metrics"
    description = "deny example user test"
    prefixes    = ["example.system.*"]
    tags {
        key = "env"
        value = "prod"
    }
    tags {
        key = "region"
        value = "us-east-1"
    }
	
    tags_anded  = true
    access_type = "BLOCK"
    account_ids    = [wavefront_user.example.id]
  }
  policy_rules {
    name        = "Allow All Metrics"
    description = "Predefined policy rule. Allows access to all metrics (timeseries, histograms, and counters) for all accounts. If this rule is removed, all accounts can access all metrics if there are no matching blocking rules."
    prefixes    = ["*"]
    tags_anded  = false
    access_type = "ALLOW"
    user_group_ids = [data.wavefront_default_user_group.everyone.group_id]
  }
}`
}

func testAccCheckWavefrontMetricsPolicyBadAccessPolicy() string {
	return `
data "wavefront_default_user_group" "everyone" {}

resource "wavefront_metrics_policy" "main" {
  policy_rules {
    name        = "Allow All Metrics"
    description = "Predefined policy rule. Allows access to all metrics (timeseries, histograms, and counters) for all accounts. If this rule is removed, all accounts can access all metrics if there are no matching blocking rules."
    prefixes    = ["*"]
    tags_anded  = false
    access_type = "NADA"
    user_group_ids = [data.wavefront_default_user_group.everyone.group_id]
  }
}`
}

func testAccCheckWavefrontMetricsPolicyNoSelectorSpecifiedPolicy() string {
	return `
data "wavefront_default_user_group" "everyone" {}

resource "wavefront_metrics_policy" "main" {
  policy_rules {
    name        = "Allow All Metrics"
    description = "Predefined policy rule. Allows access to all metrics (timeseries, histograms, and counters) for all accounts. If this rule is removed, all accounts can access all metrics if there are no matching blocking rules."
    prefixes    = ["*"]
    tags_anded  = false
    access_type = "ALLOW"
  }
}`
}
