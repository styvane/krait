package handler

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/hutsharing/krait/apiserver"
)

func DocHandle(s *apiserver.Server) http.Handler {
	ops := middleware.RedocOpts{}
	return middleware.Redoc(ops, nil)
}
