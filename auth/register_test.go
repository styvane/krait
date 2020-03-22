package auth

import (
	"testing"
)

func TestValidRequest(t *testing.T) {
	r := Request{"+73011234567", "Ru"}
	err := r.Validate()
	if err != nil {
		t.Error("Invalid request phone number")
	}
}
