package main

import (
	"testing"
	"net/http"
	"strings"
	"strconv"
)

type MockClient struct {}

func (client *MockClient) Head(url string) (resp *http.Response, err error) {
	parts := strings.Split(url, "/")
	last := parts[len(parts) - 1]
	statusCode, err := strconv.Atoi(last)
	if err != nil {
		return nil, err
	}
	resp = &http.Response{StatusCode: statusCode}
	return resp, nil
}

func TestPing(t *testing.T) {
	// client := &http.Client{}
	client := &MockClient{}
	pinger := Pinger{client}
	got := pinger.Ping("https://example.com/200")
	if !got {
		t.Errorf("expected example.com to be available")
	}
	got = pinger.Ping("https://example.com/404")
	if got {
		t.Errorf("expected example.com/404 to be unavailable")
	}
}