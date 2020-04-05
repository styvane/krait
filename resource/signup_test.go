package resource

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hutsharing/krait/handlers"
	"github.com/stretchr/testify/assert"
)

func TestRoutingInitiateSignUp(t *testing.T) {

	srv := httptest.NewServer(handlers.InitiateSignUpHandle())
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
			reader := strings.NewReader(test.payload)

			r, _ := http.Post(fmt.Sprintf("%s/signup", srv.URL), "application/json", reader)
			assert.Equalf(
				t,
				test.status,
				r.StatusCode,
				"Invalid data; expected status code: %v found: %v",
				test.status,
				r.StatusCode,
			)

			ct := r.Header.Get("Content-Type")
			assert.Equalf(t, "application/json", ct, "Expected application/json found: %v", ct)

		})
	}

}
