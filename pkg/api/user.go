package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/edgarSucre/jw/pkg/api/models"
)

func (cl *Client) CreateUser(ctx context.Context, userName string) (*models.User, error) {
	resp, err := cl.Post(ctx, "users/"+userName, nil)
	if err != nil {
		return nil, fmt.Errorf("%w:, CreateUser(%s)", err, userName)
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("%w:, Status(%v)", StatusNotCreated, resp.StatusCode)
	}

	decoder := json.NewDecoder(resp.Body)
	user := new(models.User)

	err = decoder.Decode(user)
	if err != nil {
		return nil, fmt.Errorf("%w, %s", JsonError, err.Error())
	}

	return user, nil
}

func (cl *Client) ListUsers(ctx context.Context, params models.UserListQuery) (*[]models.User, error) {
	resp, err := cl.Get(ctx, "users")
	if err != nil {
		return nil, fmt.Errorf("%w:, ListUsers", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w:, Status(%v)", StatusNotOK, resp.StatusCode)
	}

	decoder := json.NewDecoder(resp.Body)
	userList := make([]models.User, 0, 10)

	err = decoder.Decode(&userList)
	if err != nil {
		return nil, fmt.Errorf("%w, %s", JsonError, err.Error())
	}

	return &userList, nil
}
