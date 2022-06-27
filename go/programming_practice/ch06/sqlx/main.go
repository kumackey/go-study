package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Db *sqlx.DB

type Post struct {
	Id         int
	Content    string
	AuthorName string `db:"author"`
}

func (p *Post) Create() (err error) {
	err = Db.QueryRow(
		"insert into posts (content, author) values ($1, $2) returning id",
		p.Content, p.AuthorName,
	).Scan(&p.Id)

	return
}

func init() {
	var err error
	Db, err = sqlx.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func main() {
	post := Post{Content: "Hello", AuthorName: "kumaki"}
	post.Create()
	fmt.Println(post)

	readPost := Post{}
	readPost, _ = GetPost(post.Id)
	fmt.Println(readPost)
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRowx("select id, content, author from posts where id = $1", id).StructScan(&post)
	if err != nil {
		fmt.Println("errrrrrr", err)
		return
	}
	return
}
