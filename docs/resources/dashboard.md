---
layout: "wavefront"
page_title: "Wavefront: Dashboard"
description: |-
  Provides a Wavefront Dashboard resource.  This allows dashboards to be created, updated, and deleted.
---

# Resource: wavefront_dashboard

Provides a Wavefront Dashboard resource. This allows dashboards to be created, updated, and deleted.

## Example usage

```hcl

resource "wavefront_user" "basic" {
  email  = "test+tftesting@example.com"
  groups = [
    "agent_management",
    "alerts_management",
  ]
}

resource "wavefront_dashboard" "test_dashboard" {
  name                              = "Terraform Test Dashboard"
  description                       = "testing, testing"
  url                               = "tftestcreate"
  display_section_table_of_contents = true
  display_query_parameters          = true
  can_view                          = [
    wavefront_user.basic.id
  ]

  section {
    name = "section 1"
    row {
      chart {
        name        = "chart 1"
        description = "chart number 1"
        units       = "something per unit"
        source {
          name  = "source name"
          query = "ts()"
        }
        chart_setting {
          type = "linear"
        }
        summarization = "MEAN"
      }
    }
  }
  parameter_details {
    name                       = "param1"
    label                      = "param1"
    default_value              = "Label"
    hide_from_view             = false
    parameter_type             = "SIMPLE"
    values_to_readable_strings = {
      Label = "test"
    }
  }
  tags = [
    "b",
    "terraform",
    "a",
    "test"
  ]
}
```

## Argument Reference

The following arguments are supported:

