package wavefront

import (
	"fmt"

	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServiceAccount() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceAccountCreate,
		Read:   resourceServiceAccountRead,
		Update: resourceServiceAccountUpdate,
		Delete: resourceServiceAccountDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"identifier": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"active": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"permissions": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
				Set:      schema.HashString,
			},
			"user_groups": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
				Set:      schema.HashString,
			},
			"ingestion_policy": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceServiceAccountCreate(d *schema.ResourceData, meta interface{}) error {
	serviceAccounts := meta.(*wavefrontClient).client.ServiceAccounts()
	tokens := meta.(*wavefrontClient).client.Tokens()

	serviceAccount, err := serviceAccounts.Create(
		&wavefront.ServiceAccountOptions{
			ID:                d.Get("identifier").(string),
			Active:            d.Get("active").(bool),
			Description:       d.Get("description").(string),
			Permissions:       getStringSlice(d, "permissions"),
			UserGroups:        getStringSlice(d, "user_groups"),
			IngestionPolicyID: d.Get("ingestion_policy").(string),
		})
	if err != nil {
		return fmt.Errorf(
			"failed to create new Wavefront Service Account, %s",
			err)
	}
	_, err = tokens.Create(serviceAccount.ID, &wavefront.TokenOptions{Name: "main"})
	if err != nil {
		return fmt.Errorf(
			"failed to create token for new Wavefront ServiceAccount, %s",
			err)
	}
	d.SetId(serviceAccount.ID)

	return resourceServiceAccountRead(d, meta)
}

func resourceServiceAccountRead(d *schema.ResourceData, meta interface{}) error {
	serviceAccounts := meta.(*wavefrontClient).client.ServiceAccounts()
	serviceAccount, err := serviceAccounts.GetByID(d.Id())
	if wavefront.NotFound(err) {
		d.SetId("")
		return nil
	}
	if err != nil {
		return fmt.Errorf(
			"error finding Wavefront Service Account %s. %s",
			d.Id(),
			err)
	}
	if err := d.Set("identifier", serviceAccount.ID); err != nil {
		return err
	}
	if err := d.Set("active", serviceAccount.Active); err != nil {
		return err
	}
	if err := d.Set("description", serviceAccount.Description); err != nil {
		return err
	}
	if err := d.Set("ingestion_policy", serviceAccount.IngestionPolicyId()); err != nil {
		return err
	}
	err = setStringSlice(d, "permissions", serviceAccount.Permissions)
	if err != nil {
		return err
	}
	return setStringSlice(d, "user_groups", serviceAccount.UserGroupIds())
}

func resourceServiceAccountUpdate(
	d *schema.ResourceData, meta interface{}) error {
	serviceAccounts := meta.(*wavefrontClient).client.ServiceAccounts()
	serviceAccount, err := serviceAccounts.GetByID(d.Id())
	if wavefront.NotFound(err) {
		d.SetId("")
		return nil
	}
	if err != nil {
		return fmt.Errorf(""+
			"error finding Wavefront Service Account %s. %s",
			d.Id(),
			err)
	}
	options := serviceAccount.Options()
	if d.HasChange("active") {
		options.Active = d.Get("active").(bool)
	}
	if d.HasChange("description") {
		options.Description = d.Get("description").(string)
	}
	if d.HasChange("permissions") {
		options.Permissions = getStringSlice(d, "permissions")
	}
	if d.HasChange("user_groups") {
		options.UserGroups = getStringSlice(d, "user_groups")
	}
	if d.HasChange("ingestion_policy") {
		options.IngestionPolicyID = d.Get("ingestion_policy").(string)
	}
	_, err = serviceAccounts.Update(options)
	if err != nil {
		return fmt.Errorf(
			"error updating Wavefront Service Account  %s. %s",
			d.Id(),
			err,
		)
	}

	return resourceServiceAccountRead(d, meta)
}

func resourceServiceAccountDelete(
	d *schema.ResourceData, meta interface{}) error {
	serviceAccounts := meta.(*wavefrontClient).client.ServiceAccounts()
	err := serviceAccounts.DeleteByID(d.Id())
	if err != nil && !wavefront.NotFound(err) {
		return fmt.Errorf(
			"error deleting Wavefront Service Account %s. %s",
			d.Id(),
			err,
		)
	}
	d.SetId("")
	return nil
}
