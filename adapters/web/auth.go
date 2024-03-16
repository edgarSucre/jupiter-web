package web

import (
	"net/http"

	authm "github.com/edgarSucre/jw/features/auth/model"

	"github.com/labstack/echo/v4"
)

func (server *Server) authenticate(c echo.Context) error {
	var params authm.LoginParams

	if err := c.Bind(&params); err != nil {
		// TODO: handle binding err

		return c.NoContent(404)
	}

	user, err := server.authHandler.Authenticate(c.Request().Context(), params)

	if err != nil {
		ctx := c.Request().Context()
		component := server.authHandler.Login(ctx, err)

		return render(c, http.StatusOK, component)
	}

	server.sessionManager.new(c, user)

	return goToHome(c)
}

func (server *Server) login(c echo.Context) error {
	ctx := c.Request().Context()
	component := server.authHandler.Login(ctx, nil)

	return render(c, http.StatusOK, component)
}

func (server *Server) logout(c echo.Context) error {
	server.sessionManager.expire(c)

	return goToLogin(c)
}

func skipLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if session := c.Get(sessionKey); session != nil {
			return goToHome(c)
		}

		return next(c)
	}
}
