package wavefront

import (
	"fmt"
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
		latestStartTimeEpochMillis: {
			Type:     schema.TypeInt,
			Required: true,
		},
		earliestStartTimeEpochMillis: {
			Type:     schema.TypeInt,
			Required: true,
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

	earliestStartTimeEpochMillis, ok1 := d.GetOk("earliest_start_time_epoch_millis")
	if !ok1 {
		return fmt.Errorf("required parameter earliest_start_time_epoch_millis not set")
	}

	latestStartTimeEpochMillis, ok2 := d.GetOk("latest_start_time_epoch_millis")
	if !ok2 {
		return fmt.Errorf("required parameter latest_start_time_epoch_millis not set")
	}

	var earliestStartTimeEpochMillisInt64 int64
	var latestStartTimeEpochMillisInt64 int64

	earliestStartTimeEpochMillisInt64 = int64(earliestStartTimeEpochMillis.(int))
	latestStartTimeEpochMillisInt64 = int64(latestStartTimeEpochMillis.(int))

	timeRange := wavefront.TimeRange{
		StartTime: earliestStartTimeEpochMillisInt64,
		EndTime:   latestStartTimeEpochMillisInt64,
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

	if err := d.Set("events", flattenEvents(allEvents)); err != nil {
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
