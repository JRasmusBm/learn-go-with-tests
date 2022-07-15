package blogposts

import (
	"io/fs"
)

type Post struct {
}

func NewPostsFromFS(filesystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(filesystem, ".")

	if err != nil {
		return nil, err
	}

	var result []Post

	for range dir {
		result = append(result, Post{})
	}

	return result, nil
}
