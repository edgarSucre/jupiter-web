package auth

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
	"github.com/edgarSucre/jw/domain"
	"github.com/edgarSucre/jw/features/auth/view"
	"github.com/labstack/echo/v4"
)

type (
	navigator interface {
		Home(echo.Context) error
	}

	useCase interface {
		Login(context.Context, LoginParams) (User, error)
	}

	sessionManager interface {
		New(echo.Context, User) error
		Expire(echo.Context)
	}

	renderer interface {
		Render(echo.Context, int, templ.Component) error
	}
)

type Handler struct {
	navigator navigator
	renderer
	sessionManager sessionManager
	useCase        useCase
}

func NewHandler(nav navigator, renderer renderer, sm sessionManager, uc useCase) *Handler {
	return &Handler{
		navigator:      nav,
		renderer:       renderer,
		sessionManager: sm,
		useCase:        uc,
	}
}

// POST /auth/login
func (h *Handler) Authenticate(c echo.Context) error {
	ctx := c.Request().Context()
	params := new(LoginParams)

	if err := c.Bind(params); err != nil {
		errors := map[string]string{
			"title": "No se pudo validar las credenciales, intentelo mas tarde",
		}

		return h.Render(c, http.StatusUnprocessableEntity, view.LoginForm(errors, ""))
	}

	params.Sanitize()

	if errors := params.Validate(); len(errors) > 0 {
		return h.Render(c, http.StatusUnprocessableEntity, view.LoginForm(errors, params.UserName))
	}

	user, err := h.useCase.Login(ctx, *params)
	if err != nil {
		errors := domain.ViewErr(err, "No se pudo validar las credenciales, intentelo mas tarde")

		return h.Render(c, http.StatusUnprocessableEntity, view.LoginForm(errors, params.UserName))
	}

	if err := h.sessionManager.New(c, user); err != nil {
		errors := domain.ViewErr(err, "No se pudo validar las credenciales, intentelo mas tarde")

		return h.Render(c, http.StatusUnprocessableEntity, view.LoginForm(errors, params.UserName))
	}

	return h.navigator.Home(c)
}

// GET /auth/login
func (h *Handler) Login(c echo.Context) error {
	return h.Render(c, http.StatusOK, view.Login())
}

// GET /auth/logout
func (h *Handler) Logout(c echo.Context) error {
	h.sessionManager.Expire(c)

	return h.navigator.Home(c)
}
