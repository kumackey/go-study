package main

import "errors"

type Comment struct {
	Id      int
	Content string
	Author  string
	Post    *Post
}

func (c *Comment) create() (err error) {
	if c.Post == nil {
		err = errors.New("投稿が見つかりません")
		return
	}

	err = Db.QueryRow(
		"insert into comments (content, author, post_id) values ($1, $2, $3) returning id",
		c.Content, c.Author, c.Post.Id,
	).Scan(&c.Id)

	return
}
