// Code generated by rpcgen; DO NOT EDIT.
//
// Source: rpc/rpc.go
// Template: cmd/rpcgen/client.go.tmpl

package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/guibirow/apigen/api"
)

func (c *Client) Login(ctx context.Context, req api.LoginRequest) (*api.LoginResponse, error) {
	r, err := http.NewRequest(req.Method(), fmt.Sprintf("%s/%s", c.endpoint, EncodeVars(req, api.LoginEndpoint)), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.conn.Do(r)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	var result api.LoginResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, err
}

func (c *Client) CreatePost(ctx context.Context, req api.CreatePostRequest) (*api.CreatePostResponse, error) {
	r, err := http.NewRequest(req.Method(), fmt.Sprintf("%s/%s", c.endpoint, EncodeVars(req, api.CreatePostEndpoint)), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.conn.Do(r)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	var result api.CreatePostResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, err
}

func (c *Client) GetFeed(ctx context.Context, req api.GetFeedRequest) (*api.GetFeedResponse, error) {
	r, err := http.NewRequest(req.Method(), fmt.Sprintf("%s/%s", c.endpoint, EncodeVars(req, api.GetFeedEndpoint)), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.conn.Do(r)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	var result api.GetFeedResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, err
}

func (c *Client) FollowUser(ctx context.Context, req api.FollowUserRequest) (*api.FollowUserResponse, error) {
	r, err := http.NewRequest(req.Method(), fmt.Sprintf("%s/%s", c.endpoint, EncodeVars(req, api.FollowUserEndpoint)), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.conn.Do(r)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	var result api.FollowUserResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, err
}
