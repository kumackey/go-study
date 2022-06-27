package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	rand.Seed(time.Now().Unix())
	err := t.Execute(w, rand.Intn(10) > 5)
	fmt.Println(err)
}

type Post struct {
	Id      int
	Content string
	Author  string
}

var PostById map[int]*Post
var PostsByAuthor map[string][]*Post

func main() {
	PostById = make(map[int]*Post)
	PostsByAuthor = make(map[string][]*Post)

	post1 := Post{Id: 1, Content: "hello World!", Author: "Sau Sheong"}
	post2 := Post{Id: 2, Content: "hello World!2", Author: "Sau Sheong2"}
	post3 := Post{Id: 3, Content: "hello World!33", Author: "Sau Sheong3"}
	post4 := Post{Id: 4, Content: "hello World!444", Author: "Sau Sheong4"}

	store(post1)
	store(post2)
	store(post3)
	store(post4)

	fmt.Println(PostById[1])
	fmt.Println(PostById[2])
}

func store(post Post) {
	PostById[post.Id] = &post
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], &post)
}
