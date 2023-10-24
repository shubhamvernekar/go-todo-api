package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:     "todo",
		Usage:    "cli for todo server https://github.com/shubhamvernekar/go-todo-api",
		Flags:    ApiFlags,
		Commands: ApiCommands,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
