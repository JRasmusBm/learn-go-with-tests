package blogposts_test

import (
	"errors"
	"io/fs"
	"log"

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

func TestNewBlogPosts(t *testing.T) {
	t.Run("Creates new blog posts", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte("hi")},
			"hello-world2.md": {Data: []byte("hola")},
			"hello-world3.md": {Data: []byte("hola")},
		}

		got, err := blogposts.NewPostsFromFS(fs)
		if err != nil {
			t.Fatal(err)
		}

		want := len(fs)

		if len(got) != want {
			t.Errorf("got %v posts, want %v posts", got, want)
		}

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
