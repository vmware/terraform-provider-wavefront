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
			"permissions": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func resourceUserGroupCreate(d *schema.ResourceData, m interface{}) error {
	userGroups := m.(*wavefrontClient).client.UserGroups()

	ug := &wavefront.UserGroup{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	}

	resourceDecodeUserGroupPermissions(d, ug)
	if err := userGroups.Create(ug); err != nil {
		return fmt.Errorf("failed to create user group, %s", err)
	}

	d.SetId(*ug.ID)

	return resourceUserGroupRead(d, m)
}

func resourceUserGroupRead(d *schema.ResourceData, m interface{}) error {
	userGroups := m.(*wavefrontClient).client.UserGroups()
	id := d.Id()
	ug := &wavefront.UserGroup{
		ID: &id,
	}

	if err := userGroups.Get(ug); err != nil {
		return fmt.Errorf("unable to find user group %s, %s", id, err)
	}

	d.Set("name", ug.Name)
	d.Set("description", ug.Description)
	d.Set("permissions", ug.Permissions)

	return nil
}

func resourceUserGroupUpdate(d *schema.ResourceData, m interface{}) error {
	userGroups := m.(*wavefrontClient).client.UserGroups()

	id := d.Id()
	ug := &wavefront.UserGroup{
		ID: &id,
	}

	ug.Name = d.Get("name").(string)
	ug.Description = d.Get("description").(string)
	resourceDecodeUserGroupPermissions(d, ug)

	if err := userGroups.Update(ug); err != nil {
		return fmt.Errorf("unable to update user group %s, %s", id, err)
	}

	return resourceUserGroupRead(d, m)
}

func resourceUserGroupDelete(d *schema.ResourceData, m interface{}) error {
	userGroups := m.(*wavefrontClient).client.UserGroups()

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

// Decodes the permissions from the state file and returns a []string of permissions
func resourceDecodeUserGroupPermissions(d *schema.ResourceData, userGroup *wavefront.UserGroup) {
	var existingPermissions *schema.Set
	var permissions []string
	if d.HasChange("permissions") {
		n := d.Get("permissions")

		// Largely fine if new is nil, likely means we're removing the user from all explicit permissions
		if n == nil {
			n = new(schema.Set)
		}
		existingPermissions = n.(*schema.Set)
	} else {
		existingPermissions = d.Get("permissions").(*schema.Set)
	}

	for _, permission := range existingPermissions.List() {
		permissions = append(permissions, permission.(string))
	}

	userGroup.Permissions = permissions
}
