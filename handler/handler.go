package handler

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/hutsharing/krait/server"
)

func DocHandle(s *server.Server) http.Handler {
	ops := middleware.RedocOpts{}
	return middleware.Redoc(ops, nil)
}
