package web

import (
	"strings"

	"github.com/edgarSucre/jw/domain"
	"github.com/labstack/echo/v4"
)

// skipLogin redirects the user to home if a session still active
func skipLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if session := c.Get(sessionKey); session != nil {
			return goToHome(c)
		}

		return next(c)
	}
}

// hxOrBustMiddleware redirect specific non hx-requests to not found page
func hxOrBustMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if isHxRequest(c) {
			return next(c)
		}

		return goNotFound(c)
	}
}

// sessionMiddleware checks for an existing session, and redirect to login if there is none
func sessionMiddleware(sm SessionManager) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			path := c.Path()

			for _, route := range skipSessionRoutes {
				if strings.HasPrefix(path, route) {
					return next(c)
				}
			}

			if err := sm.Hidratate(c); err != nil {
				return goToLogin(c)
			}

			return next(c)
		}
	}
}

// adminMiddleware checks if the active session is Admin
func adminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sessionValue := c.Get(sessionKey)
		if sessionValue == nil {
			return goNotFound(c)
		}

		session, ok := sessionValue.(domain.Session)
		if !ok {
			return goNotFound(c)
		}

		if !session.Admin {
			return goNotFound(c)
		}

		return next(c)
	}
}
