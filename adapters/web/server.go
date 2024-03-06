package web

import (
	"net/http"

	"github.com/edgarSucre/jw/features/auth"
	"github.com/edgarSucre/jw/features/user"

	"github.com/labstack/echo/v4"
)

type Server struct {
	userHandler    user.Handler
	authHandler    auth.Handler
	sessionManager sessionManager
}

func NewServer(userHandler user.Handler, authHandler auth.Handler) *http.Server {
	srv := &Server{
		userHandler:    userHandler,
		authHandler:    authHandler,
		sessionManager: newSessionManager(),
	}

	e := echo.New()
	srv.setRoutes(e)

	return e.Server
}
