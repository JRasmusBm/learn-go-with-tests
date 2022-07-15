package blogposts

import (
	"io/fs"
)

type Post struct {
	Title string
}

func NewPostsFromFS(filesystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(filesystem, ".")

	if err != nil {
		return nil, err
	}

	var result []Post

	for _, entry := range dir {
		contents, err := fs.ReadFile(filesystem, entry.Name())
		post, err := parsePost(string(contents))

		if err != nil {
			return nil, err
		}

		result = append(result, post)
	}

	return result, nil
}

func parsePost(contents string) (Post, error) {
	return Post{Title: contents[7:]}, nil
}
