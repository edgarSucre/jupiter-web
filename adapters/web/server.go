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
	sessionManager SessionManager
}

func NewServer(userHandler user.Handler, authHandler auth.Handler) *http.Server {
	srv := &Server{
		authHandler:    authHandler,
		userHandler:    userHandler,
		sessionManager: NewSessionManager(),
	}

	e := echo.New()
	srv.setRoutes(e)

	return e.Server
}
