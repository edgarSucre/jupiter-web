package web

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/edgarSucre/jw/features/layout"
	"github.com/edgarSucre/jw/features/user"

	"github.com/labstack/echo/v4"
)

type Server struct {
	uh user.Handler
}

func NewServer(uh user.Handler) *http.Server {
	srv := &Server{
		uh: uh,
	}

	e := echo.New()
	srv.setRoutes(e)

	return e.Server
}

type tokenRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	Code         string `json:"code"`
	RedirectUri  string `json:"redirect_uri"`
}

type tokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

func (S *Server) setRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return Render(c, http.StatusOK, layout.Main())
	})

	setAdmin(S, e.Group("/admin", adminMiddleware))
	setAuth(S, e.Group("/auth"))

	e.Use(sessionMiddleware)

}

func Render(c echo.Context, statusCode int, t templ.Component) error {
	c.Response().Writer.WriteHeader(statusCode)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(c.Request().Context(), c.Response().Writer)
}

func goToLogin(c echo.Context) error {
	return c.Redirect(http.StatusSeeOther, "/auth/login")
}

func goToHome(c echo.Context) error {
	return c.Redirect(http.StatusSeeOther, "/")
}
