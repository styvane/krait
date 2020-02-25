package validation

import (
	"github.com/go-playground/validator"
	"github.com/nyaruka/phonenumbers"
)

// PhoneNumber validate the phone number
func PhoneNumber(fl validator.FieldLevel) bool {
	if p, ok := fl.Field().Interface().(phonenumbers.PhoneNumber); ok {
		return phonenumbers.IsValidNumber(&p)
	}
	return false
}
