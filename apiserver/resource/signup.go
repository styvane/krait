package resource

import (
	"net/http"

	"github.com/hutsharing/krait/apiserver"
	"github.com/hutsharing/krait/apiserver/handler"
)

// swagger:route POST /signup Signup
// Inititate registration for new user
//
// responses:
//   200: SignUpResponse
//   400: SignUpResponse
//
// swagger:response
type SignUpResponse struct {
	// in: body
	Body struct {
		// Response content
		// Required: true
		Message string

		// The HTTP status code
		// Required: true
		Code int

		// HTTP status text
		// Required: true
		CodeName string
	}
}

// RegistrationRoute register registration handler
func SignUpRoute(s *apiserver.Server) {
	s.Router.HandleFunc("/signup", handler.InitiateSignUpHandle(s)).Methods(http.MethodPost)
}
