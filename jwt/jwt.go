package jwt

import (
	"github.com/dgrijalva/jwt-go/v4"
)

type Token struct {
	Key string `json:"key" bson:"token"`
}

type claim struct {
	jwt.StandardClaims
}

func NewToken(cfg *config) (*Token, error) {
	c := claim{
		jwt.StandardClaims{
			Issuer: cfg.Issuer,
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	ss, err := t.SignedString([]byte(cfg.Secret))
	return &Token{Key: ss}, err
}

func DecodeToken(ss string, cfg *config) (*jwt.Token, error) {
	return &jwt.Token{}, nil
}

func RevokeToken() {}

func (t *Token) IsValid() bool {
	return true
}
