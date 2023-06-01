package main

import (
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "Load configuration from `FILE`",
		},
	}
	app.Commands = []*cli.Command{
		{
			Name:  "list",
			Usage: "List students",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "json",
					Usage: "output as JSON",
					Value: false,
				},
			},
			Action: cmdList,
		},
	}
	app.Name = "score"
	app.Usage = "Show student's score"
	app.Run(os.Args)
}

func cmdList(ctx *cli.Context) error {
	return nil
}
