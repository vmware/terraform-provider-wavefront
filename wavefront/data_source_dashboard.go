package wavefront

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/WavefrontHQ/go-wavefront-management-api/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	chartTitleBgColorKey                  = "chart_title_bg_color"
	chartTitleColorKey                    = "chart_title_color"
	chartTitleScalarKey                   = "chart_title_scalar"
	defaultEndTimeKey                     = "default_end_time"
	defaultStartTimeKey                   = "default_start_time"
	defaultTimeWindowKey                  = "default_time_window"
	displayDescriptionKey                 = "display_description"
	displayQueryParametersKey             = "display_query_parameters"
	displaySectionTableOfContentsKey      = "display_section_table_of_contents"
	eventFilterTypeKey                    = "event_filter_type"
	eventQueryKey                         = "event_query"
	favoriteKey                           = "favorite"
	hiddenKey                             = "hidden"
	numChartsKey                          = "num_charts"
	numFavoritesKey                       = "num_favorites"
	creatorIDKey                          = "creator_id"
	updaterIDKey                          = "updater_id"
	systemOwnedKey                        = "system_owned"
	viewsLastDayKey                       = "views_last_day"
	viewsLastMonthKey                     = "views_last_month"
	viewsLastWeekKey                      = "views_last_week"
	autoColumnTagsKey                     = "auto_column_tags"
	columnTagsKey                         = "column_tags"
	customTagsKey                         = "custom_tags"
	expectedDataSpacingKey                = "expected_data_spacing"
	fixedLegendDisplayStatsKey            = "fixed_legend_display_stats"
	fixedLegendEnabledKey                 = "fixed_legend_enabled"
	fixedLegendFilterFieldKey             = "fixed_legend_filter_field"
	fixedLegendFilterLimitKey             = "fixed_legend_filter_limit"
	fixedLegendHideLabelKey               = "fixed_legend_hide_label"
	fixedLegendPositionKey                = "fixed_legend_position"
	fixedLegendUseRawStatsKey             = "fixed_legend_use_raw_stats"
	groupBySourceKey                      = "group_by_source"
	invertDynamicLegendHoverControlKey    = "invert_dynamic_legend_hover_control"
	lineTypeKey                           = "line_type"
	maxKey                                = "max"
	minKey                                = "min"
	numTagsKey                            = "num_tags"
	plainMarkdownContentKey               = "plain_markdown_content"
	showHostsKey                          = "show_hosts"
	showLabelsKey                         = "show_labels"
	showRawValuesKey                      = "show_raw_values"
	sortValuesDescendingKey               = "sort_values_descending"
	sparklineDecimalPrecisionKey          = "sparkline_decimal_precision"
	sparklineDisplayColorKey              = "sparkline_display_color"
	sparklineDisplayFontSizeKey           = "sparkline_display_font_size"
	sparklineDisplayHorizontalPositionKey = "sparkline_display_horizontal_position"
	sparklineDisplayPostfixKey            = "sparkline_display_postfix"
	sparklineDisplayPrefixKey             = "sparkline_display_prefix"
	sparklineDisplayValueTypeKey          = "sparkline_display_value_type"
	sparklineDisplayVerticalPositionKey   = "sparkline_display_vertical_position"
	sparklineFillColorKey                 = "sparkline_fill_color"
	sparklineLineColorKey                 = "sparkline_line_color"
	sparklineSizeKey                      = "sparkline_size"
	sparklineValueColorMapApplyToKey      = "sparkline_value_color_map_apply_to"
	sparklineValueColorMapColorsKey       = "sparkline_value_color_map_colors"
	sparklineValueColorMapValuesKey       = "sparkline_value_color_map_values"
	sparklineValueColorMapValuesV2Key     = "sparkline_value_color_map_values_v2"
	sparklineValueTextMapTextKey          = "sparkline_value_text_map_text"
	sparklineValueTextMapThresholdsKey    = "sparkline_value_text_map_thresholds"
	stackTypeKey                          = "stack_type"
	tagModeKey                            = "tag_mode"
	timeBasedColoringKey                  = "time_based_coloring"
	typeKey                               = "type"
	windowingKey                          = "windowing"
	windowSizeKey                         = "window_size"
	xMaxKey                               = "xmax"
	xMinKey                               = "xmin"
	y0ScaleSIBy1024Key                    = "y0_scale_si_by1024"
	y0UnitAutoscalingKey                  = "y0_unit_autoscaling"
	y1MaxKey                              = "y1_max"
	y1MinKey                              = "y1_min"
	y1ScaleSIBy1024Key                    = "y1_scale_si_by1024"
	y1UnitAutoscalingKey                  = "y1_unit_autoscaling"
	y1UnitsKey                            = "y1_units"
	yMaxKey                               = "ymax"
	yMinKey                               = "ymin"
	fixedLegendFilterSortKey              = "fixed_legend_filter_sort"

	disabledKey            = "disabled"
	scatterPlotSourceKey   = "scatter_plot_source"
	queryBuilderEnabledKey = "querybuilder_enabled"
	sourceDescriptionKey   = "source_description"
	sourceColorKey         = "source_color"
	secondaryAxisKey       = "secondary_axis"

	baseKey              = "base"
	interpolatePointsKey = "interpolate_points_key"
	noDefaultEventsKey   = "no_default_events"
	summarizationKey     = "summarization"
	sourcesKey           = "sources"
	unitsKey             = "units"
	chartSettingsKey     = "chart_settings"
	chartAttributesKey   = "chart_attributes"

	heightFactorKey            = "height_factor"
	chartsKey                  = "charts"
	rowsKey                    = "rows"
	labelKey                   = "label"
	defaultValueKey            = "default_value"
	hideFromViewKey            = "hide_from_view"
	parameterTypeKey           = "parameter_type"
	valuesToReadableStringsKey = "values_to_readable_strings"
	queryValueKey              = "query_value"
	tagKey                     = "tag_key"
	dynamicFieldTypeKey        = "dynamic_field_type"
	sectionsKey                = "sections"
	parameterDetailsKey        = "parameter_details"
	canViewKey                 = "can_view"
	canModifyKey               = "can_modify"
)

