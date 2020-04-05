package jwt

import (
	"testing"

	"github.com/hutsharing/krait/config"
	"github.com/stretchr/testify/assert"
)

func TestNewToken(t *testing.T) {
	c := &config.JWT{}

	_, err := NewToken(c)
	assert.Nilf(t, err, "Unable to generate token: %v", err)
}

func TestDecodeToken(t *testing.T) {
	var input string
	c := &config.JWT{}
	_, err := DecodeToken(input, c)
	assert.Nilf(t, err, "Unable to decode token: %v", err)
}

func TestRevokeToken(t *testing.T) {
	RevokeToken()
}

func TestIsValidToken(t *testing.T) {
	tok := &Token{}
	tok.IsValid()
}
