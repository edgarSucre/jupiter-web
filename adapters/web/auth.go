package web

import (
	"net/http"

	"github.com/edgarSucre/jw/domain"
	"github.com/edgarSucre/jw/features/login"
	"github.com/labstack/echo/v4"
)

func setAuth(server *Server, authGroup *echo.Group) {
	authGroup.GET("/login", func(c echo.Context) error {
		return Render(c, http.StatusOK, login.Login())
	})

	authGroup.POST("/login", func(c echo.Context) error {

		// TODO: handle login - return the user

		session := domain.Session{
			Name:  "Edgar",
			Admin: true,
		}

		c.SetCookie(createNewCookie(session))

		return goToHome(c)
	}, skipLogin)

	authGroup.POST("/logout", func(c echo.Context) error {
		c.Set(sessionKey, nil)
		c.SetCookie(createExpireCookie())

		return c.Redirect(http.StatusSeeOther, "/auth/login")
	})
}

func skipLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if session := c.Get(sessionKey); session != nil {
			return goToHome(c)
		}

		return next(c)
	}
}
