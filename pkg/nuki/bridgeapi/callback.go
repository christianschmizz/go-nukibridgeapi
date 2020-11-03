package bridgeapi

import (
	"net/http"
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
	resp, err := c.request("/callback/list", nil)
	if err != nil {
		return nil, err
	}

	if resp.Is(http.StatusUnauthorized) {
		return nil, ErrInvalidToken
	}

	var data ListCallbacksResponse
	if err := resp.Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
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

	resp, err := c.request("/callback/remove", options)
	if err != nil {
		return nil, err
	}

	if resp.Is(http.StatusBadRequest) {
		return nil, ErrInvalidURL
	} else if resp.Is(http.StatusUnauthorized) {
		return nil, ErrInvalidToken
	}

	var data RemoveCallbackResponse
	if err := resp.Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
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

	resp, err := c.request("/callback/add", options)
	if err != nil {
		return nil, err
	}

	if resp.Is(http.StatusBadRequest) {
		return nil, ErrInvalidURL
	} else if resp.Is(http.StatusUnauthorized) {
		return nil, ErrInvalidToken
	}

	var data AddCallbackResponse
	if err := resp.Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
