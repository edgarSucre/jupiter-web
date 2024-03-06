package web

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/edgarSucre/jw/domain"
	authm "github.com/edgarSucre/jw/features/auth/model"
	"github.com/gorilla/securecookie"
	"github.com/labstack/echo/v4"
)

const (
	cookieName  = "jupyter-web"
	sessionKey  = "session"
	redirectKey = "redirect"

	redirectCookieErr = "cookie-error"
	redirectNoCookie  = "no-cookie"
)

type (
	sessionManager struct {
		cookieStore *securecookie.SecureCookie
	}
)

func newSessionManager() sessionManager {
	cookieHashKey := os.Getenv("COOKIE_HASH_KEY")
	cookieBlockKey := os.Getenv("COOKIE_BLOCK_KEY")

	return sessionManager{
		cookieStore: securecookie.New([]byte(cookieHashKey), []byte(cookieBlockKey)),
	}
}

func (sm sessionManager) new(c echo.Context, user authm.User) error {
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

func (sm sessionManager) hidratate(c echo.Context) error {
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

func (sm sessionManager) expire(c echo.Context) {
	cookie := &http.Cookie{
		Name:     cookieName,
		Value:    "",
		MaxAge:   -1,
		HttpOnly: true,
	}

	c.Set(sessionKey, nil)
	c.SetCookie(cookie)
}

var skipSessionRoutes = []string{"/auth", "/static"}

func sessionMiddleware(sm sessionManager) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			path := c.Path()

			for _, route := range skipSessionRoutes {
				if strings.HasPrefix(path, route) {
					return next(c)
				}
			}

			if err := sm.hidratate(c); err != nil {
				return goToLogin(c)
			}

			return next(c)
		}
	}
}
