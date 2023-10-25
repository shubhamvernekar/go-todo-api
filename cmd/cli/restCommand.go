package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/urfave/cli/v2"
)

var APIFlags = []cli.Flag{
	&cli.StringFlag{
		Name: "title",
	},
	&cli.IntFlag{
		Name: "id",
	},
}

var APICommands = []*cli.Command{
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
	ID     int    `json:"id"`
	Desc   string `json:"desc"`
	IsDone bool   `json:"is_done"`
}

type Todos []struct {
	Todo
}

func getAllTodos(_ *cli.Context) error {
	responseData, err := makeHTTPCall(http.MethodGet, GetAllTodo, "")
	if err != nil {
		return fmt.Errorf("Error : %w", err)
	}

	todos := Todos{}
	err = json.Unmarshal(responseData, &todos)

	if err != nil {
		return fmt.Errorf("Error : %w", err)
	}

	PrintTable(todos)
	return nil
}

func getTodo(ctx *cli.Context) error {
	id := ctx.Int("id")

	if id == 0 {
		return fmt.Errorf("Error : %w", errors.New("Invalid todo id provided"))
	}

	url := fmt.Sprintf("%s/%d", GetTodo, id)

	responseData, err := makeHTTPCall(http.MethodGet, url, "")
	if err != nil {
		return fmt.Errorf("Error : %w", err)
	}

	todo := Todo{}
	err = json.Unmarshal(responseData, &todo)

	if err != nil {
		return fmt.Errorf("Error : %w", err)
	}

	todos := Todos{struct{ Todo }{todo}}
	PrintTable(todos)
	return nil
}

func createTodo(ctx *cli.Context) error {
	title := ctx.String("title")

	if len(title) == 0 {
		return fmt.Errorf("Error : Invalid todo title provided")
	}

	jsonString := fmt.Sprintf(`{"desc":"%s"}`, title)

	responseData, err := makeHTTPCall(http.MethodPost, PostTodo, jsonString)
	if err != nil {
		return fmt.Errorf("Error : %w", err)
	}

	todo := Todo{}
	err = json.Unmarshal(responseData, &todo)

	if err != nil {
		return fmt.Errorf("Error : %w", err)
	}

	todos := Todos{struct{ Todo }{todo}}
	PrintTable(todos)
	return nil
}

func deleteTodo(ctx *cli.Context) error {
	id := ctx.Int("id")

	if id == 0 {
		return fmt.Errorf("Error : Invalid todo id provided")
	}

	url := fmt.Sprintf("%s/%d", DeleteTodo, id)

	_, err := makeHTTPCall(http.MethodDelete, url, "")
	if err != nil {
		return fmt.Errorf("Error : %w", err)
	}

	log.Printf("Todo %d Deleted successfully", id)
	return nil
}

func markDone(ctx *cli.Context) error {
	id := ctx.Int("id")

	if id == 0 {
		return fmt.Errorf("Error : Invalid todo id provided")
	}

	url := fmt.Sprintf("%s/%d", MarkDone, id)

	responseData, err := makeHTTPCall(http.MethodGet, url, "")
	if err != nil {
		return fmt.Errorf("Error : %w", err)
	}

	todo := Todo{}
	err = json.Unmarshal(responseData, &todo)
	if err != nil {
		return fmt.Errorf("Error : %w", err)
	}

	todos := Todos{struct{ Todo }{todo}}
	PrintTable(todos)
	return nil
}

func makeHTTPCall(requestType string, route string, jsonString string) ([]byte, error) {
	client := &http.Client{}
	url := cfg.BaseURL + cfg.Port + route

	var bodyReader io.Reader

	if len(jsonString) > 0 {
		jsonBody := []byte(jsonString)
		bodyReader = bytes.NewReader(jsonBody)
	}

	request, err := http.NewRequest(requestType, url, bodyReader)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("http statusCode %d, responseData %s", response.StatusCode, responseData)
		return nil, err
	}

	return responseData, nil
}
