package user

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
	"github.com/edgarSucre/jw/features/user/view"
	"github.com/labstack/echo/v4"
)

type useCase interface {
	Create(context.Context, CreateParams) (User, error)
	List(context.Context) ([]User, error)
}

type (
	Handler struct {
		renderer
		useCase useCase
	}

	renderer interface {
		Render(echo.Context, int, templ.Component) error
	}
)

func NewHandler(renderer renderer, uc useCase) *Handler {
	return &Handler{
		renderer: renderer,
		useCase:  uc,
	}
}

// POST /admin/users
func (h *Handler) Create(c echo.Context) error {
	ctx := c.Request().Context()
	params := new(CreateParams)

	if err := c.Bind(params); err != nil {
		errors := map[string]string{"title": "no se pudo crear el usuario, intentelo mas tarde"}

		return h.Render(c, http.StatusUnprocessableEntity, view.NewUserForm(view.UserForm{}, errors))
	}

	params.Sanitize()

	if errors := params.Validate(); len(errors) > 0 {
		return h.Render(c, http.StatusUnprocessableEntity, view.NewUserForm(params.View(), errors))
	}

	_, err := h.useCase.Create(ctx, *params)
	if err != nil {
		errors := map[string]string{"title": "no se pudo crear el usuario, intentelo mas tarde"}

		return h.Render(c, http.StatusUnprocessableEntity, view.NewUserForm(params.View(), errors))
	}

	return c.Redirect(http.StatusSeeOther, "/admin/users")
}

// GET /admin/users
func (h *Handler) Index(c echo.Context) error {
	if c.Request().Header.Get("Hx-Request") == "true" {
		return h.Render(c, http.StatusOK, view.IndexCmp())
	}

	return h.Render(c, http.StatusOK, view.IndexFull())
}

// GET /admin/user/list HxRequest only
func (h *Handler) List(c echo.Context) error {
	mUsers, err := h.useCase.List(c.Request().Context())
	if err != nil {
		// return view.ListErr()
		//TODO: handle error
	}

	return h.Render(c, http.StatusOK, view.List(usersToView(mUsers)))
}

// GET /admin/users/new
func (h *Handler) New(c echo.Context) error {
	return h.Render(c, http.StatusOK, view.NewUser(view.UserForm{}, map[string]string{}))
}
