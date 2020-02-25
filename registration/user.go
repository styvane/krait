package registration

import (
	"github.com/go-playground/validator"
	"github.com/nyaruka/phonenumbers"
)

// Passport describes required passport data.
type Passport struct {
	Number         string `json:"number" validate:"required"`
	DateIssue      string `json:"dateIssue" validate:"required"`
	DateExpiration string `json:"expirationDate validate:"required"`
	BirthPlace     string `json:"birthDate" validate:"required"`
}

// Profile describes the information to create profile.
type Profile struct {
	FirstName   string                    `json:"firstname" validate:"required"`
	Surname     string                    `json:"surname" validate:"required"`
	Patronym    string                    `json:"patronym,omitempty"`
	BirthDate   string                    `json:"birthDate" validate:"required"`
	Email       string                    `json:"email" validate:"required,email"`
	Passort     *Passport                 `json:"passport" validate:"required,dive"`
	PhoneNumber *phonenumbers.PhoneNumber `json:"phoneNumber" validate:"required"`
}

// Validate validates Profile.
// It returns nil if the struct is valid or the error if the struct is invalid.
func (p *Profile) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}

// NewProfile creates and return new Profile
func NewProfile() (*Profile, error) {
	p := &Profile{}
	return p, nil
}
