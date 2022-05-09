//package wavefront
//
//import (
//	"fmt"
//	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
//	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
//	"testing"
//	"time"
//)
//
//func TestAccWavefrontDataSourceUser_Ok(t *testing.T) {
//	resource.Test(t, resource.TestCase{
//		PreCheck:  func() { testAccPreCheck(t) },
//		Providers: testAccProviders,
//		Steps: []resource.TestStep{
//			{
//				Config: testAccCheckWavefrontDataSourceUserOkPolicy(),
//				Check:  testAccCheckWavefrontDataSourceUser("data.wavefront_metrics_policy.main"),
//			},
//		},
//	})
//}
//
//func testAccCheckWavefrontDataSourceUser(n string) resource.TestCheckFunc {
//	return func(s *terraform.State) error {
//		rs, ok := s.RootModule().Resources[n]
//
//		if !ok {
//			return fmt.Errorf("not found: %s", n)
//		}
//
//		// Validate key is unix timestamp time.Time.Format("2006-01-02 15:04:05.999999999 -0700 MST")
//		if _, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", rs.Primary.ID); err != nil {
//			return fmt.Errorf("data source id should be timestamp: %v", err)
//		}
//
//		if rs.Primary.Attributes[emailKey] != "example@example.com" {
//			return fmt.Errorf("%s key value '%s' does not match desired '%s", emailKey, rs.Primary.Attributes[emailKey], "example@example.com")
//		}
//
//
//		if rs.Primary.Attributes[userGroupsKey] != string([]string{}) {
//			return fmt.Errorf("%s key value '%s' does not match desired '%s", emailKey, rs.Primary.Attributes[emailKey], "example@example.com")
//		}
//
//		return nil
//	}
//}
//
//func testAccCheckWavefrontDataSourceUserOkPolicy() string {
//	return `
//
////data "wavefront_default_user_group" "everyone" {}
////
////resource "wavefront_user" "example" {
//// email = "example@example.com"
//// user_groups = [data.wavefront_default_user_group.everyone.group_id]
////}
////
////resource "wavefront_role" "test" {
////  name = "test-role"
////  assignees = [data.wavefront_user.example.id]
////}
//
//data "wavefront_user" "example" {
//  id = "sean.norris@woven-planet.global"
//}
//`
//}
