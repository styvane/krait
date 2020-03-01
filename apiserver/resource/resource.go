package resource

import (
	"github.com/hutsharing/krait/apiserver"
)

// Route register allow
func Route(s *apiserver.Server) {
	RegistrationRoute(s)
}
