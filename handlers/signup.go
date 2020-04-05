package handlers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/hutsharing/krait/server"
	"github.com/hutsharing/krait/validation"
)

// SignUpRequest describes the request posted to initialize sign up.
// swagger:parameters SignUp
type SignUpRequest struct {
	// in: body

	// Phone number
	// Required: true
	PhoneNumber string `json:"phoneNumber" validate:"required"`

	// Country code
	// Required: true
	CountryCode string `json:"countryCode" validate:"required"`
}

// Validate validates SignUpRequest.
// It returns nil if the struct is valid or the error if the struct is invalid
func (r *SignUpRequest) Validate() error {
	v := validator.New()
	err := v.Struct(r)
	if err != nil {
		return err
	}
	return validation.PhoneNumber(r.PhoneNumber, r.CountryCode)
}

func InitiateSignUpHandle(s *server.Server) http.HandlerFunc {

	type response struct {
		Message  string `json:"message"`
		Code     int    `json:"code"`
		CodeName string `json:"codeName"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var req SignUpRequest

		err := s.Decode(w, r, &req)
		w.Header().Set("Content-Type", "application/json")
		var resp *response
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			resp = &response{
				"Cannot process request",
				http.StatusUnprocessableEntity,
				http.StatusText(http.StatusUnprocessableEntity),
			}
		}

		if err = req.Validate(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			resp = &response{"Invalid number and/or country code",
				http.StatusBadRequest,
				http.StatusText(http.StatusBadRequest),
			}
		}
		if resp != nil {
			s.Encode(w, resp)
			return
		}
		resp = &response{
			"registration initiated",
			http.StatusOK,
			http.StatusText(http.StatusOK),
		}
		s.Encode(w, resp)
	}
}
