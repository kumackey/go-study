package main

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"log"
	"net/mail"
	"net/smtp"
	"os"
	"time"
)

type Todo struct {
	bun.BaseModel `bun:"table:todos,alias:t"`

	ID        int64     `bun:"id,pk,autoincrement"`
	Content   string    `bun:"content,notnull"`
	Done      bool      `bun:"done"`
	Until     time.Time `bun:"until,nullzero"`
	CreatedAt time.Time
	UpdatedAt time.Time `bun:",nullzero"`
	DeletedAt time.Time `bun:",soft_delete,nullzero"`
}

func main() {
	sqldb, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer sqldb.Close()

	db := bun.NewDB(sqldb, pgdialect.New())
	defer db.Close()

	var todos []Todo
	ctx := context.Background()
	err = db.NewSelect().Model(&todos).
		Order("created_at").
		Where("until is not null").
		Where("done is false").Scan(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if len(todos) == 0 {
		return
	}

	from := mail.Address{Name: "Todo App", Address: os.Getenv("MAIL_FROM")}
	var buf bytes.Buffer
	buf.WriteString("From: " + from.String() + "\r\n")
	buf.WriteString("To: " + os.Getenv("MAIL_TO") + "\r\n")
	buf.WriteString("Subject: Todo App Notification\r\n")
	buf.WriteString("\r\n")
	buf.WriteString("You have " + string(len(todos)) + " todos.\r\n")
	for _, todo := range todos {
		fmt.Fprintf(&buf, "%s: %s\r\n", todo.Until.Format("2006-01-02"), todo.Content)
	}

	smtpAuth := smtp.PlainAuth(
		os.Getenv("MAIL_DOMAIN"),
		os.Getenv("MAIL_USER"),
		os.Getenv("MAIL_PASSWORD"),
		os.Getenv("MAIL_AUTHSERVER"),
	)

	err = smtp.SendMail(
		os.Getenv("MAIL_SERVER"),
		smtpAuth,
		from.Address,
		[]string{os.Getenv("MAIL_TO")},
		buf.Bytes(),
	)
	if err != nil {
		log.Fatal(err)
	}
}
