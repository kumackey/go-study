package main

import (
	"database/sql"
	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"
	"golang.org/x/text/encoding/japanese"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "datebase.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	queries := []string{
		`CREATE TABLE IF NOT EXISTS authors(author_id TEXT, author TEXT, PRIMARY KEY(author_id))`,
		`CREATE TABLE IF NOT EXISTS contents(author_id TEXT, title_id TEXT, title TEXT, content TEXT, PRIMARY KEY(author_id, title_id))`,
		`CREATE VIRTUAL TABLE IF NOT EXISTS contents_fts USING fts4(words)`,
	}

	for _, q := range queries {
		_, err := db.Exec(q)
		if err != nil {
			log.Fatal(err)
		}
	}

	_, err = db.Exec(`INSERT INTO authors(author_id, author) VALUES(?, ?)`,
		"000879",
		"芥川 竜之介",
	)
	if err != nil {
		log.Fatal(err)
	}

	b, err := os.ReadFile("ababababa.txt")
	if err != nil {
		log.Fatal(err)
	}
	b, err = japanese.ShiftJIS.NewDecoder().Bytes(b)
	if err != nil {
		log.Fatal(err)
	}
	content := string(b)
	res, err := db.Exec(`INSERT INTO contents(author_id, title_id, title, content) VALUES(?, ?, ?, ?)`,
		"000879",
		"14",
		"あばばばば",
		content,
	)
	if err != nil {
		log.Fatal(err)
	}
	docID, err := res.LastInsertId()

	t, err := tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
	if err != nil {
		log.Fatal(err)
	}

	seg := t.Wakati(content)
	_, err = db.Exec(`INSERT INTO contents_fts(docid, words) VALUES(?, ?)`, docID, strings.Join(seg, " "))
	if err != nil {
		log.Fatal(err)
	}

	query := "虫 AND ココア"
	rows, err := db.Query(`
	SELECT 
    	a.author, c.title 
	FROM contents c 
	INNER JOIN authors a 
	    ON c.author_id = a.author_id
	INNER JOIN contents_fts cf 
	    ON c.rowid = cf.docid
	   AND words MATCH ?
	   `, query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var author, title string
		if err := rows.Scan(&author, &title); err != nil {
			log.Fatal(err)
		}
		log.Println(author, title)
	}
}
