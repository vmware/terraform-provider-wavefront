package wavefront_plugin

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strings"
)

func resourceCloudIntegrationAwsExternalId() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudIntegrationAwsExternalIdCreate,
		Read:   resourceCloudIntegrationAwsExternalIdRead,
		Delete: resourceCloudIntegrationAwsExternalIdDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{},
	}
}

func resourceCloudIntegrationAwsExternalIdCreate(d *schema.ResourceData, meta interface{}) error {
	cloudIntegrations := meta.(*wavefrontClient).client.CloudIntegrations()

	extId, err := cloudIntegrations.CreateAwsExternalID()
	if err != nil {
		return fmt.Errorf("error creating AWS External ID. %s", err)
	}

	d.SetId(extId)
	return nil
}

func resourceCloudIntegrationAwsExternalIdRead(d *schema.ResourceData, meta interface{}) error {
	cloudIntegrations := meta.(*wavefrontClient).client.CloudIntegrations()
	extId := d.Id()
	err := cloudIntegrations.VerifyAwsExternalID(extId)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			d.SetId("")
			return nil
		} else {
			return fmt.Errorf("unable to find AWS External ID %s. %s", d.Id(), err)
		}
	}
	d.SetId(extId)

	return nil
}

func resourceCloudIntegrationAwsExternalIdDelete(d *schema.ResourceData, meta interface{}) error {
	cloudIntegrations := meta.(*wavefrontClient).client.CloudIntegrations()
	extId := d.Id()
	err := cloudIntegrations.DeleteAwsExternalID(&extId)
	if err != nil {
		return fmt.Errorf("error deleting AWS External ID. %s", err)
	}
	d.SetId("")
	return nil
}
