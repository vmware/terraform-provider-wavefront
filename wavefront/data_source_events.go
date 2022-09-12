package wavefront

import (
	"time"

	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEvents() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceEventsRead,
		Schema: dataSourceEventsSchema(),
	}
}

func dataSourceEventsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		eventsKey: {
			Type:     schema.TypeSet,
			Computed: true,
			Elem: &schema.Resource{
				Schema: dataSourceEventEntitySchema(),
			},
		},
	}
}

func dataSourceEventEntitySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		nameKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		idKey: {
			Type:     schema.TypeString,
			Computed: true,
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

func dataSourceEventsRead(d *schema.ResourceData, m interface{}) error {

	var allEvents []*wavefront.Event
	eventClient := m.(*wavefrontClient).client.Events()

	// Data Source ID is set to current time to always refresh
	d.SetId(time.Now().UTC().String())

	cont := true
	offset := 0
	earliestStartTimeEpochMillis := d.Get("earliestStartTimeEpochMillis").(int64)
	latestStartTimeEpochMillis := d.Get("latestStartTimeEpochMillis").(int64)
	timeRange := wavefront.TimeRange{
		StartTime: earliestStartTimeEpochMillis,
		EndTime:   latestStartTimeEpochMillis,
	}

	for cont {
		filter := []*wavefront.SearchCondition{
			{Key: "limit", Value: string(rune(pageSize)), MatchingMethod: exactMatching},
			{Key: "offset", Value: string(rune(offset)), MatchingMethod: exactMatching},
		}

		events, err := eventClient.Find(filter, &timeRange)
		if err != nil {
			return err
		}

		allEvents = append(allEvents, events...)

		if len(allEvents) < pageSize {
			cont = false
		} else {
			offset += pageSize
		}
	}

	if err := d.Set("alerts", flattenEvents(allEvents)); err != nil {
		return err
	}
	return nil
}

func flattenEvents(events []*wavefront.Event) interface{} {
	tfMaps := make([]map[string]interface{}, len(events))
	for i, v := range events {
		tfMaps[i] = flattenEvent(v)
	}
	return tfMaps
}

func flattenEvent(event *wavefront.Event) map[string]interface{} {
	tfMap := make(map[string]interface{})
	tfMap[idKey] = *event.ID
	tfMap[nameKey] = event.Name
	tfMap[startTimeKey] = event.StartTime
	tfMap[endTimeKey] = event.EndTime
	tfMap[severityKey] = event.Severity
	tfMap[typeKey] = event.Type
	tfMap[detailsKey] = event.Details
	tfMap[isEphemeralKey] = event.Instantaneous
	tfMap[annotationsKey] = event.Annotations
	tfMap[tagsKey] = event.Tags
	return tfMap
}
