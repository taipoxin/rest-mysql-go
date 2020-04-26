package handlers

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/taipoxin/rest-mysql-go/internal/api/models"
)

// Env - container for hanlders & Datastore
type Env struct {
	Db models.Datastore
}

// RootHandler - handler for /
func (env *Env) RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound, nil)
		return
	}

	fmt.Fprintf(w, "Hello from root: %q", html.EscapeString(r.URL.Path))
}

// GetPosts - handler for GET /posts
func (env *Env) GetPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := env.Db.AllPosts()
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError, err)
		return
	}
	fmt.Fprint(w, posts, "\n")
}

// GetPost - handler for GET /post/{id}
func (env *Env) GetPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// id is int64
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		errorHandler(w, r, http.StatusBadRequest, err)
		return
	}
	post, err := env.Db.GetPost(id)
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError, err)
		return
	}
	if post.Title == "" {
		fmt.Fprint(w, "no posts with id: ", id, "\n")
		return
	}
	fmt.Fprint(w, post, "\n")
}

// AddPost - handler for POST /post
func (env *Env) AddPost(w http.ResponseWriter, r *http.Request) {

	var p models.Post
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = env.Db.AddPost(p.Title)
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError, err)
		return
	}
	fmt.Fprint(w, "inserted successfully", "\n")
}

// UpdatePost - handler for PUT /post
func (env *Env) UpdatePost(w http.ResponseWriter, r *http.Request) {

	var p models.Post
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	isUpdated, err := env.Db.UpdatePost(p.ID, p.Title)
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError, err)
		return
	}
	if isUpdated {
		fmt.Fprint(w, "updated successfully", "\n")
	} else {
		fmt.Fprint(w, "nothing to update", "\n")
	}
}

// DeletePost - handler for Delete /post/{id}
func (env *Env) DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// id is int64
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		errorHandler(w, r, http.StatusBadRequest, err)
		return
	}

	isDeleted, err := env.Db.DeletePost(id)
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError, err)
		return
	}

	if isDeleted {
		fmt.Fprint(w, "deleted successfully", "\n")
	} else {
		fmt.Fprint(w, "nothing to delete", "\n")
	}
}

// GetWelcomeHandler - handler for /welcome
func (env *Env) GetWelcomeHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError, err)
		return
	}
	fmt.Fprint(w, "Welcome to my rest api! \nYour params are: ", r.Form, "\n")
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int, err error) {
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(status)

	switch status {
	case http.StatusNotFound:
		fmt.Fprint(w, "404: not found")
	case http.StatusInternalServerError:
		fmt.Fprint(w, "500: internal error")
	case http.StatusBadRequest:
		fmt.Fprint(w, "400: bad request")
	case http.StatusMethodNotAllowed:
		fmt.Fprint(w, "405: method not allowed")
	default:
		fmt.Fprint(w, status, " unexpected")
	}

}
