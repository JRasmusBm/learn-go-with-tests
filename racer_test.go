package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))

	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	slowUrl := slowServer.URL
	fastUrl := fastServer.URL

	t.Run("Returns the faster url", func(t *testing.T) {

		got := Racer(slowUrl, fastUrl)
		want := fastUrl

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Input order does not matter", func(t *testing.T) {
		slowUrl := "http://www.facebook.com"
		fastUrl := "http://www.quii.co.uk"

		got := Racer(fastUrl, slowUrl)
		want := fastUrl

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	slowServer.Close()
	fastServer.Close()
}
