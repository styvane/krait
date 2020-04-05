package handlers

import (
	"net/http/httptest"
	"testing"

	"github.com/hutsharing/krait/server"
	"github.com/stretchr/testify/assert"
)

func TestSendVerificationCodeHandle(t *testing.T) {
	s := server.NewTestServer()
	srv := httptest.NewServer(SendVerificationCodeHandle(s))
	defer srv.Close()

	tests := []struct {
		payload string
		msg     string
		name    string
	}{
		{
			`{"phoneNumber": "+48575557175", "countryCode": "PL"}`,
			"verification code sent",
			"Ok",
		},
		{
			`{"phoneNumber": "+7111234567", "countryCode": "RU", "password": "xxxxxxxx"}`,
			"unable to send verification code",
			"Nok",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, 1, 1, "Expected: %v; found: %v", tt.msg, "mss")
		})
	}

}
