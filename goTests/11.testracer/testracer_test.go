package testracer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRacer(t *testing.T) {
	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(0 * time.Millisecond)
	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL
	want := fastURL

	got, err := Racer(slowURL, fastURL)

	assert.Equal(t, want, got)
	assert.NoError(t, err)
}

func TestRacerError(t *testing.T) {
	slowServer := makeDelayedServer(11 * time.Second)
	fastServer := makeDelayedServer(12 * time.Second)
	defer slowServer.Close()
	defer fastServer.Close()

	_, err := Racer(slowServer.URL, fastServer.URL)

	assert.Error(t, err)
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(delay)
			w.WriteHeader(http.StatusOK)
		}))
}
