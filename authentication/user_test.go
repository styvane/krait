package authentication

import (
	"testing"
)

func TestCheckValidLoginRequest(t *testing.T) {

	input := &LoginRequest{"+48123456789", "PL", "xxxxxxxxx"}
	if err := input.Validate(); err != nil {
		t.Errorf("Invalid login parameters")
	}
}
