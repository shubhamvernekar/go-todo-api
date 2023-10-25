package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var (
	GetAllTodo = "/tasks"
	GetTodo    = "/task"
	PostTodo   = "/task"
	DeleteTodo = "/task"
	MarkDone   = "/task/markDone"
)

var cfg Config

type Config struct {
	BaseURL string `default:"http://localhost" envconfig:"BASEURL"`
	Port    string `default:":4000"            envconfig:"PORT"`
}

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Couldn't load env error : %v", err)
	}

	err = envconfig.Process("TODO", &cfg)
	if err != nil {
		log.Fatalf("Couldn't load config error : %v", err)
	}
}
