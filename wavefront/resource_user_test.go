package wavefront

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"testing"

	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccWavefrontUser_BasicUser(t *testing.T) {
	var record wavefront.User
	config1, customer_name1 := testAccCheckWavefrontUserBasic()

	fmt.Printf("Record is %v \n", record)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: config1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontUserExists("wavefront_user.basic", &record),
					testAccCheckWavefrontUserAttributes(&record, []string{"agent_management", "alerts_management"}, []string{}),

					// Check against state that the attributes are as we expect
					resource.TestCheckResourceAttr(
						"wavefront_user.basic", "id", fmt.Sprintf("test+%s@example.com", customer_name1)),
					resource.TestCheckResourceAttr(
						"wavefront_user.basic", "permissions.#", "2"),
				),
			},
		},
	})
}

func TestAccWavefrontUser_BasicUserChangeGroups(t *testing.T) {
	var record wavefront.User

	config1, customer_name1 := testAccCheckWavefrontUserBasic()
	config2, customer_name2 := testAccCheckWavefrontUserChangeGroups()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: config1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontUserExists("wavefront_user.basic", &record),
					testAccCheckWavefrontUserAttributes(&record, []string{"agent_management", "alerts_management"}, []string{}),

					// Check against state that the attributes are as we expect
					resource.TestCheckResourceAttr(
						"wavefront_user.basic", "id", fmt.Sprintf("test+%s@example.com", customer_name1)),
					resource.TestCheckResourceAttr(
						"wavefront_user.basic", "permissions.#", "2"),
				),
			},
			{
				Config: config2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontUserExists("wavefront_user.basic", &record),
					testAccCheckWavefrontUserAttributes(&record, []string{"agent_management", "events_management"}, []string{}),

					// Check against state that the attributes are as we expect
					resource.TestCheckResourceAttr(
						"wavefront_user.basic", "id", fmt.Sprintf("test+%s@example.com", customer_name2)),
					resource.TestCheckResourceAttr(
						"wavefront_user.basic", "permissions.#", "2"),
				),
			},
		},
	})
}

func TestAccWavefrontUser_BasicUserChangeEmail(t *testing.T) {
	var record wavefront.User
	config1, customer_name1 := testAccCheckWavefrontUserBasic()
	config2, customer_name2 := testAccCheckWavefrontUserChangeEmail()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWavefrontUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: config1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontUserExists("wavefront_user.basic", &record),
					testAccCheckWavefrontUserAttributes(&record, []string{"agent_management", "alerts_management"}, []string{}),

					// Check against state that the attributes are as we expect
					resource.TestCheckResourceAttr(
						"wavefront_user.basic", "id", fmt.Sprintf("test+%s@example.com", customer_name1)),
					resource.TestCheckResourceAttr(
						"wavefront_user.basic", "permissions.#", "2"),
				),
			},
			{
				Config: config2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWavefrontUserExists("wavefront_user.basic", &record),
					testAccCheckWavefrontUserAttributes(&record, []string{"agent_management", "alerts_management"}, []string{}),

					// Check against state that the attributes are as we expect
					resource.TestCheckResourceAttr(
						"wavefront_user.basic", "id", fmt.Sprintf("test+%s@example.com", customer_name2)),
					resource.TestCheckResourceAttr(
						"wavefront_user.basic", "permissions.#", "2"),
				),
			},
		},
	})
}

func testAccCheckWavefrontUserDestroy(s *terraform.State) error {

	users := testAccProvider.Meta().(*wavefrontClient).client.Users()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "wavefront_user" {
			continue
		}

		results, err := users.Find(
			[]*wavefront.SearchCondition{
				{
					Key:            "id",
					Value:          rs.Primary.ID,
					MatchingMethod: "EXACT",
				},
			})
		if err != nil {
			return fmt.Errorf("error finding Wavefront User. %s", err)
		}
		if len(results) > 0 {
			return fmt.Errorf("user still exists")
		}
	}

	return nil
}

