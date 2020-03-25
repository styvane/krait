package jwt

import (
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/hutsharing/krait/config"
)

type Token struct {
	Key string `json:"key"`
}

type claim struct {
	jwt.StandardClaims
}

func NewToken(cfg *config.JWT) (*Token, error) {
	c := claim{
		jwt.StandardClaims{
			Issuer: cfg.Issuer,
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	ss, err := t.SignedString([]byte(cfg.Secret))
	return &Token{Key: ss}, err
}

func DecodeToken(ss string, cfg *config.JWT) (*jwt.Token, error) {
	return &jwt.Token{}, nil
}

func RevokeToken() {}

func (t *Token) IsValid() bool {
	return true
}
