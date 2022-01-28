package server

import (
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"golang.org/x/net/context"
)

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string

		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("Spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}

		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func TestServer(t *testing.T) {
	t.Run("When the request is successful", func(t *testing.T) {
		data := "hello world"
		store := &SpyStore{response: data}
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
	})

	t.Run("When the request is cancelled", func(t *testing.T) {
		store := &SpyStore{response: "hello world"}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingContext, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingContext)

		response := &SpyResponseWriter{}

		server.ServeHTTP(response, request)

		t.Run("Does not write a response", func(t *testing.T) {
			if response.written {
				t.Error("No response should have been written")
			}
		})
	})
}
