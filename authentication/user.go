package authentication

import (
	"github.com/go-playground/validator"
	"github.com/hutsharing/krait/validation"
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
	v := validator.New()
	err := v.Struct(l)
	if err != nil {
		return err
	}
	return validation.PhoneNumber(l.PhoneNumber, l.CountryCode)
}
