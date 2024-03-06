package web

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
	"github.com/edgarSucre/jw/features/layout"
	"github.com/labstack/echo/v4"
)

func (server *Server) setRoutes(e *echo.Echo) {
	// static assets
	e.Static("/static", "static")

	// dashboard
	e.GET("/", func(c echo.Context) error {
		return render(c, http.StatusOK, layout.Main())
	})

	// solo admin users
	admin := e.Group("/admin", adminMiddleware)

	// Users
	users := admin.Group("/users")
	users.GET("", componentHandler(server.userHandler.Index))
	users.GET("/list", componentHandler(server.userHandler.List))
	users.GET("/new", componentHandler(server.userHandler.New))
	users.POST("/create", server.userCreate)

	// Auth
	auth := e.Group("auth")
	//auth.GET("/login", componentHandler(server.authHandler.Login))
	auth.POST("/login", server.authenticate)
	auth.GET("/logout", server.logout)

	// Verifica la session, excepto en rutas especificas
	//e.Use(sessionMiddleware(server.sessionManager))
}

func goToLogin(c echo.Context) error {
	return c.Redirect(http.StatusSeeOther, "/auth/login")
}

func goToHome(c echo.Context) error {
	return c.Redirect(http.StatusSeeOther, "/")
}

func goNotFound(c echo.Context) error {
	return c.Redirect(http.StatusSeeOther, "/notfound")
}

func render(c echo.Context, statusCode int, t templ.Component) error {
	c.Response().Writer.WriteHeader(statusCode)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return t.Render(c.Request().Context(), c.Response().Writer)
}

func componentHandler(fn func(context.Context) templ.Component) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		return render(c, http.StatusOK, fn(ctx))
	}
}
