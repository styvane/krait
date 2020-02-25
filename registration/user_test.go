package registration

import (
	"testing"

	"github.com/nyaruka/phonenumbers"
)

func TestCreateNewUserProfile(t *testing.T) {
	_, err := NewProfile()
	if err != nil {
		t.Error("Cannot create profile")
	}
}

func TestCheckValidProfile(t *testing.T) {
	passport := &Passport{"s555051", "2020-01-01", "2020-01-01", "Russia"}
	p := &Profile{
		"Tom",
		"Jerry",
		"",
		"1918-01-01",
		"xxx@example.com",
		passport,
		phonenumbers.GetExampleNumber("RU"),
	}
	if err := p.Validate(); err != nil {
		t.Error("Invalid profile information")
	}
}
