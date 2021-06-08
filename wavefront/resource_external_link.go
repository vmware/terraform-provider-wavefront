package wavefront

import (
	"fmt"

	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	elNameKey                  = "name"
	elDescriptionKey           = "description"
	elTemplateKey              = "template"
	elMetricFilterRegexKey     = "metric_filter_regex"
	elSourceFilterRegexKey     = "source_filter_regex"
	elPointTagFilterRegexesKey = "point_tag_filter_regexes"
	elIsLogIntegrationKey      = "is_log_integration"
)

func resourceExternalLink() *schema.Resource {
	return &schema.Resource{
		Create: resourceExternalLinkCreate,
		Read:   resourceExternalLinkRead,
		Update: resourceExternalLinkUpdate,
		Delete: resourceExternalLinkDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			elNameKey: {
				Type:     schema.TypeString,
				Required: true,
			},
			elDescriptionKey: {
				Type:     schema.TypeString,
				Required: true,
			},
			elTemplateKey: {
				Type:     schema.TypeString,
				Required: true,
			},
			elMetricFilterRegexKey: {
				Type:     schema.TypeString,
				Optional: true,
			},
			elSourceFilterRegexKey: {
				Type:     schema.TypeString,
				Optional: true,
			},
			elPointTagFilterRegexesKey: {
				Type:     schema.TypeMap,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			elIsLogIntegrationKey: {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func resourceExternalLinkCreate(
	d *schema.ResourceData, meta interface{}) error {
	externalLinks := meta.(*wavefrontClient).client.ExternalLinks()

	externalLink := wavefront.ExternalLink{
		Name:                  d.Get(elNameKey).(string),
		Description:           d.Get(elDescriptionKey).(string),
		Template:              d.Get(elTemplateKey).(string),
		MetricFilterRegex:     d.Get(elMetricFilterRegexKey).(string),
		SourceFilterRegex:     d.Get(elSourceFilterRegexKey).(string),
		PointTagFilterRegexes: getStringMap(d, elPointTagFilterRegexesKey),
		IsLogIntegration:      d.Get(elIsLogIntegrationKey).(bool),
	}
	err := externalLinks.Create(&externalLink)
	if err != nil {
		return fmt.Errorf(
			"failed to create new Wavefront External Link, %s", err)
	}
	d.SetId(*externalLink.ID)
	return resourceExternalLinkRead(d, meta)
}

func resourceExternalLinkRead(
	d *schema.ResourceData, meta interface{}) error {
	externalLinks := meta.(*wavefrontClient).client.ExternalLinks()
	id := d.Id()
	el := wavefront.ExternalLink{ID: &id}
	err := externalLinks.Get(&el)
	if wavefront.NotFound(err) {
		d.SetId("")
		return nil
	}
	if err != nil {
		return fmt.Errorf(
			"error finding Wavefront External Link, %s. %s",
			d.Id(),
			err)
	}
	if err := d.Set(elNameKey, el.Name); err != nil {
		return err
	}
	if err := d.Set(elDescriptionKey, el.Description); err != nil {
		return err
	}
	if err := d.Set(elTemplateKey, el.Template); err != nil {
		return err
	}
	if err := d.Set(elMetricFilterRegexKey, el.MetricFilterRegex); err != nil {
		return err
	}
	if err := d.Set(elSourceFilterRegexKey, el.SourceFilterRegex); err != nil {
		return err
	}
	err = setStringMap(d, elPointTagFilterRegexesKey, el.PointTagFilterRegexes)
	if err != nil {
		return err
	}
	return d.Set(elIsLogIntegrationKey, el.IsLogIntegration)
}

func resourceExternalLinkUpdate(
	d *schema.ResourceData, meta interface{}) error {
	externalLinks := meta.(*wavefrontClient).client.ExternalLinks()
	id := d.Id()
	el := wavefront.ExternalLink{ID: &id}
	err := externalLinks.Get(&el)
	if wavefront.NotFound(err) {
		d.SetId("")
		return nil
	}
	if err != nil {
		return fmt.Errorf(""+
			"error finding Wavefront External Link, %s. %s",
			d.Id(),
			err)
	}
	if d.HasChange(elNameKey) {
		el.Name = d.Get(elNameKey).(string)
	}
	if d.HasChange(elDescriptionKey) {
		el.Description = d.Get(elDescriptionKey).(string)
	}
	if d.HasChange(elTemplateKey) {
		el.Template = d.Get(elTemplateKey).(string)
	}
	if d.HasChange(elMetricFilterRegexKey) {
		el.MetricFilterRegex = d.Get(elMetricFilterRegexKey).(string)
	}
	if d.HasChange(elSourceFilterRegexKey) {
		el.SourceFilterRegex = d.Get(elSourceFilterRegexKey).(string)
	}
	if d.HasChange(elPointTagFilterRegexesKey) {
		el.PointTagFilterRegexes = getStringMap(d, elPointTagFilterRegexesKey)
	}
	if d.HasChange(elIsLogIntegrationKey) {
		el.IsLogIntegration = d.Get(elIsLogIntegrationKey).(bool)
	}
	err = externalLinks.Update(&el)
	if err != nil {
		return fmt.Errorf(
			"error updating Wavefront External Link,  %s. %s",
			d.Id(),
			err,
		)
	}
	return resourceExternalLinkRead(d, meta)
}

func resourceExternalLinkDelete(
	d *schema.ResourceData, meta interface{}) error {
	externalLinks := meta.(*wavefrontClient).client.ExternalLinks()
	id := d.Id()
	el := wavefront.ExternalLink{ID: &id}
	err := externalLinks.Delete(&el)
	if err != nil && !wavefront.NotFound(err) {
		return fmt.Errorf(
			"error deleting Wavefront External Link, %s. %s",
			d.Id(),
			err,
		)
	}
	d.SetId("")
	return nil
}
