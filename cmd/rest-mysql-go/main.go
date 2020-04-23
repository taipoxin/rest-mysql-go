package main

import (
	"log"

	"github.com/taipoxin/rest-mysql-go/internal/api"
)

func main() {
	log.Println("Hello world")
	api.Start("localhost", "8080")
}
