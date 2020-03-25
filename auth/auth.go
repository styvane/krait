package auth

import (
	"github.com/go-playground/validator/v10"
	"github.com/hutsharing/krait/validation"
)

// LoginRequest describes the request posted to login.
// swagger:parameters Login
type LoginRequest struct {
	// in: body

	// Phone number
	// Required: true
	PhoneNumber string `json:"phoneNumber" validate:"required"`

	// Country code
	// Required: true
	CountryCode string `json:"countryCode" validate:"required"`

	// Password
	// Required: true
	Password string `json:"password" validate:"required"`
}

// Validate validates LoginRequest.
// It returns nil if the struct is valid or the error if the struct is invalid
func (l *LoginRequest) Validate() error {
	v := validator.New()
	err := v.Struct(l)
	if err != nil {
		return err
	}
	return validation.PhoneNumber(l.PhoneNumber, l.CountryCode)
}
