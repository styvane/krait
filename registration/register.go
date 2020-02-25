package registration

import (
	"github.com/go-playground/validator"
	"github.com/nyaruka/phonenumbers"
)

// SignUpRequest describes the request posted to initialize sign up.
type SignUpRequest struct {
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	CountryCode string `json:"countryCode" validate:"required"`
}

// Validate validates LoginRequest.
// It returns nil if the struct is valid or the error if the struct is invalid
func (s *SignUpRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(s)
	if err != nil {
		return err
	}
	_, err = phonenumbers.Parse(s.PhoneNumber, s.CountryCode)
	if err != nil {
		return err
	}
	return nil
}
