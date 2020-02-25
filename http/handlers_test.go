package http

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestInitiateRegistrationHandle(t *testing.T) {
	s := NewServer()
	srv := httptest.NewServer(s.InitiateRegistrationHandle())
	data := strings.NewReader(`{"phoneNumber": "+71234567890", "countryCode": "ru"}`)
	r, _ := srv.Client().Post(srv.URL, "application/json", data)
	if r.StatusCode != 200 {
		t.Error("Bad request")
	}
}