func testAccCheckWavefrontUserExists(n string, user *wavefront.User) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Record ID is set")
		}

		users := testAccProvider.Meta().(*wavefrontClient).client.Users()

		results, err := users.Find(
			[]*wavefront.SearchCondition{
				{
					Key:            "id",
					Value:          rs.Primary.ID,
					MatchingMethod: "EXACT",
				},
			})
		if err != nil {
			return fmt.Errorf("error finding Wavefront User %s", err)
		}
		// resource has been deleted out of band. So unset ID
		if len(results) != 1 {
			return fmt.Errorf("no Users Found")
		}
		if *results[0].ID != rs.Primary.ID {
			return fmt.Errorf("user not found")
		}

		*user = *results[0]

		return nil
	}
}

func testAccCheckWavefrontUserAttributes(user *wavefront.User, permissions []string, groups []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, v := range permissions {
			if sort.SearchStrings(user.Permissions, v) == len(user.Permissions) {
				return fmt.Errorf("permission not found or present on user. %s", v)
			}
		}

		for _, v := range groups {
			found := false
			for _, g := range user.Groups.UserGroups {
				if *g.ID == v {
					found = true
				}
			}
			if !found {
				return fmt.Errorf("group not found or present on user. %s", v)
			}
		}
		return nil
	}
}

type User struct {
	Identifier          string   `json:"identifier"`
	Customer            string   `json:"customer"`
	LastSuccessfulLogin int      `json:"lastSuccessfulLogin"`
	Groups              []string `json:"groups"`
	IngestionPolicies   []string `json:"ingestionPolicies"`
	Roles               []string `json:"roles"`
}

type ResponseObj struct {
	Status   Status `json:"status"`
	Response []User `json:"response"`
}

type Status struct {
	Result  string `json:"result"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func getCustomerName() string {
	var isAccTestsEnabled string
	var wf_token string
	var system_url string
	var customerName string

	isAccTestsEnabled = os.Getenv("TF_ACC")
	wf_token = os.Getenv("WAVEFRONT_TOKEN")
	system_url = os.Getenv("WAVEFRONT_ADDRESS")
	customerName = ""

	if isAccTestsEnabled == "1" {

		var url string = fmt.Sprintf("https://%s/api/v2/account/user", system_url)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Errorf("Error Creating New Request to find Customer Name!")
		}

		// Header -> Authorization: Bearer <TOKEN>
		// URL: https://cluster.wavefront.com/api/v2/account/user

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", wf_token))

		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			fmt.Errorf("Error Finding Customer Name!")
		}

		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Errorf("Error")
			}

			var m ResponseObj
			if err = json.Unmarshal(bodyBytes, &m); err != nil {
				panic(err)
			}

			return m.Response[0].Customer
		}

	}

	return customerName
}

func testAccCheckWavefrontUserBasic() (string, string) {
	var customer_name string
	var tf_resource string
	var temp_customer string = getCustomerName()

	customer_name = "tftesting"

	if len(temp_customer) > 0 {
		customer_name = temp_customer
	}

	tf_resource = fmt.Sprintf("resource \"wavefront_user\" \"basic\" { \n email = \"test+%s@example.com\" \n permissions = [ \n \"agent_management\", \n \"alerts_management\", \n ] \n} \n", customer_name)
	return tf_resource, customer_name
}

func testAccCheckWavefrontUserChangeGroups() (string, string) {
	var customer_name string
	var tf_resource string

	customer_name = getCustomerName()
	if len(customer_name) == 0 {
		customer_name = "tftesting"
	}

	tf_resource = fmt.Sprintf("resource \"wavefront_user\" \"basic\" { \n email = \"test+%s@example.com\" \n permissions = [ \n \"agent_management\", \n \"events_management\", \n] \n} \n", customer_name)
	return tf_resource, customer_name
}

func testAccCheckWavefrontUserChangeEmail() (string, string) {
	var customer_name string
	var tf_resource string

	customer_name = getCustomerName()
	if len(customer_name) == 0 {
		customer_name = "tftesting2"
	}

	tf_resource = fmt.Sprintf("resource \"wavefront_user\" \"basic\" { \n email = \"test+%s@example.com\" \n permissions = [ \n \"agent_management\", \n \"alerts_management\", \n ] \n} \n", customer_name)
	return tf_resource, customer_name
}
