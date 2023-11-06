package wavefront

import (
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

func TestParseAlertTriageDashboards(t *testing.T) {
	var dashboards []wavefront.AlertTriageDashboard

	dashboard1 := wavefront.AlertTriageDashboard{}
	dashboard1.DashboardId = testDashboardID1
	dashboard1.Description = testDashboardDesc1
	dashboard1.Parameters = map[string]map[string]string{constantsKey: {testKey1: testVal1, testKey2: testVal2}}
	dashboards = append(dashboards, dashboard1)

	dashboard2 := wavefront.AlertTriageDashboard{}
	dashboard2.DashboardId = testDashboardID2
	dashboard2.Description = testDashboardDesc2
	dashboard2.Parameters = map[string]map[string]string{constantsKey: {testKey3: testVal3, testKey4: testVal4}}
	dashboards = append(dashboards, dashboard2)

	result := parseAlertTriageDashboards(dashboards)

	expected := []map[string]interface{}{
		{
			dashboardIDKey: testDashboardID1,
			descriptionKey: testDashboardDesc1,
			parametersKey: []map[string]interface{}{
				{
					constantsKey: map[string]string{
						testKey1: testVal1,
						testKey2: testVal2,
					},
				},
			},
		},
		{
			dashboardIDKey: testDashboardID2,
			descriptionKey: testDashboardDesc2,
			parametersKey: []map[string]interface{}{
				{
					constantsKey: map[string]string{
						testKey3: testVal3,
						testKey4: testVal4,
					},
				},
			},
		},
	}

	assert.Equal(t, expected, result)

}

func TestFlattenHostLabelPair(t *testing.T) {
	// Create a sample SourceLabelPair
	pair := wavefront.SourceLabelPair{
		Host:   testHost1,
		Firing: testFiring1,
	}

	// Call the function to flatten the SourceLabelPair
	flattened := flattenHostLabelPair(pair)

	// Check if the expected values are present in the flattened map
	expected := map[string]interface{}{
		hostKey:   testHost1,
		firingKey: testFiring1,
	}

	for key, value := range expected {
		if flattened[key] != value {
			t.Errorf("Expected %s to be %v, but got %v", key, value, flattened[key])
		}
	}
}

func TestFlattenHostLabelPairs(t *testing.T) {
	// Create a slice of sample SourceLabelPairs
	pairs := []wavefront.SourceLabelPair{
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
	flattened := flattenHostLabelPairs(pairs)

	// Check if the flattened result is of the expected length
	expectedLength := len(pairs)
	if len(flattened.([]map[string]interface{})) != expectedLength {
		t.Errorf("Expected flattened length to be %d, but got %d", expectedLength, len(flattened.([]map[string]interface{})))
	}

	// Check if each flattened map matches the expected values
	expected := []map[string]interface{}{
		{
			hostKey:   testHost1,
			firingKey: testFiring1,
		},
		{
			hostKey:   testHost2,
			firingKey: testFiring0,
		},
	}

	for i, expectedMap := range expected {
		flattenedMap := flattened.([]map[string]interface{})[i]
		for key, value := range expectedMap {
			if flattenedMap[key] != value {
				t.Errorf("Expected %s to be %v, but got %v", key, value, flattenedMap[key])
			}
		}
	}
}

func TestSetAlertAttributes(t *testing.T) {
	resource := dataSourceAlert()
	var testAlertID = "test-id"
	var testLinks = []string{testLink1, testLink2}
	var testParameters = map[string]map[string]string{constantsKey: {testKey1: testVal1, testKey2: testVal2}}
	var testAlertTriageDashboards = []wavefront.AlertTriageDashboard{
		{
			DashboardId: testDashboardID1,
			Description: testDashboardDesc1,
			Parameters:  testParameters,
		},
	}

	// Create a sample schema.ReourceData
	d := resource.TestResourceData()
	alert := wavefront.Alert{
		Name:                  testAlertName,
		ID:                    &testAlertID,
		RunbookLinks:          testLinks,
		AlertTriageDashboards: testAlertTriageDashboards,
	}

	err := setAlertAttributes(d, alert)

	// Check if there's no error
	assert.NoError(t, err)

	// Check if the fields in schema.ResourceData have been set correctly
	assert.Equal(t, testAlertName, d.Get("name"))
	assert.Equal(t, testAlertID, d.Get("id"))

	// Check that the runbook_links field was set correctly
	runbookLinksData := d.Get(runbookLinksKey).([]interface{})
	assert.Equal(t, len(testLinks), len(runbookLinksData))
	foundCount := 0
	for _, linkData := range runbookLinksData {
		for _, testLink := range testLinks {
			if linkData == testLink {
				foundCount++
				break
			}
		}
	}
	assert.Equal(t, len(testLinks), foundCount)

	// Check that the alert_triage_dashboards field was set correctly
	alertTriageDashboardsData := d.Get(alertTriageDashboardsKey).([]interface{})
	assert.Equal(t, len(testAlertTriageDashboards), len(alertTriageDashboardsData))
	dashboardData := alertTriageDashboardsData[0].(map[string]interface{})
	assert.Equal(t, testAlertTriageDashboards[0].DashboardId, dashboardData[dashboardIDKey].(string))
	assert.Equal(t, testAlertTriageDashboards[0].Description, dashboardData[descriptionKey].(string))
	parameterData := dashboardData[parametersKey].([]interface{})
	assert.Equal(t, len(testParameters), len(parameterData))
	constantsData := parameterData[0].(map[string]interface{})[constantsKey].(map[string]interface{})
	assert.Equal(t, len(testParameters[constantsKey]), len(constantsData))
	for key, value := range constantsData {
		if key == testKey1 {
			assert.Equal(t, testVal1, value)
		}
		if key == testKey2 {
			assert.Equal(t, testVal2, value)
		}
	}
}

const testAccAlertIDRequiredFailConfig = `
data "wavefront_alert" "test_alert" {
}
`
