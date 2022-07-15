package blogposts

import (
	"io"
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

	for _, file := range dir {
		post, err := readPostFile(filesystem, file)

		if err != nil {
			return nil, err
		}

		result = append(result, post)
	}

	return result, nil
}

func readPostFile(filesystem fs.FS, entry fs.DirEntry) (Post, error) {
	file, err := filesystem.Open(entry.Name())

	if err != nil {
		return Post{}, err
	}
	defer file.Close()

	rawContents, err := io.ReadAll(file)
	if err != nil {
		return Post{}, err
	}

	return Post{Title: string(rawContents)[7:]}, nil
}
