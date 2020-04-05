package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/hutsharing/krait/config"
	log "github.com/sirupsen/logrus"
)

type Validator interface {
	Validate() error
}

// Server define the Server type
type Server struct {
	Router *mux.Router
	Config *config.Config
	Log    *log.Logger
}

// create a new Server
func New(c *config.Config) *Server {
	l := log.New()
	l.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	m := mux.NewRouter()
	return &Server{Router: m, Config: c, Log: l}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

// start HTTP Server
func (s *Server) Start() {
	addr := fmt.Sprintf("%s:%d", s.Config.Host, s.Config.Port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      s,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	go func() {
		s.Log.Printf("Starting Server on port %d\n", s.Config.Port)
		err := srv.ListenAndServe()
		if err != nil {
			s.Log.Printf("Exiting Server due to %v", err)
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
