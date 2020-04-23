package api

import (
	"log"
	"net/http"

	"github.com/taipoxin/rest-mysql-go/internal/api/handlers"
	"github.com/taipoxin/rest-mysql-go/internal/api/models"
)

// Start - start the server on host:port
func Start(addr string) {

	db := models.EstablishConnection()
	defer db.Close()

	// use DI scheme from https://www.alexedwards.net/blog/organising-database-access
	env := &handlers.HandlersEnv{
		Db: db,
	}

	http.HandleFunc("/", env.RootHandler)
	http.HandleFunc("/welcome", env.GetWelcomeHandler)

	log.Printf("server is listening on addr: %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))

}
