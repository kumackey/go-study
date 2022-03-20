package main

import "database/sql"

type Post struct {
	Db      *sql.DB
	Id      int
	Content string
	Author  string
}

func (post *Post) create() error {
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	stmt, err := post.Db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return err
}

func (post *Post) update() error {
	_, err := post.Db.Exec("update posts set content = $2, author = $3 where id = $1", post.Id, post.Content, post.Author)
	return err
}

func (post *Post) delete() error {
	_, err := post.Db.Exec("delete from posts where id = $1", post.Id)
	return err
}

func (post *Post) fetch(id int) error {
	err := post.Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return err
}
