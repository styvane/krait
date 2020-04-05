package resource

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hutsharing/krait/handlers"
)

// Route register allow
func Route(r *mux.Router) {
	SignUpRoute(r)
	r.Handle("/docs", handlers.DocHandle())
	r.Handle("/swagger.json", http.FileServer(http.Dir("./specs/")))
}
