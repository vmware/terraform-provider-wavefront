package wavefront_plugin

import (
	"fmt"
	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceUserGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserGroupCreate,
		Read:   resourceUserGroupRead,
		Update: resourceUserGroupUpdate,
		Delete: resourceUserGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceUserGroupCreate(d *schema.ResourceData, meta interface{}) error {
	userGroups := meta.(*wavefrontClient).client.UserGroups()

	ug := &wavefront.UserGroup{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	}

	if err := userGroups.Create(ug); err != nil {
		return fmt.Errorf("failed to create user group, %s", err)
	}

	d.SetId(*ug.ID)

	return resourceUserGroupRead(d, meta)
}

func resourceUserGroupRead(d *schema.ResourceData, meta interface{}) error {
	userGroups := meta.(*wavefrontClient).client.UserGroups()
	id := d.Id()
	ug := &wavefront.UserGroup{
		ID: &id,
	}

	if err := userGroups.Get(ug); err != nil {
		return fmt.Errorf("unable to find user group %s, %s", id, err)
	}

	d.Set("name", ug.Name)
	d.Set("description", ug.Description)

	return nil
}

func resourceUserGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	userGroups := meta.(*wavefrontClient).client.UserGroups()

	id := d.Id()
	ug := &wavefront.UserGroup{
		ID: &id,
	}

	ug.Name = d.Get("name").(string)
	ug.Description = d.Get("description").(string)

	if err := userGroups.Update(ug); err != nil {
		return fmt.Errorf("unable to update user group %s, %s", id, err)
	}

	return resourceUserGroupRead(d, meta)
}

func resourceUserGroupDelete(d *schema.ResourceData, meta interface{}) error {
	userGroups := meta.(*wavefrontClient).client.UserGroups()

	id := d.Id()
	ug := &wavefront.UserGroup{
		ID: &id,
	}

	if err := userGroups.Delete(ug); err != nil {
		return fmt.Errorf("unable to delete user group %s, %s", id, err)
	}

	d.SetId("")
	return nil
}
