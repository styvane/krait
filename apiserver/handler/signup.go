package handler

import (
	"net/http"

	"github.com/hutsharing/krait/apiserver"
	"github.com/hutsharing/krait/registration"
)

func InitiateSignUpHandle(s *apiserver.Server) http.HandlerFunc {

	type response struct {
		Message  string `json:"message"`
		Code     int    `json:"status"`
		CodeName string `json:"codeName"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var req registration.Request

		err := s.Decode(w, r, &req)

		if err = req.Validate(); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			resp := &response{"Invalid number and/or country code", http.StatusBadRequest, http.StatusText(http.StatusBadRequest)}
			s.Encode(w, resp)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		s.Encode(w, &response{"registration initiated", http.StatusOK, http.StatusText(http.StatusOK)})

	}
}
