package auth

import (
	"github.com/golang-jwt/jwt"
)

type Auther interface {
	Auth(token string) (JWTAccessClaims, error)
}

type JwtAuther struct {
	Key string
}

func NewAuther(key string) Auther {
	return &JwtAuther{
		Key: key,
	}
}

func (a *JwtAuther) Auth(token string) (JWTAccessClaims, error) {
	var tk JWTAccessClaims
	_, err := jwt.ParseWithClaims(token, &tk, func(t *jwt.Token) (interface{}, error) {
		return []byte(a.Key), nil
	})
	if err != nil {
		return JWTAccessClaims{}, err
	}

	return tk, tk.Valid()
}