func dataSourceDashboard() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceDashboardRead,
		Schema: dataSourceDashboardSchema(),
	}
}

func dataSourceDashboardSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		nameKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		idKey: {
			Type:     schema.TypeString,
			Required: true,
		},
		descriptionKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		canViewKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},

		canModifyKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},

		chartTitleBgColorKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		chartTitleColorKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		chartTitleScalarKey: {
			Type:     schema.TypeInt,
			Computed: true,
		},

		defaultStartTimeKey: {
			Type:     schema.TypeInt,
			Computed: true,
		},

		defaultEndTimeKey: {
			Type:     schema.TypeInt,
			Computed: true,
		},

		defaultTimeWindowKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		displayDescriptionKey: {
			Type:     schema.TypeBool,
			Computed: true,
		},

		displayQueryParametersKey: {
			Type:     schema.TypeBool,
			Computed: true,
		},

		displaySectionTableOfContentsKey: {
			Type:     schema.TypeBool,
			Computed: true,
		},

		eventFilterTypeKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		eventQueryKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		favoriteKey: {
			Type:     schema.TypeBool,
			Computed: true,
		},

		customerKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		deletedKey: {
			Type:     schema.TypeBool,
			Computed: true,
		},

		hiddenKey: {
			Type:     schema.TypeBool,
			Computed: true,
		},

		numChartsKey: {
			Type:     schema.TypeInt,
			Computed: true,
		},

		numFavoritesKey: {
			Type:     schema.TypeInt,
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

		systemOwnedKey: {
			Type:     schema.TypeBool,
			Computed: true,
		},

		viewsLastDayKey: {
			Type:     schema.TypeInt,
			Computed: true,
		},

		viewsLastMonthKey: {
			Type:     schema.TypeInt,
			Computed: true,
		},
		viewsLastWeekKey: {
			Type:     schema.TypeInt,
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

		urlKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		sectionsKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: sectionSchema(),
			},
		},

		tagsKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},

		parameterDetailsKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: parameterDetailSchema(),
			},
		},

		parametersKey: {
			Type:     schema.TypeMap,
			Computed: true,
		},
	}
}

func chartSettingSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		autoColumnTagsKey: {
			Type:     schema.TypeBool,
			Computed: true,
		},

		columnTagsKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		customTagsKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},

		expectedDataSpacingKey: {
			Type:     schema.TypeInt,
			Computed: true,
		},

		fixedLegendDisplayStatsKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},

		fixedLegendEnabledKey: {
			Type:     schema.TypeBool,
			Computed: true,
		},

		fixedLegendFilterFieldKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		fixedLegendFilterLimitKey: {
			Type:     schema.TypeInt,
			Computed: true,
		},

		fixedLegendFilterSortKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		fixedLegendHideLabelKey: {
			Type:     schema.TypeBool,
			Computed: true,
		},

		fixedLegendPositionKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		fixedLegendUseRawStatsKey: {
			Type:     schema.TypeBool,
			Computed: true,
		},
		groupBySourceKey: {
			Type:     schema.TypeBool,
			Computed: true,
		},
		invertDynamicLegendHoverControlKey: {
			Type:     schema.TypeBool,
			Computed: true,
		},
		lineTypeKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		maxKey: {
			Type:     schema.TypeFloat,
			Computed: true,
		},

		minKey: {
			Type:     schema.TypeFloat,
			Computed: true,
		},

		numTagsKey: {
			Type:     schema.TypeInt,
			Computed: true,
		},

		plainMarkdownContentKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		showHostsKey: {
			Type:     schema.TypeBool,
			Computed: true,
		},
		showLabelsKey: {
			Type:     schema.TypeBool,
			Computed: true,
		},
		showRawValuesKey: {
			Type:     schema.TypeBool,
			Computed: true,
		},
		sortValuesDescendingKey: {
			Type:     schema.TypeBool,
			Computed: true,
		},
		sparklineDecimalPrecisionKey: {
			Type:     schema.TypeInt,
			Computed: true,
		},

		sparklineDisplayColorKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		sparklineDisplayFontSizeKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		sparklineDisplayHorizontalPositionKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		sparklineDisplayPostfixKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		sparklineDisplayPrefixKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		sparklineDisplayValueTypeKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		sparklineDisplayVerticalPositionKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		sparklineFillColorKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		sparklineLineColorKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		sparklineSizeKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		sparklineValueColorMapApplyToKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		sparklineValueColorMapColorsKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},

		sparklineValueColorMapValuesKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeInt},
		},
		sparklineValueColorMapValuesV2Key: {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeFloat},
		},
		sparklineValueTextMapTextKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		sparklineValueTextMapThresholdsKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeFloat},
		},

		stackTypeKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		tagModeKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		timeBasedColoringKey: {
			Type:     schema.TypeBool,
			Computed: true,
		},

		typeKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		windowingKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		windowSizeKey: {
			Type:     schema.TypeInt,
			Computed: true,
		},

		xMinKey: {
			Type:     schema.TypeFloat,
			Computed: true,
		},

		xMaxKey: {
			Type:     schema.TypeFloat,
			Computed: true,
		},

		y0ScaleSIBy1024Key: {
			Type:     schema.TypeBool,
			Computed: true,
		},

		y0UnitAutoscalingKey: {
			Type:     schema.TypeBool,
			Computed: true,
		},

		y1MinKey: {
			Type:     schema.TypeFloat,
			Computed: true,
		},

		y1MaxKey: {
			Type:     schema.TypeFloat,
			Computed: true,
		},

		y1ScaleSIBy1024Key: {
			Type:     schema.TypeBool,
			Computed: true,
		},

		y1UnitAutoscalingKey: {
			Type:     schema.TypeBool,
			Computed: true,
		},

		y1UnitsKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		yMinKey: {
			Type:     schema.TypeFloat,
			Computed: true,
		},

		yMaxKey: {
			Type:     schema.TypeFloat,
			Computed: true,
		},
	}
}

func sourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		nameKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		queryKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		disabledKey: {
			Type:     schema.TypeBool,
			Computed: true,
		},
		scatterPlotSourceKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		queryBuilderEnabledKey: {
			Type:     schema.TypeBool,
			Computed: true,
		},
		sourceDescriptionKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		sourceColorKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		secondaryAxisKey: {
			Type:     schema.TypeBool,
			Computed: true,
		},
	}
}

func chartSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{

		nameKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		descriptionKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		baseKey: {
			Type:     schema.TypeInt,
			Computed: true,
		},

		includeObsoleteMetricsKey: {
			Type:     schema.TypeBool,
			Computed: true,
		},

		interpolatePointsKey: {
			Type:     schema.TypeBool,
			Computed: true,
		},

		noDefaultEventsKey: {
			Type:     schema.TypeBool,
			Computed: true,
		},

		summarizationKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		unitsKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		chartAttributesKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		chartSettingsKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: chartSettingSchema(),
			},
		},

		sourcesKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: sourceSchema(),
			},
		},
	}
}

func rowSchema() map[string]*schema.Schema {

	return map[string]*schema.Schema{
		nameKey: {
			Type:     schema.TypeString,
			Computed: true,
		},

		heightFactorKey: {
			Type:     schema.TypeInt,
			Computed: true,
		},

		chartsKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: chartSchema(),
			},
		},
	}
}

func sectionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		nameKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		rowsKey: {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: rowSchema(),
			},
		},
	}
}

