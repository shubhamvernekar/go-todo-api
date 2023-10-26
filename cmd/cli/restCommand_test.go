package main

import (
	"flag"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

var ctx *cli.Context

type DummyHTTPRequestCaller struct {
	responseType string
}

func (h *DummyHTTPRequestCaller) makeHTTPCall(requestType string, route string, jsonString string) ([]byte, error) {
	var responseString string

	if requestType == http.MethodPost {
		if h.responseType == "fail" {
			responseString = `{"hello":--2323}`
		} else if h.responseType == "pass" {
			responseString = `{"id":1,"desc":"dummy todo","is_done":false}`
		}
	}

	if requestType == http.MethodDelete {
		if h.responseType == "fail" {
			return nil, fmt.Errorf("%w http statusCode %d, responseData %s", ErrServerError, 404, "404 page not found")
		} else if h.responseType == "pass" {
			responseString = `{"id":1,"desc":"dummy todo","is_done":false}`
		}
	}

	return []byte(responseString), nil
}

func init() {
	flag := flag.NewFlagSet("title", 1)
	tmp := "hello"
	id := 1
	flag.StringVar(&tmp, "title", "hello", "")
	flag.IntVar(&id, "id", 1, "")

	ctx = cli.NewContext(nil, flag, nil)
}

func TestCreateTodoFailure(t *testing.T) {
	ctx.Set("title", "world")

	requestCaller = &DummyHTTPRequestCaller{
		responseType: "fail",
	}

	err := createTodo(ctx)
	assert.Contains(t, err.Error(), "json unmarshal")
}

func TestCreateTodoSuccess(t *testing.T) {
	ctx.Set("title", "world")

	requestCaller = &DummyHTTPRequestCaller{
		responseType: "pass",
	}

	err := createTodo(ctx)
	assert.NoError(t, err)
}

func TestDeleteTodoFail(t *testing.T) {
	ctx.Set("id", "1")

	requestCaller = &DummyHTTPRequestCaller{
		responseType: "fail",
	}

	err := deleteTodo(ctx)
	assert.Contains(t, err.Error(), "404 page not found")
}

func TestDeleteTodoSuccess(t *testing.T) {
	ctx.Set("id", "1")

	requestCaller = &DummyHTTPRequestCaller{
		responseType: "success",
	}

	err := deleteTodo(ctx)
	assert.NoError(t, err)
}
