package mocks

import (
	"bufio"
	"io/ioutil"
	"net/http"
	"os"
)

// The FileResponseClient fakes the execution of a request and returns always
// the contents of the given file with the given HTTP status.
type FileResponseClient struct {
	name   string
	status int
}

// NewFileResponseClient instantiates a new client
func NewFileResponseClient(name string, status int) *FileResponseClient {
	return &FileResponseClient{
		name:   name,
		status: status,
	}
}

// Do fakes the given request and returns the defined file content and status code
func (c *FileResponseClient) Do(*http.Request) (*http.Response, error) {
	f, err := os.Open(c.name)
	if err != nil {
		return nil, err
	}
	r := ioutil.NopCloser(bufio.NewReader(f))
	return &http.Response{
		StatusCode: c.status,
		Body:       r,
	}, nil
}
