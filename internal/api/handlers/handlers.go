package handlers

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

// RootHandler - handler for /
func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound, nil)
		return
	}

	fmt.Fprintf(w, "Hello from root: %q", html.EscapeString(r.URL.Path))
}

// GetWelcomeHandler - handler for /welcome
func GetWelcomeHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError, nil)
		return
	}
	fmt.Fprint(w, "Welcome to my rest api! \nYour params are: ", r.Form)
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
