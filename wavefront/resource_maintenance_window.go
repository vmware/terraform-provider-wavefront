package wavefront

import (
	"fmt"

	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	mwReasonKey                          = "reason"
	mwTitleKey                           = "title"
	mwStartTimeInSecondsKey              = "start_time_in_seconds"
	mwEndTimeInSecondsKey                = "end_time_in_seconds"
	mwRelevantCustomerTagsKey            = "relevant_customer_tags"
	mwRelevantHostTagsKey                = "relevant_host_tags"
	mwRelevantHostNamesKey               = "relevant_host_names"
	mwRelevantHostTagsAndedKey           = "relevant_host_tags_anded"
	mwHostTagGroupHostNamesGroupAndedKey = "host_tag_group_host_names_group_anded"
)

func resourceMaintenanceWindow() *schema.Resource {
	return &schema.Resource{
		Create: resourceMaintenanceWindowCreate,
		Read:   resourceMaintenanceWindowRead,
		Update: resourceMaintenanceWindowUpdate,
		Delete: resourceMaintenanceWindowDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			mwReasonKey: {
				Type:     schema.TypeString,
				Required: true,
			},
			mwTitleKey: {
				Type:     schema.TypeString,
				Required: true,
			},
			mwStartTimeInSecondsKey: {
				Type:     schema.TypeInt,
				Required: true,
			},
			mwEndTimeInSecondsKey: {
				Type:     schema.TypeInt,
				Required: true,
			},
			mwRelevantCustomerTagsKey: {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set:      schema.HashString,
			},
			mwRelevantHostTagsKey: {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set:      schema.HashString,
			},
			mwRelevantHostNamesKey: {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set:      schema.HashString,
			},
			mwRelevantHostTagsAndedKey: {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			mwHostTagGroupHostNamesGroupAndedKey: {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourceMaintenanceWindowCreate(
	d *schema.ResourceData, meta interface{}) error {
	maintenanceWindows := meta.(*wavefrontClient).client.MaintenanceWindows()

	mw, err := maintenanceWindows.Create(
		&wavefront.MaintenanceWindowOptions{
			Reason:                          d.Get(mwReasonKey).(string),
			Title:                           d.Get(mwTitleKey).(string),
			StartTimeInSeconds:              int64(d.Get(mwStartTimeInSecondsKey).(int)),
			EndTimeInSeconds:                int64(d.Get(mwEndTimeInSecondsKey).(int)),
			RelevantCustomerTags:            getStringSlice(d, mwRelevantCustomerTagsKey),
			RelevantHostTags:                getStringSlice(d, mwRelevantHostTagsKey),
			RelevantHostNames:               getStringSlice(d, mwRelevantHostNamesKey),
			RelevantHostTagsAnded:           d.Get(mwRelevantHostTagsAndedKey).(bool),
			HostTagGroupHostNamesGroupAnded: d.Get(mwHostTagGroupHostNamesGroupAndedKey).(bool),
		})
	if err != nil {
		return fmt.Errorf(
			"failed to create new Wavefront Maintenance Window, %s",
			err)
	}
	d.SetId(mw.ID)

	return resourceMaintenanceWindowRead(d, meta)
}

func resourceMaintenanceWindowRead(
	d *schema.ResourceData, meta interface{}) error {
	maintenanceWindows := meta.(*wavefrontClient).client.MaintenanceWindows()
	mw, err := maintenanceWindows.GetByID(d.Id())
	if wavefront.NotFound(err) {
		d.SetId("")
		return nil
	}
	if err != nil {
		return fmt.Errorf(
			"error finding Wavefront Maintenance Window %s. %s",
			d.Id(),
			err)
	}
	if err := d.Set(mwReasonKey, mw.Reason); err != nil {
		return err
	}
	if err := d.Set(mwTitleKey, mw.Title); err != nil {
		return err
	}
	if err := d.Set(mwStartTimeInSecondsKey, int(mw.StartTimeInSeconds)); err != nil {
		return err
	}
	if err := d.Set(mwEndTimeInSecondsKey, int(mw.EndTimeInSeconds)); err != nil {
		return err
	}
	err = setStringSlice(d, mwRelevantCustomerTagsKey, mw.RelevantCustomerTags)
	if err != nil {
		return err
	}
	err = setStringSlice(d, mwRelevantHostTagsKey, mw.RelevantHostTags)
	if err != nil {
		return err
	}
	err = setStringSlice(d, mwRelevantHostNamesKey, mw.RelevantHostNames)
	if err != nil {
		return err
	}
	err = d.Set(mwRelevantHostTagsAndedKey, mw.RelevantHostTagsAnded)
	if err != nil {
		return err
	}
	return d.Set(mwHostTagGroupHostNamesGroupAndedKey, mw.HostTagGroupHostNamesGroupAnded)
}

func resourceMaintenanceWindowUpdate(
	d *schema.ResourceData, meta interface{}) error {
	maintenanceWindows := meta.(*wavefrontClient).client.MaintenanceWindows()
	mw, err := maintenanceWindows.GetByID(d.Id())
	if wavefront.NotFound(err) {
		d.SetId("")
		return nil
	}
	if err != nil {
		return fmt.Errorf(""+
			"error finding Wavefront Maintenance Window %s. %s",
			d.Id(),
			err)
	}
	options := mw.Options()
	if d.HasChange(mwReasonKey) {
		options.Reason = d.Get(mwReasonKey).(string)
	}
	if d.HasChange(mwTitleKey) {
		options.Title = d.Get(mwTitleKey).(string)
	}
	if d.HasChange(mwStartTimeInSecondsKey) {
		options.StartTimeInSeconds = int64(d.Get(mwStartTimeInSecondsKey).(int))
	}
	if d.HasChange(mwEndTimeInSecondsKey) {
		options.EndTimeInSeconds = int64(d.Get(mwEndTimeInSecondsKey).(int))
	}
	if d.HasChange(mwRelevantCustomerTagsKey) {
		options.RelevantCustomerTags = getStringSlice(d, mwRelevantCustomerTagsKey)
	}
	if d.HasChange(mwRelevantHostTagsKey) {
		options.RelevantHostTags = getStringSlice(d, mwRelevantHostTagsKey)
	}
	if d.HasChange(mwRelevantHostNamesKey) {
		options.RelevantHostNames = getStringSlice(d, mwRelevantHostNamesKey)
	}
	if d.HasChange(mwRelevantHostTagsAndedKey) {
		options.RelevantHostTagsAnded = d.Get(mwRelevantHostTagsAndedKey).(bool)
	}
	if d.HasChange(mwHostTagGroupHostNamesGroupAndedKey) {
		options.HostTagGroupHostNamesGroupAnded = d.Get(mwHostTagGroupHostNamesGroupAndedKey).(bool)
	}
	_, err = maintenanceWindows.Update(mw.ID, options)
	if err != nil {
		return fmt.Errorf(
			"error updating Wavefront Maintenance Window  %s. %s",
			d.Id(),
			err,
		)
	}
	return resourceMaintenanceWindowRead(d, meta)
}

func resourceMaintenanceWindowDelete(
	d *schema.ResourceData, meta interface{}) error {
	maintenanceWindows := meta.(*wavefrontClient).client.MaintenanceWindows()
	err := maintenanceWindows.DeleteByID(d.Id())
	if err != nil && !wavefront.NotFound(err) {
		return fmt.Errorf(
			"error deleting Wavefront Maintenance Window %s. %s",
			d.Id(),
			err,
		)
	}
	d.SetId("")
	return nil
}
