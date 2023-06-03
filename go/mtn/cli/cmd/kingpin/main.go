package main

import (
	"github.com/alecthomas/kingpin/v2"
	"os"
	"strings"
)

var (
	app      = kingpin.New("score", "Show student's score")
	debug    = app.Flag("debug", "Enable debug mode.").Bool()
	serverIP = app.Flag("server", "Server IP address.").Default("127.0.0.1").IP()

	register     = app.Command("register", "Register a new user.")
	registerNick = register.Arg("nick", "User nickname.").Required().String()
	resisterName = register.Arg("name", "User name.").Required().String()

	post        = app.Command("post", "Post a new score.")
	postImage   = post.Flag("image", "Image file.").File()
	postChannel = post.Arg("channel", "Channel name.").Required().String()
	postText    = post.Arg("text", "Text message.").Required().Strings()
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case register.FullCommand():
		println(*registerNick)
	case post.FullCommand():
		if *postImage != nil {
			println("Image: ", (*postImage).Name())
		}
		text := strings.Join(*postText, " ")
		println("Post: ", text)
	}
}
