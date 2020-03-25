package jwt

import (
	"testing"

	"github.com/hutsharing/krait/config"
)

func TestNewToken(t *testing.T) {
	c := &config.JWT{}

	_, err := NewToken(c)
	if err != nil {
		t.Errorf("Unable to generate token: %v", err)
	}

}

func TestDecodeToken(t *testing.T) {
	var input string
	c := &config.JWT{}
	_, err := DecodeToken(input, c)
	if err != nil {
		t.Errorf("Unable to decode token: %v", err)
	}
}

func TestRevokeToken(t *testing.T) {

	RevokeToken()

}

func TestIsValidToken(t *testing.T) {
	tok := &Token{}
	tok.IsValid()
}
