package wavefront

import (
	"fmt"
	"strings"

	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEvent() *schema.Resource {
	return &schema.Resource{
		Create: resourceEventCreate,
		Read:   resourceEventRead,
		Update: resourceEventUpdate,
		Delete: resourceEventDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			nameKey: {
				Type:     schema.TypeString,
				Required: true,
			},
			startTimeKey: {
				Type:     schema.TypeInt,
				Optional: true,
			},
			endTimeKey: {
				Type:     schema.TypeInt,
				Optional: true,
			},
			annotationsKey: {
				Type:     schema.TypeMap,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			tagsKey: {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceEventRead(d *schema.ResourceData, meta interface{}) error {

	events := meta.(*wavefrontClient).client.Events()

	eventID := d.Id()
	tmpEvent, err := events.FindByID(eventID)

	if err != nil {
		if strings.Contains(err.Error(), "404") {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("unable to find Wavefront Event %s. %s", d.Id(), err)
	}

	d.SetId(*tmpEvent.ID)
	d.Set(nameKey, tmpEvent.Name)
	d.Set(startTimeKey, tmpEvent.StartTime)
	d.Set(endTimeKey, tmpEvent.EndTime)
	d.Set(tagsKey, tmpEvent.Tags)
	d.Set(annotationsKey, tmpEvent.Annotations)

	return nil

}

func resourceEventCreate(d *schema.ResourceData, meta interface{}) error {
	events := meta.(*wavefrontClient).client.Events()

	tags := decodeEventTags(d)
	event := &wavefront.Event{
		Name:        d.Get(nameKey).(string),
		StartTime:   int64(d.Get(startTimeKey).(int)),
		EndTime:     int64(d.Get(endTimeKey).(int)),
		Annotations: getStringMap(d, annotationsKey),
		Tags:        tags,
	}

	// Create the Event on Wavefront
	err := events.Create(event)
	if err != nil {
		return fmt.Errorf("error creating Event %s. %s", d.Get(nameKey), err)
	}

	d.SetId(*event.ID)
	return nil

}

func resourceEventUpdate(d *schema.ResourceData, meta interface{}) error {
	events := meta.(*wavefrontClient).client.Events()

	eventID := d.Id()
	newEvent, err := events.FindByID(eventID)

	if err != nil {
		d.SetId("")
		return nil
	}

	if d.HasChange(nameKey) {
		newEvent.Name = d.Get(nameKey).(string)
	}
	if d.HasChange(startTimeKey) {
		newEvent.StartTime = d.Get(startTimeKey).(int64)
	}
	if d.HasChange(endTimeKey) {
		newEvent.StartTime = d.Get(endTimeKey).(int64)
	}

	if d.HasChange(tagsKey) {
		newEvent.Tags = decodeEventTags(d)
	}

	if d.HasChange(annotationsKey) {
		newEvent.Annotations = getStringMap(d, annotationsKey)
	}

	err = events.Update(newEvent)
	if err != nil {
		return fmt.Errorf("unable to update Wavefront Event %s, %s", d.Get(nameKey), err)
	}

	return resourceEventRead(d, meta)
}

func resourceEventDelete(d *schema.ResourceData, meta interface{}) error {

	events := meta.(*wavefrontClient).client.Events()
	var err error

	eventID := d.Id()
	newEvent, err := events.FindByID(eventID)

	if err != nil {
		if strings.Contains(err.Error(), "404") {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("unable to find Wavefront Event %s. %s", d.Get(nameKey), err)
	}

	err = events.Delete(newEvent)
	if err != nil {
		return fmt.Errorf("error trying to delete Wavefront Event %s. %s", d.Get(nameKey), err)
	}

	d.SetId("")
	return nil
}

// Decodes the tags from the state and returns a []string of tags
func decodeEventTags(d *schema.ResourceData) []string {
	var tags []string
	for _, tag := range d.Get("tags").([]interface{}) {
		tags = append(tags, tag.(string))
	}
	return tags
}
