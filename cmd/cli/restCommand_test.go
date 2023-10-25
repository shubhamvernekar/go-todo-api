package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestCreateTodoFailure(t *testing.T) {
	ctx := new(cli.Context)
	ctx.Set("title", "demo 1")

	err := createTodo(ctx)
	assert.Equal(t, "server unaviable", err)
}

func TestCreateTodoSuccess(t *testing.T) {
	err := createTodo(nil)
	assert.NoError(t, err)
}
