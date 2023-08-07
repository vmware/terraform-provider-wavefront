---
layout: "wavefront"
page_title: "Wavefront: Dashboard"
description: |-
    Get the information about a specific Wavefront dashboard.
---

# Data Source: wavefront_dashboard

Use this data source to get information about a certain Wavefront dashboard by its ID.

## Argument Reference

* `id` - (Required) The ID associated with the dashboard data to be fetched.

## Example Usage

```hcl
# Get the information about a dashboard.
data "wavefront_dashboard" "example" {
  id = "dashboard-id"
}
```

## Attribute Reference

The following attributes are supported:

* `name` - Name of the dashboard.
* `tags` - A set of tags to assign to this resource.
* `description` - Human-readable description of the dashboard.
* `url` - Unique identifier, also a URL slug of the dashboard.
* `section` - Dashboard chart sections. See [dashboard sections](#dashboard-sections).
* `display_query_parameters` - Whether the dashboard parameters section is opened by default when the dashboard
  is shown.
* `parameter_details` - The current JSON representation of dashboard parameters. See [parameter details](#parameter-details).
* `display_section_table_of_contents` - Whether the "pills" quick-linked the sections of the dashboard are
  displayed by default when the dashboard is shown.
* `can_modify` - A list of users that have modify ACL access to the dashboard.
* `can_view` - A list of users that have view ACL access to the dashboard.
* `event_filter_type` - How charts belonging to this dashboard should display events. `BYCHART` is default if
  unspecified. Valid options are: `BYCHART`, `AUTOMATIC`, `ALL`, `NONE`, `BYDASHBOARD`, and `BYCHARTANDDASHBOARD`.

### Dashboard Sections

The `section` mapping supports the following:

* `name` - Name of this section.
* `row` - See [dashboard section rows](#dashboard-section-rows).

### Dashboard Section Rows

The `row` mapping supports the following:

* `chart` - Charts in this section. See [dashboard chart](#dashboard-chart).

### Dashboard Chart

The `chart` mapping supports the following:

* `source` - Query expression to plot on the chart. See [chart source queries](#chart-source-queries).
* `chart_setting` - Chart settings. See [chart settings](#chart-settings).
* `name` - Name of the source.
* `units` - String to label the units of the chart on the Y-Axis.
* `summarization` - Summarization strategy for the chart. MEAN is default.
* `description` - Description of the chart.
* `base` - The base of logarithmic scale charts. Omit or set to 0 for the default linear scale. Usually set to 10 for the traditional logarithmic scale.

### Chart Source Queries

The `source` mapping supports the following:

* `name` - Name of the source.
* `query` - Query expression to plot on the chart.
* `source_description` - A description for the purpose of this source.
* `disabled` - Whether the source is disabled.
* `scatter_plot_source` - For scatter plots, does this query source the X-axis or the Y-axis, `X`, or `Y`.
* `query_builder_enabled` - Whether or not this source line should have the query builder enabled.

### Chart Settings

The `chart_setting` mapping supports the following:

* `type` - Chart Type. `line` refers to the Line Plot, `scatter` to the Point Plot, `stacked-area` to the Stacked Area plot, `table` to the Tabular View, `scatterplot-xy` to Scatter Plot, `markdown-widget` to the Markdown display, and `sparkline` to the Single Stat view. Valid options are`line`, `scatterplot`,
  `stacked-area`, `stacked-column`, `table`, `scatterplot-xy`, `markdown-widget`, `sparkline`, `globe`, `nodemap`, `top-k`, `status-list`, and `histogram`.
* `max` - Max value of the Y-axis. Set to null or leave blank for auto.
* `line_type` - Plot interpolation type.  `linear` is default. Valid options are `linear`, `step-before`, `step-after`, `basis`, `cardinal`, and `monotone`.
* `stack_type` - Type of stacked chart (applicable only if chart type is `stacked`). `zero` (default) means stacked from y=0. `expand` means normalized from 0 to 1.  `wiggle` means minimize weighted changes. `silhouette` means to center the stream. Valid options are `zero`, `expand`, `wiggle`, `silhouette`, and `bars`.
* `windowing` - For the tabular view, whether to use the full time window for the query or the last X minutes. Valid options are `full` or `last`.
* `window_size` - Width, in minutes, of the time window to use for `last` windowing.
* `show_hosts` - For the tabular view, whether to display sources. Default is `true`.
* `show_labels` - For the tabular view, whether to display labels. Default is `true`.
* `show_raw_values` - For the tabular view, whether to display raw values. Default is `false`.
* `auto_column_tags` - This setting is deprecated.
* `column_tags` - This setting is deprecated.
* `tag_mode` - For the tabular view, which mode to use to determine which point tags to display. Valid options are `all`, `top`, or `custom`.
* `num_tags` - For the tabular view defines how many point tags to display.
* `custom_tags` - For the tabular view, a list of point tags to display when using the `custom` tag display mode.
* `group_by_source` - For the tabular view, whether to group multi metrics into a single row by a common source. If set to `false`, each source is displayed in its own row. If set to `true`, multiple metrics for the same host are displayed as different columns in the same row.
* `sort_values_descending` - For the tabular view, whether to display values in descending order. Default is `false`.
* `y1max` - For plots with multiple Y-axes, max value for the right side Y-axis. Set null for auto.
* `y1min` - For plots with multiple Y-axes, min value for the right side Y-axis. Set null for auto.
* `y1_units` - For plots with multiple Y-axes, units for right side Y-axis.
* `y0_scale_si_by_1024` - (Optional) Whether to scale numerical magnitude labels for left Y-axis by 1024 in the IEC/Binary manner (instead of by 1000 like SI).
* `y1_scale_si_by_1024` - (Optional) Whether to scale numerical magnitude labels for right Y-axis by 1024 in the IEC/Binary manner (instead of by 1000 like SI).
* `y0_unit_autoscaling` - (Optional) Whether to automatically adjust magnitude labels and units for the left Y-axis to favor smaller magnitudes and larger units.
* `y1_unit_autoscaling` - (Optional) Whether to automatically adjust magnitude labels and units for the right Y-axis to favor smaller magnitudes and larger units.
* `invert_dynamic_legend_hover_control` - (Optional) Whether to disable the display of the floating legend (but reenable it when the ctrl-key is pressed).
* `fixed_legend_enabled` - (Optional) Whether to enable a fixed tabular legend adjacent to the chart.
* `fixed_legend_use_raw_stats` - (Optional) If `true`, the legend uses non-summarized stats instead of summarized.
* `fixed_legend_position` - (Optional)  Where the fixed legend should be displayed with respect to the chart. Valid options are `RIGHT`, `TOP`, `LEFT`, `BOTTOM`.
* `fixed_legend_display_stats` - (Optional) For a chart with a fixed legend, a list of statistics to display in the legend.
* `fixed_legend_filter_sort` - (Optional) Whether to display `TOP` or `BOTTOM` ranked series in a fixed legend. Valid options are `TOP`, and `BOTTOM`.
* `fixed_legend_filter_limit` - (Optional) Number of series to include in the fixed legend.
* `fixed_legend_filter_field` - (Optional) Statistic to use for determining whether a series is displayed on the fixed legend. Valid options are `CURRENT`, `MEAN`, `MEDIAN`, `SUM`, `MIN`, `MAX`, and `COUNT`.
* `fixed_legend_hide_label` - (Optional) This setting is deprecated.
* `xmax` - For x-y scatterplots, max value for the X-axis. Set to null for auto.
* `ymax` - For x-y scatterplots, max value for the Y-axis. Set to null for auto.
* `xmin` - For x-y scatterplots, min value for the X-axis. Set to null for auto.
* `ymin` - For x-y scatterplots, min value for the Y-axis. Set to null for auto.
* `time_based_coloring` - For x-y scatterplots, whether to color more recent points as darker than older points.
* `sparkline_display_value_type` - For the single stat view, where to display the name of the query or the value of the query. Valid options are `VALUE` or `LABEL`.
* `sparkline_display_color` - For the single stat view, the color of the displayed text (when not dynamically determined). Values should be in RGBA format.
* `sparkline_display_vertical_position` - This setting is deprecated.
* `sparkline_display_horizontal_position` - For the single stat view, the horizontal position of the displayed text. Valid options are `MIDDLE`, `LEFT`, `RIGHT`.
* `sparkline_display_font_size` - For the single stat view, the font size of the displayed text, in percent.
* `sparkline_display_prefix` - For the single stat view, a string to add before the displayed text.
* `sparkline_display_postfix` - For the single stat view, a string to append to the displayed text.
* `sparkline_size` - For the single stat view, this determines whether the sparkline of the statistic is displayed in the chart. Valid options are `BACKGROUND`, `BOTTOM`, `NONE`.
* `sparkline_line_color` - For the single stat view, the color of the line. Values should be in RGBA format.
* `min` - Min value of the Y-axis. Set to null or leave blank for auto.
* `plain_markdown_content` - The markdown content for a Markdown display, in plain text.
* `sparkline_fill_color` - For the single stat view, the color of the background fill. Values should be in RGBA format.
* `sparkline_value_color_map_colors` - For the single stat view, a list of colors that differing query values map to. Must contain one more element than `sparkline_value_color_map_values_v2`. Values should be in RGBA format.
* `sparkline_value_color_map_values_v2` - For the single stat view, a list of boundaries for mapping different query values to colors. Must contain one element less than `sparkline_value_color_map_colors`.
* `sparkline_value_color_map_values` - This setting is deprecated.
* `sparkline_value_color_map_apply_to` - For the single stat view, whether to apply dynamic color settings to the displayed `TEXT` or `BACKGROUND`. Valid options are `TEXT` or `BACKGROUND`.
* `sparkline_decimal_precision` - For the single stat view, the decimal precision of the displayed number.
* `sparkline_value_text_map_text` - For the single stat view, a list of display text values that different query values map to. Must contain one more element than `sparkline_value_text_map_thresholds`.
* `sparkline_value_text_map_thresholds` - For the single stat view, a list of threshold boundaries for mapping different query values to display text. Must contain one element less than `sparkline_value_text_map_text`.
* `expected_data_spacing` - Threshold (in seconds) for time delta between consecutive points in a series above which a dotted line will replace a solid in line plots. Default is 60.

### Parameter Details

The `parameter_details` mapping supports the following:

* `label` - The label for the parameter.
* `values_to_readable_strings` - A string to string map. At least one of the keys must match the value of
  `default_value`.
* `name` - The name of the parameters.
* `default_value` - The default value of the parameter.
* `hide_from_view` - If `true` the parameter will only be shown on the edit view of the dashboard.
* `parameter_type` - The type of the parameter. `SIMPLE`, `LIST`, or `DYNAMIC`.
* `tag_key` - For `TAG_KEY` dynamic field types, the tag key to return.
* `query_value` - For `DYNAMIC` parameter types, the query to execute to return values.
* `dynamic_field_type` - For `DYNAMIC` parameter types, the type of the field. Valid options are `SOURCE`,
  `SOURCE_TAG`, `METRIC_NAME`, `TAG_KEY`, and `MATCHING_SOURCE_TAG`.
