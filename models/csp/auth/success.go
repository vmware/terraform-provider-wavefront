package auth

import (
	"encoding/json"
	"fmt"
)

// Success Entity used when authentication is successful, holds the authenticated access token.
// The access token represents the authorization of a specific application to access specific parts of a user's data.
type Success struct {
	// Type of the token.
	// Example: Bearer
	// Required: true
	TokenType *string `json:"token_type,omitempty"`

	// Base64 encoded auth token.
	// The access token. This is a JWT token that grants access to resources.
	// Required: true
	AccessToken *string `json:"access_token,omitempty"`

	// The value of the Refresh token. (aka token used in request)
	// Required false
	RefreshToken *string `json:"refresh_token"`

	// Access token expiration in seconds.
	// Required false
	ExpiresIn int `json:"expires_in"`

	// An identifier for the representation of the issued security token.
	// Required false
	IssuedTokenType *string `json:"issued_token_type"`

	// The scope of access needed for the token
	// Required false
	Scope *string `json:"scope"`

	// The ID Token is a signed JWT token returned from the authorization server and
	// contains the user's profile information, including the domain of the identity
	// provider. This domain is used to obtain the identity provider URL. This token
	// is used for optimization so the application can know the identity of the user,
	// without having to make any additional network requests. This token can be
	// generated via the Authorization Code flow only.
	// Required false
	IdToken *string `json:"id_token"`
}

// Validate validates this auth response
func (m *Success) Validate() error {
	if err := m.validateToken(); err != nil {
		return err
	}
	if err := m.validateTokenType(); err != nil {
		return err
	}
	return nil
}

// validateToken that the access_token is returned in the body
func (m *Success) validateToken() error {
	if m.AccessToken == nil {
		return fmt.Errorf("AccessToken was not returned by server")
	}
	return nil
}

// validateTokenType validates the token_type is returned in the body
func (m *Success) validateTokenType() error {
	if m.TokenType == nil || *m.TokenType != "bearer" {
		return fmt.Errorf("invalid TokenType returned %v", m.TokenType)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Success) MarshalBinary() ([]byte, error) {
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
func (m *Success) UnmarshalBinary(b []byte) error {
	var res Success
	if err := json.Unmarshal(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
