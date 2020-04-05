package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckValidLoginRequest(t *testing.T) {

	tests := []*LoginRequest{
		{"+48123456789", "PL", "xxxxxxxxx"},
	}
	for _, test := range tests {
		err := test.Validate()
		assert.Nilf(t, err, "Invalid login parameters %v", err)
	}
}
