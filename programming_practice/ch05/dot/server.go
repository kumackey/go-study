package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(
		"programming_practice/ch05/dot/tmpl.html",
		"programming_practice/ch05/dot/tmpl2.html",
	)

	t.Execute(w, "Hello")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
