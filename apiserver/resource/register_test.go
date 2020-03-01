package resource

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hutsharing/krait/apiserver"
)

func TestRoutingToRegister(t *testing.T) {
	s := apiserver.NewServer()
	srv := httptest.NewServer(s.InitiateRegistrationHandle())

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
			r := strings.NewReader(d.payload)
			resp, _ := http.Post(fmt.Sprintf("%s/register/", srv.URL), "application/json", r)
			fmt.Println(srv.URL)
			if resp.StatusCode != d.status {
				t.Errorf("Invalid data; expected status code %v found %v", d.status, resp.StatusCode)
			}
			ct := resp.Header["Content-Type"][0]
			if ct != "application/json" && resp.StatusCode == http.StatusOK {
				t.Errorf("Expected application/json found: %v", ct)
			}

			b, _ := ioutil.ReadAll(resp.Body)
			if resp.StatusCode == http.StatusOK && !bytes.Contains(b, []byte("success")) {
				t.Error("Invalid response body")
			}
		})
	}

}
