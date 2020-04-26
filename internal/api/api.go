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
	handlersEnv := &handlers.Env{
		Db: db,
	}

	http.HandleFunc("/", handlersEnv.RootHandler)
	http.HandleFunc("/welcome", handlersEnv.GetWelcomeHandler)

	http.HandleFunc("/posts", handlersEnv.GetPosts)
	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			handlersEnv.GetPost(w, r)
		case "POST":
			handlersEnv.AddPost(w, r)
		case "PUT":
			handlersEnv.UpdatePost(w, r)
		case "DELETE":
			handlersEnv.DeletePost(w, r)
		}
	})

	log.Printf("server is listening on addr: %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))

}
