package auth

import (
	"encoding/json"
)

// ClientError Entity used when the server cannot or will not process a request,
// due to something that is perceived to be a client error. (4xx)
type ClientError struct {
	Uri      string `json:"uri"`
	CspError struct {
		ErrorCode         int      `json:"errorCode"`
		Metadata          struct{} `json:"metadata"`
		NumericModuleCode int      `json:"numericModuleCode"`
		ModuleCode        string   `json:"moduleCode"`
		ErrorMessageCode  string   `json:"errorMessageCode"`
	} `json:"cspError"`
	Causes []string `json:"causes"`
	// ClientError message
	// Example: Failed to validate credentials.
	Message   string `json:"message,omitempty"`
	RequestId string `json:"requestId"`
	// status code
	StatusCode int `json:"statusCode,omitempty"`
}

// Validate validates this auth response
func (m *ClientError) Validate() error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClientError) MarshalBinary() ([]byte, error) {
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
func (m *ClientError) UnmarshalBinary(b []byte) error {
	var res ClientError
	if err := json.Unmarshal(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
