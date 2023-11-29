package auth

import (
	"encoding/json"
)

// ServerError Entity used when the server encounters an unexpected condition,
// preventing it from fulfilling the request. (5xx)
type ServerError struct {
	Uri      string `json:"uri"`
	CspError struct {
		ErrorCode         int      `json:"errorCode"`
		Metadata          struct{} `json:"metadata"`
		NumericModuleCode int      `json:"numericModuleCode"`
		ModuleCode        string   `json:"moduleCode"`
		ErrorMessageCode  string   `json:"errorMessageCode"`
	} `json:"cspError"`
	Causes []string `json:"causes"`
	// ServerError message
	// Example: unknown error occurred.
	Message   string `json:"message,omitempty"`
	RequestId string `json:"requestId"`
	// status code
	StatusCode int `json:"statusCode,omitempty"`
}

// Validate validates this auth response
func (m *ServerError) Validate() error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServerError) MarshalBinary() ([]byte, error) {
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
func (m *ServerError) UnmarshalBinary(b []byte) error {
	var res ServerError
	if err := json.Unmarshal(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
