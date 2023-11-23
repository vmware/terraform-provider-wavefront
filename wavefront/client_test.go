package wavefront

import (
	"testing"
)

func TestGetAPIClient(t *testing.T) {
	address := "example.com"
	token := "your_token"
	proxy := "https://proxy.example.com"
	cspAddress := "csp.example.com"

	// Test case for successful configuration
	client, err := newWavefrontClient(address, token, proxy, cspAddress)
	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}

	// Assert that the returned client is of type wavefrontClient
	_, ok := client.(*wavefrontClient)
	if !ok {
		t.Errorf("Expected the returned client to be a wavefrontClient")
	}

	// Test case for error in configuration
	_, err = newWavefrontClient("", "", "", "")
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	// Test case for error in configuration (address)
	_, err = newWavefrontClient("", token, "", "")
	if err == nil {
		t.Errorf("Expected an address error, but got nil")
	}

	// Test case for error in configuration (token)
	_, err = newWavefrontClient(address, "", "", "")
	if err == nil {
		t.Errorf("Expected token error, but got nil")
	}

	// Test case for optional configuration (proxy)
	_, err = newWavefrontClient(address, token, "", cspAddress)
	if err != nil {
		t.Errorf("Optional parameter doesn't throw error, but go an error")
	}

	// Test case for optional configuration (cspAddress)
	_, err = newWavefrontClient(address, token, proxy, "")
	if err != nil {
		t.Errorf("Optional parameter doesn't throw error, but go an error")
	}
}
