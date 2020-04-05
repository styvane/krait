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

func TestRoutingInitiateSignUp(t *testing.T) {
	r := mux.NewRouter()
	signUpRouting(r)
	srv := httptest.NewServer(r)
	defer srv.Close()
	tests := []struct {
		payload string
		body    string
		name    string
	}{
		{
			`{"phoneNumber": "+73011234567", "countryCode":"ru"}`,
			"verification",
			"correct payload",
		},
		{
			`{"phoneNumber": "+61234567890", "countryCode":"ru"}`,
			"invalid",
			"invalid phone number",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			reader := strings.NewReader(test.payload)

			resp, _ := srv.Client().Post(fmt.Sprintf("%s/signup", srv.URL), "application/json", reader)

			body, err := ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()

			assert.Nilf(t, err, "could not read response: %v", err)

			ss := strings.ToLower(string(body))

			assert.Containsf(
				t,
				ss,
				test.body,
				"Invalid data; expected %v in body; found: %v",
				test.body,
				ss,
			)

			ct := resp.Header.Get("Content-Type")
			assert.Equalf(t, "application/json", ct, "Expected application/json found: %v", ct)

		})
	}

}
