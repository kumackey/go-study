package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"github.com/mattn/go-runewidth"
	"log"
	"strings"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.ASCII = true
	g.SetManagerFunc(layout)

	// CTRL-Cで終了
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func quit(_ *gocui.Gui, _ *gocui.View) error {
	return gocui.ErrQuit
}

func layout(g *gocui.Gui) error {
	s := `
Go is an open source programming language 
that makes it easy to build simple, 
reliable, and efficient software.
`
	maxX, maxY := g.Size()
	if v, err := g.SetView("hello", maxX/2-21, maxY/2-2, maxX/2+21, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, s)
	}

	return nil
}

func goRunewidth() {
	s := "Hello, 世界だよね。君の世界はすごい"
	fmt.Println(s)
	width := runewidth.StringWidth(s)
	fmt.Println(strings.Repeat("~", width))
	fmt.Println(runewidth.Wrap(s, 11))
}
