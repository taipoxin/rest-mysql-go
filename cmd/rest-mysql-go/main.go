package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/taipoxin/rest-mysql-go/internal/api"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func main() {

	log.Println("Hello world")
	addr := os.Getenv("SERVER_HOST") + ":" + os.Getenv("SERVER_PORT")
	api.Start(addr)
}
