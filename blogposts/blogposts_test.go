package blogposts_test

import (
	"errors"
	"io/fs"
	"log"
	"reflect"

	blogposts "github.com/jrasmusbm/learn-go-with-tests/reading-files"

	"testing"
	"testing/fstest"
)

var givenError error = errors.New("Failed to open!")

type StubFailingFS struct {
}

func (s *StubFailingFS) Open(name string) (fs.File, error) {
	log.Printf("jrasmusbm")

	return nil, givenError
}

func assertPostsEqual(t testing.TB, want, got []blogposts.Post) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v posts, want %v posts", got, want)
	}
}

func TestNewBlogPosts(t *testing.T) {
	t.Run("Creates new blog posts", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte("Title: Hi")},
			"hello-world2.md": {Data: []byte("Title: Rewriting Git History")},
			"hello-world3.md": {Data: []byte("Title: Hello")},
		}

		got, err := blogposts.NewPostsFromFS(fs)
		if err != nil {
			t.Fatal(err)
		}

		want := []blogposts.Post{
			{Title: "Hi"},
			{Title: "Rewriting Git History"},
			{Title: "Hello"},
		}

		assertPostsEqual(t, got, want)
	})

	t.Run("Propagates errors from fs.Open", func(t *testing.T) {
		_, err := blogposts.NewPostsFromFS(&StubFailingFS{})

		if err == nil {
			t.Fatalf("Expected an error!")
		}

		if err.Error() != givenError.Error() {
			t.Errorf("Expected error %v got %v", err, givenError)
		}
	})

}
