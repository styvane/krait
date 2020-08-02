package useraccount

import (
	"github.com/go-playground/validator/v10"
	"github.com/ttacon/libphonenumber"
)

// Passport describes required passport data.
type Passport struct {
	Number         string `json:"number" validate:"required" bson:"number"`
	DateIssue      string `json:"dateIssue" validate:"required" bson:"dateIssue"`
	DateExpiration string `json:"expirationDate validate:"required" bson:"expirationDate"`
	BirthPlace     string `json:"birthDate" validate:"required" bson:"birthPlace"`
}

// profile describes the information to create profile.
type profile struct {
	FirstName                   string `json:"firstname" validate:"required" bson:"firstName"`
	Surname                     string `json:"surname" validate:"required" bson:"surName"`
	Patronym                    string `json:"patronym,omitempty" bson:"patronym"`
	*Passport                   `json:"passport" validate:"required,dive"`
	Email                       string `json:"email" validate:"required,email"`
	*libphonenumber.PhoneNumber `json:"phoneNumber" validate:"required" bson:"phoneNumber"`
}

// New create a new profile
func New() *profile {
	p := &profile{}
	return p
}

// Validate validates Profile.
// It returns nil if the struct is valid or the error if the struct is invalid.
func (p *profile) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}
