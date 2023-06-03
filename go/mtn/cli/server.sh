#!/bin/sh
# -d dockerオプション
if [ "$1" = "-d" ]; then
  docker-compose build --no-cache app
  docker-compose up -d
  exit 0
fi

docker-compose up -d postgres
DATABASE_URL="postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable"
DATABASE_URL=$DATABASE_URL go run ./cmd/todo-webapp/main.go

