package wavefront

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceExternalLinks() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceExternalLinksRead,
		Schema: dataSourceExternalLinksSchema(),
	}
}

func dataSourceExternalLinksSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		// Computed Values
		externalLinksKey: {
			Type:     schema.TypeSet,
			Computed: true,
			Elem: &schema.Resource{
				Schema: dataSourceExternalLinkSchema(),
			},
		},
		limitKey: {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  100,
		},
		offsetKey: {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
	}

}

func dataSourceExternalLinksRead(d *schema.ResourceData, m interface{}) error {
	var allExternalLinks []*wavefront.ExternalLink

	limit := d.Get(limitKey).(int)
	offset := d.Get(offsetKey).(int)

	if err := json.Unmarshal(searchAll(limit, offset, "extlink", nil, nil, m), &allExternalLinks); err != nil {
		return fmt.Errorf("Response is invalid JSON")
	}

	// Data Source ID is set to current time to always refresh
	d.SetId(time.Now().UTC().String())

	return d.Set(externalLinksKey, flattenExternalLinks(allExternalLinks))
}

func flattenExternalLinks(externalLinks []*wavefront.ExternalLink) []map[string]interface{} {
	tfMaps := make([]map[string]interface{}, len(externalLinks))
	for i, v := range externalLinks {
		tfMaps[i] = flattenExternalLink(v)
	}
	return tfMaps
}

func flattenExternalLink(externalLink *wavefront.ExternalLink) map[string]interface{} {
	tfMap := make(map[string]interface{})
	tfMap[idKey] = *externalLink.ID
	tfMap[nameKey] = externalLink.Name
	tfMap[descriptionKey] = externalLink.Description
	tfMap[creatorIDKey] = externalLink.CreatorId
	tfMap[updaterIDKey] = externalLink.UpdaterId
	tfMap[createdEpochMillisKey] = externalLink.CreatedEpochMillis
	tfMap[updatedEpochMillisKey] = externalLink.UpdatedEpochMillis
	tfMap[templateKey] = externalLink.Template
	tfMap[isLogIntegrationKey] = externalLink.IsLogIntegration
	tfMap[metricFilterRegexKey] = externalLink.MetricFilterRegex
	tfMap[sourceFilterRegexKey] = externalLink.SourceFilterRegex
	tfMap[pointTagFilterRegexesKey] = externalLink.PointTagFilterRegexes

	return tfMap
}
