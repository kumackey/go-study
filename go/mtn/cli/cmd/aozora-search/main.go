package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
)

const usage = `Usage: aozora-search [options] <command> [<args>]`

func main() {
	var dsn string
	flag.StringVar(&dsn, "d", "datebase.sqlite", "datebase")
	flag.Usage = func() {
		fmt.Print(usage)
	}
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(2)
	}

	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	switch flag.Arg(0) {
	case "authors":
		err = showAuthors(db)
	case "titles":
		if flag.NArg() != 2 {
			flag.Usage()
			os.Exit(2)
		}

		err = showTitles(db, flag.Arg(1))
	case "contents":
		if flag.NArg() != 3 {
			flag.Usage()
			os.Exit(2)
		}
		err = showContents(db, flag.Arg(1), flag.Arg(2))
	case "query":
		if flag.NArg() != 2 {
			flag.Usage()
			os.Exit(2)
		}
		err = query(db, flag.Arg(1))
	}
	if err != nil {
		log.Fatal(err)
	}
}

func query(db *sql.DB, arg string) error {
	return nil
}

func showContents(db *sql.DB, arg string, arg2 string) error {
	return nil
}

func showTitles(db *sql.DB, arg string) error {
	return nil
}

func showAuthors(db *sql.DB) error {
	return nil
}
