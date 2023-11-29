package auth

import (
	"encoding/json"
	"fmt"
)

// RequestBody contains all the parameters required to retrieve a new auth token.
type RequestBody struct {
	// Access token: https://docs.wavefront.com/using_wavefront_api.html#make-api-calls-by-using-a-user-account
	// Example: whatM9vIgu4sa2bytesTRows6NZR8QxAL3vQJ6QcGLaTRKvT0jLDogfishFJRA32
	// Required: true
	ApiToken *string `json:"api_token,omitempty"`

	// TODO: probably need to support an mfa authentication technique.
	// What is a standard way to get user input when using terraform?
	// Seems like we need to start caching the access token so they don't have to repeat mfa.
	// I am not building with this in mind currently. Trying to stay simple and not over build.
	// Required true if using an MFA device.
	// Passcode *string `json:"passcode,omitempty"`
}

// Validate validates this csp login specification
func (m *RequestBody) Validate() error {
	if err := m.validateRefreshToken(); err != nil {
		return err
	}
	return nil
}

func (m *RequestBody) validateRefreshToken() error {
	if m.ApiToken == nil {
		return fmt.Errorf("ApiToken required for CSP authorizaiton request")
	}
	return nil
}

// MarshalBinary interface implementation
func (m *RequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	body, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// UnmarshalBinary interface implementation
func (m *RequestBody) UnmarshalBinary(b []byte) error {
	var res RequestBody
	if err := json.Unmarshal(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
