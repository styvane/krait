package handler

import (
	"net/http"

	"github.com/hutsharing/krait/auth"
	"github.com/hutsharing/krait/server"
)

func LoginHandle(s *server.Server) http.HandlerFunc {

	type response struct {
		Message  string `json:"message"`
		Code     int    `json:"code"`
		CodeName string `json:"codeName"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var req auth.LoginRequest

		err := s.Decode(w, r, &req)
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			resp := &response{
				"Invalid JSON structure",
				http.StatusUnprocessableEntity,
				http.StatusText(http.StatusUnprocessableEntity),
			}
			s.Encode(w, resp)
			return
		}

		if err = req.Validate(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			resp := &response{
				"Invalid payload",
				http.StatusBadRequest,
				http.StatusText(http.StatusBadRequest),
			}
			s.Encode(w, resp)
			return
		}

		resp := &response{
			"Verification code sent",
			http.StatusOK,
			http.StatusText(http.StatusOK),
		}
		s.Encode(w, resp)
	}
}
