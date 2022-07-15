package blogposts

import (
	"io/fs"
)

type Post struct {
}

func NewPostsFromFS(filesystem fs.FS) []Post {
	dir, _ := fs.ReadDir(filesystem, ".")
	var result []Post

	for range dir {
		result = append(result, Post{})
	}

	return result
}
