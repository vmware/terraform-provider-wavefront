package wavefront

import (
	"fmt"

	"github.com/WavefrontHQ/go-wavefront-management-api/v2"
)

type wavefrontClient struct {
	client wavefront.Client
}

func newWavefrontClient(address, token string, proxy string, cspAddress string) (interface{}, error) {
	if address == "" {
		return nil, fmt.Errorf("wavefront address required")
	}

	if token == "" {
		return nil, fmt.Errorf("token required")
	}

	if cspAddress == "" {
		cspAddress = "console.cloud.vmware.com"
	}

	//TODO: Add details about the CSP staging environment:
	//	https://console-stg.cloud.vmware.com/csp/gateway/discovery and the new Org there "Terraform Provider"
	//token, err := auth.GetCspBearerToken(cspAddress, token, false)
	//if err != nil {
	//	return nil, fmt.Errorf("Could not get CSP bearer token: %v\n", err)
	//}

	wFClient, err := wavefront.NewClient(&wavefront.Config{
		Address:   address,
		Token:     token,
		HttpProxy: proxy,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to configure Wavefront Client %s", err)
	}

	return &wavefrontClient{client: *wFClient}, nil
}
