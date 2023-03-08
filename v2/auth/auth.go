package auth

import (
	"github.com/golang-jwt/jwt"
)

type Auther interface {
	Auth(token string) (JWTAccessClaims, error)
}

type JwtAuther struct {
	Key    string
	Parser *jwt.Parser
}

func NewAuther(key string) Auther {
	parse := new(jwt.Parser)
	parse.SkipClaimsValidation = true
	return &JwtAuther{
		Key:    key,
		Parser: parse,
	}
}

func (a *JwtAuther) Auth(token string) (JWTAccessClaims, error) {
	var tk JWTAccessClaims
	_, err := a.Parser.ParseWithClaims(token, &tk, func(t *jwt.Token) (interface{}, error) {
		return []byte(a.Key), nil
	})
	if err != nil {
		return JWTAccessClaims{}, err
	}

	return tk, tk.Valid()
}
