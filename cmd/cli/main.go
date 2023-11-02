package main

import (
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	LoadConfig()
	client = Client{
		baseUrl: cfg.BaseURL + cfg.Port,
		cli:     http.DefaultClient,
	}

	app := &cli.App{
		Name:     "todo",
		Usage:    "cli for todo server https://github.com/shubhamvernekar/go-todo-api",
		Flags:    APIFlags,
		Commands: APICommands,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
