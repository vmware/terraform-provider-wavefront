package wavefront

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccWavefrontIngestionPolicy_Accounts(t *testing.T) {
	var record wavefront.IngestionPolicyResponse

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontIngestionPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: `
					resource "wavefront_user" "basic" {
						email  = "test+tftesting@example.com"
					}
					resource wavefront_ingestion_policy tester {
						name = "Test Ingestion Policy"
						description = "Ingestion policy for Terraform test"
						scope = "ACCOUNT"
						accounts  = [wavefront_user.basic.id]
					}`,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontIngestionPolicyExists("wavefront_ingestion_policy.tester", &record),
					testAccWavefrontIngestionPolicyEquals(
						&wavefront.IngestionPolicyResponse{
							Name:        "Test Ingestion Policy",
							Description: "Ingestion policy for Terraform test",
							Scope:       "ACCOUNT",
							Accounts: []wavefront.IngestionPolicyAccount{
								{
									Name: "test+tftesting@example.com",
									ID:   "test+tftesting@example.com",
								},
							},
						},
						&record,
					),
					// Check against state that the attributes are as we expect
					resource.TestCheckResourceAttr(
						"wavefront_ingestion_policy.tester", "name", "Test Ingestion Policy"),
					resource.TestCheckResourceAttr(
						"wavefront_ingestion_policy.tester", "description", "Ingestion policy for Terraform test"),
					resource.TestCheckResourceAttr(
						"wavefront_ingestion_policy.tester", "scope", "ACCOUNT"),
					resource.TestCheckResourceAttr(
						"wavefront_ingestion_policy.tester", "accounts.0", "test+tftesting@example.com"),
				),
			},
		},
	})
}

func TestAccWavefrontIngestionPolicy_Groups(t *testing.T) {
	var record wavefront.IngestionPolicyResponse

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontIngestionPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: `
					resource "wavefront_user_group" "basic" {
						name = "test group"
						description = "test group"
					}
					resource wavefront_ingestion_policy tester {
						name = "Test Ingestion Policy"
						description = "Ingestion policy for Terraform test"
						scope = "GROUP"
						groups  = [wavefront_user_group.basic.id]
					}`,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontIngestionPolicyExists("wavefront_ingestion_policy.tester", &record),
					testAccWavefrontIngestionPolicyEquals(
						&wavefront.IngestionPolicyResponse{
							Name:        "Test Ingestion Policy",
							Description: "Ingestion policy for Terraform test",
							Scope:       "GROUP",
							Groups: []wavefront.IngestionPolicyGroup{
								{
									Name:        "test group",
									Description: "test group",
								},
							},
						},
						&record,
					),
					// Check against state that the attributes are as we expect
					resource.TestCheckResourceAttr(
						"wavefront_ingestion_policy.tester", "name", "Test Ingestion Policy"),
					resource.TestCheckResourceAttr(
						"wavefront_ingestion_policy.tester", "description", "Ingestion policy for Terraform test"),
					resource.TestCheckResourceAttr(
						"wavefront_ingestion_policy.tester", "scope", "GROUP"),
					resource.TestMatchResourceAttr(
						"wavefront_ingestion_policy.tester", "groups.0", regexp.MustCompile("[a-z0-9-]{36}")),
				),
			},
		},
	})
}

