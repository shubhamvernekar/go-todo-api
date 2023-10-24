package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func PrintTable(todos Todos) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)

	fmt.Fprintln(w, "ID\tTitle\tIsDone")

	for _, todo := range todos {
		fmt.Fprintf(w, "%d\t%s\t%t\n", todo.ID, todo.Desc, todo.IsDone)
	}

	w.Flush()
}
