package web

import (
	"encoding/json"
	"net/http"

	"github.com/a-h/templ"
	"github.com/edgarSucre/jw/features/components"
	"github.com/edgarSucre/jw/features/layout"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	hxLocationHeaderName = "HX-Location"
	hxRequestHeaderName  = "Hx-Request"
	hxRedirectHeaderName = "HX-Redirect"
)

func (server *Server) setRoutes(e *echo.Echo) {
	// static assets
	e.Static("/static", "static")
	e.File("/favicon.ico", "static/images/favicon.ico")

	// dashboard
	e.GET("/", func(c echo.Context) error {
		if isHxRequest(c) {
			return render(c, http.StatusOK, layout.Main())
		}

		return render(c, http.StatusOK, layout.Index())
	})

	// solo admin users
	//admin := e.Group("/admin", adminMiddleware)
	admin := e.Group("/admin")

	// Users
	userHandler := server.userHandler
	users := admin.Group("/users")
	users.GET("", userHandler.Index)
	users.POST("", userHandler.Create, hxOrBustMiddleware)
	users.GET("/list", userHandler.List, hxOrBustMiddleware)
	users.GET("/new", userHandler.New)
	users.DELETE("/:id", userHandler.Delete)

	// Auth
	auth := e.Group("auth")
	auth.GET("/login", server.authHandler.Login)
	auth.POST("/login", server.authHandler.Authenticate)
	auth.GET("/logout", server.authHandler.Logout)

	e.GET("/notfound", func(c echo.Context) error {
		title := "Lo sentimos, no se puede encontrar la pagina"
		msg := "La pagina que busca pudo haber sido movida, borrada o no existe"

		if isHxRequest(c) {
			return render(c, 200, components.Error(title, msg))
		}

		return render(c, 200, components.ErrorReload(title, msg))
	})

	e.GET("*", func(c echo.Context) error {
		return goNotFound(c)
	})

	// Verifica la session, excepto en rutas especificas
	//e.Use(sessionMiddleware(server.sessionManager))
	e.Use(middleware.Logger())
}

type (
	Navigator struct {
	}

	Redirector interface {
		Redirect(code int, url string) error
	}
)

func (nav Navigator) Home(c echo.Context) error {
	return goToHome(c)
}

func (nav Navigator) Login(c echo.Context) error {
	return goToLogin(c)
}

func (nav Navigator) NotFound(c echo.Context) error {
	return goNotFound(c)
}

func (nav Navigator) IsHxRequest(c echo.Context) bool {
	return isHxRequest(c)
}

func goToLogin(c echo.Context) error {
	if isHxRequest(c) {
		location := locationString(Location{Path: "/auth/login", Target: "body"})
		c.Response().Header().Set(hxLocationHeaderName, location)

		return c.NoContent(http.StatusOK)
	}

	return c.Redirect(http.StatusSeeOther, "/auth/login")
}

func goToHome(c echo.Context) error {
	if isHxRequest(c) {
		// location := locationString(Location{Path: "/", Target: "#mainSection"})
		c.Response().Header().Set(hxRedirectHeaderName, "/")
		return c.NoContent(http.StatusOK)
	}

	return c.Redirect(http.StatusSeeOther, "/")
}

func goNotFound(c echo.Context) error {
	if isHxRequest(c) {
		location := locationString(Location{Path: "/notfound", Target: "#mainSection"})
		c.Response().Header().Set(hxLocationHeaderName, location)

		return c.NoContent(http.StatusNotFound)
	}

	return c.Redirect(http.StatusSeeOther, "/notfound")
}

func isHxRequest(c echo.Context) bool {
	val := c.Request().Header.Get(hxRequestHeaderName)

	return val == "true"
}

type ViewRenderer struct{}

func (ViewRenderer) Render(c echo.Context, code int, view templ.Component) error {
	return render(c, code, view)
}

func render(c echo.Context, code int, view templ.Component) error {
	c.Response().Writer.WriteHeader(code)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return view.Render(c.Request().Context(), c.Response().Writer)
}

type Location struct {
	Path   string `json:"path"`
	Target string `json:"target"`
}

func locationString(l Location) string {
	ed, _ := json.Marshal(l)
	return string(ed)
}
