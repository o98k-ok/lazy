package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	ErrExpiredAccessToken    = errors.New("token has expired")
	ErrNotExpiredAccessToken = errors.New("token is not expired")
	ErrUserDismissed         = errors.New("user has dismissed")
)

type UserEntity struct {
	UserId   string // custom user id
	UserName string
	Avatar   string
	Fid      string // lark user id
}

// JWTAccessClaims jwt claims
type JWTAccessClaims struct {
	jwt.StandardClaims
	UserId   string `json:"user_id,omitempty"`
	UserName string `json:"username,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
}

// Valid claims verification
func (a *JWTAccessClaims) Valid() error {
	if time.Unix(a.ExpiresAt, 0).Before(time.Now()) {
		return ErrExpiredAccessToken
	}
	return nil
}
