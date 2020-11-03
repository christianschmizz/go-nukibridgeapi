package bridgeapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type apiResponse struct {
	path         string
	url          string
	httpResponse *http.Response
}

func NewApiResponse(resp *http.Response) (*apiResponse, error) {
	readBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body for '%s': %w", resp.Request.URL, err)
	}
	resp.Body = ioutil.NopCloser(bytes.NewReader(readBody))
	return &apiResponse{
		httpResponse: resp,
	}, nil
}

func (r *apiResponse) Decode(v interface{}) error {
	if err := json.NewDecoder(r.httpResponse.Body).Decode(v); err != nil {
		return fmt.Errorf("decoding failed: %w\n\n%+v", err, r.httpResponse.Body)
	}
	return nil
}

func (r *apiResponse) Is(statusCode int) bool {
	return r.httpResponse.StatusCode == statusCode
}
