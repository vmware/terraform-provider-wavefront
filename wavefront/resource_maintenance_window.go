package wavefront

import (
	"fmt"

	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	reasonKey                          = "reason"
	titleKey                           = "title"
	startTimeInSecondsKey              = "start_time_in_seconds"
	endTimeInSecondsKey                = "end_time_in_seconds"
	relevantCustomerTagsKey            = "relevant_customer_tags"
	relevantHostTagsKey                = "relevant_host_tags"
	relevantHostNamesKey               = "relevant_host_names"
	relevantHostTagsAndedKey           = "relevant_host_tags_anded"
	hostTagGroupHostNamesGroupAndedKey = "host_tag_group_host_names_group_anded"
)

func resourceMaintenanceWindow() *schema.Resource {
	return &schema.Resource{
		Create: resourceMaintenanceWindowCreate,
		Read:   resourceMaintenanceWindowRead,
		Update: resourceMaintenanceWindowUpdate,
		Delete: resourceMaintenanceWindowDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			reasonKey: {
				Type:     schema.TypeString,
				Required: true,
			},
			titleKey: {
				Type:     schema.TypeString,
				Required: true,
			},
			startTimeInSecondsKey: {
				Type:     schema.TypeInt,
				Required: true,
			},
			endTimeInSecondsKey: {
				Type:     schema.TypeInt,
				Required: true,
			},
			relevantCustomerTagsKey: {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set:      schema.HashString,
			},
			relevantHostTagsKey: {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set:      schema.HashString,
			},
			relevantHostNamesKey: {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set:      schema.HashString,
			},
			relevantHostTagsAndedKey: {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			hostTagGroupHostNamesGroupAndedKey: {
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
			Reason:                          d.Get(reasonKey).(string),
			Title:                           d.Get(titleKey).(string),
			StartTimeInSeconds:              int64(d.Get(startTimeInSecondsKey).(int)),
			EndTimeInSeconds:                int64(d.Get(endTimeInSecondsKey).(int)),
			RelevantCustomerTags:            getStringSlice(d, relevantCustomerTagsKey),
			RelevantHostTags:                getStringSlice(d, relevantHostTagsKey),
			RelevantHostNames:               getStringSlice(d, relevantHostNamesKey),
			RelevantHostTagsAnded:           d.Get(relevantHostTagsAndedKey).(bool),
			HostTagGroupHostNamesGroupAnded: d.Get(hostTagGroupHostNamesGroupAndedKey).(bool),
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
	if err := d.Set(reasonKey, mw.Reason); err != nil {
		return err
	}
	if err := d.Set(titleKey, mw.Title); err != nil {
		return err
	}
	if err := d.Set(startTimeInSecondsKey, int(mw.StartTimeInSeconds)); err != nil {
		return err
	}
	if err := d.Set(endTimeInSecondsKey, int(mw.EndTimeInSeconds)); err != nil {
		return err
	}
	err = setStringSlice(d, relevantCustomerTagsKey, mw.RelevantCustomerTags)
	if err != nil {
		return err
	}
	err = setStringSlice(d, relevantHostTagsKey, mw.RelevantHostTags)
	if err != nil {
		return err
	}
	err = setStringSlice(d, relevantHostNamesKey, mw.RelevantHostNames)
	if err != nil {
		return err
	}
	err = d.Set(relevantHostTagsAndedKey, mw.RelevantHostTagsAnded)
	if err != nil {
		return err
	}
	return d.Set(hostTagGroupHostNamesGroupAndedKey, mw.HostTagGroupHostNamesGroupAnded)
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
	if d.HasChange(reasonKey) {
		options.Reason = d.Get(reasonKey).(string)
	}
	if d.HasChange(titleKey) {
		options.Title = d.Get(titleKey).(string)
	}
	if d.HasChange(startTimeInSecondsKey) {
		options.StartTimeInSeconds = int64(d.Get(startTimeInSecondsKey).(int))
	}
	if d.HasChange(endTimeInSecondsKey) {
		options.EndTimeInSeconds = int64(d.Get(endTimeInSecondsKey).(int))
	}
	if d.HasChange(relevantCustomerTagsKey) {
		options.RelevantCustomerTags = getStringSlice(d, relevantCustomerTagsKey)
	}
	if d.HasChange(relevantHostTagsKey) {
		options.RelevantHostTags = getStringSlice(d, relevantHostTagsKey)
	}
	if d.HasChange(relevantHostNamesKey) {
		options.RelevantHostNames = getStringSlice(d, relevantHostNamesKey)
	}
	if d.HasChange(relevantHostTagsAndedKey) {
		options.RelevantHostTagsAnded = d.Get(relevantHostTagsAndedKey).(bool)
	}
	if d.HasChange(hostTagGroupHostNamesGroupAndedKey) {
		options.HostTagGroupHostNamesGroupAnded = d.Get(hostTagGroupHostNamesGroupAndedKey).(bool)
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
