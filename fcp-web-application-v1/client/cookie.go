package client

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
)

func GetClientWithCookie(token string, cookies ...*http.Cookie) (*http.Client, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	cookies = append(cookies, &http.Cookie{
		Name:  "session_token",
		Value: token,
	})

	var cookieHost string = os.Getenv("RAILWAY_STATIC_URL")
	cookiesSchame := "https"
	if cookieHost == "" {
		cookieHost = "http"
		cookieHost = "localhost"
	}

	// jar.SetCookies(&url.URL{
	// 	Scheme: "http",
	// 	Host:   "localhost:8080",
	// }, cookies)

	jar.SetCookies(&url.URL{
		Scheme: cookiesSchame,
		Host:   cookieHost,
	}, cookies)

	c := &http.Client{
		Jar: jar,
	}

	return c, nil
}
