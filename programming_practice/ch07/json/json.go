package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

func main() {
	jsonFile, err := os.Open("programming_practice/ch07/json/post.json")
	if err != nil {
		fmt.Println("Error opening", err)
		return
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading json", err)
		return
	}

	var post Post
	json.Unmarshal(jsonData, &post)
	fmt.Println(post)
}
