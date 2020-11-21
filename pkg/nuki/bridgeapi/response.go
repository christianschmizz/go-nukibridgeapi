package bridgeapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// APIResponseHandler eases working with a http response
type APIResponseHandler struct {
	httpResponse *http.Response
}

// NewAPIResponseHandler returns a new response handler
func NewAPIResponseHandler(resp *http.Response) (*APIResponseHandler, error) {
	readBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body for '%s': %w", resp.Request.URL, err)
	}
	resp.Body = ioutil.NopCloser(bytes.NewReader(readBody))
	return &APIResponseHandler{
		httpResponse: resp,
	}, nil
}

// Decode decodes the response's body to the given value
func (r *APIResponseHandler) Decode(v interface{}) error {
	if err := json.NewDecoder(r.httpResponse.Body).Decode(v); err != nil {
		return fmt.Errorf("decoding failed: %w\n\n%+v", err, r.httpResponse.Body)
	}
	return nil
}

// Is checks the response for the given status code
func (r *APIResponseHandler) Is(statusCode int) bool {
	return r.httpResponse.StatusCode == statusCode
}
