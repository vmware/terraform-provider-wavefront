package wavefront

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"

	"github.com/WavefrontHQ/go-wavefront-management-api/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/assert"
)

const (
	testAlertName = "test-name"
	testHost1     = "foo.com"
	testHost2     = "bar.com"
	testFiring0   = 0
	testFiring1   = 1
	testLink1     = "test-link-1"
	testLink2     = "test-link-2"
	testKey       = "test-key"
)

func TestAccAlertIDRequired(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccAlertIDRequiredFailConfig,
				ExpectError: regexp.MustCompile("The argument \"id\" is required, but no definition was found."),
			},
		},
	})
}

func TestParseAlertTriageDashboardParameters(t *testing.T) {
	// Create some test parameters
	expected := map[string]map[string]string{
		constantsKey: {testKey1: testVal1, testKey2: testVal2},
		testKey:      {testKey3: testVal3, testKey4: testVal4},
	}

	// Parse the test parameters
	actual := parseAlertTriageDashboardParameters(expected)

	// Assert that the parsed parameters match the input
	assert.True(t, parsedParametersMatch(expected, actual))
}

func TestParseAlertTriageDashboardWithParameters(t *testing.T) {
	// Create some test parameters
	expectedParams := map[string]map[string]string{
		constantsKey: {testKey1: testVal1, testKey2: testVal2},
		testKey:      {testKey3: testVal3, testKey4: testVal4},
	}

	// Create a test dashboard
	expected := wavefront.AlertTriageDashboard{
		DashboardId: testDashboardID1,
		Description: testDashboardDesc1,
		Parameters:  expectedParams,
	}

	// Parse the test dashboard
	actual := parseAlertTriageDashboard(expected)

	// Assert that the parsed dashboard matches the input
	assert.True(t, parsedAlertTriageDashboardMatch(expected, actual))
}

func TestParseAlertTriageDashboardWithoutParameters(t *testing.T) {
	// Create a test dashboard
	expected := wavefront.AlertTriageDashboard{
		DashboardId: testDashboardID1,
		Description: testDashboardDesc1,
		Parameters:  map[string]map[string]string{},
	}

	// Parse the test dashboard
	actual := parseAlertTriageDashboard(expected)

	// Assert that the parsed dashboard matches the input
	assert.True(t, parsedAlertTriageDashboardMatch(expected, actual))
}

func TestParseAlertTriageDashboards(t *testing.T) {
	// Create some test parameters
	expectedParams := map[string]map[string]string{
		constantsKey: {testKey1: testVal1, testKey2: testVal2},
		testKey:      {testKey3: testVal3, testKey4: testVal4},
	}

	// Create an array of test dashboards
	expected := []wavefront.AlertTriageDashboard{
		{
			DashboardId: testDashboardID1,
			Description: testDashboardDesc1,
			Parameters:  expectedParams,
		},
		{
			DashboardId: testDashboardID2,
			Description: testDashboardDesc2,
			Parameters:  expectedParams,
		},
	}

	// Parse the test dashboards
	actual := parseAlertTriageDashboards(expected)

	// Assert that the parsed result matches the input
	assert.Equal(t, len(expected), len(actual))
	for _, eachExpected := range expected {
		foundMatch := false
		for _, eachActual := range actual {
			if parsedAlertTriageDashboardMatch(eachExpected, eachActual) {
				foundMatch = true
				break
			}
		}
		assert.True(t, foundMatch)
	}
}

func TestFlattenHostLabelPair(t *testing.T) {
	// Create a test SourceLabelPair
	expected := wavefront.SourceLabelPair{
		Host:   testHost1,
		Firing: testFiring1,
	}

	// Call the function to flatten the SourceLabelPair
	actual := flattenHostLabelPair(expected)

	// Assert the result matches the input
	assert.True(t, flattenedHostLabelPairMatch(expected, actual))
}

func TestFlattenHostLabelPairs(t *testing.T) {
	// Create a slice of test SourceLabelPairs
	expected := []wavefront.SourceLabelPair{
		{
			Host:   testHost1,
			Firing: testFiring1,
		},
		{
			Host:   testHost2,
			Firing: testFiring0,
		},
	}

	// Call the function to flatten the slice of SourceLabelPairs
	actual := flattenHostLabelPairs(expected).([]map[string]interface{})

	// First compare the count of host label pairs
	assert.Equal(t, len(expected), len(actual), "did not get the correct number of flattened host label pairs")

	// Now compare the contents
	for _, eachExpected := range expected {
		foundMatch := false
		for _, eachActual := range actual {
			if flattenedHostLabelPairMatch(eachExpected, eachActual) {
				foundMatch = true
				break
			}
		}
		assert.True(t, foundMatch, fmt.Sprintf("did not find a match for: wavefront.SourceLabelPair%v\n", eachExpected))
	}
}

func TestSetAlertAttributes(t *testing.T) {
	var testAlertID = "test-id"
	testAlert := wavefront.Alert{
		Name:         testAlertName,
		ID:           &testAlertID,
		RunbookLinks: []string{testLink1, testLink2},
		AlertTriageDashboards: []wavefront.AlertTriageDashboard{
			{
				DashboardId: testDashboardID1,
				Description: testDashboardDesc1,
				Parameters:  map[string]map[string]string{constantsKey: {testKey1: testVal1, testKey2: testVal2}},
			},
		},
	}

	err := setAlertAttributes(dataSourceAlert().TestResourceData(), testAlert)

	// Assert that there was no error setting these values
	assert.NoError(t, err)
}

const testAccAlertIDRequiredFailConfig = `
data "wavefront_alert" "test_alert" {
}
`

func parsedParametersMatch(expected map[string]map[string]string, actual []map[string]interface{}) bool {
	// First compare the lengths
	if len(expected) != len(actual) {
		return false
	}

	// Then compare the contents
	for _, eachParam := range actual {
		for paramKey, paramValue := range eachParam {
			expectedParam := expected[paramKey]
			actualParam := paramValue.(map[string]string)
			if !reflect.DeepEqual(expectedParam, actualParam) {
				return false
			}
		}
	}
	return true
}

func parsedAlertTriageDashboardMatch(expected wavefront.AlertTriageDashboard, actual map[string]interface{}) bool {

	if expected.DashboardId != actual[dashboardIDKey].(string) {
		return false
	}

	if expected.Description != actual[descriptionKey].(string) {
		return false
	}

	return parsedParametersMatch(expected.Parameters, actual[parametersKey].([]map[string]interface{}))
}

func flattenedHostLabelPairMatch(expected wavefront.SourceLabelPair, actual map[string]interface{}) bool {
	if expected.Host != actual[hostKey].(string) {
		return false
	}

	if expected.Firing != actual[firingKey].(int) {
		return false
	}

	return true
}
