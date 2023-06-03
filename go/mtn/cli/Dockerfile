# ベースイメージを指定
FROM golang:latest

# 作業ディレクトリを設定
WORKDIR /app

# Goモジュールの依存関係をコピー
COPY go.mod go.sum ./

# 依存関係をダウンロード
RUN go mod download

# アプリケーションのソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN go build -o todo-webapp cmd/todo-webapp/main.go

# コンテナ起動時に実行するコマンドを指定
CMD ["./todo-webapp"]