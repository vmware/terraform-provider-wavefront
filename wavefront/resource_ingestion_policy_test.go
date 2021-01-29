package wavefront

import (
	"fmt"
	"testing"

	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccWavefrontIngestionPolicy_Basic(t *testing.T) {
	var record wavefront.IngestionPolicy

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontIngestionPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: `
resource wavefront_ingestion_policy tester {
    name = "Test Ingestion Policy"
    description = "Ingestion policy for Terraform test"
}`,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontIngestionPolicyExists("wavefront_ingestion_policy.tester", &record),
					testAccWavefrontIngestionPolicyEquals(
						&wavefront.IngestionPolicy{
							Name:        "Test Ingestion Policy",
							Description: "Ingestion policy for Terraform test",
						},
						&record,
					),
					// Check against state that the attributes are as we expect
					resource.TestCheckResourceAttr(
						"wavefront_ingestion_policy.tester", "name", "Test Ingestion Policy"),
					resource.TestCheckResourceAttr(
						"wavefront_ingestion_policy.tester", "description", "Ingestion policy for Terraform test"),
				),
			},
		},
	})
}

func testAccCheckWavefrontIngestionPolicyDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*wavefrontClient).client.IngestionPolicies()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "wavefront_ingestion_policy" {
			continue
		}

		id := rs.Primary.ID
		policy := wavefront.IngestionPolicy{ID: id}
		err := client.Get(&policy)

		if wavefront.NotFound(err) {
			continue
		}

		if err != nil {
			return fmt.Errorf("error finding ingestion policy, %s", err)
		}

		return fmt.Errorf("ingestion policy still exists, %s", rs.Primary.ID)
	}
	return nil
}

func testAccCheckWavefrontIngestionPolicyExists(
	n string, policy *wavefront.IngestionPolicy) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Record ID is set")
		}

		client := testAccProvider.Meta().(*wavefrontClient).client.IngestionPolicies()
		id := rs.Primary.ID
		*policy = wavefront.IngestionPolicy{ID: id}
		err := client.Get(policy)

		if wavefront.NotFound(err) {
			return fmt.Errorf("ingestion policy not found %s", rs.Primary.ID)
		}

		if err != nil {
			return fmt.Errorf("error finding ingestion policy %s", err)
		}

		return nil
	}
}

func testAccWavefrontIngestionPolicyEquals(
	expected *wavefront.IngestionPolicy,
	policy *wavefront.IngestionPolicy) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		result := cmp.Diff(
			ingestionPolicyZeroExtraFields(expected),
			ingestionPolicyZeroExtraFields(policy),
			cmpopts.EquateEmpty())

		if result == "" {
			return nil
		}

		return fmt.Errorf("options differ: %s", result)
	}
}

func ingestionPolicyZeroExtraFields(policy *wavefront.IngestionPolicy) *wavefront.IngestionPolicy {
	policyCopy := *policy
	policyCopy.ID = ""
	return &policyCopy
}
