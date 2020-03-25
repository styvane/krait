package resource

import (
	"net/http"

	"github.com/hutsharing/krait/apiserver"
	"github.com/hutsharing/krait/apiserver/handler"
)

// swagger:route POST /signup SignUp
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
		CodeName string `json:"codeName"`
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
		CodeName string `json:"codeName"`
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
		CodeName string `json:"codeName"`
	}
}

// RegistrationRoute register registration handler
func SignUpRoute(s *apiserver.Server) {
	s.Router.HandleFunc("/signup", handler.InitiateSignUpHandle(s)).Methods(http.MethodPost)
}
