package wavefront

import (
	"fmt"
	"time"

	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
	ID                              string   `json:"id"`
	RunningState                    string   `json:"runningState"`
	SortAttr                        int      `json:"sortAttr"`
	Reason                          string   `json:"reason"`
	CustomerId                      string   `json:"customerId"`
	RelevantCustomerTags            []string `json:"relevantCustomerTags"`
	Title                           string   `json:"title"`
	StartTimeInSeconds              int64    `json:"startTimeInSeconds"`
	EndTimeInSeconds                int64    `json:"endTimeInSeconds"`
	RelevantHostTags                []string `json:"relevantHostTags"`
	RelevantHostNames               []string `json:"relevantHostNames"`
	CreatorId                       string   `json:"creatorId"`
	UpdaterId                       string   `json:"updaterId"`
	CreatedEpochMillis              int64    `json:"createdEpochMillis"`
	UpdatedEpochMillis              int64    `json:"updatedEpochMillis"`
	RelevantHostTagsAnded           bool     `json:"relevantHostTagsAnded"`
	HostTagGroupHostNamesGroupAnded bool     `json:"hostTagGroupHostNamesGroupAnded"`
	EventName                       string   `json:"eventName"`

*/

const (
	runningStateKey                    = "runningState"
	sortAttrKey                        = "sortAttr"
	reasonKey                          = "reason"
	customerIdKey                      = "customerId"
	relevantCustomerTagsKey            = "relevantCustomerTags"
	titleKey                           = "title"
	startTimeInSecondsKey              = "startTimeInSeconds"
	endTimeInSecondsKey                = "endTimeInSeconds"
	relevantHostTagsKey                = "relevantHostTags"
	relevantHostNamesKey               = "relevantHostNames"
	relevantHostTagsAndedKey           = "relevantHostTagsAnded"
	hostTagGroupHostNamesGroupAndedKey = "hostTagGroupHostNamesGroupAnded"
	eventNameKey                       = "eventName"
	maintenanceWindowsKey              = "maintenance_windows"
)

func dataSourceMaintenanceWindow() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceMaintenanceWindowRead,
		Schema: dataSourceMaintenanceWindowSchema(),
	}
}

func dataSourceMaintenanceWindowSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		idKey: {
			Type:     schema.TypeString,
			Required: true,
		},

		runningStateKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		sortAttrKey: {
			Type:     schema.TypeInt,
			Computed: true,
		},

		reasonKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		customerIdKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		relevantCustomerTagsKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},

		titleKey: {
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

		startTimeInSecondsKey: {
			Type:     schema.TypeInt,
			Computed: true,
		},

		endTimeInSecondsKey: {
			Type:     schema.TypeInt,
			Computed: true,
		},

		relevantHostTagsKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},

		relevantHostTagsAndedKey: {
			Type:     schema.TypeBool,
			Computed: true,
		},

		hostTagGroupHostNamesGroupAndedKey: {
			Type:     schema.TypeBool,
			Computed: true,
		},

		eventNameKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		relevantHostNamesKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
}

func dataSourceMaintenanceWindowRead(d *schema.ResourceData, m interface{}) error {
	maintenanceWindowClient := m.(*wavefrontClient).client.MaintenanceWindows()
	id, ok := d.GetOk("id")
	if !ok {
		return fmt.Errorf("required parameter '%s' not set", idKey)
	}

	idStr := fmt.Sprintf("%s", id)
	maintenanceWindow, err := maintenanceWindowClient.GetByID(idStr)

	if err != nil {
		return fmt.Errorf("error finding maintenance window with id %s", idStr)
	}

	// Data Source ID is set to current time to always refresh
	d.SetId(time.Now().UTC().String())
	return setMaintenanceWindowAttributes(d, maintenanceWindow)
}

func setMaintenanceWindowAttributes(d *schema.ResourceData, maintenanceWindow *wavefront.MaintenanceWindow) error {
	if err := d.Set(idKey, maintenanceWindow.ID); err != nil {
		return err
	}
	if err := d.Set(eventNameKey, maintenanceWindow.EventName); err != nil {
		return err
	}
	if err := d.Set(runningStateKey, maintenanceWindow.RunningState); err != nil {
		return err
	}
	if err := d.Set(reasonKey, maintenanceWindow.Reason); err != nil {
		return err
	}
	if err := d.Set(sortAttrKey, maintenanceWindow.SortAttr); err != nil {
		return err
	}
	if err := d.Set(createdEpochMillisKey, maintenanceWindow.CreatedEpochMillis); err != nil {
		return err
	}
	if err := d.Set(updatedEpochMillisKey, maintenanceWindow.UpdatedEpochMillis); err != nil {
		return err
	}
	if err := d.Set(customerIdKey, maintenanceWindow.CustomerId); err != nil {
		return err
	}
	if err := d.Set(titleKey, maintenanceWindow.Title); err != nil {
		return err
	}
	if err := d.Set(startTimeInSecondsKey, maintenanceWindow.StartTimeInSeconds); err != nil {
		return err
	}
	if err := d.Set(endTimeInSecondsKey, maintenanceWindow.EndTimeInSeconds); err != nil {
		return err
	}
	if err := d.Set(creatorIDKey, maintenanceWindow.EndTimeInSeconds); err != nil {
		return err
	}
	if err := d.Set(updaterIDKey, maintenanceWindow.EndTimeInSeconds); err != nil {
		return err
	}
	if err := d.Set(relevantHostTagsAndedKey, maintenanceWindow.EndTimeInSeconds); err != nil {
		return err
	}
	if err := d.Set(hostTagGroupHostNamesGroupAndedKey, maintenanceWindow.HostTagGroupHostNamesGroupAnded); err != nil {
		return err
	}
	if err := d.Set(relevantCustomerTagsKey, maintenanceWindow.RelevantCustomerTags); err != nil {
		return err
	}
	if err := d.Set(relevantHostNamesKey, maintenanceWindow.RelevantHostNames); err != nil {
		return err
	}
	if err := d.Set(relevantHostTagsKey, maintenanceWindow.RelevantHostTags); err != nil {
		return err
	}

	return nil
}
