package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func setAdmin(server *Server, adminGroup *echo.Group) {
	userGroup := adminGroup.Group("/users")

	userGroup.GET("", func(c echo.Context) error {
		ctx := c.Request().Context()
		return Render(c, http.StatusOK, server.uh.ListUsers(ctx))
	})
	userGroup.GET("/new", func(c echo.Context) error {
		ctx := c.Request().Context()
		return Render(c, http.StatusOK, server.uh.New(ctx))
	})
}

func adminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session, err := getSessionFromContext(c)
		if err != nil {
			return goToLogin(c)
		}

		if !session.Admin {
			// TODO: handle not authorized
		}

		return next(c)
	}
}
