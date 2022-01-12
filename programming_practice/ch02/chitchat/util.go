package main

import (
	"errors"
	"example.com/go-mod-test/programming_practice/ch02/chitchat/data"
	"fmt"
	"html/template"
	"net/http"
)

func session(w http.ResponseWriter, r *http.Request) (sess data.Session, error) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		sess := data.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}

	return
}

func generateHTML(w http.ResponseWriter, data interface{}, fn ...string) {
	var files []string
	for _, file := range fn {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}
