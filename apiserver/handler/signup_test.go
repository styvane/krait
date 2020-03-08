package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hutsharing/krait/apiserver"
)

func TestInitiateRegistrationHandle(t *testing.T) {
	s := apiserver.NewServer()
	srv := httptest.NewServer(InitiateSignUpHandle(s))
	defer srv.Close()

	data := []struct {
		payload string
		status  int
		name    string
	}{
		{
			`{"phoneNumber": "+73011234567", "countryCode":"ru"}`, http.StatusOK, "correct payload"},
		{`{"phoneNumber": "+61234567890", "countryCode":"ru"}`, http.StatusBadRequest, "invalid phone number"},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			r := strings.NewReader(d.payload)
			resp, _ := srv.Client().Post(srv.URL, "application/json", r)
			if resp.StatusCode != d.status {
				t.Errorf("Invalid data; expected status code %v found %v", d.status, resp.StatusCode)
			}

		})
	}
}
