package wavefront

import (
	"fmt"

	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	ipNameKey        = "name"
	ipDescriptionKey = "description"
)

func resourceIngestionPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceIngestionPolicyCreate,
		Read:   resourceIngestionPolicyRead,
		Update: resourceIngestionPolicyUpdate,
		Delete: resourceIngestionPolicyDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			ipNameKey: {
				Type:     schema.TypeString,
				Required: true,
			},
			ipDescriptionKey: {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceIngestionPolicyCreate(
	d *schema.ResourceData, meta interface{}) error {
	client := meta.(*wavefrontClient).client.IngestionPolicies()

	policy := wavefront.IngestionPolicy{
		Name:        d.Get(ipNameKey).(string),
		Description: d.Get(ipDescriptionKey).(string),
	}

	err := client.Create(&policy)

	if err != nil {
		return fmt.Errorf(
			"failed to create ingestion policy, %s", err)
	}

	d.SetId(policy.ID)

	return resourceIngestionPolicyRead(d, meta)
}

func resourceIngestionPolicyRead(
	d *schema.ResourceData, meta interface{}) error {
	client := meta.(*wavefrontClient).client.IngestionPolicies()
	policy := wavefront.IngestionPolicy{ID: d.Id()}
	err := client.Get(&policy)

	if wavefront.NotFound(err) {
		d.SetId("")
		return nil
	}

	if err != nil {
		return fmt.Errorf("error finding ingestion policy, %s. %s", d.Id(), err)
	}

	if err := d.Set(ipNameKey, policy.Name); err != nil {
		return err
	}

	if err := d.Set(ipDescriptionKey, policy.Description); err != nil {
		return err
	}

	return nil
}

func resourceIngestionPolicyUpdate(
	d *schema.ResourceData, meta interface{}) error {
	client := meta.(*wavefrontClient).client.IngestionPolicies()
	policy := wavefront.IngestionPolicy{ID: d.Id()}
	err := client.Get(&policy)

	if wavefront.NotFound(err) {
		d.SetId("")
		return nil
	}

	if err != nil {
		return fmt.Errorf(""+"error finding ingestion policy, %s. %s", d.Id(), err)
	}

	if d.HasChange(ipNameKey) {
		policy.Name = d.Get(ipNameKey).(string)
	}
	if d.HasChange(ipDescriptionKey) {
		policy.Description = d.Get(ipDescriptionKey).(string)
	}

	err = client.Update(&policy)

	if err != nil {
		return fmt.Errorf("error updating ingestion policy,  %s. %s", d.Id(), err)
	}

	return resourceIngestionPolicyRead(d, meta)
}

func resourceIngestionPolicyDelete(
	d *schema.ResourceData, meta interface{}) error {
	client := meta.(*wavefrontClient).client.IngestionPolicies()
	policy := wavefront.IngestionPolicy{ID: d.Id()}
	err := client.Delete(&policy)

	if err != nil && !wavefront.NotFound(err) {
		return fmt.Errorf("error deleting ingestion policy, %s. %s", d.Id(), err)
	}

	d.SetId("")

	return nil
}
