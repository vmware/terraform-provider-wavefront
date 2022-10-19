package wavefront

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMaintenanceWindows() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceMaintenanceWindowsRead,
		Schema: dataSourceMaintenanceWindowsSchema(),
	}
}

func dataSourceMaintenanceWindowsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		// Computed Values
		maintenanceWindowsKey: {
			Type:     schema.TypeSet,
			Computed: true,
			Elem: &schema.Resource{
				Schema: dataSourceMaintenanceWindowSchema(),
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

func dataSourceMaintenanceWindowsRead(d *schema.ResourceData, m interface{}) error {
	var allMaintenanceWindows []*wavefront.MaintenanceWindow

	limit := d.Get(limitKey).(int)
	offset := d.Get(offsetKey).(int)

	if err := json.Unmarshal(searchAll(limit, offset, "maintenancewindow", nil, nil, m), &allMaintenanceWindows); err != nil {
		return fmt.Errorf("Response is invalid JSON")
	}

	// Data Source ID is set to current time to always refresh
	d.SetId(time.Now().UTC().String())

	if err := d.Set(maintenanceWindowsKey, flattenMaintenanceWindows(allMaintenanceWindows)); err != nil {
		return err
	}
	return nil
}

func flattenMaintenanceWindows(maintenanceWindows []*wavefront.MaintenanceWindow) []map[string]interface{} {
	tfMaps := make([]map[string]interface{}, len(maintenanceWindows))
	for i, v := range maintenanceWindows {
		tfMaps[i] = flattenMaintenanceWindow(v)
	}
	return tfMaps
}

func flattenMaintenanceWindow(maintenanceWindow *wavefront.MaintenanceWindow) map[string]interface{} {
	tfMap := make(map[string]interface{})
	tfMap[idKey] = maintenanceWindow.ID
	tfMap[eventNameKey] = maintenanceWindow.EventName
	tfMap[runningStateKey] = maintenanceWindow.RunningState
	tfMap[reasonKey] = maintenanceWindow.Reason
	tfMap[sortAttrKey] = maintenanceWindow.SortAttr
	tfMap[createdEpochMillisKey] = maintenanceWindow.CreatedEpochMillis
	tfMap[updatedEpochMillisKey] = maintenanceWindow.UpdatedEpochMillis
	tfMap[customerIdKey] = maintenanceWindow.CustomerId
	tfMap[titleKey] = maintenanceWindow.Title
	tfMap[startTimeInSecondsKey] = maintenanceWindow.StartTimeInSeconds
	tfMap[endTimeInSecondsKey] = maintenanceWindow.EndTimeInSeconds
	tfMap[creatorIDKey] = maintenanceWindow.CreatorId
	tfMap[updaterIDKey] = maintenanceWindow.UpdaterId
	tfMap[relevantHostTagsAndedKey] = maintenanceWindow.RelevantHostTagsAnded
	tfMap[hostTagGroupHostNamesGroupAndedKey] = maintenanceWindow.HostTagGroupHostNamesGroupAnded
	tfMap[relevantCustomerTagsKey] = maintenanceWindow.RelevantCustomerTags
	tfMap[relevantHostNamesKey] = maintenanceWindow.RelevantHostNames
	tfMap[relevantHostTagsKey] = maintenanceWindow.RelevantHostTags

	return tfMap
}
