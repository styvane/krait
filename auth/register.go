package auth

import (
	"github.com/go-playground/validator/v10"
	"github.com/hutsharing/krait/validation"
)

// SignUpRequest describes the request posted to initialize sign up.
// swagger:parameters Signup
type SignUpRequest struct {
	// in: body

	// Phone number
	// Required: true
	PhoneNumber string `json:"phoneNumber" validate:"required"`

	// Country code
	// Required: true
	CountryCode string `json:"countryCode" validate:"required"`
}

// Validate validates SignUpRequest.
// It returns nil if the struct is valid or the error if the struct is invalid
func (r *SignUpRequest) Validate() error {
	v := validator.New()
	err := v.Struct(r)
	if err != nil {
		return err
	}
	return validation.PhoneNumber(r.PhoneNumber, r.CountryCode)
}
