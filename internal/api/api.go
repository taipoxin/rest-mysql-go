package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/taipoxin/rest-mysql-go/internal/api/handlers"
	"github.com/taipoxin/rest-mysql-go/internal/api/models"
)

// Start - start the server on host:port
func Start(addr string) {

	db := models.EstablishConnection()
	defer db.Close()

	// use DI scheme from https://www.alexedwards.net/blog/organising-database-access
	handlersEnv := &handlers.Env{
		Db: db,
	}

	r := mux.NewRouter()

	r.HandleFunc("/", handlersEnv.RootHandler).Methods("GET")
	r.HandleFunc("/welcome", handlersEnv.GetWelcomeHandler).Methods("GET")

	r.HandleFunc("/posts", handlersEnv.GetPosts).Methods("GET")

	r.HandleFunc("/post/{id}", handlersEnv.GetPost).Methods("GET")
	r.HandleFunc("/post", handlersEnv.AddPost).Methods("POST")
	r.HandleFunc("/post", handlersEnv.UpdatePost).Methods("PUT")
	r.HandleFunc("/post/{id}", handlersEnv.DeletePost).Methods("DELETE")

	log.Printf("server is listening on addr: %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, r))

}
