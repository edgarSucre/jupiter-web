package web

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
	"github.com/edgarSucre/jw/features/layout"
	"github.com/labstack/echo/v4"
)

const hxRequestHeaderName = "Hx-Request"

func (server *Server) setRoutes(e *echo.Echo) {
	// static assets
	e.Static("/static", "static")

	// dashboard
	e.GET("/", func(c echo.Context) error {
		return render(c, http.StatusOK, layout.Main())
	})

	// solo admin users
	//admin := e.Group("/admin", adminMiddleware)
	admin := e.Group("/admin")

	// Users
	userHandler := server.userHandler
	users := admin.Group("/users")
	users.GET("", hxComponentHandler(userHandler.IndexFull, userHandler.IndexCmp))
	users.GET("/list", componentHandler(userHandler.List))
	users.GET("/new", componentHandler(userHandler.New))
	users.POST("/create", server.userCreate)

	// Auth
	auth := e.Group("auth")
	auth.GET("/login", server.login)
	auth.POST("/login", server.authenticate)
	auth.GET("/logout", server.logout)

	// Verifica la session, excepto en rutas especificas
	e.Use(sessionMiddleware(server.sessionManager))

	e.GET("/notfound", func(c echo.Context) error {
		return c.String(200, "no se encontro lo que buscas")
	})
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

type templHandler func(context.Context) templ.Component

func componentHandler(fn templHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		return render(c, http.StatusOK, fn(ctx))
	}
}

// hxComponentHandler render fullpage version of partial if not an hx-request
func hxComponentHandler(
	fullPageHandler templHandler,
	partialHandler templHandler,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		hxRequestHeaderValue := c.Request().Header.Get(hxRequestHeaderName)
		if hxRequestHeaderValue == "true" {
			return render(c, http.StatusOK, partialHandler(ctx))
		}

		return render(c, http.StatusOK, fullPageHandler(ctx))
	}
}

// prevents non hx-requests to be handle
func hxOrBustMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		hxRequestHeaderValue := c.Request().Header.Get(hxRequestHeaderName)
		if hxRequestHeaderValue == "true" {
			return next(c)
		}

		return goNotFound(c)
	}
}
