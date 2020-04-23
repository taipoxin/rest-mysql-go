package handlers

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/taipoxin/rest-mysql-go/internal/api/models"
)

type HandlersEnv struct {
	Db models.Datastore
}

// RootHandler - handler for /
func (env *HandlersEnv) RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound, nil)
		return
	}

	fmt.Fprintf(w, "Hello from root: %q", html.EscapeString(r.URL.Path))
}

// GetWelcomeHandler - handler for /welcome
func (env *HandlersEnv) GetWelcomeHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError, err)
		return
	}
	fmt.Fprint(w, "Welcome to my rest api! \nYour params are: ", r.Form, "\n")

	posts, err := env.Db.AllPosts()
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError, err)
		return
	}

	fmt.Fprint(w, posts, "\n")
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int, err error) {
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "custom 404")
	}
	if status == http.StatusInternalServerError {
		fmt.Fprint(w, "custom 500")
	}

}
