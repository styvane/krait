package validation

import (
	"errors"
	"strings"

	"github.com/nyaruka/phonenumbers"
)

// PhoneNumber validate the phone number
func PhoneNumber(phoneNumber, countryCode string) error {
	p, err := phonenumbers.Parse(phoneNumber, countryCode)
	if err != nil {
		return err
	}
	c := int(p.GetCountryCode())
	countryCode = strings.ToUpper(countryCode)
	if !phonenumbers.IsValidNumber(p) || phonenumbers.GetRegionCodeForCountryCode(c) != countryCode {
		return errors.New("Invalid phone number")
	}
	return nil
}
