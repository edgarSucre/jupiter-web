package user

import (
	"context"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/edgarSucre/jw/features/user/view"
	"github.com/labstack/echo/v4"
)

type useCase interface {
	Create(context.Context, CreateParams) (User, error)
	Delete(context.Context, int) error
	Get(context.Context, int) (User, error)
	List(context.Context) ([]User, error)
}

type (
	navigator interface {
		Home(echo.Context) error
		Location(echo.Context, string, string) error
		Login(echo.Context) error
		IsHxRequest(echo.Context) bool
	}

	Handler struct {
		navigator navigator
		renderer
		useCase useCase
	}

	renderer interface {
		Render(echo.Context, int, templ.Component) error
	}

	smsd interface {
		navigator
		renderer
	}
)

func NewHandler(nav navigator, renderer renderer, uc useCase) *Handler {
	return &Handler{
		navigator: nav,
		renderer:  renderer,
		useCase:   uc,
	}
}

// POST /admin/users
func (h *Handler) Create(c echo.Context) error {
	ctx := c.Request().Context()
	params := new(CreateParams)

	if err := c.Bind(params); err != nil {
		errors := map[string]string{"title": "no se pudo crear el usuario, intentelo mas tarde"}

		return h.Render(c, http.StatusUnprocessableEntity, view.New(view.UserForm{}, errors))
	}

	params.Sanitize()

	if errors := params.Validate(); len(errors) > 0 {
		return h.Render(c, http.StatusUnprocessableEntity, view.New(params.View(), errors))
	}

	_, err := h.useCase.Create(ctx, *params)
	if err != nil {
		errors := map[string]string{"title": "no se pudo crear el usuario, intentelo mas tarde"}

		return h.Render(c, http.StatusUnprocessableEntity, view.New(params.View(), errors))
	}

	return h.navigator.Location(c, "/admin/users", "#mainSection")
}

// GET admin/users/:id/edit
func (h *Handler) Edit(c echo.Context) error {
	ctx := c.Request().Context()
	idParam := c.Param("id")

	userID, err := strconv.Atoi(idParam)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	user, err := h.useCase.Get(ctx, userID)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	viewData, errors := user.viewUpdateForm(), map[string]string{}

	if h.navigator.IsHxRequest(c) {
		return h.Render(c, http.StatusOK, view.Edit(viewData, errors))
	}

	return h.Render(c, http.StatusOK, view.EditReload(viewData, errors))
}

// GET /admin/users
func (h *Handler) Index(c echo.Context) error {
	if h.navigator.IsHxRequest(c) {
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
	user, errors := view.UserForm{}, map[string]string{}

	if h.navigator.IsHxRequest(c) {
		return h.Render(c, http.StatusOK, view.New(user, errors))
	}

	return h.Render(c, http.StatusOK, view.NewReload(user, errors))
}

// DELETE /admin/users/userID
func (h *Handler) Delete(c echo.Context) error {
	idParam := c.Param("id")

	userID, err := strconv.Atoi(idParam)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	if err = h.useCase.Delete(ctx, userID); err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.NoContent(http.StatusOK)
}
