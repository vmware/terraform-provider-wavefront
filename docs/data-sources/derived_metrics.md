---
layout: "wavefront"
page_title: "Wavefront: Derived Metrics"
description: |-
Get the information for all derived metrics from Wavefront
---

# Data Source: wavefront_derived_metrics

Use this data source to get information for all derived metrics from Wavefront.

## Example Usage

```hcl
# Get the info for all derived metrics
data "wavefront_derived_metrics" "example" {
}
```

## Attribute Reference

* `derived_metrics` - List of all derived metrics in Wavefront.
    * `name` - The name of the Derived Metric in Wavefront.
    * `id` - The id of the Derived Metric in Wavefront.
    * `query`- A Wavefront query that is evaluated at regular intervals (default is 1 minute).
    * `minutes` - How frequently the query generating the derived metric is run.
    * `in_trash` - A Boolean variable indicating trash status.
    * `tags` - A set of tags assigned to the Derived Metric.
    * `query_failing` - A Boolean variable indicating whether query is failing for derived metric.
    * `last_error_message` - Last error message occurred.
    * `last_failed_time` - Timestamp of last failed derived metric.
    * `additional_information` - User-supplied additional explanatory information for the derived metric.
    * `create_user_id` - The id of user who created derived metric.
    * `update_user_id` - The id of user who updated derived metric.
    * `status` - The status of derived metric.
    * `hosts_used` - A list of hosts used in derived metric.
    * `last_processed_millis` - The last processed timestamp.
    * `process_rate_minutes` -  The specified query is executed every `process_rate_minutes` minutes.
    * `points_scanned_at_last_query` - The number of points scanned when last query was executed.
    * `include_obsolete_metrics` - A boolean flag indicating whether to include obsolete metrics or not.
    * `last_query_time` - The timestamp when query was executed last time.
    * `metrics_used` - A list of metrics used in a derived metric.
    * `query_qb_enabled` - A boolean flag for enabling `query_qb`
    * `deleted` - A Boolean flag indicating derived metric deleted or not.
    * `created_epoch_millis` - The timestamp in epoch millis when derived metrics is created.
    * `updated_epoch_millis` - The timestamp in epoch millis when derived metrics is updated.


	
