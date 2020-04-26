package handlers

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"

	"github.com/taipoxin/rest-mysql-go/internal/api/models"
)

type HandlersEnv struct {
	Db models.Datastore
}

// RootHandler - handler for /
func (env *HandlersEnv) RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		errorHandler(w, r, http.StatusMethodNotAllowed, nil)
		return
	}
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound, nil)
		return
	}

	fmt.Fprintf(w, "Hello from root: %q", html.EscapeString(r.URL.Path))
}

// GetPosts - handler for GET /posts
func (env *HandlersEnv) GetPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		errorHandler(w, r, http.StatusMethodNotAllowed, nil)
		return
	}
	posts, err := env.Db.AllPosts()
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError, err)
		return
	}
	fmt.Fprint(w, posts, "\n")
}

// GetPost - handler for GET /post?id=x
func (env *HandlersEnv) GetPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		errorHandler(w, r, http.StatusMethodNotAllowed, nil)
		return
	}
	err := r.ParseForm()
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError, err)
		return
	}
	if len(r.Form["id"]) == 0 {
		errorHandler(w, r, http.StatusBadRequest, err)
		return
	}
	// id is int64
	id, err := strconv.ParseInt(r.Form["id"][0], 10, 64)
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
func (env *HandlersEnv) AddPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		errorHandler(w, r, http.StatusMethodNotAllowed, nil)
		return
	}

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

// UpdatePost - handler for PUT /updatepost
func (env *HandlersEnv) UpdatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		errorHandler(w, r, http.StatusMethodNotAllowed, nil)
		return
	}

	var p models.Post
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = env.Db.UpdatePost(p.ID, p.Title)
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError, err)
		return
	}
	fmt.Fprint(w, "updated successfully", "\n")
}

// DeletePost - handler for Delete /deletepost
func (env *HandlersEnv) DeletePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		errorHandler(w, r, http.StatusMethodNotAllowed, nil)
		return
	}
	err := r.ParseForm()
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError, err)
		return
	}
	if len(r.Form["id"]) == 0 {
		errorHandler(w, r, http.StatusBadRequest, err)
		return
	}
	// id is int64
	id, err := strconv.ParseInt(r.Form["id"][0], 10, 64)
	if err != nil {
		errorHandler(w, r, http.StatusBadRequest, err)
		return
	}

	err = env.Db.DeletePost(id)
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError, err)
		return
	}
	fmt.Fprint(w, "deleted successfully", "\n")
}

// GetWelcomeHandler - handler for /welcome
func (env *HandlersEnv) GetWelcomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		errorHandler(w, r, http.StatusMethodNotAllowed, nil)
		return
	}
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
