package blogposts

import (
	"io/fs"
)

type Post struct {
}

func NewPostsFromFS(filesystem fs.FS) []Post {
	return []Post{{}, {}}
}
