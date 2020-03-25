package auth

import (
	"testing"
)

func TestValidSignUpRequest(t *testing.T) {
	r := SignUpRequest{"+73011234567", "Ru"}
	err := r.Validate()
	if err != nil {
		t.Error("Invalid request phone number")
	}
}
