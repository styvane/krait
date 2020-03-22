package validation

import (
	"errors"
	"strings"

	"github.com/ttacon/libphonenumber"
)

// PhoneNumber validate the phone number
func PhoneNumber(phoneNumber, countryCode string) error {
	p, err := libphonenumber.Parse(phoneNumber, countryCode)
	if err != nil {
		return err
	}
	c := int(p.GetCountryCode())
	countryCode = strings.ToUpper(countryCode)
	if !libphonenumber.IsValidNumber(p) || libphonenumber.GetRegionCodeForCountryCode(c) != countryCode {
		return errors.New("Invalid phone number")
	}
	return nil
}
