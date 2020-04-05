package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidSignUpRequest(t *testing.T) {
	tests := []*SignUpRequest{
		{"+73011234567", "Ru"},
	}

	for _, test := range tests {
		err := test.Validate()
		assert.Nilf(t, err, "Invalid request phone number", err)
	}
}
