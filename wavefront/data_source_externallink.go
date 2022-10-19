package wavefront

import (
	"fmt"
	"time"

	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	templateKey              = "template"
	isLogIntegrationKey      = "is_log_integration"
	metricFilterRegexKey     = "metric_filter_regex"
	sourceFilterRegexKey     = "source_filter_regex"
	pointTagFilterRegexesKey = "point_tag_filter_regexes"
	externalLinksKey         = "external_links"
)

func dataSourceExternalLink() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceExternalLinkRead,
		Schema: dataSourceExternalLinkSchema(),
	}
}

func dataSourceExternalLinkSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		nameKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		idKey: {
			Type:     schema.TypeString,
			Required: true,
		},

		descriptionKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		creatorIDKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		updaterIDKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		createdEpochMillisKey: {
			Type:     schema.TypeInt,
			Computed: true,
		},

		updatedEpochMillisKey: {
			Type:     schema.TypeInt,
			Computed: true,
		},

		isLogIntegrationKey: {
			Type:     schema.TypeBool,
			Computed: true,
		},

		templateKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		metricFilterRegexKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		sourceFilterRegexKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		pointTagFilterRegexesKey: {
			Type:     schema.TypeMap,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
}

func dataSourceExternalLinkRead(d *schema.ResourceData, m interface{}) error {
	externalLinkClient := m.(*wavefrontClient).client.ExternalLinks()
	id, ok := d.GetOk("id")
	if !ok {
		return fmt.Errorf("required parameter '%s' not set", idKey)
	}

	idStr := fmt.Sprintf("%s", id)
	extLink := wavefront.ExternalLink{ID: &idStr}
	if err := externalLinkClient.Get(&extLink); err != nil {
		return err
	}

	// Data Source ID is set to current time to always refresh
	d.SetId(time.Now().UTC().String())
	return setExternalLinkAttributes(d, extLink)
}

func setExternalLinkAttributes(d *schema.ResourceData, extLink wavefront.ExternalLink) error {
	if err := d.Set(idKey, *extLink.ID); err != nil {
		return err
	}
	if err := d.Set(nameKey, extLink.Name); err != nil {
		return err
	}
	if err := d.Set(descriptionKey, extLink.Description); err != nil {
		return err
	}
	if err := d.Set(creatorIDKey, extLink.CreatorId); err != nil {
		return err
	}
	if err := d.Set(updaterIDKey, extLink.UpdaterId); err != nil {
		return err
	}
	if err := d.Set(createdEpochMillisKey, extLink.CreatedEpochMillis); err != nil {
		return err
	}
	if err := d.Set(updatedEpochMillisKey, extLink.UpdatedEpochMillis); err != nil {
		return err
	}
	if err := d.Set(templateKey, extLink.Template); err != nil {
		return err
	}
	if err := d.Set(isLogIntegrationKey, extLink.IsLogIntegration); err != nil {
		return err
	}
	if err := d.Set(metricFilterRegexKey, extLink.MetricFilterRegex); err != nil {
		return err
	}
	if err := d.Set(sourceFilterRegexKey, extLink.SourceFilterRegex); err != nil {
		return err
	}
	if err := d.Set(pointTagFilterRegexesKey, extLink.PointTagFilterRegexes); err != nil {
		return err
	}

	return nil
}
