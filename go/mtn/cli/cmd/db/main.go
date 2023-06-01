package main

import (
	_ "github.com/mattn/go-sqlite3"
)

//func main() {
//	query := "虫 AND ココア"
//	rows, err := db.Query(`
//	SELECT
//   	a.author, c.title
//	FROM contents c
//	INNER JOIN authors a
//	    ON c.author_id = a.author_id
//	INNER JOIN contents_fts cf
//	    ON c.rowid = cf.docid
//	   AND words MATCH ?
//	   `, query)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer rows.Close()
//
//	for rows.Next() {
//		var author, title string
//		if err := rows.Scan(&author, &title); err != nil {
//			log.Fatal(err)
//		}
//		log.Println(author, title)
//	}
//}