* `tags` - (Required) A set of tags to assign to this resource.
* `name` - (Required) Name of the dashboard.
* `description` - (Required) Human-readable description of the dashboard.
* `url` - (Required) Unique identifier, also a URL slug of the dashboard.
* `section` - (Required) Dashboard chart sections. See [dashboard sections](#dashboard-sections).
* `display_query_parameters` - (Optional) Whether the dashboard parameters section is opened by default when the dashboard
  is shown.
* `parameter_details` - (Optional) The current JSON representation of dashboard parameters. See [parameter details](#parameter-details).
* `display_section_table_of_contents` - (Optional) Whether the "pills" quick-linked the sections of the dashboard are
  displayed by default when the dashboard is shown.
* `can_modify` - (Optional) A list of users that have modify ACL access to the dashboard.
* `can_view` - (Optional) A list of users that have view ACL access to the dashboard.
* `event_filter_type` - (Optional) How charts belonging to this dashboard should display events. `BYCHART` is default if
  unspecified. Valid options are: `BYCHART`, `AUTOMATIC`, `ALL`, `NONE`, `BYDASHBOARD`, and `BYCHARTANDDASHBOARD`.

### Dashboard Sections

The `section` mapping supports the following:

* `name` - (Required) Name of this section.
* `row` - (Required) See [dashboard section rows](#dashboard-section-rows).

### Dashboard Section Rows

The `row` mapping supports the following:

* `chart` - (Required) Charts in this section. See [dashboard chart](#dashboard-chart).

### Dashboard Chart

The `chart` mapping supports the following:

* `source` - (Required) Query expression to plot on the chart. See [chart source queries](#chart-source-queries).
* `chart_setting` - (Required) Chart settings. See [chart settings](#chart-settings).
* `name` - (Required) Name of the source.
* `units` - (Required)  String to label the units of the chart on the Y-Axis.
* `summarization` - (Required) Summarization strategy for the chart. MEAN is default. Valid options are, `MEAN`,
  `MEDIAN`, `MIN`, `MAX`, `SUM`, `COUNT`, `LAST`, `FIRST`.
* `description` - (Optional) Description of the chart.
* `base` - (Optional) The base of logarithmic scale charts. Omit or set to 0 for the default linear scale. Usually set to 10 for the traditional logarithmic scale.
* `no_default_events` - (Optional) Show events related to the sources included in queries

### Chart Source Queries

The `source` mapping supports the following:

* `name` - (Required) Name of the source.
* `query` - (Required) Query expression to plot on the chart.
* `source_description` - (Optional) A description for the purpose of this source.
* `disabled` - (Optional)  Whether the source is disabled.
* `scatter_plot_source` - (Optional) For scatter plots, does this query source the X-axis or the Y-axis, `X`, or `Y`.
* `query_builder_enabled` - (Optional)  Whether or not this source line should have the query builder enabled.

### Chart Settings

The `chart_setting` mapping supports the following:

* `type` - (Required) Chart Type. `line` refers to the Line Plot, `scatter` to the Point Plot, `stacked-area` to
  the Stacked Area plot, `table` to the Tabular View, `scatterplot-xy` to Scatter Plot, `markdown-widget` to the
  Markdown display, and `sparkline` to the Single Stat view. Valid options are`line`, `scatterplot`,
  `stacked-area`, `stacked-column`, `table`, `scatterplot-xy`, `markdown-widget`, `sparkline`, `globe`, `nodemap`,
  `top-k`, `status-list`, and `histogram`.
* `max` - (Optional) Max value of the Y-axis. Set to null or leave blank for auto.
* `line_type` - (Optional) Plot interpolation type.  `linear` is default. Valid options are `linear`, `step-before`,
  `step-after`, `basis`, `cardinal`, and `monotone`.
* `stack_type` - (Optional) Type of stacked chart (applicable only if chart type is `stacked`). `zero` (default) means
  stacked from y=0. `expand` means normalized from 0 to 1.  `wiggle` means minimize weighted changes. `silhouette` means to
  center the stream. Valid options are `zero`, `expand`, `wiggle`, `silhouette`, and `bars`.
* `windowing` - (Optional) For the tabular view, whether to use the full time window for the query or the last X minutes.
  Valid options are `full` or `last`.
* `window_size` - (Optional) Width, in minutes, of the time window to use for `last` windowing.
* `show_hosts` - (Optional) For the tabular view, whether to display sources. Default is `true`.
* `show_labels` - (Optional) For the tabular view, whether to display labels. Default is `true`.
* `show_raw_values` - (Optional) For the tabular view, whether to display raw values. Default is `false`.
* `auto_column_tags` - (Optional) This setting is deprecated.
* `column_tags` - (Optional) This setting is deprecated.
* `tag_mode` - (Optional) For the tabular view, which mode to use to determine which point tags to display.
  Valid options are `all`, `top`, or `custom`.
* `num_tags` - (Optional) For the tabular view defines how many point tags to display.
* `custom_tags` - (Optional) For the tabular view, a list of point tags to display when using the `custom` tag display mode.
* `group_by_source` - (Optional) For the tabular view, whether to group multi metrics into a single row by a common source.
  If `false`, each source is displayed in its own row. if `true`, multiple metrics for the same host are displayed as different
  columns in the same row.
* `sort_values_descending` - (Optional) For the tabular view, whether to display values in descending order. Default is `false`.
* `y1max` - (Optional) For plots with multiple Y-axes, max value for the right side Y-axis. Set null for auto.
* `y1min` - (Optional) For plots with multiple Y-axes, min value for the right side Y-axis. Set null for auto.
* `y1_units` - (Optional) For plots with multiple Y-axes, units for right side Y-axis.
* `y0_scale_si_by_1024` - (Optional) Whether to scale numerical magnitude labels for left Y-axis by 1024 in the IEC/Binary manner (instead of by 1000 like SI).
* `y1_scale_si_by_1024` - (Optional) Whether to scale numerical magnitude labels for right Y-axis by 1024 in the IEC/Binary manner (instead of by 1000 like SI).
* `y0_unit_autoscaling` - (Optional) Whether to automatically adjust magnitude labels and units for the left Y-axis to favor smaller magnitudes and larger units.
* `y1_unit_autoscaling` - (Optional) Whether to automatically adjust magnitude labels and units for the right Y-axis to favor smaller magnitudes and larger units.
* `invert_dynamic_legend_hover_control` - (Optional) Whether to disable the display of the floating legend (but
  reenable it when the ctrl-key is pressed).
* `fixed_legend_enabled` - (Optional) Whether to enable a fixed tabular legend adjacent to the chart.
* `fixed_legend_use_raw_stats` - (Optional) If `true`, the legend uses non-summarized stats instead of summarized.
* `fixed_legend_position` - (Optional)  Where the fixed legend should be displayed with respect to the chart.
  Valid options are `RIGHT`, `TOP`, `LEFT`, `BOTTOM`.
* `fixed_legend_display_stats` - (Optional) For a chart with a fixed legend, a list of statistics to display in the legend.
* `fixed_legend_filter_sort` - (Optional) Whether to display `TOP` or `BOTTOM` ranked series in a fixed legend. Valid options
  are `TOP`, and `BOTTOM`.
* `fixed_legend_filter_limit` - (Optional) Number of series to include in the fixed legend.
* `fixed_legend_filter_field` - (Optional) Statistic to use for determining whether a series is displayed on the fixed legend.
  Valid options are `CURRENT`, `MEAN`, `MEDIAN`, `SUM`, `MIN`, `MAX`, and `COUNT`.
* `fixed_legend_hide_label` - (Optional) This setting is deprecated.
* `xmax` - (Optional) For x-y scatterplots, max value for the X-axis. Set to null for auto.
* `ymax` - (Optional) For x-y scatterplots, max value for the Y-axis. Set to null for auto.
* `xmin` - (Optional) For x-y scatterplots, min value for the X-axis. Set to null for auto.
* `ymin` - (Optional) For x-y scatterplots, min value for the Y-axis. Set to null for auto.
* `time_based_coloring` - (Optional) For x-y scatterplots, whether to color more recent points as darker than older points.
* `sparkline_display_value_type` - (Optional) For the single stat view, where to display the name of the query or the value of the query.
  Valid options are `VALUE` or `LABEL`.
* `sparkline_display_color` - (Optional) For the single stat view, the color of the displayed text (when not dynamically determined).
  Values should be in `rgba(,,,,)` format.
* `sparkline_display_vertical_position` - (Optional) This setting is deprecated.
* `sparkline_display_horizontal_position` - (Optional) For the single stat view, the horizontal position of the displayed text.
  Valid options are `MIDDLE`, `LEFT`, `RIGHT`.
* `sparkline_display_font_size` - (Optional) For the single stat view, the font size of the displayed text, in percent.
* `sparkline_display_prefix` - (Optional) For the single stat view, a string to add before the displayed text.
* `sparkline_display_postfix` - (Optional) For the single stat view, a string to append to the displayed text.
* `sparkline_size` - (Optional) For the single stat view, this determines whether the sparkline of the statistic is displayed in the chart.
  Valid options are `BACKGROUND`, `BOTTOM`, `NONE`.
* `sparkline_line_color` - (Optional) For the single stat view, the color of the line. Values should be in `rgba(,,,,)` format.
* `min` - (Optional) Min value of the Y-axis. Set to null or leave blank for auto.
* `plain_markdown_content` - (Optional)  The markdown content for a Markdown display, in plain text.
* `sparkline_fill_color` - (Optional) For the single stat view, the color of the background fill. Values should be
  in `rgba(,,,,)`.
* `sparkline_value_color_map_colors` - (Optional) For the single stat view, A list of colors that differing query values map to.
  Must contain one more element than `sparkline_value_color_map_values_v2`. Values should be in `rgba(,,,,)`.
* `sparkline_value_color_map_values_v2` - (Optional) For the single stat view, a list of boundaries for mapping different
  query values to colors. Must contain one element less than `sparkline_value_color_map_colors`.
* `sparkline_value_color_map_values` - (Optional) This setting is deprecated.
* `sparkline_value_color_map_apply_to` - (Optional) For the single stat view, whether to apply dynamic color settings to
  the displayed `TEXT` or `BACKGROUND`. Valid options are `TEXT` or `BACKGROUND`.
* `sparkline_decimal_precision` - (Optional) For the single stat view, the decimal precision of the displayed number.
* `sparkline_value_text_map_text` - (Optional) For the single stat view, a list of display text values that different query
  values map to. Must contain one more element than `sparkline_value_text_map_thresholds`.
* `sparkline_value_text_map_thresholds` - (Optional) For the single stat view, a list of threshold boundaries for
  mapping different query values to display text. Must contain one element less than `sparkline_value_text_map_text`.
* `expected_data_spacing` - (Optional) Threshold (in seconds) for time delta between consecutive points in a series
  above which a dotted line will replace a solid in in line plots. Default is 60.

### Parameter Details

The `parameter_details` mapping supports the following:

* `label` - (Required) The label for the parameter.
* `values_to_readable_strings` - (Required) A string->string map. At least one of the keys must match the value of
  `default_value`.
* `name` - (Required) The name of the parameters.
* `default_value` - (Required) The default value of the parameter.
* `hide_from_view` - (Required) If `true` the parameter will only be shown on the edit view of the dashboard.
* `parameter_type` - (Required) The type of the parameter. `SIMPLE`, `LIST`, or `DYNAMIC`.
* `tag_key` - (Optional) for `TAG_KEY` dynamic field types, the tag key to return.
* `query_value` - (Optional) For `DYNAMIC` parameter types, the query to execute to return values.
* `dynamic_field_type` - (Optional) For `DYNAMIC` parameter types, the type of the field. Valid options are `SOURCE`,
  `SOURCE_TAG`, `METRIC_NAME`, `TAG_KEY`, and `MATCHING_SOURCE_TAG`.

### Chart Attributes

The `chart_attribute` mapping is a raw JSON object that supports the full API.
The easiest way to identify the configuration you want, is to edit your dashboard in the UI,
view it as JSON and copy the chartAttributes section.

### Example

```hcl
resource "wavefront_dashboard" "chart_settings_dash" {
  name        = "Terraform Chart Settings"
  description = "testing, testing"
  url         = "tftestcreate"
  section {
    name = "section 1"
    row {
      chart {
        name        = "chart 1"
        description = "chart number 1"
        units       = "something per unit"
        source {
          name  = "source name"
          query = "ts()"
        }
        summarization = "MEAN"
        chart_setting {
          auto_column_tags                      = false
          column_tags                           = "deprecated"
          custom_tags                           = ["tag1", "tag2"]
          expected_data_spacing                 = 120
          fixed_legend_display_stats            = ["stat1", "stat2"]
          fixed_legend_enabled                  = true
          fixed_legend_filter_field             = "CURRENT"
          fixed_legend_filter_limit             = 1
          fixed_legend_filter_sort              = "TOP"
          fixed_legend_hide_label               = false
          fixed_legend_position                 = "RIGHT"
          fixed_legend_use_raw_stats            = true
          line_type                             = "linear"
          max                                   = 100
          min                                   = 0
          sparkline_decimal_precision           = 1
          sparkline_display_color               = "rgba(1,1,1,1)"
          sparkline_display_font_size           = 14
          sparkline_display_horizontal_position = "LEFT"
          sparkline_display_postfix             = "postfix"
          sparkline_display_prefix              = "prefix"
          sparkline_display_value_type          = "VALUE"
          sparkline_display_vertical_position   = "deprecated"
          sparkline_fill_color                  = "rgba(1,1,1,1)"
          sparkline_line_color                  = "rgba(1,1,1,1)"
          sparkline_size                        = "BOTTOM"
          sparkline_value_color_map_apply_to    = "TEXT"
          sparkline_value_color_map_colors      = ["rgba(1,1,1,1)", "rgba(2,2,2,2)", "rgba(3,3,3,3)"]
          sparkline_value_color_map_values      = [1, 2]
          sparkline_value_text_map_text         = ["a"]
          sparkline_value_text_map_thresholds   = [1]
          type                                  = "line"
          y0_scale_si_by_1024                   = true
          y0_unit_autoscaling                   = true
          y1max                                 = 100
          y1_scale_si_by_1024                   = true
          y1_unit_autoscaling                   = true
          y1_units                              = "units"
        }
        chart_attribute = <<-EOT
          {
            "dashboardLinks": {
              "*": {
                "variables": {
                  "xxx": "xxx"
                },
                "destination": "/dashboards/xxxx"
              }
            }
          }
        EOT
      }
    }
  }
  parameter_details {
    name                       = "param1"
    label                      = "param1"
    default_value              = "defaultQuery"
    hide_from_view             = false
    parameter_type             = "DYNAMIC"
    values_to_readable_strings = {
      defaultQuery = "dev-elasticsearch"
    }
    dynamic_field_type = "MATCHING_SOURCE_TAG"
    query_value        = "ts(aws.ec2.diskwritebytes.average)"
  }
  tags = [
    "terraform",
    "test"
  ]
}
```

## Import

Dashboards can be imported by using the `id`, e.g.:

```
$ terraform import wavefront_dashboard.dashboard tftestimport
```
