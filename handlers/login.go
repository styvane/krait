package handlers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/hutsharing/krait/validation"
)

// LoginRequest describes the request posted to login.
// swagger:parameters Login
type LoginRequest struct {
	// in: body

	// Phone number
	// Required: true
	PhoneNumber string `json:"phoneNumber" validate:"required"`

	// Country code
	// Required: true
	CountryCode string `json:"countryCode" validate:"required"`

	// Password
	// Required: true
	Password string `json:"password" validate:"required"`
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

func LoginHandle() http.HandlerFunc {

	type response struct {
		Message  string `json:"message"`
		Code     int    `json:"code"`
		CodeName string `json:"codeName"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var req LoginRequest

		err := decode(w, r, &req)
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			resp := &response{
				"Invalid JSON structure",
				http.StatusUnprocessableEntity,
				http.StatusText(http.StatusUnprocessableEntity),
			}
			encode(w, resp)
			return
		}

		if err = req.Validate(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			resp := &response{
				"Invalid payload",
				http.StatusBadRequest,
				http.StatusText(http.StatusBadRequest),
			}
			encode(w, resp)
			return
		}

		resp := &response{
			"Verification code sent",
			http.StatusOK,
			http.StatusText(http.StatusOK),
		}
		encode(w, resp)
	}
}
