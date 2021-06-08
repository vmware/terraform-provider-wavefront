package wavefront

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func suppressCase(k, old, new string, d *schema.ResourceData) bool {
	return strings.EqualFold(old, new)
}

func suppressSpaces(k, old, new string, d *schema.ResourceData) bool {
	return strings.TrimSpace(old) == strings.TrimSpace(new)
}

func isJSONForFieldTheSame(_, old, new string, _ *schema.ResourceData) bool {
	var oldJSON interface{}
	var newJSON interface{}

	if err := json.Unmarshal([]byte(old), &oldJSON); err != nil {
		return false
	}
	if err := json.Unmarshal([]byte(new), &newJSON); err != nil {
		return false
	}
	return reflect.DeepEqual(oldJSON, newJSON)
}

func trimSpaces(d interface{}) string {
	if s, ok := d.(string); ok {
		return strings.TrimSpace(s)
	}

	return ""
}

func trimSpacesMap(m map[string]interface{}) map[string]string {
	trimmed := map[string]string{}
	for key, v := range m {
		trimmed[key] = trimSpaces(v)
	}
	return trimmed
}

// Decodes the ACL from the state
func decodeAccessControlList(d *schema.ResourceData) (canView, canModify []string) {
	for _, cv := range d.Get("can_view").(*schema.Set).List() {
		canView = append(canView, cv.(string))
	}

	for _, cv := range d.Get("can_modify").(*schema.Set).List() {
		canModify = append(canModify, cv.(string))
	}

	return canView, canModify
}

// Decodes the tags from the state and returns a []string of tags
func decodeTags(d *schema.ResourceData) (tags []string) {
	for _, tag := range d.Get("tags").(*schema.Set).List() {
		tags = append(tags, tag.(string))
	}
	return tags
}

// Decodes a TypeList of []interface{} to []string
func decodeTypeListToString(d *schema.ResourceData, field string) []string {
	var decoded []string
	encoded := d.Get(field).([]interface{})

	for _, v := range encoded {
		decoded = append(decoded, fmt.Sprint(v))
	}

	return decoded
}

// Decodes a TypeMap of map[string]interface{} into map[string]string for binding to the API
func decodeTypeMapToStringMap(d *schema.ResourceData, field string) map[string]string {
	decoded := map[string]string{}
	if encoded, ok := d.GetOk(field); ok {
		for k, v := range encoded.(map[string]interface{}) {
			decoded[k] = fmt.Sprint(v)
		}
	}
	return decoded
}
