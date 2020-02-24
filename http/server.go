package http

import (
	"log"
	"net/http"
)

type server struct {
	*http.Server
}

func NewServer() *server {
	return &server{&http.Server{Addr: ":8080"}}
}

func (s *server) Run() {
	log.Fatal(s.ListenAndServe())
}
