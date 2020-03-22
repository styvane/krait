package handler

import (
	"net/http"

	"github.com/hutsharing/krait/apiserver"
	"github.com/hutsharing/krait/auth"
)

func InitiateSignUpHandle(s *apiserver.Server) http.HandlerFunc {

	type response struct {
		Message  string `json:"message"`
		Code     int    `json:"code"`
		CodeName string `json:"codeName"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var req auth.SignUpRequest

		err := s.Decode(w, r, &req)
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			resp := &response{
				"Cannot process request",
				http.StatusUnprocessableEntity,
				http.StatusText(http.StatusUnprocessableEntity),
			}
			s.Encode(w, resp)
		}

		if err = req.Validate(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			resp := &response{
				"Invalid number and/or country code",
				http.StatusBadRequest,
				http.StatusText(http.StatusBadRequest),
			}
			s.Encode(w, resp)
			return
		}
		s.Encode(w, &response{"registration initiated", http.StatusOK, http.StatusText(http.StatusOK)})
	}
}
