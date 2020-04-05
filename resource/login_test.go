package resource

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestLoginRoute(t *testing.T) {
	r := mux.NewRouter()
	loginRoute(r)
	srv := httptest.NewServer(r)
	defer srv.Close()
	tests := []struct {
		payload string
		body    string
		name    string
	}{
		{
			`{"phoneNumber": "+48575557175", "countryCode": "PL", "password": "xxxxxxxx"}`,
			"verification",
			"Ok",
		},
		{
			`{"phoneNumber": "+7111234567", "countryCode": "RU", "password": "xxxxxxxx"}`,
			"invalid",
			"Nok",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := strings.NewReader(test.payload)

			resp, _ := srv.Client().Post(fmt.Sprintf("%s/login", srv.URL), "application/json", r)
			body, err := ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()

			assert.Nilf(t, err, "could not read response: %v", err)

			ss := strings.ToLower(string(body))
			assert.Containsf(
				t,
				ss,
				test.body,
				"Expected %v in body; found: %v",
				test.body,
				ss,
			)
		})
	}

}
