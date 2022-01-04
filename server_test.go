package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"golang.org/x/net/context"
)

type SpyStore struct {
	response  string
	cancelled bool
}

func (s *SpyStore) Fetch() string {
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func TestServer(t *testing.T) {
	t.Run("Returns the response", func(t *testing.T) {
		want := "hello world"

		server := Server(&SpyStore{response: want, cancelled: false})

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Tells the store to cancel work if request was cancelled", func(t *testing.T) {
		store := &SpyStore{response: "hello world", cancelled: false}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingContext, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingContext)

		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if !store.cancelled {
			t.Error("Store was not told to cancel")
		}
	})
}
