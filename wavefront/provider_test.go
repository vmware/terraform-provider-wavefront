package wavefront

import (
	"testing"

	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"wavefront": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(_ *testing.T) {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("WAVEFRONT_TOKEN"); v == "" {
		t.Fatal("WAVEFRONT_TOKEN must be set for acceptance tests")
	}
	if v := os.Getenv("WAVEFRONT_ADDRESS"); v == "" {
		t.Fatal("WAVEFRONT_ADDRESS must be set for acceptance tests")
	}
}
