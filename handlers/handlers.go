package handlers

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/hutsharing/krait/server"
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

func SendVerificationCodeHandle(s *server.Server) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {}
}

// swagger documentation handler
func DocHandle(s *server.Server) http.Handler {
	ops := middleware.RedocOpts{}
	return middleware.Redoc(ops, nil)
}
