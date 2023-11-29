package auth

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	CspTokenLength = 64
)

// GetCspBearerToken GetAuthToken retrieves auth token for csp users
// When accessing and endpoint secured by CSP,
// the received `token` must be provided in the
// `Authorization` request header field as follows:
// `Authorization: Bearer {token}`
func GetCspBearerToken(cspAddress, token string, skipTLS bool) (string, error) {
	// Refresh Token, Access Token and Wavefront Service Account Tokens are different lengths.
	// If CSP Refresh token:
	if len(token) == CspTokenLength {
		return "", nil
	}
	fmt.Printf("Getting CSP bearer token\n")

	httpClient := &http.Client{
		Transport: &http.Transport{
			Proxy:           http.ProxyFromEnvironment,
			TLSNextProto:    map[string]func(authority string, c *tls.Conn) http.RoundTripper{},
			TLSClientConfig: &tls.Config{InsecureSkipVerify: skipTLS},
		},
	}

	request, err := cspTokenRequest(cspAddress, token)
	if err != nil {
		return "", err
	}

	rawResp, err := httpClient.Do(request)
	if err != nil {
		return "", err
	}

	result, err := ReadResponse(rawResp)
	if err != nil {
		return "", err
	}

	response, ok := result.(*OK200)
	if !ok {
		return "", response
	}

	if response != nil && !strings.EqualFold(*response.Payload.TokenType, "bearer") {
		return "", fmt.Errorf("expected a `bearer` token type, got: %s", *response.Payload.TokenType)
	}

	return *response.Payload.AccessToken, nil
}

// cspTokenRequest creates a valid request for communicating with the CSP API.
func cspTokenRequest(address, apiToken string) (*http.Request, error) {
	if !strings.HasPrefix(address, "http://") && !strings.HasPrefix(address, "https://") {
		address = "https://" + address
	}

	currentURL, err := url.Parse(address + "/api/v2/" + "/am/api/auth/api-tokens/authorize")
	if err != nil {
		return nil, err
	}

	requestBody := &RequestBody{ApiToken: &apiToken}
	err = requestBody.validateRefreshToken()
	if err != nil {
		return nil, err
	}

	jsonBody, err := requestBody.MarshalBinary()
	if err != nil {
		return nil, err
	}

	body := io.NopCloser(bytes.NewReader(jsonBody))

	request, err := http.NewRequest("POST", currentURL.String(), body)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")

	return request, nil
}
