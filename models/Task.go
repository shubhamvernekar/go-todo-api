package models

type Task struct {
	ID     int    `json:id`
	Desc   string `json:desc`
	IsDone bool   `json:isDone`
}
