package http

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/hutsharing/krait/config"
)

// server define the server type
type server struct {
	router *mux.Router
	config *config.Config
}

// create a new server
func NewServer() *server {
	return &server{router: mux.NewRouter()}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// start HTTP server
func (s *server) Start() {
	srv := &http.Server{Addr: ":8080", Handler: s}
	go func() {
		log.Println("Starting server on port 8080")
		err := srv.ListenAndServe()
		if err != nil {
			log.Println("Exiting server...")
			os.Exit(1)
		}

	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, os.Kill)
	<-signals

	// graceful shutdown server and allow 30 secs for ongoing operation to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	srv.Shutdown(ctx)
}
