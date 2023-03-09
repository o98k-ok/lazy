package auth

import "net/http"

type SsoManager interface {
	Login(w http.ResponseWriter, u UserEntity) error
	Logout(w http.ResponseWriter) error
	Refresh(w http.ResponseWriter, token string) error
}

type CommonManagerImp struct {
	Cookie    CookieJar
	So        Sso
	TokenName string
}

func NewManager(key string, host string) SsoManager {
	return &CommonManagerImp{
		Cookie:    NewCookieJar([]string{host}),
		So:        NewSso(key),
		TokenName: "access_token",
	}
}

func (imi *CommonManagerImp) Login(w http.ResponseWriter, u UserEntity) error {
	token, err := imi.So.GenerateToken(u)
	if err != nil {
		return err
	}
	return imi.Cookie.SetCookie(w, imi.TokenName, token)
}

func (imi *CommonManagerImp) Logout(w http.ResponseWriter) error {
	return imi.Cookie.ClearCookie(w, imi.TokenName)
}

func (imi *CommonManagerImp) Refresh(w http.ResponseWriter, token string) error {
	tk, err := imi.So.Refresh(token)
	if err != nil {
		return err
	}

	return imi.Cookie.SetCookie(w, imi.TokenName, tk)
}
