## 2.4.0 (Unreleased)
## 2.3.1 (July 02, 2020)

BUG FIXES:

* resource/wavefront_cloud_integration_*: Fixed an issue where resource lookup would crash when no results were returned.

ENHANCEMENTS:

* provider/wavefront: Support for `http_proxy` in provider and support for environment variables `http_proxy` and `https_proxy`
* resource/wavefront_alert_target: new `target_id` computed value which prefixes `target:` onto the id for joining into wavefront alerts.
* resource/wavefront_alert: added validation on `target` field


## 2.3.0 (June 02, 2020)

BREAKING CHANGES:

* resource/wavefront_alert: `threshold_conditions` renamed to `conditions` based on API naming.
* resource/wavefront_user: `groups` aptly renamed to `permissions` to reflect the user/group/role permissions model
* resource/wavefront_group: `permissions` has been removed as they are no longer supported directly on groups

FEATURES:

* **New Resource:** `wavefront_role`

BUG FIXES:

* resource/wavefront_cloud_integration_tesla: Fixed issue where `force_new` changes caused new resource
* resource/wavefront_cloud_integration_gcp: Fixed `categories` not being properly persisted to state
* resource/wavefront_cloud_integration_ec2: Fixed `hostname_tags` not being properly persisted to state

ENHANCEMENTS:

* resource/wavefront_dashboard: `parameter_details` will no longer always show changes
* resource/wavefront_dashboard: `tags` set on a dashboard will update properly

## 2.2.0 (May 03, 2020)

NOTES:

* Updated to latest go-wavefront-management-api
* Temporarily fixed some failing tests by skipping non-empty plans
* Cleaned up text in README/CONTRIBUTING where stale repo pointers existed

FEATURES:

* **New Resource:** `wavefront_cloud_integration_cloudwatch`
* **New Resource:** `wavefront_cloud_integration_cloudtrail`
* **New Resource:** `wavefront_cloud_integration_ec2`
* **New Resource:** `wavefront_cloud_integration_gcp`
* **New Resource:** `wavefront_cloud_integration_gcp_billing`
* **New Resource:** `wavefront_cloud_integration_azure`
* **New Resource:** `wavefront_cloud_integration_azure_activity_log`
* **New Resource:** `wavefront_cloud_integration_newrelic`
* **New Resource:** `wavefront_cloud_integration_app_dynamics`
* **New Resource:** `wavefront_cloud_integration_tesla`

ENHANCEMENTS:

* resource/wavefront_alert: On delete will call `skipTrash` to prevent cluttering trashcan
* resource/wavefront_dashboard: On delete will call `skipTrash` to prevent cluttering trashcan
* resource/wavefront_dashboard_json: On delete will call `skipTrash` to prevent cluttering trashcan
* resource/wavefront_derived_metric: On delete will call `skipTrash` to prevent cluttering trashcan

## 2.1.3 (February 10, 2020)

NOTES:

* Consistent error messages casing across all errors


FEATURES:

* **New Resource:** `wavefront_default_user_group`

ENHANCEMENTS:

* resource/wavefront_alert: Support for `can_view` and `can_modify` ACL
* resource/wavefront_dashboard: Support for `can_view` and `can_modify` ACL
* resource/wavefront_dashboard_json: Support for `can_view` and `can_modify` ACL


## 2.1.2 (January 13, 2020)

NOTES:

* resource/wavefront_user: adding missing import tests
* resource/wavefront_derived_metrics: added missing import tests

FEATURES:

* **New Resource:** `wavefront_user_group`

BUG FIXES:

* resource/wavefront_derived_metrics: Fixed issue where Derived Metrics were not reading tags

## 2.1.1 (December 19, 2019)

FEATURES:

* **New Resource:** `wavefront_derived_metrics`
* **New Resource:** `wavefront_user`
* **New Resource:** `wavefront_alert_target`

BUG FIXES:

* Fixed issue where deleted resources would not properly detect to recreate resource on plan/apply

## 2.1.0 (July 03, 2019)

FEATURES:

* **New Resource:** `wavefront_dashboard_json`

ENHANCEMENTS:

* resource/wavefront_alert: Added support for threshold alerts

## 2.0.0 (June 11, 2019)

NOTES:

* Upgrade to Terraform 0.12 to support new language features
* May cause breaking changes due to new syntax ([See Upgrading to 0.12](https://www.terraform.io/upgrade-guides/0-12.html))
* In testing `values_to_readable_strings {` needed to change to `values_to_readable_strings = {` and `is_html_content = 1` changed to `is_html_content = true`

## 1.0.1 (January 08, 2018)

BUG FIXES:

* resource/dashbaord: Sort parameter details alphabetically to ensure no changes they are always evaluated in the correct order*

## 1.0.0 (December 29, 2017)

BREAKING CHANGES:

* resource/dashboard: Added support for Dynamic and List parameter types*
* string_key and string_value have been removed from parameter_detail
* values_to_readable_strings replaces string_key and string_value as a map[string]string. Each key in the map is 
effectively a separate string_key and the value is a separate string_value.
* The value of default_value must equal one of the keys (not value) within the values_to_readable_string map.

ENHANCEMENTS:

* resource/dashboard: Add missing fields to source. Allow disabled, `scatter_plot_source`, `query_builder_enabled`, `source_description` to optionally be applied to sources

## 0.2.0 (January 03, 2018)

NOTES:

* Updated README section on handling the creation of multiple alerts

ENHANCEMENTS:

* resource/alert: Fixed #11 - `condition`, `display_expression`, and `additional_information` have been updated to call TrimSpaces. Preventing multiple plan/applies from re-applying the same state.

## 0.1.2 (October 13, 2017)

ENHANCEMENTS:

* resource/alert:  Allow optional Alert attributes (as defined by the API) to be omitted from Terraform. `display_expression` and `resolve_after_minutes` are now optional.

## 0.1.1 (October 12, 2017)

NOTES: 

* Builds both linux and darwin versions of the plugin and uploads them all to github releases.

## 0.1.0 (September 15, 2017)

NOTES: 

* First Release - Supports a limited Set of the Wavefront API*

NEW FEATURES:

* **New Resource:** `wavefront_alert`
* **New Resource:** `wavefront_dashboard`
