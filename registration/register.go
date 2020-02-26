package registration

import (
	"github.com/go-playground/validator"
	"github.com/hutsharing/krait/validation"
)

// Request describes the request posted to initialize sign up.
type Request struct {
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	CountryCode string `json:"countryCode" validate:"required"`
}

// Validate validates Request.
// It returns nil if the struct is valid or the error if the struct is invalid
func (s *Request) Validate() error {
	v := validator.New()
	err := v.Struct(s)
	if err != nil {
		return err
	}
	return validation.PhoneNumber(s.PhoneNumber, s.CountryCode)
}
