package wavefront

import (
	"fmt"
	"time"

	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDerivedMetric() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceDerivedMetricRead,
		Schema: derivedMetricResponseSchema(),
	}
}

func derivedMetricResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		idKey: {
			Type:     schema.TypeString,
			Required: true,
		},
		nameKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		queryKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		minutesKey: {
			Type:     schema.TypeInt,
			Computed: true,
		},
		inTrashKey: {
			Type:     schema.TypeBool,
			Computed: true,
		},
		tagsKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		queryFailingKey: {
			Type:     schema.TypeBool,
			Computed: true,
		},

		lastErrorMessageKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		lastFailedTimeKey: {
			Type:     schema.TypeInt,
			Computed: true,
		},

		additionalInformationKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		createUserIDKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		updateUserIDKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		statusKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},

		hostsUsedKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},

		lastProcessedMillisKey: {
			Type:     schema.TypeInt,
			Computed: true,
		},

		processRateMinutesKey: {
			Type:     schema.TypeInt,
			Computed: true,
		},

		pointsScannedAtLastQueryKey: {
			Type:     schema.TypeInt,
			Computed: true,
		},

		includeObsoleteMetricsKey: {
			Type:     schema.TypeBool,
			Computed: true,
		},

		lastQueryTimeKey: {
			Type:     schema.TypeInt,
			Computed: true,
		},

		metricsUsedKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},

		queryQBEnabledKey: {
			Type:     schema.TypeBool,
			Computed: true,
		},

		deletedKey: {
			Type:     schema.TypeBool,
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
	}
}

func dataSourceDerivedMetricRead(d *schema.ResourceData, m interface{}) error {
	derivedMetricClient := m.(*wavefrontClient).client.DerivedMetrics()
	id, ok := d.GetOk("id")
	if !ok {
		return fmt.Errorf("required parameter '%s' not set", idKey)
	}

	idStr := fmt.Sprintf("%s", id)
	derivedMetric := wavefront.DerivedMetric{ID: &idStr}
	if err := derivedMetricClient.Get(&derivedMetric); err != nil {
		return err
	}

	// Data Source ID is set to current time to always refresh
	d.SetId(time.Now().UTC().String())
	return setDerivedMetricAttributes(d, derivedMetric)
}

func setDerivedMetricAttributes(d *schema.ResourceData, derivedMetric wavefront.DerivedMetric) error {
	if err := d.Set(idKey, *derivedMetric.ID); err != nil {
		return err
	}
	if err := d.Set(nameKey, derivedMetric.Name); err != nil {
		return err
	}
	if err := d.Set(queryKey, derivedMetric.Query); err != nil {
		return err
	}
	if err := d.Set(minutesKey, derivedMetric.Minutes); err != nil {
		return err
	}
	if err := d.Set(inTrashKey, derivedMetric.InTrash); err != nil {
		return err
	}
	if err := d.Set(tagsKey, derivedMetric.Tags.CustomerTags); err != nil {
		return err
	}
	if err := d.Set(queryFailingKey, derivedMetric.QueryFailing); err != nil {
		return err
	}
	if err := d.Set(lastErrorMessageKey, derivedMetric.LastErrorMessage); err != nil {
		return err
	}
	if err := d.Set(lastFailedTimeKey, derivedMetric.LastFailedTime); err != nil {
		return err
	}
	if err := d.Set(createUserIDKey, derivedMetric.CreateUserId); err != nil {
		return err
	}
	if err := d.Set(updateUserIDKey, derivedMetric.UpdateUserId); err != nil {
		return err
	}
	if err := d.Set(additionalInformationKey, derivedMetric.AdditionalInformation); err != nil {
		return err
	}
	if err := d.Set(statusKey, derivedMetric.Status); err != nil {
		return err
	}
	if err := d.Set(hostsUsedKey, derivedMetric.HostsUsed); err != nil {
		return err
	}
	if err := d.Set(lastProcessedMillisKey, derivedMetric.LastProcessedMillis); err != nil {
		return err
	}
	if err := d.Set(processRateMinutesKey, derivedMetric.ProcessRateMinutes); err != nil {
		return err
	}
	if err := d.Set(pointsScannedAtLastQueryKey, derivedMetric.PointsScannedAtLastQuery); err != nil {
		return err
	}
	if err := d.Set(includeObsoleteMetricsKey, derivedMetric.IncludeObsoleteMetrics); err != nil {
		return err
	}
	if err := d.Set(lastQueryTimeKey, derivedMetric.LastQueryTime); err != nil {
		return err
	}
	if err := d.Set(metricsUsedKey, derivedMetric.MetricsUsed); err != nil {
		return err
	}
	if err := d.Set(queryQBEnabledKey, derivedMetric.QueryQBEnabled); err != nil {
		return err
	}
	if err := d.Set(deletedKey, derivedMetric.Deleted); err != nil {
		return err
	}
	if err := d.Set(createdEpochMillisKey, derivedMetric.CreatedEpochMillis); err != nil {
		return err
	}
	return d.Set(updatedEpochMillisKey1, derivedMetric.UpdatedEpochMillis)
}
