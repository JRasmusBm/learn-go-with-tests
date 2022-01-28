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
	t         *testing.T
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func (s *SpyStore) assertWasCancelled() {
	s.t.Helper()
	if !s.cancelled {
		s.t.Error("Store was not told to cancel")
	}
}

func (s *SpyStore) assertWasNotCancelled() {
	s.t.Helper()
	if s.cancelled {
		s.t.Error("Store was told to cancel")
	}
}

func TestServer(t *testing.T) {
	t.Run("When the request is successful", func(t *testing.T) {
		data := "hello world"
		store := &SpyStore{response: data, cancelled: false, t: t}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		t.Run("Returns the data", func(t *testing.T) {
			got := response.Body.String()
			want := data

			if got != want {
				t.Errorf("got %v want %v", got, want)
			}
		})

		t.Run("Does not get cancelled", func(t *testing.T) {
			store.assertWasNotCancelled()
		})
	})

	t.Run("When the request is cancelled", func(t *testing.T) {
		store := &SpyStore{response: "hello world", cancelled: false, t: t}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingContext, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingContext)

		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		t.Run("Tells the store to cancel", func(t *testing.T) {
			store.assertWasCancelled()
		})
	})
}
