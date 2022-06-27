package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type FakePost struct {
	Id      int
	Content string
	Author  string
}

func (f *FakePost) fetch(id int) error {
	f.Id = id
	return nil
}

func (f *FakePost) create() error {
	//TODO implement me
	panic("implement me")
}

func (f *FakePost) update() error {
	//TODO implement me
	panic("implement me")
}

func (f *FakePost) delete() error {
	//TODO implement me
	panic("implement me")
}

func TestHanleGet(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest(&FakePost{}))
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/post/1", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.Id != 1 {
		t.Error("Cannnot fetch JSON POST")
	}
}