func parameterDetailSchema() map[string]*schema.Schema {

	//		ValuesToReadableStrings map[string]string `json:"valuesToReadableStrings"`
	return map[string]*schema.Schema{
		labelKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		defaultValueKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		hideFromViewKey: {
			Type:     schema.TypeBool,
			Computed: true,
		},
		parameterTypeKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		queryValueKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		tagKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		dynamicFieldTypeKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		valuesToReadableStringsKey: {
			Type:     schema.TypeMap,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
	}
}

func dataSourceDashboardRead(d *schema.ResourceData, m interface{}) error {

	dashboardClient := m.(*wavefrontClient).client.Dashboards()
	id, ok := d.GetOk("id")
	if !ok {
		return fmt.Errorf("required parameter '%s' not set", idKey)
	}

	idStr := fmt.Sprintf("%s", id)
	dashboard := wavefront.Dashboard{ID: idStr}
	if err := dashboardClient.Get(&dashboard); err != nil {
		return err
	}

	// Data Source ID is set to current time to always refresh
	d.SetId(time.Now().UTC().String())
	return setDashboardAttributes(d, dashboard)

}

func setDashboardAttributes(d *schema.ResourceData, dashboard wavefront.Dashboard) error {
	if err := d.Set(idKey, dashboard.ID); err != nil {
		return err
	}
	if err := d.Set(nameKey, dashboard.Name); err != nil {
		return err
	}
	if err := d.Set(tagsKey, dashboard.Tags); err != nil {
		return err
	}
	if err := d.Set(descriptionKey, dashboard.Description); err != nil {
		return err
	}
	if err := d.Set(urlKey, dashboard.Url); err != nil {
		return err
	}
	if err := d.Set(chartTitleBgColorKey, dashboard.ChartTitleBgColor); err != nil {
		return err
	}
	if err := d.Set(chartTitleColorKey, dashboard.ChartTitleColor); err != nil {
		return err
	}
	if err := d.Set(chartTitleScalarKey, dashboard.ChartTitleScalar); err != nil {
		return err
	}
	if err := d.Set(defaultStartTimeKey, dashboard.DefaultStartTime); err != nil {
		return err
	}
	if err := d.Set(defaultEndTimeKey, dashboard.DefaultEndTime); err != nil {
		return err
	}
	if err := d.Set(defaultTimeWindowKey, dashboard.DefaultTimeWindow); err != nil {
		return err
	}
	if err := d.Set(displayDescriptionKey, dashboard.DisplayDescription); err != nil {
		return err
	}
	if err := d.Set(displayQueryParametersKey, dashboard.DisplayQueryParameters); err != nil {
		return err
	}
	if err := d.Set(displaySectionTableOfContentsKey, dashboard.DisplaySectionTableOfContents); err != nil {
		return err
	}
	if err := d.Set(eventFilterTypeKey, dashboard.EventFilterType); err != nil {
		return err
	}
	if err := d.Set(eventQueryKey, dashboard.EventQuery); err != nil {
		return err
	}
	if err := d.Set(favoriteKey, dashboard.Favorite); err != nil {
		return err
	}
	if err := d.Set(customerKey, dashboard.Customer); err != nil {
		return err
	}
	if err := d.Set(deletedKey, dashboard.Deleted); err != nil {
		return err
	}
	if err := d.Set(hiddenKey, dashboard.Hidden); err != nil {
		return err
	}
	if err := d.Set(numChartsKey, dashboard.NumCharts); err != nil {
		return err
	}
	if err := d.Set(numFavoritesKey, dashboard.NumFavorites); err != nil {
		return err
	}
	if err := d.Set(creatorIDKey, dashboard.CreatorId); err != nil {
		return err
	}
	if err := d.Set(updaterIDKey, dashboard.UpdaterId); err != nil {
		return err
	}
	if err := d.Set(systemOwnedKey, dashboard.SystemOwned); err != nil {
		return err
	}
	if err := d.Set(viewsLastMonthKey, dashboard.ViewsLastMonth); err != nil {
		return err
	}
	if err := d.Set(viewsLastWeekKey, dashboard.ViewsLastWeek); err != nil {
		return err
	}
	if err := d.Set(viewsLastDayKey, dashboard.ViewsLastDay); err != nil {
		return err
	}
	if err := d.Set(createdEpochMillisKey, dashboard.CreatedEpochMillis); err != nil {
		return err
	}
	if err := d.Set(updatedEpochMillisKey, dashboard.UpdatedEpochMillis); err != nil {
		return err
	}
	if err := d.Set(canModifyKey, dashboard.ACL.CanModify); err != nil {
		return err
	}
	if err := d.Set(canViewKey, dashboard.ACL.CanView); err != nil {
		return err
	}
	if err := d.Set(sectionsKey, flattenSections(dashboard.Sections)); err != nil {
		return err
	}
	if err := d.Set(parametersKey, convertStructToMap(dashboard.Parameters)); err != nil {
		return err
	}
	return d.Set(parameterDetailsKey, flattenParameterDetails(dashboard.ParameterDetails))
}

func convertStructToMap(parameters struct{}) map[string]interface{} {
	var mapResult map[string]interface{}

	paramsStruct := parameters
	data, _ := json.Marshal(paramsStruct)
	err := json.Unmarshal(data, &mapResult)
	if err != nil {
		panic(err)
	}
	return mapResult
}

func flattenParameterDetails(details map[string]wavefront.ParameterDetail) []map[string]interface{} {
	paramsDetails := make([]map[string]interface{}, 0, len(details))

	for _, v := range details {
		paramsDetails = append(paramsDetails, flattenParameterDetail(v))
	}
	return paramsDetails
}

func flattenParameterDetail(detail wavefront.ParameterDetail) map[string]interface{} {
	tfMap := make(map[string]interface{})
	tfMap[labelKey] = detail.Label
	tfMap[defaultValueKey] = detail.DefaultValue
	tfMap[hideFromViewKey] = detail.HideFromView
	tfMap[parameterTypeKey] = detail.ParameterType
	tfMap[queryValueKey] = detail.QueryValue
	tfMap[tagKey] = detail.TagKey
	tfMap[dynamicFieldTypeKey] = detail.DynamicFieldType
	tfMap[valuesToReadableStringsKey] = flattenValuesToReadableString(detail.ValuesToReadableStrings)
	return tfMap
}

func flattenValuesToReadableString(valueToReadableStrings map[string]string) map[string]interface{} {
	tfMap := make(map[string]interface{})

	for k, v := range valueToReadableStrings {
		tfMap[k] = v
	}

	return tfMap
}

func flattenSections(sections []wavefront.Section) []map[string]interface{} {

	tfMaps := make([]map[string]interface{}, len(sections))
	for i, v := range sections {
		tfMaps[i] = flattenSection(v)
	}
	return tfMaps

}

func flattenSection(v wavefront.Section) map[string]interface{} {
	tfMap := make(map[string]interface{})
	tfMap[nameKey] = v.Name
	tfMap[rowsKey] = flattenRows(v.Rows)
	return tfMap
}

func flattenRows(rows []wavefront.Row) []map[string]interface{} {
	rwMaps := make([]map[string]interface{}, len(rows))
	for i, v := range rows {
		rwMaps[i] = flattenRow(v)
	}
	return rwMaps
}

func flattenRow(row wavefront.Row) map[string]interface{} {
	tfMap := make(map[string]interface{})
	tfMap[nameKey] = row.Name
	tfMap[heightFactorKey] = row.HeightFactor
	tfMap[chartsKey] = flattenCharts(row.Charts)
	return tfMap
}

func flattenCharts(charts []wavefront.Chart) []map[string]interface{} {
	chMaps := make([]map[string]interface{}, len(charts))
	for i, v := range charts {
		chMaps[i] = flattenChart(v)
	}
	return chMaps
}

func flattenChart(chart wavefront.Chart) map[string]interface{} {
	tfMap := make(map[string]interface{})
	tfMap[nameKey] = chart.Name
	tfMap[descriptionKey] = chart.Description
	tfMap[baseKey] = chart.Base
	tfMap[includeObsoleteMetricsKey] = chart.IncludeObsoleteMetrics
	tfMap[interpolatePointsKey] = chart.InterpolatePoints
	tfMap[noDefaultEventsKey] = chart.NoDefaultEvents
	tfMap[summarizationKey] = chart.Summarization
	tfMap[unitsKey] = chart.Units
	tfMap[chartAttributesKey] = marshalJSON(chart.ChartAttributes)
	tfMap[chartSettingsKey] = flattenChartSettings(chart.ChartSettings)
	tfMap[sourcesKey] = flattenChartSources(chart.Sources)
	return tfMap
}

func marshalJSON(attributes json.RawMessage) string {
	value, err := json.Marshal(&attributes)
	if err != nil {
		panic(err)
	}
	return string(value)
}

func flattenChartSources(sources []wavefront.Source) []map[string]interface{} {
	srcMaps := make([]map[string]interface{}, len(sources))
	for i, v := range sources {
		srcMaps[i] = flattenChartSource(v)
	}
	return srcMaps
}

func flattenChartSource(source wavefront.Source) map[string]interface{} {
	srcMap := make(map[string]interface{})
	srcMap[nameKey] = source.Name
	srcMap[queryKey] = source.Query
	srcMap[disabledKey] = source.Disabled
	srcMap[scatterPlotSourceKey] = source.ScatterPlotSource
	srcMap[queryBuilderEnabledKey] = source.QuerybuilderEnabled
	srcMap[sourceDescriptionKey] = source.SourceDescription
	srcMap[sourceColorKey] = source.SourceColor
	srcMap[secondaryAxisKey] = source.SecondaryAxis
	return srcMap
}

func flattenChartSettings(settings wavefront.ChartSetting) []map[string]interface{} {
	chartSettingMap := make(map[string]interface{})
	chartSettingMap[autoColumnTagsKey] = settings.AutoColumnTags
	chartSettingMap[columnTagsKey] = settings.ColumnTags
	chartSettingMap[customTagsKey] = settings.CustomTags // check if getting flattend or not
	chartSettingMap[expectedDataSpacingKey] = settings.ExpectedDataSpacing
	chartSettingMap[fixedLegendDisplayStatsKey] = settings.FixedLegendDisplayStats
	chartSettingMap[fixedLegendEnabledKey] = settings.FixedLegendEnabled
	chartSettingMap[fixedLegendFilterFieldKey] = settings.FixedLegendFilterField
	chartSettingMap[fixedLegendFilterLimitKey] = settings.FixedLegendFilterLimit
	chartSettingMap[fixedLegendFilterSortKey] = settings.FixedLegendFilterSort
	chartSettingMap[fixedLegendHideLabelKey] = settings.FixedLegendHideLabel
	chartSettingMap[fixedLegendPositionKey] = settings.FixedLegendPosition
	chartSettingMap[fixedLegendUseRawStatsKey] = settings.FixedLegendUseRawStats
	chartSettingMap[groupBySourceKey] = settings.GroupBySource
	chartSettingMap[invertDynamicLegendHoverControlKey] = settings.InvertDynamicLegendHoverControl
	chartSettingMap[lineTypeKey] = settings.LineType
	chartSettingMap[maxKey] = settings.Max
	chartSettingMap[minKey] = settings.Min
	chartSettingMap[numTagsKey] = settings.NumTags
	chartSettingMap[plainMarkdownContentKey] = settings.PlainMarkdownContent
	chartSettingMap[showHostsKey] = settings.ShowHosts
	chartSettingMap[showLabelsKey] = settings.ShowLabels
	chartSettingMap[showRawValuesKey] = settings.ShowRawValues
	chartSettingMap[sortValuesDescendingKey] = settings.SortValuesDescending
	chartSettingMap[sparklineDecimalPrecisionKey] = settings.SparklineDecimalPrecision
	chartSettingMap[sparklineDisplayColorKey] = settings.SparklineDisplayColor
	chartSettingMap[sparklineDisplayFontSizeKey] = settings.SparklineDisplayFontSize
	chartSettingMap[sparklineDisplayHorizontalPositionKey] = settings.SparklineDisplayHorizontalPosition
	chartSettingMap[sparklineDisplayPostfixKey] = settings.SparklineDisplayPostfix
	chartSettingMap[sparklineDisplayPrefixKey] = settings.SparklineDisplayPrefix
	chartSettingMap[sparklineDisplayValueTypeKey] = settings.SparklineDisplayValueType
	chartSettingMap[sparklineDisplayVerticalPositionKey] = settings.SparklineDisplayVerticalPosition
	chartSettingMap[sparklineFillColorKey] = settings.SparklineFillColor
	chartSettingMap[sparklineLineColorKey] = settings.SparklineLineColor
	chartSettingMap[sparklineSizeKey] = settings.SparklineSize
	chartSettingMap[sparklineValueColorMapApplyToKey] = settings.SparklineValueColorMapApplyTo
	chartSettingMap[sparklineValueColorMapColorsKey] = settings.SparklineValueColorMapColors       //list
	chartSettingMap[sparklineValueColorMapValuesKey] = settings.SparklineValueColorMapValues       //liist
	chartSettingMap[sparklineValueColorMapValuesV2Key] = settings.SparklineValueColorMapValuesV2   //list
	chartSettingMap[sparklineValueTextMapTextKey] = settings.SparklineValueTextMapText             //list
	chartSettingMap[sparklineValueTextMapThresholdsKey] = settings.SparklineValueTextMapThresholds //list
	chartSettingMap[stackTypeKey] = settings.StackType
	chartSettingMap[tagModeKey] = settings.TagMode
	chartSettingMap[timeBasedColoringKey] = settings.TimeBasedColoring
	chartSettingMap[typeKey] = settings.Type
	chartSettingMap[windowingKey] = settings.Windowing
	chartSettingMap[windowSizeKey] = settings.WindowSize
	chartSettingMap[xMinKey] = settings.Xmin
	chartSettingMap[xMaxKey] = settings.Xmax
	chartSettingMap[y0ScaleSIBy1024Key] = settings.Y0ScaleSIBy1024
	chartSettingMap[y0UnitAutoscalingKey] = settings.Y0UnitAutoscaling
	chartSettingMap[y1MinKey] = settings.Y1Min
	chartSettingMap[y1MaxKey] = settings.Y1Max
	chartSettingMap[y1ScaleSIBy1024Key] = settings.Y1ScaleSIBy1024
	chartSettingMap[y1UnitAutoscalingKey] = settings.Y1UnitAutoscaling
	chartSettingMap[y1UnitsKey] = settings.Y1Units
	chartSettingMap[yMinKey] = settings.Y1Min
	chartSettingMap[yMaxKey] = settings.Y1Max

	chartSettingsMap := make([]map[string]interface{}, 1)
	chartSettingsMap[0] = chartSettingMap
	return chartSettingsMap
}
