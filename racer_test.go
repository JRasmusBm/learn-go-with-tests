package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))

}

func TestRacer(t *testing.T) {
	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(0 * time.Millisecond)

	defer slowServer.Close()
	defer fastServer.Close()

	t.Run("Returns the faster url", func(t *testing.T) {
		got := Racer(slowServer.URL, fastServer.URL)
		want := fastServer.URL

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Input order does not matter", func(t *testing.T) {
		got := Racer(fastServer.URL, slowServer.URL)
		want := fastServer.URL

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
