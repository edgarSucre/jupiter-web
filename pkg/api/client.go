package api

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

const contentType = "application/json"

type Client struct {
	baseUrl string
	c       http.Client
	token   string
}

func NewClient() *Client {
	client := &Client{
		baseUrl: "jupiter-host:port",
		c:       http.Client{},
	}

	return client
}

func (cl *Client) URL(path string) string {
	return fmt.Sprintf("%s/%s", cl.baseUrl, path)
}

func (cl *Client) addAuth(req *http.Request) {
	req.Header.Set("Authorization", "Bearer "+cl.token)
}

func (cl *Client) Post(ctx context.Context, path string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, cl.URL(path), body)
	if err != nil {
		return nil, fmt.Errorf("%w: POST", err)
	}

	cl.addAuth(req)

	return cl.c.Do(req)
}

func (cl *Client) Get(ctx context.Context, path string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, cl.URL(path), nil)
	if err != nil {
		return nil, fmt.Errorf("%w: GET", err)
	}

	cl.addAuth(req)

	return cl.c.Do(req)
}
