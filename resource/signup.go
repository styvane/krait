package resource

import (
	"github.com/gorilla/mux"
	"github.com/hutsharing/krait/handlers"
)

// swagger:route POST /auth/signup SignUp
// Inititate registration for new user
//
// responses:
//   200: okResponseWrapper
//   400: badRequestResponseWrapper
//   422: unprocessableEntityResponseWrapper
//
// swagger:response
type okResponseWrapper struct {
	// in: body
	Body struct {
		// Response content
		// Required: true
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
type badRequestResponseWrapper struct {
	// in: body
	Body struct {
		// Response content
		// Required: true
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
type unprocessableEntityResponseWrapper struct {
	// in: body
	Body struct {
		// Response content
		// Required: true
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

// signUpRouting register registration handlers
func signUpRouting(r *mux.Router) {
	r.HandleFunc("/signup", handlers.InitiateSignUpHandle())
}
