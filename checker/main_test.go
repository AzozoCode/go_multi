package main

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
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

	mockClient := &MockHttpClient{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString("ok")),
		},
		Error: nil,
	}

	cfg := SiteConfig{
		URL:           "https://localhost.com",
		Frequency:     1,
		AcceptedCodes: []int{200},
	}

	results := make(chan Result, 1)

	check(cfg, mockClient, results)

	result := <-results

	if !result.Up || result.Status != 200 {
		t.Error("site should be up with 200 status code")
	}
}

func Test_check_Error(t *testing.T) {

	mockClient := &MockHttpClient{
		Error: errors.New("error"),
	}

	cfg := SiteConfig{
		URL: "https://localhost.com",
	}

	results := make(chan Result, 1)

	check(cfg, mockClient, results)

	result := <-results

	if result.Up {
		t.Error("site should be down with 0 status code")
	}
}

func Test_DefaultClient_Get(t *testing.T) {

	client := DefaultClient{}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	}))

	defer srv.Close()

	resp, err := client.Get(srv.URL)

	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	statusCode := resp.StatusCode

	if statusCode != http.StatusOK {
		t.Errorf("expected status code to be %d but got %d", http.StatusOK, statusCode)
	}

	expectedBody := "ok"

	bs, err := io.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if expectedBody != string(bs) {
		t.Errorf("expected %v; got %v", expectedBody, string(bs))
	}
}
