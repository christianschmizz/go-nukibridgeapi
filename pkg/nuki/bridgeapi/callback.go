package bridgeapi

import (
	"fmt"
)

// Callback is an item at the list of callbacks at a bridge
type Callback struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
}

// ListCallbacksResponse represents the result of an listing request
type ListCallbacksResponse struct {
	Callbacks []Callback `json:"callbacks"`
}

// ListCallbacks request a list of registered callbacks from the bridge
func (c *Connection) ListCallbacks() (*ListCallbacksResponse, error) {
	var response ListCallbacksResponse
	if err := c.get(c.hashedURL("/callback/list", nil), &response); err != nil {
		return nil, fmt.Errorf("could not fetch list of callbacks: %w", err)
	}
	return &response, nil
}

// RemoveCallbackResponse represents the result of an callback removal request
type RemoveCallbackResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

// RemoveCallback requests the removal of the given callback from the bridge
func (c *Connection) RemoveCallback(callbackID int) (*RemoveCallbackResponse, error) {
	options := &struct {
		CallbackID int `url:"id"`
	}{callbackID}
	var response RemoveCallbackResponse
	if err := c.get(c.hashedURL("/callback/remove", options), &response); err != nil {
		return nil, fmt.Errorf("could not remove callback: %w", err)
	}
	return &response, nil
}

// AddCallbackResponse represents the result of an request
type AddCallbackResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

// AddCallback requests the addition of the given callback
func (c *Connection) AddCallback(callbackURL string) (*AddCallbackResponse, error) {
	options := &struct {
		CallbackURL string `url:"url"`
	}{callbackURL}
	var response AddCallbackResponse
	if err := c.get(c.hashedURL("/callback/add", options), &response); err != nil {
		return nil, fmt.Errorf("could not add callback: %w", err)
	}
	return &response, nil
}
