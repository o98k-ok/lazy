package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type Sso interface {
	Auther
	GenerateToken(user UserEntity) (string, error)
	Refresh(token string) (string, error)
	UserValid(user string) bool
}

type SsoImp struct {
	Auther
	AuthKey      string
	Expire       time.Duration
	UserValidFnc func(user string) bool
}

func NewSso(key string) Sso {
	return &SsoImp{
		Auther:  NewAuther(key),
		AuthKey: key,
		Expire:  time.Hour * 2,
	}
}

func (s *SsoImp) Auth(token string) (JWTAccessClaims, error) {
	return s.Auther.Auth(token)
}

func (s *SsoImp) GenerateToken(user UserEntity) (string, error) {
	if !s.UserValid(user.Fid) {
		return "", ErrUserDismissed
	}

	claims := JWTAccessClaims{
		StandardClaims: jwt.StandardClaims{
			Audience:  user.UserId,
			ExpiresAt: time.Now().Add(s.Expire).Unix(),
			Id:        user.Fid,
			Subject:   user.Fid,
		},
		UserId:   user.UserId,
		UserName: user.UserName,
		Avatar:   user.Avatar,
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, &claims).SignedString([]byte(s.AuthKey))
}

func (s *SsoImp) Refresh(token string) (string, error) {
	user, err := s.Auth(token)
	if !errors.Is(err, ErrExpiredAccessToken) {
		return "", ErrNotExpiredAccessToken
	}

	if !s.UserValid(user.Id) {
		return "", ErrUserDismissed
	}
	user.ExpiresAt = time.Now().Add(s.Expire).Unix()
	return jwt.NewWithClaims(jwt.SigningMethodHS256, &user).SignedString([]byte(s.AuthKey))
}

func (s *SsoImp) UserValid(user string) bool {
	if s.UserValidFnc != nil {
		return s.UserValidFnc(user)
	}
	return true
}

func (s *SsoImp) SetValidHandler(arg func(string) bool) {
	s.UserValidFnc = arg
}
