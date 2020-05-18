package wavefront_plugin

import (
	"fmt"
	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strings"
)

func resourceRole() *schema.Resource {
	return &schema.Resource{
		Create: resourceRoleCreate,
		Read:   resourceRoleRead,
		Update: resourceRoleUpdate,
		Delete: resourceRoleDelete,
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
				Optional: true,
			},
			"permissions": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"assignees": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func getPermissions(d *schema.ResourceData) ([]string, []string) {
	return getChangedLists(d, "permissions")
}

func getAssignees(d *schema.ResourceData) ([]string, []string) {
	return getChangedLists(d, "assignees")
}

func getChangedLists(d *schema.ResourceData, key string) ([]string, []string) {
	old, new := d.GetChange(key)
	var oldP, newP []string
	for _, o := range old.(*schema.Set).List() {
		oldP = append(oldP, fmt.Sprint(o))
	}
	for _, n := range new.(*schema.Set).List() {
		newP = append(newP, fmt.Sprint(n))
	}

	return oldP, newP
}

func resourceRoleCreate(d *schema.ResourceData, meta interface{}) error {
	r := meta.(*wavefrontClient).client.Roles()

	_, permissions := getPermissions(d)
	_, assignees := getAssignees(d)

	role := &wavefront.Role{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Permissions: permissions,
	}

	err := r.Create(role)
	if err != nil {
		return fmt.Errorf("error trying to create role %s. %s", role.Name, err)
	}
	d.SetId(role.ID)

	if len(assignees) > 0 {
		err = r.AddAssignees(assignees, role)
		if err != nil {
			return fmt.Errorf("error trying to add assignees %v on role %s. %s", assignees, role.ID, err)
		}
	}

	return resourceRoleRead(d, meta)
}

func resourceRoleRead(d *schema.ResourceData, meta interface{}) error {
	r := meta.(*wavefrontClient).client.Roles()
	roles, err := r.Find([]*wavefront.SearchCondition{
		{
			Key:            "id",
			Value:          d.Id(),
			MatchingMethod: "EXACT",
		},
	})
	if err != nil {
		if strings.Contains(err.Error(), "406") {
			d.SetId("")
			return nil
		}
		return err
	}

	if len(roles) == 0 {
		d.SetId("")
		return nil
	}

	role := roles[0]

	log.Printf("[INFO] permissions identified on role: %s", role.Permissions)
	d.SetId(role.ID)
	d.Set("name", role.Name)
	d.Set("description", role.Description)
	d.Set("permissions", role.Permissions)

	return nil
}

func resourceRoleUpdate(d *schema.ResourceData, meta interface{}) error {
	r := meta.(*wavefrontClient).client.Roles()

	_, np := getPermissions(d)
	oa, na := getAssignees(d)

	var removeAssignees []string
	var found bool

	for _, o := range oa {
		found = false
		for _, n := range na {
			if o == n {
				found = true
				break
			}
		}
		if !found {
			removeAssignees = append(removeAssignees, o)
		}
	}

	role := &wavefront.Role{
		ID:          d.Id(),
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Permissions: np,
	}

	err := r.Update(role)
	if err != nil {
		return err
	}

	if len(na) > 0 {
		err = r.AddAssignees(na, role)
		if err != nil {
			return fmt.Errorf("error trying to add assignees %v on role %s. %s", na, role.ID, err)
		}
	}

	if len(removeAssignees) > 0 {
		err = r.RemoveAssignees(removeAssignees, role)
		if err != nil {
			// Endpoint will swallow errors if some are bad and others are not, but otherwise will throw an error
			// when all assignees to remove are bad...
			if !strings.Contains(err.Error(), "No valid user or user group IDs were found") {
				return fmt.Errorf("error trying to remove assignees %v on role %s. %s", removeAssignees, role.ID, err)
			}
		}
	}

	for _, p := range np {
		err = r.GrantPermission(p, []*wavefront.Role{role})
		if err != nil {
			return fmt.Errorf("error trying to grant permission %s on role %s. %s", p, role.ID, err)
		}
	}

	return resourceRoleRead(d, meta)
}

func resourceRoleDelete(d *schema.ResourceData, meta interface{}) error {
	r := meta.(*wavefrontClient).client.Roles()
	roles, err := r.Find([]*wavefront.SearchCondition{
		{
			Key:            "id",
			Value:          d.Id(),
			MatchingMethod: "EXACT",
		},
	})
	if err != nil {
		if strings.Contains(err.Error(), "406") {
			d.SetId("")
			return nil
		}
		return err
	}

	if len(roles) == 0 {
		d.SetId("")
		return nil
	}

	role := roles[0]
	err = r.Delete(role)
	if err != nil {
		return err
	}

	return resourceRoleRead(d, meta)
}
