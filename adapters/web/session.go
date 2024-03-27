package web

import (
	"net/http"
	"os"
	"time"

	"github.com/edgarSucre/jw/domain"
	"github.com/edgarSucre/jw/features/auth"

	"github.com/gorilla/securecookie"
	"github.com/labstack/echo/v4"
)

var skipSessionRoutes = []string{"/auth", "/static", "/favicon.ico"}

const (
	cookieName  = "jupyter-web"
	sessionKey  = "session"
	redirectKey = "redirect"

	redirectCookieErr = "cookie-error"
	redirectNoCookie  = "no-cookie"
)

type (
	SessionManager struct {
		cookieStore *securecookie.SecureCookie
	}
)

func NewSessionManager() SessionManager {
	cookieHashKey := os.Getenv("COOKIE_HASH_KEY")
	cookieBlockKey := os.Getenv("COOKIE_BLOCK_KEY")

	return SessionManager{
		cookieStore: securecookie.New([]byte(cookieHashKey), []byte(cookieBlockKey)),
	}
}

func (sm SessionManager) New(c echo.Context, user auth.User) error {
	session := domain.Session{
		Name:  user.Name,
		Admin: user.Admin,
	}

	cookieValue, err := sm.cookieStore.Encode(cookieName, session)
	if err != nil {
		return err
	}

	cookie := &http.Cookie{
		Name:     cookieName,
		Value:    cookieValue,
		Expires:  time.Now().Add(time.Hour * 3),
		Secure:   true,
		HttpOnly: true,
		Path:     "/",
	}

	c.SetCookie(cookie)

	return nil
}

func (sm SessionManager) Hidratate(c echo.Context) error {
	all := c.Request().Cookies()
	_ = all
	cookie, err := c.Cookie(cookieName)
	if err != nil {
		return err
	}

	var session domain.Session

	err = sm.cookieStore.Decode(cookieName, cookie.Value, &session)
	if err != nil {
		return err
	}

	c.Set(sessionKey, session)

	return nil
}

func (sm SessionManager) Expire(c echo.Context) {
	cookie := &http.Cookie{
		Name:     cookieName,
		Value:    "",
		MaxAge:   -1,
		HttpOnly: true,
	}

	c.Set(sessionKey, nil)
	c.SetCookie(cookie)
}
