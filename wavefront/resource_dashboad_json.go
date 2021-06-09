package wavefront

import (
	"fmt"
	"log"
	"strings"

	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDashboardJSON() *schema.Resource {
	return &schema.Resource{
		Create: resourceDashboardJSONCreate,
		Read:   resourceDashboardJSONRead,
		Update: resourceDashboardJSONUpdate,
		Delete: resourceDashboardJSONDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"dashboard_json": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: ValidateDashboardJSON,
				StateFunc:    NormalizeDashboardJSON,
			},
		},
	}

}

func buildDashboardJSON(d *schema.ResourceData) (*wavefront.Dashboard, error) {
	var dashboard wavefront.Dashboard
	dashboardJSONString := d.Get("dashboard_json").(string)
	// json is already validated during resource Validation
	_ = dashboard.UnmarshalJSON([]byte(dashboardJSONString))

	// set url name as the resource ID
	dashboard.ID = dashboard.Url
	return &dashboard, nil
}

func resourceDashboardJSONRead(d *schema.ResourceData, meta interface{}) error {
	dashboards := meta.(*wavefrontClient).client.Dashboards()
	dash := wavefront.Dashboard{
		ID: d.Id(),
	}
	err := dashboards.Get(&dash)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("error finding Wavefront Dashboard %s. %s", d.Id(), err)
	}
	bytes, _ := dash.MarshalJSON()
	// Use the Wavefront url as the Terraform ID
	d.SetId(dash.ID)
	err = d.Set("dashboard_json", NormalizeDashboardJSON(string(bytes)))
	if err != nil {
		return fmt.Errorf("failed to set dashboard json %s. %s", d.Id(), err)
	}
	return nil
}

func resourceDashboardJSONCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Create Wavefront Dashboard %s", d.Id())
	dashboards := meta.(*wavefrontClient).client.Dashboards()
	dashboard, err := buildDashboardJSON(d)

	if err != nil {
		return fmt.Errorf("failed to parse dashboard, %s", err)
	}

	err = dashboards.Create(dashboard)
	if err != nil {
		return fmt.Errorf("failed to create dashboard, %s", err)
	}
	d.SetId(dashboard.ID)
	log.Printf("[INFO] Wavefront Dashboard %s Created", d.Id())
	return resourceDashboardJSONRead(d, meta)
}

func resourceDashboardJSONUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Update Wavefront Dashboard %s", d.Id())
	dashboards := meta.(*wavefrontClient).client.Dashboards()
	dashboard, err := buildDashboardJSON(d)

	if err != nil {
		return fmt.Errorf("failed to parse dashboard, %s", err)
	}

	err = dashboards.Update(dashboard)
	if err != nil {
		return fmt.Errorf("failed to create dashboard, %s", err)
	}

	log.Printf("[INFO] Wavefront Dashboard %s Updated", d.Id())
	return resourceDashboardJSONRead(d, meta)
}

func resourceDashboardJSONDelete(d *schema.ResourceData, meta interface{}) error {
	dashboards := meta.(*wavefrontClient).client.Dashboards()
	dash := wavefront.Dashboard{
		ID: d.Id(),
	}

	err := dashboards.Get(&dash)
	if err != nil {
		// Dashboard has already been deleted, so we'll mark it as such...
		if strings.Contains(err.Error(), "404") {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("error finding Wavefront Dashboard %s. %s", d.Id(), err)
	}

	// Delete the Dashboard
	err = dashboards.Delete(&dash, true)
	if err != nil {
		return fmt.Errorf("failed to delete Dashboard %s. %s", d.Id(), err)
	}
	d.SetId("")
	return nil
}

func ValidateDashboardJSON(val interface{}, key string) ([]string, []error) {
	dashboardJSONString := val.(string)
	var dashboard wavefront.Dashboard
	err := dashboard.UnmarshalJSON([]byte(dashboardJSONString))
	if err != nil {
		return nil, []error{err}
	}
	return nil, nil
}

func NormalizeDashboardJSON(val interface{}) string {
	dashboardJSONString := val.(string)
	var dashboard wavefront.Dashboard
	_ = dashboard.UnmarshalJSON([]byte(dashboardJSONString))

	// set url name as the resource ID
	dashboard.ID = dashboard.Url

	// remove keys which are not needed for diff
	dashboard.CreatedEpochMillis = 0
	dashboard.UpdatedEpochMillis = 0
	dashboard.CreatorId = ""
	dashboard.UpdaterId = ""
	dashboard.Customer = ""
	dashboard.ViewsLastDay = 0
	dashboard.ViewsLastWeek = 0
	dashboard.ViewsLastMonth = 0
	dashboard.NumCharts = 0
	dashboard.NumFavorites = 0
	dashboard.Favorite = false

	ret, _ := dashboard.MarshalJSON()
	return string(ret)
}
