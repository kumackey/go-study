package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"net/http"
	"path"
	"strconv"
)

func main() {
	var err error
	db, err := sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/post/", handleRequest(&Post{Db: db}))
	server.ListenAndServe()
}

type Text interface {
	fetch(id int) error
	create() error
	update() error
	delete() error
}

func handleRequest(t Text) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		switch r.Method {
		case "GET":
			err = handleGet(w, r, t)
		case "POST":
			err = handlePost(w, r, t)
		case "PUT":
			err = handlePut(w, r, t)
		case "DELETE":
			err = handleDelete(w, r, t)

		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func handleDelete(w http.ResponseWriter, r *http.Request, t Text) error {
	//id, err := strconv.Atoi(path.Base(r.URL.Path))
	//if err != nil {
	//	return err
	//}
	//post, err := fetch(id)
	//if err != nil {
	//	return err
	//}
	//
	//err = post.delete()
	//if err != nil {
	//	return err
	//}
	//w.WriteHeader(200)
	return nil
}

func handlePut(w http.ResponseWriter, r *http.Request, t Text) error {
	//id, err := strconv.Atoi(path.Base(r.URL.Path))
	//if err != nil {
	//	return err
	//}
	//err = t.fetch(id)
	//if err != nil {
	//	return err
	//}
	//len := r.ContentLength
	//body := make([]byte, len)
	//r.Body.Read(body)
	//json.Unmarshal(body, &post)
	//err = t.update()
	//if err != nil {
	//	return err
	//}
	//w.WriteHeader(200)
	return nil
}

func handlePost(w http.ResponseWriter, r *http.Request, t Text) error {
	//len := r.ContentLength
	//body := make([]byte, len)
	//r.Body.Read(body)
	//var post Post
	//json.Unmarshal(body, &post)
	//post.create()
	//w.WriteHeader(200)
	//id := strconv.Itoa(post.Id)
	//w.Write([]byte(id))
	return nil
}

func handleGet(w http.ResponseWriter, r *http.Request, t Text) error {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return err
	}

	err = t.fetch(id)
	if err != nil {
		return err
	}

	output, err := json.MarshalIndent(t, "", "\t\t")
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)

	return nil
}
