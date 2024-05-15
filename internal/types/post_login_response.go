// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostLoginResponse post login response
//
// swagger:model postLoginResponse
type PostLoginResponse struct {

	// Access token required for accessing protected API endpoints
	// Example: c1247d8d-0d65-41c4-bc86-ec041d2ac437
	// Required: true
	// Format: uuid4
	AccessToken *strfmt.UUID4 `json:"access_token"`

	// Access token expiry in seconds
	// Example: 86400
	// Required: true
	ExpiresIn *int64 `json:"expires_in"`

	// Refresh token for refreshing the access token once it expires
	// Example: 1dadb3bd-50d8-485d-83a3-6111392568f0
	// Required: true
	// Format: uuid4
	RefreshToken *strfmt.UUID4 `json:"refresh_token"`

	// Type of access token, will always be `bearer`
	// Example: bearer
	// Required: true
	TokenType *string `json:"token_type"`
}

// Validate validates this post login response
func (m *PostLoginResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiresIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRefreshToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTokenType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostLoginResponse) validateAccessToken(formats strfmt.Registry) error {

	if err := validate.Required("access_token", "body", m.AccessToken); err != nil {
		return err
	}

	if err := validate.FormatOf("access_token", "body", "uuid4", m.AccessToken.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PostLoginResponse) validateExpiresIn(formats strfmt.Registry) error {

	if err := validate.Required("expires_in", "body", m.ExpiresIn); err != nil {
		return err
	}

	return nil
}

func (m *PostLoginResponse) validateRefreshToken(formats strfmt.Registry) error {

	if err := validate.Required("refresh_token", "body", m.RefreshToken); err != nil {
		return err
	}

	if err := validate.FormatOf("refresh_token", "body", "uuid4", m.RefreshToken.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PostLoginResponse) validateTokenType(formats strfmt.Registry) error {

	if err := validate.Required("token_type", "body", m.TokenType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post login response based on context it is used
func (m *PostLoginResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostLoginResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostLoginResponse) UnmarshalBinary(b []byte) error {
	var res PostLoginResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
