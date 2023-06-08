package main

import (
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"todo/ent"
)

func main() {
	execEnt()
}

func execEnt() {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	err = client.Schema.Create(context.Background())
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	for _, e := range client.Todo.Query().AllX(context.Background()) {
		fmt.Println(e)
	}

	_, err = client.Todo.Create().SetText("Buy milk").SetText("Go言語を学ぶ").Save(context.Background())
	if err != nil {
		log.Fatalf("failed creating todo: %v", err)
	}
}
