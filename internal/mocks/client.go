package mocks

import (
	"bufio"
	"io/ioutil"
	"net/http"
	"os"
)

type fileResponseClient struct {
	name   string
	status int
}

func NewFileResponseClient(name string, status int) *fileResponseClient {
	return &fileResponseClient{
		name:   name,
		status: status,
	}
}

func (c *fileResponseClient) Do(*http.Request) (*http.Response, error) {
	f, _ := os.Open(c.name)
	r := ioutil.NopCloser(bufio.NewReader(f))
	return &http.Response{
		StatusCode: c.status,
		Body:       r,
	}, nil
}
