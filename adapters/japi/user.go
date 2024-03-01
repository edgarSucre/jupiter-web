package japi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/edgarSucre/jw/domain"
)

func (cl *Client) CreateUser(ctx context.Context, userName string) (domain.JupyterUser, error) {
	var user domain.JupyterUser

	resp, err := cl.Post(ctx, "users/"+userName, nil)
	if err != nil {
		return user, fmt.Errorf("%w:, CreateUser(%s)", err, userName)
	}

	if resp.StatusCode != http.StatusCreated {
		return user, fmt.Errorf("%w:, Status(%v)", StatusNotCreated, resp.StatusCode)
	}

	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(&user)
	if err != nil {
		return user, fmt.Errorf("%w, %s", JsonError, err.Error())
	}

	return user, nil
}

func (cl *Client) UserList(ctx context.Context) ([]domain.JupyterUser, error) {
	resp, err := cl.Get(ctx, "users")
	if err != nil {
		return nil, fmt.Errorf("%w:, userList", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w:, Status(%v)", StatusNotOK, resp.StatusCode)
	}

	decoder := json.NewDecoder(resp.Body)
	userList := make([]domain.JupyterUser, 0, 10)

	err = decoder.Decode(&userList)
	if err != nil {
		return nil, fmt.Errorf("%w, %s", JsonError, err.Error())
	}

	return userList, nil
}
