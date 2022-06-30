package blogposts_test

import (
	blogposts "github.com/jrasmusbm/learn-go-with-tests/reading-files"

	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	t.Run("Creates new blog posts", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte("hi")},
			"hello-world2.md": {Data: []byte("hola")},
		}

		got := len(blogposts.NewPostsFromFS(fs))
		want := len(fs)

		if got != want {
			t.Errorf("got %v posts, want %v posts", got, want)
		}
	})

}
