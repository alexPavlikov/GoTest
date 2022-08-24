package models

type Post struct {
	id      string
	title   string
	content string
}

func NewPost(id, title, content string) *Post {
	return &Post{id, title, content}
}
