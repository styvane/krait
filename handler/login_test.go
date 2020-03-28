package handler

import (
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hutsharing/krait/server"
	"github.com/stretchr/testify/assert"
)

func TestLoginHandle(t *testing.T) {
	s := server.NewTestServer()
	srv := httptest.NewServer(LoginHandle(s))
	defer srv.Close()
	tests := []struct {
		payload string
		status  int
		msg     string
		name    string
	}{
		{
			`{"phoneNumber": "+48575557175", "countryCode": "PL", "password": "xxxxxxxx"}`,
			200,
			"verification code sent",
			"Ok",
		},
		{
			`{"phoneNumber": "+7111234567", "countryCode": "RU", "password": "xxxxxxxx"}`,
			400,
			"invalid payload",
			"Nok",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := strings.NewReader(test.payload)
			resp, _ := srv.Client().Post(srv.URL, "application/json", r)
			assert.Equal(
				t,
				test.status,
				resp.StatusCode,
				"Expected status code %v; found %v",
				test.status,
				resp.StatusCode,
			)

			body, err := ioutil.ReadAll(resp.Body)
			assert.Nilf(t, err, "Couldn't read response, %v", err)

			ss := strings.ToLower(string(body))
			assert.Truef(
				t,
				strings.Contains(ss, test.msg),
				"Expected %q in body found %q",
				test.msg,
				ss,
			)

		})
	}

}
