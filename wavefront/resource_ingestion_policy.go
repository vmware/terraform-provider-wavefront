package wavefront

import (
	"errors"
	"fmt"

	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	ipNameKey        = "name"
	ipDescriptionKey = "description"
	ipScopeKey       = "scope"

	ipAccountsKey   = "accounts"
	ipGroupsKey     = "groups"
	ipSourcesKey    = "sources"
	ipNamespacesKey = "namespaces"
	ipTagsKey       = "tags"
	ipTagKey        = "key"
	ipTagValue      = "value"

	ipValueKey = "value"
	ipKeyKey   = "key"
)

// Schema
func resourceIngestionPolicy() *schema.Resource {

	ingestionPolicyTagSchema := ingestionPolicyTagSchema()

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
			ipScopeKey: {
				Type:     schema.TypeString,
				Required: true,
			},
			ipAccountsKey: {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			ipGroupsKey: {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			ipSourcesKey: {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			ipNamespacesKey: {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			ipTagsKey: {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: ingestionPolicyTagSchema,
				},
			},
		},
	}
}

func ingestionPolicyTagSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		ipKeyKey: {
			Type:     schema.TypeString,
			Required: true,
		},
		ipValueKey: {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

func flattenIngestionPolicyAccountIDs(accounts []wavefront.IngestionPolicyAccount) []string {
	var tfMaps []string
	if accounts != nil && len(accounts) > 0 {
		for _, v := range accounts {
			tfMaps = append(tfMaps, v.ID)
		}
	}
	return tfMaps
}

func flattenIngestionPolicyGroupIDs(groups []wavefront.IngestionPolicyGroup) []string {
	var tfMaps []string
	if groups != nil && len(groups) > 0 {
		for _, v := range groups {
			tfMaps = append(tfMaps, v.ID)
		}
	}
	return tfMaps
}

func convertIngestionPolicyTagsToMap(raw []wavefront.IngestionPolicyTag) []map[string]string {
	var tags = make([]map[string]string, 0)
	if raw != nil {
		for _, r := range raw {
			tag := make(map[string]string)
			tag[ipTagKey] = r.Key
			tag[ipTagValue] = r.Value
			tags = append(tags, tag)
		}
	}
	return tags
}

// Helpers
func parseIngestionPolicyTags(raw interface{}) []map[string]string {
	var tags = make([]map[string]string, 0)
	if raw != nil {
		for _, r := range raw.([]interface{}) {
			v := r.(map[string]interface{})
			tag := make(map[string]string)
			tag[ipTagKey] = v[ipTagKey].(string)
			tag[ipTagValue] = v[ipTagValue].(string)
			tags = append(tags, tag)
		}
	}
	return tags
}

// CRUD
func resourceIngestionPolicyCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*wavefrontClient).client.IngestionPolicies()
	newPolicyRequest := wavefront.IngestionPolicyRequest{
		Name:        d.Get(ipNameKey).(string),
		Description: d.Get(ipDescriptionKey).(string),
		Scope:       d.Get(ipScopeKey).(string),
		Accounts:    parseStrArr(d.Get(ipAccountsKey)),
		Groups:      parseStrArr(d.Get(ipGroupsKey)),
		Sources:     parseStrArr(d.Get(ipSourcesKey)),
		Namespaces:  parseStrArr(d.Get(ipNamespacesKey)),
		Tags:        parseIngestionPolicyTags(d.Get(ipTagsKey)),
	}
	ingestionPolicy, err := client.Create(&newPolicyRequest)

	if err != nil {
		return fmt.Errorf("failed to create ingestion policy, %s", err)
	}

	d.SetId(ingestionPolicy.ID)
	return resourceIngestionPolicyRead(d, meta)
}

func resourceIngestionPolicyRead(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*wavefrontClient).client.IngestionPolicies()
	ingestionPolicy, err := client.GetByID(d.Id())

	if wavefront.NotFound(err) {
		d.SetId("")
		return nil
	}

	if err != nil {
		return fmt.Errorf("error finding ingestion policy, %s. %s", d.Id(), err)
	}

	if err = d.Set(ipNameKey, ingestionPolicy.Name); err != nil {
		return err
	}

	if err = d.Set(ipDescriptionKey, ingestionPolicy.Description); err != nil {
		return err
	}

	if err = d.Set(ipScopeKey, ingestionPolicy.Scope); err != nil {
		return err
	}

	scope := ingestionPolicy.Scope

	if scope == "ACCOUNT" {
		accounts := flattenIngestionPolicyAccountIDs(ingestionPolicy.Accounts)
		if len(accounts) < 1 {
			return errors.New("ingestion policy account scope must have at least one associated account")
		} else {
			if err = d.Set(ipAccountsKey, accounts); err != nil {
				return err
			}
		}
	}

	if scope == "GROUP" {
		groups := flattenIngestionPolicyGroupIDs(ingestionPolicy.Groups)
		if len(groups) < 1 {
			return errors.New("ingestion policy group scope must have at least one associated group")
		} else {
			if err = d.Set(ipGroupsKey, groups); err != nil {
				return err
			}
		}
	}

	if scope == "SOURCES" {
		if err = d.Set(ipSourcesKey, ingestionPolicy.Sources); err != nil {
			return err
		}
	}

	if scope == "NAMESPACES" {
		if err = d.Set(ipNamespacesKey, ingestionPolicy.Namespaces); err != nil {
			return err
		}
	}

	if scope == "TAGS" {
		tags := convertIngestionPolicyTagsToMap(ingestionPolicy.Tags)
		if err = d.Set(ipTagsKey, tags); err != nil {
			return err
		}
	}

	return nil
}

func resourceIngestionPolicyUpdate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*wavefrontClient).client.IngestionPolicies()
	policy, err := client.GetByID(d.Id())

	if wavefront.NotFound(err) {
		d.SetId("")
		return nil
	}

	policy.Name = d.Get(ipNameKey).(string)
	policy.Description = d.Get(ipDescriptionKey).(string)
	policy.Scope = d.Get(ipScopeKey).(string)
	policy.Accounts = d.Get(ipAccountsKey).([]wavefront.IngestionPolicyAccount)
	policy.Groups = d.Get(ipGroupsKey).([]wavefront.IngestionPolicyGroup)
	policy.Sources = d.Get(ipSourcesKey).([]string)
	policy.Namespaces = d.Get(ipNamespacesKey).([]string)
	policy.Tags = d.Get(ipTagsKey).([]wavefront.IngestionPolicyTag)

	if err != nil {
		return fmt.Errorf(""+"error finding ingestion policy, %s. %s", d.Id(), err)
	}

	err = client.Update(policy)

	if err != nil {
		return fmt.Errorf("error updating ingestion policy,  %s. %s", d.Id(), err)
	}

	return resourceIngestionPolicyRead(d, meta)
}

func resourceIngestionPolicyDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*wavefrontClient).client.IngestionPolicies()
	err := client.DeleteByID(d.Id())

	if err != nil && !wavefront.NotFound(err) {
		return fmt.Errorf("error deleting ingestion policy, %s. %s", d.Id(), err)
	}

	d.SetId("")

	return nil
}
