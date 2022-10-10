package wavefront

import (
	"encoding/json"
	"fmt"

	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	accessTypeKey                = "access_type"
	accountsKey                  = "account_ids"
	customerKey                  = "customer"
	descriptionKey               = "description"
	emailKey                     = "email"
	exactMatching                = "EXACT"
	idKey                        = "id"
	lastSuccessfulLoginKey       = "last_successful_login"
	nameKey                      = "name"
	pageSize                     = 100
	permissionsKey               = "permissions"
	prefixesKey                  = "prefixes"
	policyRulesKey               = "policy_rules"
	policyTagKey                 = "key"
	policyTagValue               = "value"
	roleIdsTagKey                = "role_ids"
	rolesKey                     = "roles"
	tagsAndedKey                 = "tags_anded"
	tagsKey                      = "tags"
	updatedEpochMillisKey        = "updated_epoch_millis"
	userGroupsKey                = "user_group_ids"
	userGroupsListKey            = "user_groups"
	usersKey                     = "users"
	queryKey                     = "query"
	minutesKey                   = "minutes"
	inTrashKey                   = "in_trash"
	queryFailingKey              = "query_failing"
	lastFailedTimeKey            = "last_failed_time"
	lastErrorMessageKey          = "last_error_message"
	additionalInformationKey     = "additional_information"
	updateUserIDKey              = "update_user_id"
	createUserIDKey              = "create_user_id"
	statusKey                    = "status"
	hostsUsedKey                 = "hosts_used"
	lastProcessedMillisKey       = "last_processed_millis"
	processRateMinutesKey        = "process_rate_minutes"
	pointsScannedAtLastQueryKey  = "points_scanned_at_last_query"
	includeObsoleteMetricsKey    = "include_obsolete_metrics"
	lastQueryTimeKey             = "last_query_time"
	metricsUsedKey               = "metrics_used"
	queryQBEnabledKey            = "query_qb_enabled"
	updatedEpochMillisKey1       = "updated_epoch_millis"
	createdEpochMillisKey        = "created_epoch_millis"
	deletedKey                   = "deleted"
	derivedMetricsKey            = "derived_metrics"
	urlKey                       = "url"
	startTimeKey                 = "start_time"
	endTimeKey                   = "endtime_key"
	severityKey                  = "severity"
	detailsKey                   = "details"
	isEphemeralKey               = "is_ephemeral"
	annotationsKey               = "annotations"
	eventsKey                    = "events"
	latestStartTimeEpochMillis   = "latest_start_time_epoch_millis"
	earliestStartTimeEpochMillis = "earliest_start_time_epoch_millis"
	limitKey                     = "limit"
	offsetKey                    = "offset"
	dashboardsKey                = "dashboards"
)

// compareStringSliceAnyOrder compares two string slices in any order. It returns
// all the strings appearing only in the left slice followed by all the strings
// appearing only in the right slice.
func compareStringSliceAnyOrder(left, right []string) (onlyLeft, onlyRight []string) {
	leftMap := stringSliceAnyOrderAsMap(left)
	rightMap := stringSliceAnyOrderAsMap(right)
	for str, leftCount := range leftMap {
		rightCount := rightMap[str]
		for leftCount < rightCount {
			onlyRight = append(onlyRight, str)
			rightCount--
		}
		for leftCount > rightCount {
			onlyLeft = append(onlyLeft, str)
			leftCount--
		}
		delete(rightMap, str)
	}
	for str, rightCount := range rightMap {
		for rightCount > 0 {
			onlyRight = append(onlyRight, str)
			rightCount--
		}
	}
	return
}

func stringSliceAnyOrderAsMap(strs []string) map[string]int {
	result := make(map[string]int)
	for _, s := range strs {
		result[s]++
	}
	return result
}

// setStringSlice stores a slice of strings as a set under a particular key.
func setStringSlice(d *schema.ResourceData, key string, strs []string) error {
	result := make([]interface{}, 0, len(strs))
	for _, str := range strs {
		result = append(result, str)
	}
	return d.Set(key, schema.NewSet(schema.HashString, result))
}

// getStringSlice returns a set as a slice of strings.
func getStringSlice(d *schema.ResourceData, key string) []string {
	interfaceList := d.Get(key).(*schema.Set).List()
	result := make([]string, 0, len(interfaceList))
	for _, val := range interfaceList {
		result = append(result, val.(string))
	}
	return result
}

// setStringMap stores a map[string]string under a particular key.
func setStringMap(
	d *schema.ResourceData, key string, strMap map[string]string) error {
	result := make(map[string]interface{}, len(strMap))
	for k, v := range strMap {
		result[k] = v
	}
	return d.Set(key, result)
}

// getStringMap retrieves a map[string]string under a particular key
func getStringMap(d *schema.ResourceData, key string) map[string]string {
	interfaceMap := d.Get(key).(map[string]interface{})
	result := make(map[string]string, len(interfaceMap))
	for k, v := range interfaceMap {
		result[k] = v.(string)
	}
	return result
}

// parseStrArr parses a raw interface from d *schema.ResourceData that contains an array of strings
func parseStrArr(raw interface{}) []string {
	var arr []string
	if raw != nil && len(raw.([]interface{})) > 0 {
		for _, v := range raw.([]interface{}) {
			arr = append(arr, v.(string))
		}
	}
	return arr
}

func searchAll(limit int, offset int, typ string, timeRange *wavefront.TimeRange, filter []*wavefront.SearchCondition, m interface{}) json.RawMessage {
	searchParams := &wavefront.SearchParams{
		Conditions: filter,
		Limit:      limit,
		Offset:     offset,
		TimeRange:  timeRange,
	}

	searchClient := m.(*wavefrontClient).client.NewSearch(typ, searchParams)
	var searchResponse *wavefront.SearchResponse
	var err error
	searchResponse, err = searchClient.Execute()

	if err != nil {
		fmt.Errorf("Failed to search for dashboards")
	}

	return searchResponse.Response.Items
}
