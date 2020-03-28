package resource

import (
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hutsharing/krait/handler"
	"github.com/hutsharing/krait/server"
	"github.com/stretchr/testify/assert"
)

func TestLoginRoute(t *testing.T) {
	s := server.NewTestServer()
	srv := httptest.NewServer(handler.LoginHandle(s))
	defer srv.Close()
	tests := []struct {
		payload string
		status  int
		name    string
	}{
		{
			`{"phoneNumber": "+48575557175", "countryCode": "PL", "password": "xxxxxxxx"}`,
			200,
			"Ok",
		},
		{
			`{"phoneNumber": "+7111234567", "countryCode": "RU", "password": "xxxxxxxx"}`,
			400,
			"Nok",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := strings.NewReader(test.payload)
			resp, _ := srv.Client().Post(fmt.Sprintf("%s/login", srv.URL), "application/json", r)
			assert.Equalf(
				t,
				test.status,
				test.status,
				"Expected status code: %v; found: %v",
				test.status,
				resp.StatusCode,
			)
		})
	}

}
