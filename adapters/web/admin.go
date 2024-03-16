package web

import (
	"net/http"

	"github.com/edgarSucre/jw/domain"
	userModel "github.com/edgarSucre/jw/features/user/model"
	"github.com/labstack/echo/v4"
)

func (server *Server) userCreate(c echo.Context) error {
	ctx := c.Request().Context()
	var params userModel.CreateParams

	if err := c.Bind(&params); err != nil {
		// TODO: handle parsing error
	}

	_, err := server.userHandler.Create(ctx, params)
	if err != nil {
		// TODO: handle creation error
	}

	return c.Redirect(http.StatusSeeOther, "/admin/users")
}

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