func TestAccWavefrontIngestionPolicy_Sources(t *testing.T) {
	var record wavefront.IngestionPolicyResponse

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
						scope = "SOURCE"
						sources  = ["wf-proxy"]
					}`,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontIngestionPolicyExists("wavefront_ingestion_policy.tester", &record),
					testAccWavefrontIngestionPolicyEquals(
						&wavefront.IngestionPolicyResponse{
							Name:        "Test Ingestion Policy",
							Description: "Ingestion policy for Terraform test",
							Scope:       "SOURCE",
							Sources:     []string{"wf-proxy"},
						},
						&record,
					),
					// Check against state that the attributes are as we expect
					resource.TestCheckResourceAttr(
						"wavefront_ingestion_policy.tester", "name", "Test Ingestion Policy"),
					resource.TestCheckResourceAttr(
						"wavefront_ingestion_policy.tester", "description", "Ingestion policy for Terraform test"),
					resource.TestCheckResourceAttr(
						"wavefront_ingestion_policy.tester", "scope", "SOURCE"),
					resource.TestCheckResourceAttr(
						"wavefront_ingestion_policy.tester", "sources.0", "wf-proxy"),
				),
			},
		},
	})
}

func TestAccWavefrontIngestionPolicy_Namespaces(t *testing.T) {
	var record wavefront.IngestionPolicyResponse

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
						scope = "NAMESPACE"
						namespaces  = ["system"]
					}`,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontIngestionPolicyExists("wavefront_ingestion_policy.tester", &record),
					testAccWavefrontIngestionPolicyEquals(
						&wavefront.IngestionPolicyResponse{
							Name:        "Test Ingestion Policy",
							Description: "Ingestion policy for Terraform test",
							Scope:       "NAMESPACE",
							Namespaces:  []string{"system"},
						},
						&record,
					),
					// Check against state that the attributes are as we expect
					resource.TestCheckResourceAttr(
						"wavefront_ingestion_policy.tester", "name", "Test Ingestion Policy"),
					resource.TestCheckResourceAttr(
						"wavefront_ingestion_policy.tester", "description", "Ingestion policy for Terraform test"),
					resource.TestCheckResourceAttr(
						"wavefront_ingestion_policy.tester", "scope", "NAMESPACE"),
					resource.TestCheckResourceAttr(
						"wavefront_ingestion_policy.tester", "namespaces.0", "system"),
				),
			},
		},
	})
}

func TestAccWavefrontIngestionPolicy_Tags(t *testing.T) {
	var record wavefront.IngestionPolicyResponse

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
						scope = "TAGS"
						tags {
							key = "user"
							value = "*"
						}
					}`,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontIngestionPolicyExists("wavefront_ingestion_policy.tester", &record),
					testAccWavefrontIngestionPolicyEquals(
						&wavefront.IngestionPolicyResponse{
							Name:        "Test Ingestion Policy",
							Description: "Ingestion policy for Terraform test",
							Scope:       "TAGS",
							Tags: []wavefront.IngestionPolicyTag{
								{
									Key:   "user",
									Value: "*",
								},
							},
						},
						&record,
					),
					// Check against state that the attributes are as we expect
					resource.TestCheckResourceAttr(
						"wavefront_ingestion_policy.tester", "name", "Test Ingestion Policy"),
					resource.TestCheckResourceAttr(
						"wavefront_ingestion_policy.tester", "description", "Ingestion policy for Terraform test"),
					resource.TestCheckResourceAttr(
						"wavefront_ingestion_policy.tester", "scope", "TAGS"),
					resource.TestCheckResourceAttr(
						"wavefront_ingestion_policy.tester", "tags.0.key", "user"),
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
		_, err := client.GetByID(id)

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

func testAccCheckWavefrontIngestionPolicyExists(n string, policy *wavefront.IngestionPolicyResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Record ID is set")
		}

		var err error
		client := testAccProvider.Meta().(*wavefrontClient).client.IngestionPolicies()
		id := rs.Primary.ID
		var response *wavefront.IngestionPolicyResponse
		response, err = client.GetByID(id)
		*policy = *response

		if wavefront.NotFound(err) {
			return fmt.Errorf("ingestion policy not found %s", rs.Primary.ID)
		}

		if err != nil {
			return fmt.Errorf("error finding ingestion policy %s", err)
		}

		return nil
	}
}

func testAccWavefrontIngestionPolicyEquals(expected *wavefront.IngestionPolicyResponse, policy *wavefront.IngestionPolicyResponse) resource.TestCheckFunc {

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

func ingestionPolicyZeroExtraFields(policy *wavefront.IngestionPolicyResponse) *wavefront.IngestionPolicyResponse {
	policyCopy := *policy
	policyCopy.ID = ""
	for i := range policyCopy.Groups {
		policyCopy.Groups[i].ID = ""
	}
	return &policyCopy
}
