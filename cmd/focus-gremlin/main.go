package main

import (
	"fmt"
	"os"

	"github.com/alexthestreet/focus-gremlin/internal/cli"
)

func main() {
	if err := cli.NewRootCommand().Execute(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
