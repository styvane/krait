package resource

import (
	"net/http"

	"github.com/hutsharing/krait/handler"
	"github.com/hutsharing/krait/server"
)

// Route register allow
func Route(s *server.Server) {
	SignUpRoute(s)
	s.Router.Handle("/docs", handler.DocHandle(s))
	s.Router.Handle("/swagger.json", http.FileServer(http.Dir("./specs/")))
}
