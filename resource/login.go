package resource

import (
	"net/http"

	"github.com/hutsharing/krait/handlers"
	"github.com/hutsharing/krait/server"
)

// swagger:route POST /login Login
// Login user
//
// responses:
//   200: okloginResponseWrapper
//   400: badLoginResponseWrapper
//   422: unprocessableLoginResponseWrapper
//
// swagger:response
type okloginResponseWrapper struct {
	// in: body
	Body struct {
		// Response content
		// Required: true
		// Example: Verification code sent
		Message string `json:"message"`

		// The HTTP status code
		// Required: true
		// Example: 200
		Code int `json:"code"`

		// HTTP status text
		// Required: true
		// Example: OK
		CodeName string `json:"codeName"`
	}
}

// swagger:response
type badLoginResponseWrapper struct {
	// in: body
	Body struct {
		// Response content
		// Required: true
		// Example: Invalid Data type
		Message string `json:"message"`

		// The HTTP status code
		// Required: true
		// Example: 400
		Code int `json:"code"`

		// HTTP status text
		// Required: true
		// Example: Bad Request
		CodeName string `json:"codeName"`
	}
}

// swagger:response
type unprocessableLoginResponseWrapper struct {
	// in: body
	Body struct {
		// Response content
		// Required: true
		// Example: Invalid payload
		Message string `json:"message"`
		// The HTTP status code
		// Required: true
		// Example: 422
		Code int `json:"code"`

		// HTTP status text
		// Required: true
		// Example: Unprocessable Entity
		CodeName string `json:"codeName"`
	}
}

// LoginRoute register authentication handlers
func LoginRoute(s *server.Server) {
	s.Router.HandleFunc("/login", handlers.LoginHandle(s)).Methods(http.MethodPost)
}
