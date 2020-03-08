package resource

import (
	"net/http"

	"github.com/hutsharing/krait/apiserver"
	"github.com/hutsharing/krait/apiserver/handler"
)

// Route register allow
func Route(s *apiserver.Server) {
	SignUpRoute(s)
	s.Router.Handle("/docs", handler.DocHandle(s))
	s.Router.Handle("/swagger.json", http.FileServer(http.Dir("./specs/")))
}
