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

func NewToken(cfg *config.Config) (*Token, error) {
	c := claim{
		jwt.StandardClaims{
			Issuer: cfg.JWTIssuer,
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	ss, err := t.SignedString(cfg.JWTSecret)
	return &Token{Key: ss}, err
}

func DecodeToken(ss string, cfg *config.Config) (*jwt.Token, error) {
	return &jwt.Token{}, nil
}
