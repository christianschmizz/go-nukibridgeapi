package bridge

import (
	"fmt"
)

type Callback struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
}

type ListCallbacksResponse struct {
	Callbacks []Callback `json:"callbacks"`
}

func (c *connection) ListCallbacks() (*ListCallbacksResponse, error) {
	var listCallbacksResponse ListCallbacksResponse
	if err := c.get(c.hashedURL("/callback/list", nil), &listCallbacksResponse); err != nil {
		return nil, fmt.Errorf("could not fetch list of callbacks: %w", err)
	}
	return &listCallbacksResponse, nil
}

type RemoveCallbackResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

func (c *connection) RemoveCallback(callbackID int) (*RemoveCallbackResponse, error) {
	options := &struct {
		CallbackID int `url:"id"`
	}{callbackID}
	var removeCallbackResponse RemoveCallbackResponse
	if err := c.get(c.hashedURL("/callback/remove", options), &removeCallbackResponse); err != nil {
		return nil, fmt.Errorf("could not remove callback: %w", err)
	}
	return &removeCallbackResponse, nil
}

type AddCallbackResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

func (c *connection) AddCallback(callbackURL string) (*AddCallbackResponse, error) {
	options := &struct {
		CallbackURL string `url:"url"`
	}{callbackURL}
	var addCallbackResponse AddCallbackResponse
	if err := c.get(c.hashedURL("/callback/add", options), &addCallbackResponse); err != nil {
		return nil, fmt.Errorf("could not add callback: %w", err)
	}
	return &addCallbackResponse, nil
}
