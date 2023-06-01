package main

import (
	"flag"
	"fmt"
	"os"
)

var commandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

func main() {
	os.Exit(run(os.Args[1:]))
}

func run(args []string) int {
	max := commandLine.Int("max", 255, "max value")
	name := commandLine.String("name", "default", "name value")
	if err := commandLine.Parse(args); err != nil {
		fmt.Fprintf(os.Stderr, "parse error: %v\n", err)
		return 1
	}

	if *max > 999 {
		fmt.Fprintf(os.Stderr, "max value is too big: %d\n", *max)
		return 1
	}
	if *name == "" {
		fmt.Fprintf(os.Stderr, "name value is empty\n")
		return 1
	}

	return 0
}
