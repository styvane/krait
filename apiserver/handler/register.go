package handler

import (
	"net/http"

	"github.com/hutsharing/krait/apiserver"
	"github.com/hutsharing/krait/registration"
)

func InitiateRegistrationHandle(s *apiserver.Server) http.HandlerFunc {

	type response struct {
		Message string `json:"message"`
		Status  int    `json:"status"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var sr registration.Request

		err := s.Decode(w, r, &sr)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		if err = sr.Validate(); err != nil {
			http.Error(w, "Invalid number and/or country code", http.StatusBadRequest)
			return
		}

		resp := &response{"success", http.StatusOK}
		s.Encode(w, resp)
	}
}
