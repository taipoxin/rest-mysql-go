package api

import (
	"log"
	"net/http"

	"github.com/taipoxin/rest-mysql-go/internal/api/handlers"
)

// Start - start the server on host:port
func Start(host string, port string) {
	addr := host + ":" + port
	http.HandleFunc("/", handlers.RootHandler)
	http.HandleFunc("/welcome", handlers.GetWelcomeHandler)

	log.Printf("Server is listening on addr: %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))

}
