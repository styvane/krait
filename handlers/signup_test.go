package handlers

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hutsharing/krait/server"
	"github.com/stretchr/testify/assert"
)

func TestValidSignUpRequest(t *testing.T) {
	tests := []*SignUpRequest{
		{"+73011234567", "Ru"},
	}

	for _, test := range tests {
		err := test.Validate()
		assert.Nilf(t, err, "Invalid request phone number", err)
	}
}

func TestInitiateRegistrationHandle(t *testing.T) {
	s := server.NewTestServer()
	srv := httptest.NewServer(InitiateSignUpHandle(s))
	defer srv.Close()

	tests := []struct {
		payload string
		status  int
		name    string
	}{
		{
			`{"phoneNumber": "+73011234567", "countryCode":"ru"}`,
			200,
			"correct payload",
		},
		{
			`{"phoneNumber": "+61234567890", "countryCode":"ru"}`,
			400,
			"invalid phone number",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := strings.NewReader(test.payload)
			resp, _ := srv.Client().Post(srv.URL, "application/json", r)
			assert.Equalf(
				t,
				test.status,
				resp.StatusCode,
				"Invalid data; expected status code: %v found: %v",
				test.status,
				resp.StatusCode,
			)
		})
	}
}
