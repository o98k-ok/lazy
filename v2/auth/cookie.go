package auth

import (
	"net/http"
	"time"
)

type CookieJar interface {
	SetCookie(w http.ResponseWriter, key string, value string) error
	ClearCookie(w http.ResponseWriter, key string) error
}

func NewCookieJar(domains []string) CookieJar {
	return &CookieJarImp{
		Domains: domains,
		Path:    "/",
		ExpFunc: func() time.Time {
			return time.Now().AddDate(0, 0, 1)
		},
	}
}

type CookieJarImp struct {
	Domains []string
	Path    string
	ExpFunc func() time.Time
}

func (cji *CookieJarImp) SetCookie(w http.ResponseWriter, key string, value string) error {
	exp := cji.ExpFunc()

	for _, domain := range cji.Domains {
		ck := http.Cookie{
			Name:    key,
			Value:   value,
			Path:    cji.Path,
			Domain:  domain,
			Expires: exp,
		}
		http.SetCookie(w, &ck)
	}
	return nil
}

func (cji *CookieJarImp) ClearCookie(w http.ResponseWriter, key string) error {
	for _, domain := range cji.Domains {
		ck := http.Cookie{
			Name:   key,
			Value:  "droped",
			Path:   cji.Path,
			Domain: domain,
			MaxAge: -1,
		}
		http.SetCookie(w, &ck)
	}
	return nil
}
