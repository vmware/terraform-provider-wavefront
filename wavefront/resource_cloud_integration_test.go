package wavefront_plugin

import (
	"fmt"
	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"strings"
)

func testAccCheckWavefrontCloudIntegrationDestroy(s *terraform.State) error {
	integrations := testAccProvider.Meta().(*wavefrontClient).client.CloudIntegrations()
	for _, rs := range s.RootModule().Resources {
		// Skip anything that isn't a cloud integration OR anything that is an
		// aws external id as these are not our responsibility
		if !strings.Contains(rs.Type, "wavefront_cloud_integration") ||
			strings.Contains(rs.Type, "_aws_external_id") {
			continue
		}

		tmp := wavefront.CloudIntegration{
			Id: rs.Primary.ID,
		}

		err := integrations.Get(&tmp)
		if err == nil {
			return fmt.Errorf("cloud integration still exists")
		}
	}

	return nil
}

func testAccCheckWavefrontCloudIntegrationExists(n string, integration *wavefront.CloudIntegration) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Record ID is set")
		}

		ci := testAccProvider.Meta().(*wavefrontClient).client.CloudIntegrations()
		tmp := wavefront.CloudIntegration{Id: rs.Primary.ID}

		err := ci.Get(&tmp)
		if err != nil {
			return fmt.Errorf("error finding Wavefront Cloud Integration %s", err)
		}

		*integration = tmp
		return nil
	}
}

func testAccCheckWavefrontCloudIntegrationAttributes(integration *wavefront.CloudIntegration, service string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if integration.ForceSave {
			return fmt.Errorf("bad value: %v", integration.ForceSave)
		}

		if integration.Name != "Test Integration" {
			return fmt.Errorf("bad value: %s", integration.Name)
		}

		if integration.Service != service {
			return fmt.Errorf("bad value, expected %s. got %s", service, integration.Service)
		}

		if val, ok := integration.AdditionalTags["tag1"]; !ok {
			return fmt.Errorf("key missing: %s", "tag1")
		} else {
			if val != "value1" {
				return fmt.Errorf("tag1 value is incorrect. got %s", val)
			}
		}

		if val, ok := integration.AdditionalTags["tag2"]; !ok {
			return fmt.Errorf("key missing: %s", "tag2")
		} else {
			if val != "value2" {
				return fmt.Errorf("tag2 value is incorrect. got %s", val)
			}
		}
		return nil
	}
}

func testAccCheckWavefrontCloudIntegrationResourceAttributes(resourcePrefix, service string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		err := resource.TestCheckResourceAttr(resourcePrefix, "name", "Test Integration")(s)
		if err != nil {
			return err
		}
		err = resource.TestCheckResourceAttr(resourcePrefix, "force_save", "true")(s)
		if err != nil {
			return err
		}
		err = resource.TestCheckResourceAttr(resourcePrefix, "service", service)(s)
		if err != nil {
			return err
		}
		err = resource.TestCheckResourceAttr(resourcePrefix, "additional_tags.tag1", "value1")(s)
		if err != nil {
			return err
		}
		err = resource.TestCheckResourceAttr(resourcePrefix, "additional_tags.tag2", "value2")(s)
		if err != nil {
			return err
		}
		return nil
	}
}
