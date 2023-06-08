package main

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

func main() {
	query()
}

func queryWithContext() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	id := 3
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	rows, err := db.QueryContext(ctx, "SELECT * FROM users WHERE id < $1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var age int
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name, age)
	}
}

func queryRow() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	id := 3
	var name string
	var age int
	err = db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&id, &name, &age)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(id, name, age)
}

func query() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	id := 3
	rows, err := db.Query("SELECT * FROM users WHERE id < $1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var age int
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name, age)
	}
}
