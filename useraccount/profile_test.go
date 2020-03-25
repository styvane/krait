package useraccount

import (
	"testing"

	"github.com/ttacon/libphonenumber"
)

func TestCheckValidProfile(t *testing.T) {
	passport := &Passport{"s555051", "2020-01-01", "2020-01-01", "Russia"}
	p := &Profile{
		"Tom",
		"Jerry",
		"",
		"1918-01-01",
		"xxx@example.com",
		passport,
		libphonenumber.GetExampleNumber("RU"),
	}
	if err := p.Validate(); err != nil {
		t.Error("Invalid profile information")
	}
}
