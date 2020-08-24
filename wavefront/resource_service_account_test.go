package wavefront

import (
	"fmt"
	"testing"

	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccWavefrontServiceAccount_Basic(t *testing.T) {
	var record wavefront.ServiceAccount

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontServiceAccountDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccWavefrontServiceAccountBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontServiceAccountExists("wavefront_service_account.basic", &record),
					testAccWavefrontServiceAccountAttributes(
						&record,
						[]string{"agent_management", "alerts_management"},
						"A description",
					),

					// Check against state that the attributes are as we expect
					resource.TestCheckResourceAttr(
						"wavefront_service_account.basic", "identifier", "sa::tftesting"),
					resource.TestCheckResourceAttr(
						"wavefront_service_account.basic", "permissions.#", "2"),
					resource.TestCheckResourceAttr(
						"wavefront_service_account.basic", "description", "A description"),
				),
			},
		},
	})
}

func TestAccWavefrontServiceAccount_BasicChangeGroups(t *testing.T) {
	var record wavefront.ServiceAccount

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontServiceAccountDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccWavefrontServiceAccountBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontServiceAccountExists("wavefront_service_account.basic", &record),
				),
			},
			{
				Config: testAccWavefrontServiceAccountChangeGroups(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontServiceAccountExists("wavefront_service_account.basic", &record),
					testAccWavefrontServiceAccountAttributes(
						&record,
						[]string{"agent_management", "events_management"},
						"A description"),
					// Check against state that the attributes are as we expect
					resource.TestCheckResourceAttr(
						"wavefront_service_account.basic", "identifier", "sa::tftesting"),
					resource.TestCheckResourceAttr(
						"wavefront_service_account.basic", "permissions.#", "2"),
				),
			},
		},
	})
}

func TestAccWavefrontServiceAccount_BasicChangeID(t *testing.T) {
	var record wavefront.ServiceAccount

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontServiceAccountDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccWavefrontServiceAccountBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontServiceAccountExists("wavefront_service_account.basic", &record),
					func(s *terraform.State) error {
						if record.ID != "sa::tftesting" {
							return fmt.Errorf("record.ID is %s", record.ID)
						}
						return nil
					},
				),
			},
			{
				Config: testAccWavefrontServiceAccountChangeID(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontServiceAccountExists("wavefront_service_account.basic", &record),
					func(s *terraform.State) error {
						if record.ID != "sa::tftesting2" {
							return fmt.Errorf("record.ID is %s", record.ID)
						}
						return nil
					},
				),
			},
		},
	})
}

func testAccCheckWavefrontServiceAccountDestroy(s *terraform.State) error {

	serviceAccounts := testAccProvider.Meta().(*wavefrontClient).client.ServiceAccounts()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "wavefront_service_account" {
			continue
		}
		_, err := serviceAccounts.GetByID(rs.Primary.ID)
		if wavefront.NotFound(err) {
			continue
		}
		if err != nil {
			return fmt.Errorf("error finding Wavefront User, %s", err)
		}
		return fmt.Errorf("user still exists, %s", rs.Primary.ID)
	}
	return nil
}

func testAccCheckWavefrontServiceAccountExists(
	n string, serviceAccount *wavefront.ServiceAccount) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Record ID is set")
		}

		serviceAccounts := testAccProvider.Meta().(*wavefrontClient).client.ServiceAccounts()

		result, err := serviceAccounts.GetByID(rs.Primary.ID)
		if wavefront.NotFound(err) {
			return fmt.Errorf("user not found %s", rs.Primary.ID)
		}
		if err != nil {
			return fmt.Errorf("error finding Wavefront User %s", err)
		}
		*serviceAccount = *result
		return nil
	}
}

func testAccWavefrontServiceAccountAttributes(
	serviceAccount *wavefront.ServiceAccount,
	permissions []string,
	description string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		onlyLeft, onlyRight := compareStringSliceAnyOrder(
			permissions, serviceAccount.Permissions)
		if len(onlyLeft) > 0 || len(onlyRight) > 0 {
			return fmt.Errorf(
				"missing permissions: %v, unexpected permissions: %v",
				onlyLeft,
				onlyRight)
		}
		if description != serviceAccount.Description {
			return fmt.Errorf(
				"expected description: %s, got %s",
				description,
				serviceAccount.Description)
		}
		return nil
	}
}

func testAccWavefrontServiceAccountBasic() string {
	return `
resource "wavefront_service_account" "basic" {
	identifier  = "sa::tftesting"
	permissions = [
		"agent_management",
		"alerts_management",
	]
    description = "A description"
    active = true
}`
}

func testAccWavefrontServiceAccountChangeGroups() string {
	return `
resource "wavefront_service_account" "basic" {
	identifier  = "sa::tftesting"
	permissions = [
		"agent_management",
		"events_management",
	]
    description = "A description"
    active = true
}`
}

func testAccWavefrontServiceAccountChangeID() string {
	return `
resource "wavefront_service_account" "basic" {
	identifier  = "sa::tftesting2"
	permissions = [
		"agent_management",
		"alerts_management",
	]
    description = "A description"
    active = true
}`
}
