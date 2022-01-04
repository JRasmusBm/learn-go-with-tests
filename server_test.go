package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubStore struct {
	response string
}

func (s *StubStore) Fetch() string {
	return s.response
}

func TestServer(t *testing.T) {
	t.Run("Returns the response", func(t *testing.T) {
		want := "hello world"

		server := Server(&StubStore{want})

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
