package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/urfave/cli/v2"
)

var ApiFlags = []cli.Flag{
	&cli.StringFlag{
		Name: "title",
	},
	&cli.IntFlag{
		Name: "id",
	},
}

var ApiCommands = []*cli.Command{
	{
		Name:   "get_all",
		Usage:  "Get all todos list",
		Action: getAllTodos,
	},
	{
		Name:   "get",
		Usage:  "Get specific todos by id",
		Action: getTodo,
	},
	{
		Name:   "create",
		Usage:  "Create todo",
		Action: createTodo,
	},
	{
		Name:   "mark_done",
		Usage:  "Toggle done status of todo by id",
		Action: markDone,
	},
	{
		Name:   "delete",
		Usage:  "delete todo by given id",
		Action: deleteTodo,
	},
}

type Todo struct {
	ID     int    `json:id`
	Desc   string `json:Desc`
	IsDone bool   `json:IsDone`
}

type Todos []struct {
	Todo
}

func getAllTodos(ctx *cli.Context) error {
	url := BASE_URL + PORT + GETALL

	response, err := http.Get(url)

	if err != nil || response.StatusCode != 200 {
		log.Fatal("Error : ", err)
		os.Exit(0)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal("Error : ", err)
		os.Exit(0)
	}

	todos := Todos{}
	err = json.Unmarshal(responseData, &todos)

	if err != nil {
		log.Fatal("Error : ", err)
		os.Exit(0)
	}

	PrintTable(todos)
	return nil
}

func getTodo(ctx *cli.Context) error {
	url := BASE_URL + PORT + GET
	id := ctx.Int("id")

	if id == 0 {
		log.Fatal("Error : Invalid todo id provided")
		os.Exit(0)
	}

	url += "/" + strconv.Itoa(id)

	response, err := http.Get(url)

	if err != nil || response.StatusCode != 200 {
		log.Fatal("Error : ", err)
		os.Exit(0)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}

	todo := Todo{}
	err = json.Unmarshal(responseData, &todo)

	if err != nil {
		log.Fatal("Error : ", err)
		os.Exit(0)
	}

	todos := Todos{struct{ Todo }{todo}}
	PrintTable(todos)
	return nil
}

func createTodo(ctx *cli.Context) error {
	url := BASE_URL + PORT + POST
	title := ctx.String("title")

	if len(title) == 0 {
		log.Fatal("Error : Invalid todo title provided")
		os.Exit(0)
	}

	jsonBody := []byte(fmt.Sprintf(`{"desc": "%s"}`, title))
	bodyReader := bytes.NewReader(jsonBody)
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, url, bodyReader)
	if err != nil {
		log.Fatal("Error : ", err)
		os.Exit(0)
	}

	response, err := client.Do(req)
	if err != nil || response.StatusCode != 200 {
		log.Fatal("Error : ", err)
		os.Exit(0)
	}

	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}

	todo := Todo{}
	err = json.Unmarshal(responseData, &todo)

	if err != nil {
		log.Fatal("Error : ", err)
		os.Exit(0)
	}

	todos := Todos{struct{ Todo }{todo}}
	PrintTable(todos)
	return nil
}

func deleteTodo(ctx *cli.Context) error {
	url := BASE_URL + PORT + DELETE
	id := ctx.Int("id")

	if id == 0 {
		log.Fatal("Error : Invalid todo id provided")
		os.Exit(0)
	}

	url += "/" + strconv.Itoa(id)
	client := &http.Client{}

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Fatal("Error : ", err)
		os.Exit(0)
	}

	response, err := client.Do(req)
	if err != nil || response.StatusCode != 200 {
		log.Fatal("Error : ", err)
		os.Exit(0)
	}
	defer response.Body.Close()

	fmt.Printf("Todo %d Deleted successfully", id)

	return nil
}

func markDone(ctx *cli.Context) error {
	url := BASE_URL + PORT + MARKDONE
	id := ctx.Int("id")

	if id == 0 {
		log.Fatal("Error : Invalid todo id provided")
		os.Exit(0)
	}

	url += "/" + strconv.Itoa(id)

	response, err := http.Get(url)

	if err != nil || response.StatusCode != 200 {
		log.Fatal("Error : ", err)
		os.Exit(0)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}

	todo := Todo{}
	err = json.Unmarshal(responseData, &todo)

	if err != nil {
		log.Fatal("Error : ", err)
		os.Exit(0)
	}

	todos := Todos{struct{ Todo }{todo}}
	PrintTable(todos)
	return nil
}
