package web

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/edgarSucre/jw/domain"
	"github.com/gorilla/securecookie"
	"github.com/labstack/echo/v4"
)

type ()

var (
	cookieHashKey  = os.Getenv("COOKIE_HASH_KEY")
	cookieBlockKey = os.Getenv("COOKIE_BLOCK_KEY")
	cookieManager  = securecookie.New([]byte(cookieHashKey), []byte(cookieBlockKey))
)

const (
	cookieName  = "jupyter-web"
	sessionKey  = "session"
	redirectKey = "redirect"

	redirectCookieErr = "cookie-error"
	redirectNoCookie  = "no-cookie"
)

func createNewCookie(session domain.Session) *http.Cookie {
	encoded, err := cookieManager.Encode(cookieName, session)
	if err != nil {
		// TODO: handle encoding error
		return nil
	}

	cookie := &http.Cookie{
		Name:     cookieName,
		Value:    encoded,
		Expires:  time.Now().Add(time.Hour * 3),
		Secure:   true,
		HttpOnly: true,
	}

	return cookie
}

func createExpireCookie() *http.Cookie {
	cookie := &http.Cookie{
		Name:     cookieName,
		Value:    "",
		MaxAge:   -1,
		HttpOnly: true,
	}

	return cookie
}

func getSessionFromContext(c echo.Context) (domain.Session, error) {
	sessionValue := c.Get(sessionKey)
	if sessionValue == nil {
		return domain.Session{}, errors.New("asdasd")
	}

	if session, ok := sessionValue.(domain.Session); ok {
		return session, nil
	}

	return domain.Session{}, errors.New("asdasd")

}

func sessionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if strings.HasPrefix(c.Request().URL.Path, "/auth") {
			return next(c)
		}

		cookie, err := c.Cookie(cookieName)
		if err != nil {
			c.Set(redirectKey, redirectNoCookie)

			return goToLogin(c)
		}

		var session domain.Session
		if err = cookieManager.Decode(cookieName, cookie.Value, &session); err != nil {
			c.Set(redirectKey, redirectCookieErr)

			return goToLogin(c)
		}

		c.Set(sessionKey, session)

		return next(c)
	}
}
