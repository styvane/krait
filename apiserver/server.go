package apiserver

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/hutsharing/krait/config"
)

type Validator interface {
	Validate() error
}

// Server define the Server type
type Server struct {
	Router *mux.Router
	Config *config.Config
}

// create a new Server
func NewServer() *Server {
	return &Server{Router: mux.NewRouter()}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

// start HTTP Server
func (s *Server) Start() {
	srv := &http.Server{Addr: ":8080", Handler: s}
	go func() {
		log.Println("Starting Server on port 8080")
		err := srv.ListenAndServe()
		if err != nil {
			log.Println("Exiting Server...")
			os.Exit(1)
		}

	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, os.Kill)
	<-signals

	// graceful shutdown Server and allow 30 secs for ongoing operation to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	srv.Shutdown(ctx)
}

// Decode incoming request body and return an error value
func (s *Server) Decode(w http.ResponseWriter, r *http.Request, v interface{}) error {
	d := json.NewDecoder(r.Body)
	return d.Decode(v)
}

// Encode response
func (s *Server) Encode(w http.ResponseWriter, v interface{}) error {
	e := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	return e.Encode(v)
}
