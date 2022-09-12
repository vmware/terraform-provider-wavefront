package wavefront

import (
	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"time"
)

func dataSourceDashboards() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceDashboardsRead,
		Schema: dataSourceDashboardsSchema(),
	}
}

func dataSourceDashboardsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		// Computed Values
		"dashboards": {
			Type:     schema.TypeSet,
			Computed: true,
			Elem: &schema.Resource{
				Schema: dataSourceDashboardSchema(),
			},
		},
	}

}

func dataSourceDashboardsRead(d *schema.ResourceData, m interface{}) error {
	var allDashboards []*wavefront.Dashboard
	dashboardClient := m.(*wavefrontClient).client.Dashboards()

	// Data Source ID is set to current time to always refresh
	d.SetId(time.Now().UTC().String())

	cont := true
	offset := 0

	for cont {
		filter := []*wavefront.SearchCondition{
			{Key: "limit", Value: string(rune(pageSize)), MatchingMethod: exactMatching},
			{Key: "offset", Value: string(rune(offset)), MatchingMethod: exactMatching},
		}

		dashboards, err := dashboardClient.Find(filter)
		if err != nil {
			return err
		}

		allDashboards = append(allDashboards, dashboards...)

		if len(allDashboards) < pageSize {
			cont = false
		} else {
			offset += pageSize
		}
	}

	if err := d.Set("dashboards", flattenDashboards(allDashboards)); err != nil {
		return err
	}
	return nil
}

func flattenDashboards(dashboards []*wavefront.Dashboard) interface{} {
	tfMaps := make([]map[string]interface{}, len(dashboards))
	for i, v := range dashboards {
		tfMaps[i] = flattenDashboard(v)
	}
	return tfMaps
}

func flattenDashboard(dashboard *wavefront.Dashboard) map[string]interface{} {

	tfMap := make(map[string]interface{})
	tfMap[idKey] = dashboard.ID
	tfMap[nameKey] = dashboard.Name
	tfMap[tagsKey] = dashboard.Tags
	tfMap[descriptionKey] = dashboard.Description
	tfMap[urlKey] = dashboard.Url
	tfMap[chartTitleBgColorKey] = dashboard.ChartTitleBgColor
	tfMap[chartTitleColorKey] = dashboard.ChartTitleColor
	tfMap[chartTitleScalarKey] = dashboard.ChartTitleScalar
	tfMap[defaultStartTimeKey] = dashboard.DefaultStartTime
	tfMap[defaultEndTimeKey] = dashboard.DefaultEndTime
	tfMap[defaultTimeWindowKey] = dashboard.DefaultTimeWindow
	tfMap[displayDescriptionKey] = dashboard.DisplayDescription
	tfMap[displayQueryParametersKey] = dashboard.DisplayQueryParameters
	tfMap[displaySectionTableOfContentsKey] = dashboard.DisplaySectionTableOfContents
	tfMap[eventFilterTypeKey] = dashboard.EventFilterType
	tfMap[eventQueryKey] = dashboard.EventQuery
	tfMap[favoriteKey] = dashboard.Favorite
	tfMap[customerKey] = dashboard.Customer
	tfMap[deletedKey] = dashboard.Deleted
	tfMap[hiddenKey] = dashboard.Hidden
	tfMap[numChartsKey] = dashboard.NumCharts
	tfMap[numFavoritesKey] = dashboard.NumFavorites
	tfMap[creatorIdKey] = dashboard.CreatorId
	tfMap[updaterIdKey] = dashboard.UpdaterId
	tfMap[systemOwnedKey] = dashboard.SystemOwned
	tfMap[viewsLastMonthKey] = dashboard.ViewsLastMonth
	tfMap[viewsLastWeekKey] = dashboard.ViewsLastWeek
	tfMap[viewsLastDayKey] = dashboard.ViewsLastDay
	tfMap[createdEpochMillisKey] = dashboard.CreatedEpochMillis
	tfMap[updatedEpochMillisKey] = dashboard.UpdatedEpochMillis
	tfMap[canViewKey] = dashboard.ACL.CanView
	tfMap[canModifyKey] = dashboard.ACL.CanModify
	tfMap[parametersKey] = convertStructToMap(dashboard.Parameters)
	tfMap[parameterDetailsKey] = flattenParameterDetails(dashboard.ParameterDetails)
	tfMap[sectionsKey] = flattenSections(dashboard.Sections)
	return tfMap

}
