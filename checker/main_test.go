package main

import (
	"net/http"
	"testing"
)

type MockHttpClient struct {
	Response *http.Response
	Error    error
}

func (c *MockHttpClient) Get(url string) (response *http.Response, err error) {
	return c.Response, c.Error
}

func Test_check(t *testing.T) {

}
