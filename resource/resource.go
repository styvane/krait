package resource

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hutsharing/krait/handlers"
)

// Init register register routes
func Init(r *mux.Router) {
	r.Handle("/docs", handlers.DocHandle())
	r.Handle("/swagger.json", http.FileServer(http.Dir("./specs/")))

	auth := r.PathPrefix("/auth/").Methods(http.MethodPost).Subrouter()
	signUpRouting(auth)
}
