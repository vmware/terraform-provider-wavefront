package wavefront

import (
	"fmt"
	"time"

	"github.com/WavefrontHQ/go-wavefront-management-api/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEvent() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceEventRead,
		Schema: dataSourceEventSchema(),
	}
}

func dataSourceEventSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		nameKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		idKey: {
			Type:     schema.TypeString,
			Required: true,
		},

		startTimeKey: {
			Type:     schema.TypeInt,
			Computed: true,
		},

		endTimeKey: {
			Type:     schema.TypeInt,
			Computed: true,
		},

		severityKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		typeKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		detailsKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		isEphemeralKey: {
			Type:     schema.TypeBool,
			Computed: true,
		},

		annotationsKey: {
			Type:     schema.TypeMap,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},

		tagsKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
}

func dataSourceEventRead(d *schema.ResourceData, m interface{}) error {
	eventClient := m.(*wavefrontClient).client.Events()
	id, ok := d.GetOk("id")
	if !ok {
		return fmt.Errorf("required parameter '%s' not set", idKey)
	}

	idStr := fmt.Sprintf("%s", id)

	var event *wavefront.Event
	var err error
	if event, err = eventClient.FindByID(idStr); err != nil {
		return err
	}

	d.SetId(time.Now().UTC().String())
	return setEventAttributes(d, *event)
}

func setEventAttributes(d *schema.ResourceData, event wavefront.Event) error {
	if err := d.Set(idKey, *event.ID); err != nil {
		return err
	}
	if err := d.Set(nameKey, event.Name); err != nil {
		return err
	}
	if err := d.Set(startTimeKey, event.StartTime); err != nil {
		return err
	}
	if err := d.Set(endTimeKey, event.EndTime); err != nil {
		return err
	}
	if err := d.Set(severityKey, event.Severity); err != nil {
		return err
	}
	if err := d.Set(typeKey, event.Type); err != nil {
		return err
	}
	if err := d.Set(detailsKey, event.Details); err != nil {
		return err
	}
	if err := d.Set(isEphemeralKey, event.Instantaneous); err != nil {
		return err
	}
	if err := d.Set(annotationsKey, event.Annotations); err != nil {
		return err
	}
	return d.Set(tagsKey, event.Tags)
}
