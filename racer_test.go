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
func AssertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestRacer(t *testing.T) {
	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(0 * time.Millisecond)

	defer slowServer.Close()
	defer fastServer.Close()

	t.Run("Returns the faster url", func(t *testing.T) {
		got, err := Racer(slowServer.URL, fastServer.URL, 1*time.Second)
		want := fastServer.URL

		AssertNoError(t, err)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}

	})

	t.Run("Input order does not matter", func(t *testing.T) {
		got, err := Racer(fastServer.URL, slowServer.URL, 1*time.Second)
		want := fastServer.URL

		AssertNoError(t, err)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Returns error on timeout", func(t *testing.T) {
		verySlowServer := makeDelayedServer(20 * time.Second)

		defer verySlowServer.Close()

		_, err := Racer(verySlowServer.URL, verySlowServer.URL, 0*time.Second)

		if err == nil {
			t.Error("Expected error but didn't get one")
		}
	})

}
