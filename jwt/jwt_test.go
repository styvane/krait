package jwt

import (
	"testing"

	"github.com/hutsharing/krait/config"
)

func TestGenerateToken(t *testing.T) {
	cfg := &config.Config{}

	_, err := GenerateToken(cfg)
	if err != nil {
		t.Errorf("Unable to generate token: %v", err)
	}

}

func TestDecodeToken(t *testing.T) {
	var input string
	cfg := &config.Config{}
	_, err := DecodeToken(input, cfg)
	if err != nil {
		t.Errorf("Unable to decode token: %v", err)
	}
}

func TestRevokeToken(t *testing.T) {
	RevokeToken()

}

func TestIsValidToken(t *testing.T) {
	IsValidToken
}
