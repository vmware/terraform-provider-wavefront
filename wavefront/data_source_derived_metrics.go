package wavefront

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDerivedMetrics() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceDerivedMetricsRead,
		Schema: dataSourceDerivedMetricsSchema(),
	}
}

func dataSourceDerivedMetricsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		// Computed Values
		derivedMetricsKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: derivedMetricSchema(),
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

func derivedMetricSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		idKey: {
			Type:     schema.TypeString,
			Computed: true,
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

		updatedEpochMillisKey1: {
			Type:     schema.TypeInt,
			Computed: true,
		},
	}
}

func dataSourceDerivedMetricsRead(d *schema.ResourceData, m interface{}) error {
	var allDerivedMetrics []*wavefront.DerivedMetric

	limit := d.Get(limitKey).(int)
	offset := d.Get(offsetKey).(int)

	if err := json.Unmarshal(searchAll(limit, offset, "derivedmetric", nil, nil, m), &allDerivedMetrics); err != nil {
		return fmt.Errorf("Response is invalid JSON")
	}

	// Data Source ID is set to current time to always refresh
	d.SetId(time.Now().UTC().String())

	if err := d.Set(derivedMetricsKey, flattenDerivedMetrics(allDerivedMetrics)); err != nil {
		return err
	}
	return nil
}

func flattenDerivedMetrics(derivedMetrics []*wavefront.DerivedMetric) []map[string]interface{} {
	tfMaps := make([]map[string]interface{}, len(derivedMetrics))
	for i, v := range derivedMetrics {
		tfMaps[i] = flattenDerivedMetric(v)
	}
	return tfMaps
}

func flattenDerivedMetric(derivedMetric *wavefront.DerivedMetric) map[string]interface{} {
	tfMap := make(map[string]interface{})
	tfMap[idKey] = *derivedMetric.ID
	tfMap[nameKey] = derivedMetric.Name
	tfMap[queryKey] = derivedMetric.Query
	tfMap[minutesKey] = derivedMetric.Minutes
	tfMap[inTrashKey] = derivedMetric.InTrash
	tfMap[tagsKey] = derivedMetric.Tags.CustomerTags
	tfMap[queryFailingKey] = derivedMetric.QueryFailing
	tfMap[lastErrorMessageKey] = derivedMetric.LastErrorMessage
	tfMap[lastFailedTimeKey] = derivedMetric.LastFailedTime
	tfMap[updateUserIDKey] = derivedMetric.UpdateUserId
	tfMap[createUserIDKey] = derivedMetric.CreateUserId
	tfMap[additionalInformationKey] = derivedMetric.AdditionalInformation
	tfMap[statusKey] = derivedMetric.Status
	tfMap[hostsUsedKey] = derivedMetric.HostsUsed
	tfMap[lastProcessedMillisKey] = derivedMetric.LastProcessedMillis
	tfMap[processRateMinutesKey] = derivedMetric.ProcessRateMinutes
	tfMap[pointsScannedAtLastQueryKey] = derivedMetric.PointsScannedAtLastQuery
	tfMap[includeObsoleteMetricsKey] = derivedMetric.IncludeObsoleteMetrics
	tfMap[lastQueryTimeKey] = derivedMetric.LastQueryTime
	tfMap[metricsUsedKey] = derivedMetric.MetricsUsed
	tfMap[queryQBEnabledKey] = derivedMetric.QueryQBEnabled
	tfMap[deletedKey] = derivedMetric.Deleted
	tfMap[createdEpochMillisKey] = derivedMetric.CreatedEpochMillis
	tfMap[updatedEpochMillisKey1] = derivedMetric.UpdatedEpochMillis

	return tfMap
}
