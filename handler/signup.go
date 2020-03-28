package handler

import (
	"net/http"

	"github.com/hutsharing/krait/auth"
	"github.com/hutsharing/krait/server"
)

func InitiateSignUpHandle(s *server.Server) http.HandlerFunc {

	type response struct {
		Message  string `json:"message"`
		Code     int    `json:"code"`
		CodeName string `json:"codeName"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var req auth.SignUpRequest

		err := s.Decode(w, r, &req)
		w.Header().Set("Content-Type", "application/json")
		var resp *response
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			resp = &response{
				"Cannot process request",
				http.StatusUnprocessableEntity,
				http.StatusText(http.StatusUnprocessableEntity),
			}
		}

		if err = req.Validate(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			resp = &response{"Invalid number and/or country code",
				http.StatusBadRequest,
				http.StatusText(http.StatusBadRequest),
			}
		}
		if resp != nil {
			s.Encode(w, resp)
			return
		}
		resp = &response{
			"registration initiated",
			http.StatusOK,
			http.StatusText(http.StatusOK),
		}
		s.Encode(w, resp)
	}
}
