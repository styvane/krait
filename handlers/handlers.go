package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

type request struct {
	// in: body

	// Phone number
	// Required: true
	PhoneNumber string `json:"phoneNumber" validate:"required"`

	// Country code
	// Required: true
	CountryCode string `json:"countryCode" validate:"required"`
}

func SendVerificationCodeHandle() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {}
}

// swagger documentation handler
func DocHandle() http.Handler {
	ops := middleware.RedocOpts{}
	return middleware.Redoc(ops, nil)
}

// decode incoming request body and return an error value
func decode(w http.ResponseWriter, r *http.Request, v interface{}) error {
	d := json.NewDecoder(r.Body)
	defer r.Body.Close()
	return d.Decode(v)
}

// encode response
func encode(w http.ResponseWriter, v interface{}) error {
	e := json.NewEncoder(w)
	return e.Encode(v)
}
