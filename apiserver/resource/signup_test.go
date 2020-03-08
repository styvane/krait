package resource

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hutsharing/krait/apiserver"
	"github.com/hutsharing/krait/apiserver/handler"
)

func TestRoutingInitiateSignUp(t *testing.T) {
	s := apiserver.NewServer()
	srv := httptest.NewServer(handler.InitiateSignUpHandle(s))
	defer srv.Close()
	data := []struct {
		payload string
		status  int
		name    string
	}{
		{
			`{"phoneNumber": "+73011234567", "countryCode":"ru"}`, http.StatusOK, "correct payload",
		},
		{`{"phoneNumber": "+61234567890", "countryCode":"ru"}`, http.StatusBadRequest, "invalid phone number"},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			reader := strings.NewReader(d.payload)

			r, _ := http.Post(fmt.Sprintf("%s/signup", srv.URL), "application/json", reader)
			if r.StatusCode != d.status {
				t.Errorf("Invalid data; expected status code %v found %v", d.status, r.StatusCode)
			}
			ct := r.Header.Get("Content-Type")
			if ct != "application/json" {
				t.Errorf("Expected application/json found: %v", ct)
			}

		})
	}

}
