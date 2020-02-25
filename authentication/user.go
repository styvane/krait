package authentication

import (
	"github.com/go-playground/validator"
	"github.com/nyaruka/phonenumbers"
)

// LoginRequest describes the request posted to login.
type LoginRequest struct {
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	CountryCode string `json:"countryCode" validate:"required"`
	Password    string `json:"password" validate:"required"`
}

// Validate validates LoginRequest.
// It returns nil if the struct is valid or the error if the struct is invalid
func (l *LoginRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(l)
	if err != nil {
		return err
	}
	_, err = phonenumbers.Parse(l.PhoneNumber, l.CountryCode)
	if err != nil {
		return err
	}
	return nil
}
