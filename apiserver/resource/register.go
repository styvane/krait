package resource

import (
	"net/http"

	"github.com/hutsharing/krait/apiserver"
	"github.com/hutsharing/krait/apiserver/handler"
)

// RegistrationRoute register registration handler
func RegistrationRoute(s *apiserver.Server) {
	post := s.Router.Methods(http.MethodPost).Subrouter()
	post.HandleFunc("/register/", handler.InitiateRegistrationHandle(s))
}
