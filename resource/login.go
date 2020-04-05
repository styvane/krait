package resource

import (
	"github.com/gorilla/mux"
	"github.com/hutsharing/krait/handlers"
)

// swagger:route POST /auth/login Login
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
		Reason string `json:"reason"`
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
		Reason string `json:"reason"`
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
		Reason string `json:"reason"`
	}
}

// loginRouting register authentication handlers
func loginRoute(r *mux.Router) {
	r.HandleFunc("/login", handlers.LoginHandle())
}
