package wavefront

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudIntegrationAwsExternalID() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudIntegrationAwsExternalIDCreate,
		Read:   resourceCloudIntegrationAwsExternalIDRead,
		Delete: resourceCloudIntegrationAwsExternalIDDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{},
	}
}

func resourceCloudIntegrationAwsExternalIDCreate(d *schema.ResourceData, meta interface{}) error {
	cloudIntegrations := meta.(*wavefrontClient).client.CloudIntegrations()

	extID, err := cloudIntegrations.CreateAwsExternalID()
	if err != nil {
		return fmt.Errorf("error creating AWS External ID. %s", err)
	}

	d.SetId(extID)
	return nil
}

func resourceCloudIntegrationAwsExternalIDRead(d *schema.ResourceData, meta interface{}) error {
	cloudIntegrations := meta.(*wavefrontClient).client.CloudIntegrations()
	extID := d.Id()
	err := cloudIntegrations.VerifyAwsExternalID(extID)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("unable to find AWS External ID %s. %s", d.Id(), err)
	}
	d.SetId(extID)

	return nil
}

func resourceCloudIntegrationAwsExternalIDDelete(d *schema.ResourceData, meta interface{}) error {
	cloudIntegrations := meta.(*wavefrontClient).client.CloudIntegrations()
	extID := d.Id()
	err := cloudIntegrations.DeleteAwsExternalID(&extID)
	if err != nil {
		return fmt.Errorf("error deleting AWS External ID. %s", err)
	}
	d.SetId("")
	return nil
}
